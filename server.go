// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-08-01
// @Last modification: 2021-08-13
// @Version: 1.3.2

package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
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
	e.Use(middleware.Secure())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 9,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:  "header:X-XSRF-TOKEN",
		CookieName:   "cf_sid",
		CookieMaxAge: 86400,
	}))
	// e.Use(session.Middleware(sessions.NewFilesystemStore("./", []byte(getRandomString(16)))))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(getRandomString(32)))))
	if EnableLogger {
		// e.Use(middleware.Logger())
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${remote_ip} - - [${time_rfc3339}] ${method} ${uri} ${status} ${latency_human} ${bytes_in} ${bytes_out} ${user_agent}\n",
		}))
	}

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
	e.GET("/login", Login)
	e.POST("/api/*", API)

	// 启动服务
	e.HideBanner = true
	fmt.Println("⇨ Pi Dashboard Go v" + VERSION)
	e.Logger.Fatal(e.Start(port))
}

func View(c echo.Context) error {
	username, _ := getNowUsernameAndPassword()

	sess, _ := session.Get(SessionName, c)
	//通过sess.Values读取会话数据
	userName, _ := sess.Values["id"]
	isLogin, _ := sess.Values["isLogin"]

	if userName != username || isLogin != true {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	device := Device()
	device["version"] = VERSION
	device["site_title"] = Title
	device["interval"] = Interval
	device["go_version"] = runtime.Version()

	return c.Render(http.StatusOK, "view.tmpl", device)
}

func Login(c echo.Context) error {
	username, _ := getNowUsernameAndPassword()

	sess, _ := session.Get(SessionName, c)
	//通过sess.Values读取会话数据
	userName, _ := sess.Values["id"]
	isLogin, _ := sess.Values["isLogin"]

	if userName == username && isLogin == true {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	device := make(map[string]string)
	device["version"] = VERSION
	device["site_title"] = Title
	device["go_version"] = runtime.Version()
	device["device_photo"] = Device()["device_photo"]

	return c.Render(http.StatusOK, "login.tmpl", device)
}

func API(c echo.Context) error {
	switch method := strings.Split(c.Request().URL.Path, "api/")[1]; {

	case method == "login":
		username, password := getNowUsernameAndPassword()

		sess, _ := session.Get(SessionName, c)
		//通过sess.Values读取会话数据
		userName, _ := sess.Values["id"]
		isLogin, _ := sess.Values["isLogin"]

		if userName == username && isLogin == true {
			status := map[string]bool{
				"status": true,
			}
			return c.JSON(http.StatusOK, status)
		}

		//获取登录信息
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		var loginUsername, loginPassword interface{}
		if err != nil {
			loginUsername = ""
			loginPassword = ""
		} else {
			loginUsername = json_map["username"]
			loginPassword = json_map["password"]
		}

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

			status := map[string]bool{
				"status": true,
			}
			return c.JSON(http.StatusOK, status)
		} else {

			status := map[string]bool{
				"status": false,
			}
			return c.JSON(http.StatusOK, status)
		}

	case method == "logout":
		sess, _ := session.Get(SessionName, c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: false,
		}
		sess.Values["id"] = ""
		sess.Values["isLogin"] = ""

		sess.Save(c.Request(), c.Response())

		status := map[string]bool{
			"status": true,
		}
		return c.JSON(http.StatusOK, status)

	case method == "device":
		username, _ := getNowUsernameAndPassword()

		sess, _ := session.Get(SessionName, c)
		//通过sess.Values读取会话数据
		userName, _ := sess.Values["id"]
		isLogin, _ := sess.Values["isLogin"]

		if userName != username || isLogin != true {
			status := map[string]string{
				"result": "Unauthorized",
			}
			return c.JSON(http.StatusUnauthorized, status)
		}

		device := Device()
		device["version"] = VERSION
		device["site_title"] = Title
		device["interval"] = Interval
		device["go_version"] = runtime.Version()

		return c.JSON(http.StatusOK, device)

	case method == "operation":
		username, _ := getNowUsernameAndPassword()

		sess, _ := session.Get(SessionName, c)
		//通过sess.Values读取会话数据
		userName, _ := sess.Values["id"]
		isLogin, _ := sess.Values["isLogin"]

		if userName != username || isLogin != true {
			status := map[string]string{
				"result": "Unauthorized",
			}
			return c.JSON(http.StatusUnauthorized, status)
		}

		status := map[string]bool{
			"status": true,
		}

		switch operation := c.QueryParam("action"); {
		case operation == "reboot":
			go Popen("reboot")
			return c.JSON(http.StatusOK, status)
		case operation == "shutdown":
			go Popen("shutdown -h now")
			return c.JSON(http.StatusOK, status)
		case operation == "dropcaches":
			go Popen("echo 3 > /proc/sys/vm/drop_caches")
			return c.JSON(http.StatusOK, status)
		case operation == "checknewversion":
			nowVersion, _ := getLatestVersionFromGitHub()
			result := make(map[string]string)
			if nowVersion > VERSION {
				result["new_version"] = nowVersion
				result["new_version_url"] = PROJECT + "/releases/tag/v" + nowVersion
			} else {
				result["new_version"] = ""
				result["new_version_url"] = ""
			}

			return c.JSON(http.StatusOK, result)

		}

	}

	status := map[string]string{
		"status": "UnknownMethod",
	}
	return c.JSON(http.StatusOK, status)
}

func getNowUsernameAndPassword() (username, password string) {
	username = USERNAME
	password = PASSWORD
	auth := strings.Split(Auth, ":")
	if len(auth) == 2 {
		username = auth[0]
		password = auth[1]
	}

	return username, password
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

func getLatestVersionFromGitHub() (nowVersion string, downloadURL []string) {
	url := "https://api.github.com/repos/plutobell/pi-dashboard-go/releases/latest"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	for key, value := range result {
		if key == "tag_name" {
			nowVersion = value.(string)[1:]
		}
		if key == "assets" {
			assets := value.([]interface{})
			for _, architecture := range assets {
				for key, value := range architecture.(map[string]interface{}) {
					if key == "browser_download_url" {
						downloadURL = append(downloadURL, value.(string))
					}
				}
			}
		}
	}

	return nowVersion, downloadURL
}
