package main

import (
	"github.com/DroiTaipei/droictx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setDroiHeader(c *gin.Context, req *http.Request) {
	for hk, _ := range droictx.IFieldHeaderKeyMap() {
		if v := c.Request.Header.Get(hk); len(v) > 0 {
			req.Header.Set(hk, v)
		}
	}
}
