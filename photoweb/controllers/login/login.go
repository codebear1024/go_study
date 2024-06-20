package login

import (
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

const (
	glicense = "123456"
)

type user struct {
	userName string
	password string
}

var userMap = make(map[string]user)

func (c *LoginController) renderHtml(tmpl string, key interface{}, data interface{}) {
	c.Data[key] = data
	c.TplName = tmpl
}

func (c *LoginController) Get() {
	c.renderHtml("login.tpl", "prompt", "")
}

func (c *LoginController) Post() {
	userName := c.GetString("username")
	password := c.GetString("password")
	// fmt.Println(userName + "===" + password)
	if len(userName) == 0 || password != userMap[userName].password {
		c.renderHtml("login.tpl", "prompt", "登陆失败，请检查账号和密码")
		return
	}
	c.Redirect("/", http.StatusFound)
}

func (c *LoginController) RegisterGet() {
	c.renderHtml("register.tpl", "prompt", "")
}

func (c *LoginController) RegisterPost() {
	userName := c.GetString("username")
	passWord := c.GetString("passWord")
	passwdConfirm := c.GetString("passwordconfirm")
	license := c.GetString("license")
	if len(userName) == 0 || len(passWord) == 0 {
		c.renderHtml("register.tpl", "prompt", "注册失败，账号或密码为空")
		return
	}
	if userMap[userName].userName == userName {
		c.renderHtml("register.tpl", "prompt", "注册失败，用户名已存在")
		return
	}
	if passWord != passwdConfirm {
		c.renderHtml("register.tpl", "prompt", "注册失败，两次输入密码不相同")
		return
	}
	if license != glicense {
		c.renderHtml("register.tpl", "prompt", "注册失败，许可码错误")
		return
	}
	userMap[userName] = user{userName: userName, password: passWord}
	c.Redirect("/login", http.StatusFound)
}
