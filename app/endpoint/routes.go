package endpoint

import (
	"github.com/kataras/iris"
	"github.com/dgrijalva/jwt-go"
	jwtM "github.com/iris-contrib/middleware/jwt"
	postM "github.com/youkyll/goshepherd/app/model/post"
	folderM "github.com/youkyll/goshepherd/app/model/folder"
	"encoding/base64"
)

func Routes() {

	myJwtM := jwtM.New(jwtM.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			decoded, err := base64.URLEncoding.DecodeString("cmnsb57ULnO1wEbFYMaRWjc-VNq8GuEVtRCSJuspjJ0WYkT0X36e4e75foXkTduZ")
			if err != nil {
				return nil, err
			}
			return decoded, nil
		},
	})


	api := iris.Party("/api")
	{

		//middleware
		api.Use(myJwtM)

		//投稿
		api.Get("/posts", list)
		api.Post("/posts", create)
		api.Get("/posts/:post_id", detail)
		api.Patch("/posts/:post_id", update)
		api.Delete("/posts/:post_id", delete)

		//フォルダ
		api.Get("/folders", markList)
		api.Post("/folders", markCreate)
		api.Delete("/folders/:folder_id", markDelete)
	}
}

func list(c *iris.Context) {

	if fw := c.Param(`fw`);fw != "" {
		posts := postM.SearchTitle(fw)
		c.JSON(200, posts)
	} else {
		posts := postM.All()
		c.JSON(200, posts)
	}
}

func create(c *iris.Context) {

}
func detail(c *iris.Context) {

}
func update(c *iris.Context) {

}
func delete(c *iris.Context) {

}

func markList(c *iris.Context) {
	folders := folderM.All()

	c.JSON(200, folders)
}

func markCreate(c *iris.Context) {

}

func markDelete(c *iris.Context) {

}
