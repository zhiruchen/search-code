package codesearch

type glb struct{}

func NewGitlabSearcher() CodeSearcher {
	return &glb{}
}

func (g *glb) Origin() string {
	return "Gitlab"
}

func (g *glb) Search(keywords string) ([]*Result, error) {
	return nil, nil
}
