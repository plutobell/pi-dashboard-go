# Pi Dashboard Go
**Pi Dashboard Go** is a Golang implementation of pi-dashboard

* **[中文文档](https://ojoll.com/archives/86/)**



![](./screenshots/screenshot_index.png)

![](./screenshots/screenshot_login.png)

![](./screenshots/screenshot_index_dark.png)

![](./screenshots/screenshot_login_dark.png)





## Install

Thanks to the characteristics of the Golang language, the deployment of **Pi Dashboard Go** is very simple: **single binary executable file**.

#### Download

Just download the executable file from the project **[Releases](https://github.com/plutobell/pi-dashboard-go/releases)**, **no other dependencies**.

#### Authority

Grant executable permissions

```
chmod +x pi-dashboard-go
```

**Note：Pi Dashboard Go requires root privileges.**



## Use

#### Usage

**Pi Dashboard Go** can be configured via command line parameters：

```bash
Pi Dashboard Go version: v1.5.1
Project address: https://github.com/plutobell/pi-dashboard-go

Usage: Pi Dashboard Go [-auth USR:PSW] [-disk Paths] [-help]
[-interval Seconds] [-log] [-net NIC] [-port Port]
[-session Days] [-title Title] [-version]

Options:
  -auth string
        specify username and password (default "pi:123")
  -disk string
        specify the disk (default "/")
  -help
        this help
  -interval string
        specify the update interval in seconds (default "1")
  -log
        enable log display
  -net string
        specify the network device (default "lo")
  -port string
        specify the running port (default "8080")
  -session string
        specify the login status validity in days (default "7")
  -title string
        specify the website title (default "Pi Dashboard Go")
  -version
        show version and exit
```



## Thanks

* **[Pi Dashboard](https://github.com/spoonysonny/pi-dashboard)**
* **[Golang](https://golang.org/)**
* **[echo](https://github.com/labstack/echo)**



## Changelog

* **[Changelog](./CHANGELOG.md)**