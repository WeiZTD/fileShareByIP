package handler

import (
	"fileShareByIP/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func RedirectToFile(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/file")
}

func FileUploadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.tmpl", nil)
}

func UploadFile(c *gin.Context) {

	file, err := c.FormFile("File")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprint(err.Error()))
		return
	}
	shareDir, _ := c.Get("shareDir")
	filePath := shareDir.(string) + "/" + file.Filename
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		c.HTML(http.StatusBadRequest, "upload.tmpl", gin.H{"alert": "ERROR: file already exist"})
		return
	}

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.HTML(http.StatusBadRequest, "upload.tmpl", gin.H{"alert": fmt.Sprintf("ERROR: %v", fmt.Sprint(err.Error()))})
		return
	}
	c.HTML(http.StatusOK, "upload.tmpl", gin.H{"alert": fmt.Sprintf("%s uploaded!", file.Filename)})
}

func AdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.tmpl", gin.H{"whitelist": &middleware.Whitelist})
}

func AdminAction(c *gin.Context) {
	action := c.Param("action")
	switch action {
	case "updateWhitelist":
		tempIP, receive := c.GetPostForm("IP")
		ip := strings.TrimSpace(tempIP)
		if !receive || len(ip) < 1 {
			c.HTML(http.StatusOK, "admin/index.tmpl", gin.H{"alert": "IP address is empty", "whitelist": &middleware.Whitelist})
			return
		}
		allowStr := c.PostForm("Allow")
		allow, _ := strconv.ParseBool(allowStr)
		isAdminStr := c.PostForm("IsAdmin")
		isAdmin, _ := strconv.ParseBool(isAdminStr)
		description := c.PostForm("Description")
		authInfo := middleware.AuthInfo{
			Allow:       allow,
			IsAdmin:     isAdmin,
			Description: description,
		}
		middleware.UpdateWhitelist(ip, authInfo)

		c.Redirect(http.StatusFound, "/admin")
	default:
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
