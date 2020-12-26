package codesearch

type sof struct{}

func NewSOFSearcher() CodeSearcher {
	return &sof{}
}

func (s *sof) Origin() string {
	return "Stackoverflow"
}

func (s *sof) Search(keywords string) ([]*Result, error) {
	return nil, nil
}
