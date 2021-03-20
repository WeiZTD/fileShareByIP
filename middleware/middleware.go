package middleware

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AuthInfo struct {
	Allow       bool   `json:"allow"`
	IsAdmin     bool   `json:"isAdmin"`
	Description string `json:"description"`
}

var Whitelist = make(map[string]AuthInfo)

func IPWhitelist() gin.HandlerFunc {
	return func(c *gin.Context) {
		if Whitelist[c.ClientIP()].Allow {
			return
		}
		c.AbortWithStatus(http.StatusProcessing)
	}
}

func AdminList() gin.HandlerFunc {
	return func(c *gin.Context) {
		if Whitelist[c.ClientIP()].IsAdmin {
			return
		}
		c.AbortWithStatus(http.StatusForbidden)
	}
}

func LoadWhitelist() error {
	if _, err := os.Stat("whitelist.json"); os.IsNotExist(err) {
		whitelistJSON := []byte(`{
			"::1":{
				"ip": "::1",
				"allow": true,
				"isAdmin": true,
				"description": "localhost IPv6"
			},
			"127.0.0.1":{
				"allow": true,
				"isAdmin": true,
				"description": "localhost IPv4"
			}
}`)
		if err := createWhitelistJSON(whitelistJSON); err != nil {
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

func UpdateWhitelist(ip string, authinfo AuthInfo) error {
	Whitelist[ip] = authinfo

	whitelistJSON, err := json.Marshal(Whitelist)
	if err != nil {
		return err
	}

	if err := createWhitelistJSON(whitelistJSON); err != nil {
		return err
	}

	return nil
}

func createWhitelistJSON(whitelistJSON []byte) error {
	f, err := os.OpenFile("whitelist.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	if _, err := f.Write(whitelistJSON); err != nil {
		return err
	}
	f.Close()
	return nil
}
