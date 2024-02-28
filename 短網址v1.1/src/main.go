package main

import (
	"shortner/api"
	"shortner/db"

	"github.com/gin-gonic/gin"
)

//長網址丟進來->MD5加密後變成短網址->存進資料庫OR暫存 (POST)
//long url => server => short url

// 搜尋短網址 -> 轉址->藏一段HTML方便轉譯 (GET)
// short url => short server (重定向)=> long url
// 短網址在上方欄位出現時會打GET API 到此程式 因而搜尋在資料表中的code獲得原本地址
func init() {
	if err := db.Connect("127.0.0.1", "3306", "test2", "root", "root"); err != nil {
		panic(err)
	}
}
func main() {
	router := gin.Default()
	router.POST("/shortner", api.AddLongURLToShortURL)

	router.GET("/shortner/:code", api.ToAnotherWebsite)
	router.Run(":8080")

	db.Disconnect()
}

//優化錯誤處理:
//驗證正確的網址格式
//驗證資料庫沒有重複的code
//GET code(獨一無二) 應該傳出原本url
