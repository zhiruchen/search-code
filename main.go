package codesearch

import (
	"fmt"
	"sync"
)

const defaultResultCap = 10

// Result search result
type Result struct {
	Title     string
	URL       string
	ShortDesc string
}

// CodeSearcher the searcher used to search code
type CodeSearcher interface {
	Origin() string // the origin site
	Search(keywords string) ([]*Result, error)
}

func RunSearch(keywords string) []*Result {
	searchers := []CodeSearcher{
		NewGithubSearcher(),
		NewGitlabSearcher(),
		NewSOFSearcher(),
	}

	resultCh := searchFromSites(keywords, searchers)
	fmt.Print("after searchFromSites")

	searchResults := make([]*Result, 0, defaultResultCap)
	for r := range resultCh {
		searchResults = append(searchResults, &r)
	}
	fmt.Print("Searching Done, Find ", len(searchResults), " Results")
	return searchResults
}

func searchFromSites(keywords string, searchers []CodeSearcher) chan Result {
	var wg sync.WaitGroup
	var resultCh = make(chan Result, defaultResultCap)
	for _, searcher := range searchers {
		wg.Add(1)
		go func(searcher CodeSearcher) {
			defer wg.Done()
			fmt.Println("Searching from: ", searcher.Origin())
			results, err := searcher.Search(keywords)
			if err != nil {
				fmt.Println("search err: ", err)
				return
			}

			if len(results) == 0 {
				fmt.Println("no results found")
				return
			}

			for _, r := range results {
				resultCh <- *r
			}
		}(searcher)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()
	return resultCh
}
