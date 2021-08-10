// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-08-01
// @Last modify: 2021-08-10
// @Version: 1.2.0

package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
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
	// e.Use(session.Middleware(sessions.NewFilesystemStore("./", []byte(getRandomString(16)))))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(getRandomString(16)))))

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
	e.POST("/", View)
	e.GET("/login", Login)
	e.POST("/login", Login)

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

	sess, _ := session.Get(SessionName, c)
	userName, _ := sess.Values["id"]
	isLogin, _ := sess.Values["isLogin"]

	if userName != USERNAME || isLogin != true {
		return c.Redirect(http.StatusFound, "/login")
	}

	if ajax := c.QueryParam("logout"); ajax == "true" {
		sess, _ := session.Get(SessionName, c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: false,
		}
		sess.Values["id"] = ""
		sess.Values["isLogin"] = ""

		sess.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusFound, "/login")
	}

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

func Login(c echo.Context) error {
	username := USERNAME
	password := PASSWORD
	auth := strings.Split(Auth, ":")
	if len(auth) == 2 {
		username = auth[0]
		password = auth[1]
	} else {
		fmt.Println("Auth format error, will use default value")
	}

	sess, _ := session.Get(SessionName, c)
	//通过sess.Values读取会话数据
	userName, _ := sess.Values["id"]
	isLogin, _ := sess.Values["isLogin"]

	if userName == username && isLogin == true {
		return c.Redirect(http.StatusFound, "/")
	}

	//获取登录请求参数
	loginUsername := c.FormValue("username")
	loginPassword := c.FormValue("password")

	if loginUsername == username && loginPassword == password {
		maxAge, _ := strconv.Atoi(SessionMaxAge)

		sess, _ := session.Get(SessionName, c)
		sess.Options = &sessions.Options{
			Path:     "/",            //所有页面都可以访问会话数据
			MaxAge:   86400 * maxAge, //会话有效期，单位秒
			HttpOnly: true,
		}
		//记录会话数据, sess.Values 是map类型，可以记录多个会话数据
		sess.Values["id"] = loginUsername
		sess.Values["isLogin"] = true
		//保存用户会话数据
		sess.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusFound, "/")
	} else {
		device := make(map[string]string)
		device["version"] = VERSION
		device["site_title"] = Title
		device["go_version"] = runtime.Version()
		device["device_photo"] = Device()["device_photo"]

		return c.Render(http.StatusOK, "login.tmpl", device)
	}

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

func getRandomString(len int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		result = append(result, bytes[r.Intn(62)])
	}

	return string(result)
}
