// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-08-01
// @Last modify: 2021-06-17
// @Version: 1.1.2

package main

import (
	"crypto/subtle"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"runtime"
	"strings"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed assets
var f embed.FS

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
	// e.Use(middleware.Logger())
	e.Use(middleware.BasicAuth(authFunc))

	//静态文件
	assetHandler := http.FileServer(getFileSystem(false))
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", assetHandler)))

	//初始化模版引擎
	t := &Template{
		templates: template.Must(template.New("").ParseFS(f, "assets/*.tmpl")),
	}

	//向echo实例注册模版引擎
	e.Renderer = t

	// 路由
	e.GET("/", View)

	// 启动服务
	e.HideBanner = true
	fmt.Println("⇨ Pi Dashboard Go v" + VERSION)
	e.Logger.Fatal(e.Start(port))
}

func View(c echo.Context) error {
	device := Device()
	device["version"] = VERSION
	device["site_title"] = Title
	device["interval"] = Interval
	device["go_version"] = runtime.Version()

	if ajax := c.QueryParam("ajax"); ajax == "true" {
		return c.JSON(http.StatusOK, device)
	}

	status := map[string]string{
		"status": "ok",
	}
	switch operate := c.QueryParam("operate"); {
	case operate == "reboot":
		go Popen("reboot")
		return c.JSON(http.StatusOK, status)
	case operate == "shutdown":
		go Popen("shutdown -h now")
		return c.JSON(http.StatusOK, status)
	case operate == "dropcaches":
		go Popen("echo 3 > /proc/sys/vm/drop_caches")
		return c.JSON(http.StatusOK, status)
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

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		// fmt.Println("using live mode.")
		return http.FS(os.DirFS("assets"))
	}

	// fmt.Println("using embed mode.")
	fsys, err := fs.Sub(f, "assets")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
