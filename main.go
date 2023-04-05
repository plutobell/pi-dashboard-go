// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-08-01
// @Last modification: 2023-04-05
// @Version: 1.7.0

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/plutobell/pi-dashboard-go/config"
	"github.com/plutobell/pi-dashboard-go/device"
	"github.com/plutobell/pi-dashboard-go/server"
)

func init() {
	flag.BoolVar(&config.Help, "help", false, "this help")
	flag.BoolVar(&config.Version, "version", false, "show version and exit")
	flag.StringVar(&config.Port, "port", "8080", "specify the running port")
	flag.StringVar(&config.Title, "title", "Pi Dashboard Go", "specify the website title")
	flag.StringVar(&config.Net, "net", "lo", "specify the network device")
	flag.StringVar(&config.Disk, "disk", "/", "specify the filesystem path")
	flag.StringVar(&config.Auth, "auth", config.USERNAME+":"+config.PASSWORD, "specify username and password")
	flag.StringVar(&config.Interval, "interval", "1", "specify the update interval in seconds")
	flag.StringVar(&config.SessionMaxAge, "session", "7", "specify the login status validity in days")
	flag.StringVar(&config.Theme, "theme", "light", "specify the theme between 'light' and 'dark'")
	flag.BoolVar(&config.EnableLogger, "log", false, "enable log display")

	config.SessionName = "logged_in"
	config.FileName = filepath.Base(os.Args[0])
	config.LinuxUserInfo, _ = user.Current()

	flag.Usage = usage
}

func main() {
	flag.Parse()

	if config.Help {
		flag.Usage()
		return
	}
	if config.Version {
		fmt.Println("Pi Dashboard Go v" + config.VERSION)
		fmt.Println("Project address: " + config.PROJECT)
		return
	}
	netDevs, err := device.Popen("cat /proc/net/dev")
	if err != nil {
		log.Fatal(err)
		return
	}
	if !strings.Contains(netDevs, config.Net+":") {
		fmt.Println("Network card does not exist")
		return
	}
	pathExists := false
	_, err = os.Stat(config.Disk)
	if err == nil {
		pathExists = true
	}
	if os.IsNotExist(err) {
		pathExists = false
	}

	if config.Disk != "/" {
		if !pathExists {
			fmt.Println("Disk does not exist")
			return
		}
	}
	authSlice := strings.Split(config.Auth, ":")
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
	if len([]rune(config.Title)) > 25 {
		fmt.Println("Title is too long")
		return
	}

	isDigit := true
	for _, r := range config.Interval {
		if !unicode.IsDigit(rune(r)) {
			isDigit = false
			break
		}
	}
	if !isDigit {
		fmt.Println("Interval parameter value is invalid")
		return
	}

	IntervalInt, err := strconv.Atoi(config.Interval)
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

	SessionMaxAgeInt, err := strconv.Atoi(config.SessionMaxAge)
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

	if config.Theme != "light" && config.Theme != "dark" {
		fmt.Println("Theme name not supported")
		return
	}

	server.Run()
}

func usage() {
	fmt.Fprintf(os.Stderr, `Pi Dashboard Go version: v%s
Project address: %s

Usage: %s [-auth USR:PSW] [-disk Paths] [-help]
[-interval Seconds] [-log] [-net NIC] [-port Port]
[-session Days] [-theme Theme] [-title Title] [-version]

Options:
`, config.VERSION, config.PROJECT, config.FileName)
	flag.PrintDefaults()
}
