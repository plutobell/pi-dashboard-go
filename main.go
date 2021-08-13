// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-08-01
// @Last modification: 2021-08-13
// @Version: 1.3.1

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	//AUTHOR 作者信息
	AUTHOR string = "github:plutobell"
	//VERSION 版本信息
	VERSION string = "1.3.1"
	//USERNAME 默认用户
	USERNAME string = "pi"
	//PASSWORD 默认密码
	PASSWORD string = "123"
)

var (
	help    bool
	version bool
	// Port 端口
	Port string
	// Title 网站标题
	Title string
	// Net 网卡名称
	Net string
	// Disk 硬盘路径
	Disk string
	// Auth 用户名和密码
	Auth string
	// Interval 页面更新间隔
	Interval string
	// SessionMaxAge 登录状态有效期
	SessionMaxAge string
	// 启用日志显示
	EnableLogger bool
	// SessionName Session名称
	SessionName string
)

func init() {
	flag.BoolVar(&help, "help", false, "this help")
	flag.BoolVar(&version, "version", false, "show version and exit")
	flag.StringVar(&Port, "port", "8080", "specify the running port")
	flag.StringVar(&Title, "title", "Pi Dashboard Go", "specify the website title")
	flag.StringVar(&Net, "net", "lo", "specify the network device")
	flag.StringVar(&Disk, "disk", "/", "specify the disk")
	flag.StringVar(&Auth, "auth", USERNAME+":"+PASSWORD, "specify username and password")
	flag.StringVar(&Interval, "interval", "1", "specify the update interval in seconds")
	flag.StringVar(&SessionMaxAge, "session", "7", "specify the login status validity in days")
	flag.BoolVar(&EnableLogger, "log", false, "enable log display")

	SessionName = "logged_in"

	flag.Usage = usage
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
		return
	}
	if version {
		fmt.Println("Pi Dashboard Go " + VERSION)
		fmt.Println("Project address: https://github.com/plutobell/pi-dashboard-go")
		return
	}
	netDevs, err := Popen("cat /proc/net/dev")
	if err != nil {
		log.Fatal(err)
		return
	}
	if !strings.Contains(netDevs, Net+":") {
		fmt.Println("Network card does not exist")
		return
	}
	diskLists, err := Popen("blkid")
	if err != nil {
		log.Fatal(err)
		return
	}
	if Disk != "/" {
		if !strings.Contains(diskLists, Disk+":") {
			fmt.Println("Disk does not exist")
			return
		}
	}
	authSlice := strings.Split(Auth, ":")
	if len(authSlice) != 2 {
		fmt.Println("Auth format error")
		return
	}
	if len([]rune(authSlice[0])) > 15 || len([]rune(authSlice[0])) == 0 {
		fmt.Println("Username is too long")
		return
	}
	if len([]rune(authSlice[1])) > 15 || len([]rune(authSlice[1])) == 0 {
		fmt.Println("Password is too long")
		return
	}
	if len([]rune(Title)) > 25 {
		fmt.Println("Title is too long")
		return
	}

	isDigit := true
	for _, r := range Interval {
		if !unicode.IsDigit(rune(r)) {
			isDigit = false
			break
		}
	}
	if !isDigit {
		fmt.Println("Interval parameter value is invalid")
		return
	}

	IntervalInt, err := strconv.Atoi(Interval)
	if err != nil {
		log.Fatal(err)
		return
	}
	if IntervalInt > 900 {
		fmt.Println("Interval is too long")
		return
	} else if IntervalInt < 0 {
		fmt.Println("Interval should be no less than 0")
		return
	}

	SessionMaxAgeInt, err := strconv.Atoi(SessionMaxAge)
	if err != nil {
		log.Fatal(err)
		return
	}
	if SessionMaxAgeInt > 365 {
		fmt.Println("Session days is too long")
		return
	} else if SessionMaxAgeInt < 0 {
		fmt.Println("Session days should be no less than 0")
		return
	}

	Server()
}

func usage() {
	fmt.Fprintf(os.Stderr, `Pi Dashboard Go version: %s
Project address: https://github.com/plutobell/pi-dashboard-go

Usage: Pi Dashboard Go [-auth USR:PSW] [-disk Paths] [-help]
[-interval Seconds] [-log] [-net NIC] [-port Port]
[-session Days] [-title Title] [-version]

Options:
`, VERSION)
	flag.PrintDefaults()
}
