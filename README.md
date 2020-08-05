# Pi Dashboard Go
**Pi Dashboard Go** is a Golang implementation of pi-dashboard

* [中文文档](https://ojoll.com/archives/86/)



## Install

Thanks to the characteristics of the Golang language, the deployment of **Pi Dashboard Go** is very simple: **single binary executable file**.

#### Download

Just download the executable file from the project **[Releases](https://github.com/plutobell/pi-dashboard-go/releases)**, **no other dependencies**.

#### Authority

Grant executable permissions

```
chmod +x pi-dashboard-go
```



## Use

#### Usage

**Pi Dashboard Go** can be configured via command line parameters：

```
Pi Dashboard Go  version: 1.0.0
Project address: https://github.com/plutobell/pi-dashboard-go

Usage: Pi Dashboard Go [-help] [-version] [-port port] [-title title] [-net net] [-disk disk] [-auth usr:psw]

Options:
  -auth string
        specify username and password (default "pi:123")
  -disk string
        specify the disk (default "/")
  -help
        this help
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

* [Pi Dashboard](https://github.com/spoonysonny/pi-dashboard)

* [echo](https://github.com/labstack/echo)

* [go.rice](https://github.com/GeertJohan/go.rice)
* [Golang](https://golang.org/)



## Changelog

**2020-8-5**

* v1.0.3 : 
  * Newly added time formatting function resolveTime
  * Detail optimization
* v1.0.2 : 
  * Improve command line parameter verification
  * Detail optimization
  * Add test case device_test.go
  * New page loading animation

**2020-8-4**

* v1.0.1 : Bug fixes, detail optimization
* v1.0.0