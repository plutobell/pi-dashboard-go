// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-8-1
// @Last modify: 2020-8-14
// @Version: 1.0.9

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	//AUTHOR 作者信息
	AUTHOR string = "github:plutobell"
	//VERSION 版本信息
	VERSION string = "1.0.9"
	//USERNAME 默认用户
	USERNAME string = "pi"
	//PASSWORD 默认密码
	PASSWORD string = "123"
)

var (
	help    bool
	version bool
	//Port 端口
	Port string
	//Title 网站标题
	Title string
	//Net 网卡名称
	Net string
	//Disk 硬盘路径
	Disk string
	//Auth 用户名和密码
	Auth string
)

func init() {
	flag.BoolVar(&help, "help", false, "this help")
	flag.BoolVar(&version, "version", false, "show version and exit")
	flag.StringVar(&Port, "port", "8080", "specify the running port")
	flag.StringVar(&Title, "title", "Pi Dashboard Go", "specify the website title")
	flag.StringVar(&Net, "net", "lo", "specify the network device")
	flag.StringVar(&Disk, "disk", "/", "specify the disk")
	flag.StringVar(&Auth, "auth", USERNAME+":"+PASSWORD, "specify username and password")

	flag.Usage = usage
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
		return
	}
	if version {
		fmt.Println("Pi Dashboard Go  " + VERSION)
		fmt.Println("Project address: https://github.com/plutobell/pi-dashboard-go")
		return
	}
	netDevs, err := Popen("cat /proc/net/dev")
	if err != nil {
		log.Fatal(err)
	}
	if !strings.Contains(netDevs, Net+":") {
		fmt.Println("Network card does not exist")
		return
	}
	diskLists, err := Popen("blkid")
	if err != nil {
		log.Fatal(err)
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

	Server()
}

func usage() {
	fmt.Fprintf(os.Stderr, `Pi Dashboard Go  version: %s
Project address: https://github.com/plutobell/pi-dashboard-go

Usage: Pi Dashboard Go [-help] [-version] [-port port] [-title title] [-net net] [-disk disk] [-auth usr:psw]

Options:
`, VERSION)
	flag.PrintDefaults()
}
