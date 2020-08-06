// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-8-1
// @Last modify: 2020-8-6
// @Version: 1.0.5

package main

import (
	"crypto/subtle"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//Template 模板
type Template struct {
	templates *template.Template
}

//Render 渲染器
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

//Server 实例
func Server() {
	//Echo 实例
	e := echo.New()
	port := ":" + Port

	//注册中间件
	e.Use(middleware.Recover())
	//e.Use(middleware.Logger())
	e.Use(middleware.BasicAuth(authFunc))

	//静态文件
	// e.Static("/assets", "assets")
	// e.File("/favicon.ico", "assets/favicon.ico")
	// the file server for rice. "app" is the folder where the files come from.
	assetHandler := http.FileServer(rice.MustFindBox("assets").HTTPBox())
	e.GET("/", echo.WrapHandler(assetHandler))
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", assetHandler)))

	//初始化模版引擎
	templateBox, err := rice.FindBox("assets")
	if err != nil {
		log.Fatal(err)
	}
	templateString, err := templateBox.String("view.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	t := &Template{
		templates: template.Must(template.New("view.tmpl").Parse(templateString)),
	}

	//向echo实例注册模版引擎
	e.Renderer = t

	// 路由
	e.GET("/", View)

	// 启动服务
	e.HideBanner = true
	fmt.Println("⇨ Pi Dashboard Go  v" + VERSION)
	e.Logger.Fatal(e.Start(port))
}

// View 函数
func View(c echo.Context) error {
	device := Device()
	device["version"] = VERSION
	device["site_title"] = Title

	ajax := c.QueryParam("ajax")
	if ajax == "true" {
		return c.JSON(http.StatusOK, device)
	}
	return c.Render(http.StatusOK, "view.tmpl", device)
}

func authFunc(username, password string, c echo.Context) (bool, error) {
	// Be careful to use constant time comparison to prevent timing attacks
	userName := USERNAME
	passWord := PASSWORD
	auth := strings.Split(Auth, ":")
	if len(auth) == 2 {
		userName = auth[0]
		passWord = auth[1]
	} else {
		fmt.Println("Auth格式错误")
		return false, nil
	}

	if subtle.ConstantTimeCompare([]byte(userName), []byte(username)) == 1 &&
		subtle.ConstantTimeCompare([]byte(passWord), []byte(password)) == 1 {
		return true, nil
	}
	return false, nil
}
