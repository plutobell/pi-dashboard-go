// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-08-01
// @Last modification: 2021-08-28
// @Version: 1.5.1

package config

import "os/user"

const (
	//PROJECT 项目地址
	PROJECT string = "https://github.com/plutobell/pi-dashboard-go"
	//AUTHOR 作者信息
	AUTHOR string = "github:plutobell"
	//VERSION 版本信息
	VERSION string = "1.5.1"
	//USERNAME 默认用户
	USERNAME string = "pi"
	//PASSWORD 默认密码
	PASSWORD string = "123"
)

var (
	Help    bool
	Version bool
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
	// FileName 当前文件名
	FileName string
	// LinuxUserInfo 当前Linux用户信息
	LinuxUserInfo *user.User
)
