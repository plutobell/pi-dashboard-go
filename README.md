# Pi Dashboard Go
**Pi Dashboard Go** is a Golang implementation of pi-dashboard

* [中文文档](https://ojoll.com/archives/86/)



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

**2020-8-9**

* v1.0.8 : 
  * Optimize swap display details
  * Add shortcut buttons such as shutdown and reboot

**2020-8-7**

* v1.0.7 : 
  * Optimize network card flow and curve display
  * Interface detail adjustment

**2020-8-6**

* v1.0.6 : 
  * Fix the bug that the network card data display error
  * Fixed navigation bar at the top
  * Interface detail adjustment
* v1.0.5 : 
  * Interface color adjustment
  * Data update detection and prompt
  * Optimize code for server
  * Detail adjustment
* v1.0.4 : 
  * Adjust Cached calculation method
  * Added theme-color for mobile browser
  * Added display login user statistics
  * Bug fixes and details optimization

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