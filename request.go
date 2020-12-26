package codesearch

import (
	"github.com/parnurzeal/gorequest"
)

func SendSearhRequest(url, keywords string) (resp string, err error) {
	request := gorequest.New()
	_, body, errs := request.Get(url).End()
	if len(errs) > 0 {
		return "", errs[0]
	}

	return body, nil
}
