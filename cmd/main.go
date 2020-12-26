package main

import (
  "flag"
  "fmt"

  "github.com/zhiruchen/codesearch"
)

func main() {
  keywords := flag.String("kws", "", "the keywords to search code")
  flag.Parse()

  if *keywords == ""{
    fmt.Println("empty keywords")
    return
  }

  fmt.Println("searching code from stackoverflow/github/bitbucket/gitlab")
  results := codesearch.RunSearch(*keywords)

  for i, r := range results {
    fmt.Println("result: ", i, " , ", *r)
  }
}