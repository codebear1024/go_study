package photo

import (
	"io"
	"net/http"
	"os"

	beego "github.com/beego/beego/v2/server/web"
)

type PhotoController struct {
	beego.Controller
}

const (
	uploadDir = "./upload"
	viewDir   = "./view"
)

func (c *PhotoController) Get() {
	fileInfoArr, err := os.ReadDir(uploadDir)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.WriteString(err.Error())
		return
	}
	var image []string
	for _, fileInfo := range fileInfoArr {
		image = append(image, fileInfo.Name())
	}
	c.Data["filename"] = image
	c.TplName = "list.tpl"
}

func (c *PhotoController) UploadGet() {
	c.TplName = "upload.tpl"
}

func (c *PhotoController) UploadPost() {
	f, h, err := c.GetFile("image")
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.WriteString(err.Error())
		return
	}
	filename := h.Filename
	defer f.Close()
	t, err := os.Create(uploadDir + "/" + filename)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.WriteString(err.Error())
		return
	}
	defer t.Close()
	if _, err := io.Copy(t, f); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.WriteString(err.Error())
		return
	}
	c.Redirect("/", http.StatusFound)
}

func (c *PhotoController) View() {
	imageId := c.GetString("id")
	imagePath := uploadDir + "/" + imageId
	action := c.GetString("action")
	if action == "del" {
		if err := os.Remove(imagePath); err != nil {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
			c.Ctx.WriteString(err.Error())
		}
		c.Redirect("/", http.StatusFound)
		return
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "image")
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, imagePath)
}
