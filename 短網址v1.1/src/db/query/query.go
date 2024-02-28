package query

import (
	"shortner/db"
	"shortner/db/model"
)

func SelectCode(code string) (url string) {
	var a model.Shortner
	db.Get().Take(&a, "code=?", code)
	url = a.LongURL
	return url
}
