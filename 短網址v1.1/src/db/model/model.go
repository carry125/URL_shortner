package model

const Tableshortner string = "shortner"

type Shortner struct {
	ID      int64  `json:"-" gorm:"column:id;primaryKey"`
	LongURL string `json:"long_url" gorm:"column:long_url"`
	Code    string `json:"code" gorm:"column:code"`
}
