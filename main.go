// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-8-1
// @Last modify: 2020-8-4
// @Version: 1.0.1

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	//VERSION 版本信息
	VERSION string = "1.0.1"
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
	flag.StringVar(&Auth, "auth", "pi:123", "specify username and password")

	flag.Usage = usage
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
		return
	}
	if version {
		fmt.Println("pi-dashboard-go " + VERSION)
		return
	}
	netDevs := Popen("cat /proc/net/dev")
	if !strings.Contains(netDevs, Net+":") {
		fmt.Println("网卡不存在")
		return
	}
	diskLists := Popen("blkid")
	if Disk != "/" {
		if !strings.Contains(diskLists, Disk+":") {
			fmt.Println("磁盘不存在")
			return
		}
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
