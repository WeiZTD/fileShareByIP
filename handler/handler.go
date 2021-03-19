package handler

import (
	"fileShareByIP/middleware"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func RedirectToFile(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/file")
}

func AdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.tmpl", gin.H{"Users": &middleware.Whitelist})
}

func AdminAction(c *gin.Context) {
	action := c.Param("action")
	switch action {
	case "updateWhitelist":
		tempIP, receive := c.GetPostForm("IP")
		ip := strings.TrimSpace(tempIP)
		if !receive || len(ip) < 1 {
			c.HTML(http.StatusOK, "admin/index.tmpl", gin.H{"alert": "IP address is empty", "Users": &middleware.Whitelist})
			return
		}
		allowStr := c.PostForm("Allow")
		allow, _ := strconv.ParseBool(allowStr)
		isAdminStr := c.PostForm("IsAdmin")
		isAdmin, _ := strconv.ParseBool(isAdminStr)
		description := c.PostForm("Description")
		user := middleware.User{
			IP:          ip,
			Allow:       allow,
			IsAdmin:     isAdmin,
			Description: description,
		}
		middleware.UpdateWhitelist(user)

		c.Redirect(http.StatusFound, "/admin")
	default:
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
