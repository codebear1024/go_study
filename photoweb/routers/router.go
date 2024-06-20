package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"photoweb/controllers/login"
	"photoweb/controllers/photo"
)

func init() {
	beego.Router("/login", &login.LoginController{})
	beego.Router("/register", &login.LoginController{}, "post:RegisterPost;get:RegisterGet")
	beego.Router("/", &photo.PhotoController{})
	beego.Router("/view", &photo.PhotoController{}, "get:View")
	beego.Router("/upload", &photo.PhotoController{}, "get:UploadGet;post:UploadPost")
}
