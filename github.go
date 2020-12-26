package codesearch

type ghb struct{}

func NewGithubSearcher() CodeSearcher {
	return &ghb{}
}

func (g *ghb) Origin() string {
	return "Github"
}

func (g *ghb) Search(keywords string) ([]*Result, error) {
	return nil, nil
}
