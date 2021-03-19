package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type User struct {
	IP          string `json:"ip"`
	Allow       bool   `json:"allow"`
	IsAdmin     bool   `json:"isAdmin"`
	Description string `json:"description"`
}

var Whitelist = []User{}

func IPWhitelist() gin.HandlerFunc {
	return func(c *gin.Context) {
		for i := 0; i < len(Whitelist); i++ {
			if Whitelist[i].IP == c.ClientIP() {
				if Whitelist[i].Allow {
					return
				}
				break
			}
		}
		c.AbortWithStatus(http.StatusProcessing)
	}
}

func AdminList() gin.HandlerFunc {
	return func(c *gin.Context) {
		for i := 0; i < len(Whitelist); i++ {
			if Whitelist[i].IP == c.ClientIP() {
				if Whitelist[i].IsAdmin {
					return
				}
				break
			}
		}
		c.AbortWithStatus(http.StatusForbidden)
	}
}

func LoadWhitelist() error {
	if _, err := os.Stat("whitelist.json"); os.IsNotExist(err) {
		err = createWhitelistJSON()
		if err != nil {
			return err
		}
	}

	whitelistJSON, err := os.ReadFile("whitelist.json")
	if err == os.ErrNotExist {
		return err
	}
	err = json.Unmarshal(whitelistJSON, &Whitelist)
	if err != nil {
		return err
	}
	return nil
}

func createWhitelistJSON() error {
	JsonString := `[{"ip":"::1","allow":true,"isAdmin":true,"description":"localhost IPv6"},{"ip":"127.0.0.1","allow":true,"isAdmin":true,"description":"localhost IPv4"}]`
	if err := ioutil.WriteFile("whitelist.json", []byte(JsonString), 0644); err != nil {
		return err
	}
	return nil
}

func UpdateWhitelist(user User) error {
	userExist := false
	for i := 0; i < len(Whitelist); i++ {
		if Whitelist[i].IP == user.IP {
			Whitelist[i].Description = user.Description
			Whitelist[i].Allow = user.Allow
			Whitelist[i].IsAdmin = user.IsAdmin
			userExist = true
			break
		}
	}
	if !userExist {
		Whitelist = append(Whitelist, user)
	}

	whitelistJSON, err := json.Marshal(Whitelist)
	if err != nil {
		return err
	}

	f, err := os.OpenFile("whitelist.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	f.Write(whitelistJSON)
	f.Close()
	return nil
}
