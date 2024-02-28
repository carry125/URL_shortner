package create

import (
	"fmt"
	"shortner/crypt"
	"shortner/db"
	"shortner/db/model"
)

func ToShortURL(longurl string) (code string) {
	//var md model.Shortner
	code = crypt.MD5(longurl)

	md := model.Shortner{
		LongURL: longurl,
		Code:    code,
	}

	if err := db.Get().Debug().Create(&md); err != nil {
		fmt.Println(err)
	}
	return code
}
