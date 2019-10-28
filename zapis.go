// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zapis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	GetConfigString func() string
	GetDebugMode    func() bool
	SetDebugMode    func(bool)
	GetHealth       func() bool
}

var z Handler

func Bind(r *gin.Engine, h Handler) {
	// bind health check API
	r.GET("configz", getConfig)
	r.GET("healthz", healthz)
	r.GET("modez", getMode)
	r.POST("modez", postMode)
}

func getConfig(c *gin.Context) {
	if z.GetConfigString == nil {
		c.String(http.StatusNotImplemented, "not implemented\n")
		return
	}
	c.String(http.StatusOK, z.GetConfigString())
}

func getMode(c *gin.Context) {
	if z.GetDebugMode == nil {
		c.String(http.StatusNotImplemented, "not implemented\n")
		return
	}
	if z.GetDebugMode() {
		c.String(http.StatusOK, "debug\n")
	} else {
		c.String(http.StatusOK, "release\n")
	}
}

func postMode(c *gin.Context) {
	if z.GetDebugMode == nil || z.SetDebugMode == nil {
		c.String(http.StatusNotImplemented, "not implemented\n")
		return
	}
	mode := !z.GetDebugMode()
	z.SetDebugMode(mode)
	if mode {
		c.String(http.StatusOK, "debug\n")
	} else {
		c.String(http.StatusOK, "release\n")
	}
}

func healthz(c *gin.Context) {
	if z.GetHealth == nil {
		c.String(http.StatusNotImplemented, "not implemented")
		return
	}
	if z.GetHealth() {
		c.String(http.StatusOK, "ok")
	} else {
		c.String(http.StatusServiceUnavailable, "not ok")
	}
}
