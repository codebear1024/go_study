package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"study/photoweb/sessionmanger"
)

const (
	UploadDir = "./upload"
	ViewDir   = "./view"
	glicense  = "123456"
)

type user struct {
	userName string
	password string
}

var templates = make(map[string]*template.Template)
var userMap = make(map[string]user)
var globalSessionManger *sessionmanger.SessionManger

func init() {
	fileInfoArr, err := os.ReadDir(ViewDir)
	if err != nil {
		panic(err.Error())
		return
	}
	for _, fileInfo := range fileInfoArr {
		templateName := fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath := ViewDir + "/" + templateName
		log.Println("Loading template: ", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
	}
	globalSessionManger = sessionmanger.NewSessionManger("sessionid", 3600)
	go globalSessionManger.SessionGC()
}

func renderHtml(w http.ResponseWriter, tmpl string, data interface{}) (err error) {
	t, err := template.ParseFiles(ViewDir + "/" + tmpl)
	if err != nil {
		return
	}
	err = t.Execute(w, data)
	// err = templates[tmpl].Execute(w, data)
	return
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	session := globalSessionManger.SessionGet(w, r)
	if session == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	globalSessionManger.SessionUpdate(session.SessionId())
	fileInfoArr, err := os.ReadDir(UploadDir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var image []string
	for _, fileInfo := range fileInfoArr {
		image = append(image, fileInfo.Name())
	}
	if err = renderHtml(w, "list.html", image); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	session := globalSessionManger.SessionGet(w, r)
	if session == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	globalSessionManger.SessionUpdate(session.SessionId())
	if r.Method == "GET" {
		if err := renderHtml(w, "upload.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UploadDir + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	session := globalSessionManger.SessionGet(w, r)
	if session == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	globalSessionManger.SessionUpdate(session.SessionId())
	imageId := r.FormValue("id")
	imagePath := UploadDir + "/" + imageId
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := renderHtml(w, "login.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		userName := r.FormValue("username")
		password := r.FormValue("password")
		if len(userName) == 0 || password != userMap[userName].password {
			_ = renderHtml(w, "login.html", "登陆失败，请检查账号和密码")
			return
		}
		session := globalSessionManger.SessionStart(w, r)
		if session != nil {
			globalSessionManger.SessionUpdate(session.SessionId())
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := renderHtml(w, "register.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		userName := r.FormValue("username")
		password := r.FormValue("password")
		passwdConfirm := r.FormValue("passwordconfirm")
		license := r.FormValue("license")
		if len(userName) == 0 || len(password) == 0 {
			_ = renderHtml(w, "register.html", "注册失败，账号或密码为空")
			return
		}
		if userMap[userName].userName == userName {
			_ = renderHtml(w, "register.html", "注册失败，用户名已存在")
			return
		}
		if password != passwdConfirm {
			_ = renderHtml(w, "register.html", "注册失败，两次输入密码不相同")
			return
		}
		if license != glicense {
			_ = renderHtml(w, "register.html", "注册失败，许可码错误")
			return
		}
		userMap[userName] = user{userName: userName, password: password}
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func main() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	err := http.ListenAndServe(":8999", nil)
	if err != nil {
		log.Println(err.Error())
	}
}
