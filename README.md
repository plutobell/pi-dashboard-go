# Pi Dashboard Go
**Pi Dashboard Go** is a Golang implementation of pi-dashboard

* **[中文文档](https://ojoll.com/archives/86/)**



![](./screenshot.png)

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

```
Pi Dashboard Go version: 1.1.0
Project address: https://github.com/plutobell/pi-dashboard-go

Usage: Pi Dashboard Go [-help] [-version] [-port port] [-title title] [-net net] [-disk disk] [-auth usr:psw] [-interval interval]

Options:
  -auth string
        specify username and password (default "pi:123")
  -disk string
        specify the disk (default "/")
  -help
        this help
  -interval string
        specify the update interval in seconds (default "1")
  -net string
        specify the network device (default "lo")
  -port string
        specify the running port (default "8080")
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