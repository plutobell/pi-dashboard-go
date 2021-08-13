// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-08-01
// @Last modification: 2021-08-13
// @Version: 1.3.2

package main

import (
	"errors"
	"io/ioutil"
	"math"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//Popen 函数用于执行系统命令
func Popen(command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		//fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return "", errors.New("Error:can not obtain stdout pipe for command")
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		//fmt.Println("Error:The command is err,", err)
		return "", errors.New("Error:The command is err")
	}

	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		//fmt.Println("ReadAll Stdout:", err.Error())
		return "", errors.New("ReadAll Stdout:" + err.Error())
	}

	if err := cmd.Wait(); err != nil {
		//fmt.Println("wait:", err.Error())
		return "", errors.New("wait:" + err.Error())
	}

	return string(bytes), nil
}

//Device 函数获取设备信息
func Device() map[string]string {
	//piCpuModelInfo Raspberry Pi CPU型号信息
	piCpuModelInfo := map[string]string{
		"Raspberry Pi 4 Model B":  "BCM2711",
		"Raspberry Pi 3 Model B+": "BCM2837B0",
		"Raspberry Pi 3 Model B":  "BCM2837/A0/B0",
		"Raspberry Pi 2 Model B":  "BCM2836/7",
		"Raspberry Pi Model B+":   "BCM2835",
		"Raspberry Pi Model B":    "BCM2835",
		"Raspberry Pi 3 Model A+": "BCM2837B0",
		"Raspberry Pi Model A+":   "BCM2835",
		"Raspberry Pi Zero WH":    "BCM2835",
		"Raspberry Pi Zero W":     "BCM2835",
		"Raspberry Pi Zero":       "BCM2835",
	}

	//Command 命令列表
	device := make(map[string]string)
	var command map[string]string = map[string]string{
		"ip":               "ip a | grep -w inet | grep -v inet6 | grep -v 127 | awk '{ print $2 }'",
		"uptime":           "cat /proc/uptime | awk '{ print $1}'",
		"login_user_count": "who -q | awk 'NR==2{print $2}'",
		"load_average":     "cat /proc/loadavg | awk '{print $1,$2,$3,$4}'",
		"current_user":     "whoami",
		"hostname":         "cat /etc/hostname",
		"os":               "uname -o",
		"system":           "cat /etc/os-release | grep PRETTY_NAME=",
		"arch":             "arch",
		"uname":            "uname -a",
		"model":            "cat /proc/cpuinfo | grep -i Model |sort -u |head -1",
		"cpu_revision":     "cat /proc/cpuinfo | grep Revision | awk '{ print $3}'",
		"cpu_model_name":   "lscpu | grep 'Model name' | awk '{ print $3}'",
		"cpu_cores":        "lscpu | grep 'CPU(s):' | awk '{ print $2}'",
		"cpu_status":       "top -bn1 | grep -w '%Cpu(s):' | awk '{ print $2,$4,$6,$8,$10,$12,$14,$16}'",
		"cpu_used":         "cat <(grep 'cpu ' /proc/stat) <(sleep 1 && grep 'cpu ' /proc/stat) | awk -v RS='' '{print ($13-$2+$15-$4)*100/($13-$2+$15-$4+$16-$5)}'",
		"cpu_temperature":  "cat /sys/class/thermal/thermal_zone0/temp",
		"cpu_freq":         "cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq",
		"memory_total":     "cat /proc/meminfo | grep -w MemTotal: | awk '{ print $2}'",
		"memory_free":      "cat /proc/meminfo | grep -w MemFree: | awk '{ print $2}'",
		"memory_available": "cat /proc/meminfo | grep -w MemAvailable: | awk '{ print $2}'",
		"memory_buffers":   "cat /proc/meminfo | grep -w Buffers: | awk '{ print $2}'",
		"memory_cached":    "cat /proc/meminfo | grep -w Cached: | awk '{ print $2}'",
		"swap_total":       "cat /proc/meminfo | grep SwapTotal: | awk '{ print $2}'",
		"swap_free":        "cat /proc/meminfo | grep SwapFree: | awk '{ print $2}'",
		"disk":             "df " + Disk + " | awk 'NR==2{print $3,$4}'",
		"net_status_lo":    "cat /proc/net/dev | grep lo: | awk '{ print $2,$3,$10,$11}'",
		"net_status":       "cat /proc/net/dev | grep " + Net + ": | awk '{ print $2,$3,$10,$11}'",
	}

	arch := runtime.GOARCH
	if strings.Contains(arch, "arm") {
		arch = "arm"
	}

	if arch != "arm" {
		command["cpu_cores"] = "cat /proc/cpuinfo |grep 'cores' | uniq | awk '{ print $4 }'"
		command["cpu_model_name"] = "lscpu | grep 'Model name'"
		command["cpu_freq"] = "cat /proc/cpuinfo |grep MHz|uniq | awk '{ print $4 }'"
	}

	for k, v := range command {
		res, err := Popen(v)
		if err != nil {
			// log.Fatal(err)
			device[k] = "NaN"
		}
		res = strings.Replace(res, "\n", "", -1)
		device[k] = res
	}

	cpuTemperature, _ := strconv.Atoi(device["cpu_temperature"])
	device["uptime"] = resolveTime(device["uptime"])
	device["hostname"] = strings.Replace(strings.Replace(device["hostname"], " ", "", -1), "\n", "", -1)
	if strings.Contains(device["system"], "\"") {
		device["system"] = strings.Replace(strings.Replace(strings.Split(device["system"], "\"")[1], " GNU/Linux ", " ", -1), "\"", "", -1)
	} else {
		device["system"] = "NaN"
	}

	if arch == "arm" {
		device["cpu_temperature"] = strconv.FormatFloat(float64(cpuTemperature)/1000, 'f', 1, 64)
		if strings.Contains(device["model"], ":") {
			device["model"] = strings.Split(device["model"], ":")[1]
		} else {
			device["model"] = "NaN"
		}

		for key, value := range piCpuModelInfo {
			if strings.Contains(device["model"], key) {
				device["cpu_revision"] = device["cpu_model_name"]
				device["cpu_model_name"] = value
				break
			}
		}

	} else {
		device["cpu_temperature"] = "NaN"
		device["model"] = "Linux Computer"
		if strings.Contains(device["cpu_model_name"], ":") {
			device["cpu_model_name"] = strings.Trim(strings.Split(device["cpu_model_name"], ":")[1], " ")
		} else {
			device["cpu_model_name"] = "NaN"
		}
	}
	if strings.Contains(strings.ToLower(device["model"]), "raspberry") {
		device["device_photo"] = "raspberry.png"
	} else {
		device["device_photo"] = "linux.png"
	}

	device["login_user_count"] = strings.Split(device["login_user_count"], "=")[1]
	device["ip"] = strings.Split(device["ip"], "/")[0]
	device["now_time"] = time.Now().Format("2006-01-02 15:04:05")

	loadAverage := strings.Split(device["load_average"], " ")
	device["load_average_1m"] = loadAverage[0]
	device["load_average_5m"] = loadAverage[1]
	device["load_average_15m"] = loadAverage[2]
	device["load_average_process_running"] = strings.Split(loadAverage[3], "/")[0]
	device["load_average_process_total"] = strings.Split(loadAverage[3], "/")[1]
	delete(device, "load_average")

	cpuStatusRaw := device["cpu_status"]
	exceptionSituation := []string{"id,", "wa,", "hi,", "si,"}
	for _, exception := range exceptionSituation {
		if exception == "id" {
			cpuStatusRaw = strings.Replace(cpuStatusRaw, exception, "100.0", -1)
		} else {
			cpuStatusRaw = strings.Replace(cpuStatusRaw, exception, "0.0", -1)
		}
	}
	cpuStatus := strings.Split(cpuStatusRaw, " ")
	device["cpu_status_user"] = cpuStatus[0]
	device["cpu_status_nice"] = cpuStatus[1]
	device["cpu_status_system"] = cpuStatus[2]
	device["cpu_status_idle"] = cpuStatus[3]
	device["cpu_status_iowait"] = cpuStatus[4]
	device["cpu_status_irq"] = cpuStatus[5]
	device["cpu_status_softirq"] = cpuStatus[6]
	delete(device, "cpu_status")
	cpuUsed, _ := strconv.ParseFloat(device["cpu_used"], 64)
	device["cpu_used"] = strconv.FormatFloat(cpuUsed, 'f', 1, 64)
	if arch == "arm" {
		cpuFreq, _ := strconv.ParseInt(device["cpu_freq"], 10, 64)
		device["cpu_freq"] = strconv.FormatInt(cpuFreq/1000, 10)
	} else {
		device["cpu_freq"] = strings.Split(device["cpu_freq"], ".")[0]
	}

	nowTime := strings.Split(device["now_time"], " ")
	device["now_time_ymd"] = nowTime[0]
	device["now_time_hms"] = nowTime[1]
	delete(device, "now_time")

	diskSlice := strings.Split(device["disk"], " ")
	device["disk_used"] = diskSlice[0]
	device["disk_free"] = diskSlice[1]
	delete(device, "disk")
	device["disk_name"] = strings.ToUpper(Disk)
	diskUsed, _ := strconv.ParseFloat(device["disk_used"], 64)
	diskFree, _ := strconv.ParseFloat(device["disk_free"], 64)
	device["disk_total"] = strconv.FormatFloat((diskFree+diskUsed)/1024/1024, 'f', 1, 64)
	device["disk_used_percent"] = strconv.FormatFloat(100*diskUsed/(diskFree+diskUsed), 'f', 1, 64)
	device["disk_used"] = strconv.FormatFloat(diskUsed/1024/1024, 'f', 1, 64)
	device["disk_free"] = strconv.FormatFloat(diskFree/1024/1024, 'f', 1, 64)

	device["net_dev_name"] = Net
	netStatus := strings.Split(device["net_status"], " ")
	device["net_status_in_data"] = netStatus[0]
	device["net_status_in_package"] = netStatus[1]
	device["net_status_out_data"] = netStatus[2]
	device["net_status_out_package"] = netStatus[3]
	delete(device, "net_status")
	netInData, _ := strconv.ParseFloat(device["net_status_in_data"], 64)
	netOutData, _ := strconv.ParseFloat(device["net_status_out_data"], 64)
	device["net_status_in_data_format"] = bytesRound(netInData, 2)
	device["net_status_out_data_format"] = bytesRound(netOutData, 2)
	// device["net_status_in_data"] = strconv.FormatFloat(netInData/1024/1024, 'f', 1, 64)
	// device["net_status_out_data"] = strconv.FormatFloat(netOutData/1024/1024, 'f', 1, 64)
	inPackage, _ := strconv.ParseFloat(device["net_status_in_package"], 64)
	outPackage, _ := strconv.ParseFloat(device["net_status_out_package"], 64)
	netLoadAverage := (inPackage + outPackage) / 2
	device["net_status_load_average"] = strconv.FormatFloat(netLoadAverage, 'f', 1, 64)

	netStatusLo := strings.Split(device["net_status_lo"], " ")
	device["net_status_lo_in_data"] = netStatusLo[0]
	device["net_status_lo_in_package"] = netStatusLo[1]
	device["net_status_lo_out_data"] = netStatusLo[2]
	device["net_status_lo_out_package"] = netStatusLo[3]
	delete(device, "net_status_lo")
	netLoInData, _ := strconv.ParseFloat(device["net_status_lo_in_data"], 64)
	netLoOutData, _ := strconv.ParseFloat(device["net_status_lo_out_data"], 64)
	device["net_status_lo_in_data_format"] = bytesRound(netLoInData, 2)
	device["net_status_lo_out_data_format"] = bytesRound(netLoOutData, 2)
	// device["net_status_lo_in_data"] = strconv.FormatFloat(netLoInData/1024/1024, 'f', 1, 64)
	// device["net_status_lo_out_data"] = strconv.FormatFloat(netLoOutData/1024/1024, 'f', 1, 64)
	inLoPackage, _ := strconv.ParseFloat(device["net_status_lo_in_package"], 64)
	outLoPackage, _ := strconv.ParseFloat(device["net_status_lo_out_package"], 64)
	netLoLoadAverage := (inLoPackage + outLoPackage) / 2
	device["net_status_lo_load_average"] = strconv.FormatFloat(netLoLoadAverage, 'f', 1, 64)

	memoryCached, _ := strconv.ParseFloat(device["memory_cached"], 64)
	memoryCached = memoryCached / 1024
	device["memory_cached"] = strconv.FormatFloat(memoryCached, 'f', 1, 64)

	memoryFree, _ := strconv.ParseFloat(device["memory_free"], 64)
	memoryFree = memoryFree / 1024
	device["memory_free"] = strconv.FormatFloat(memoryFree, 'f', 1, 64)

	memoryAvailable, _ := strconv.ParseFloat(device["memory_available"], 64)
	memoryAvailable = memoryAvailable / 1024
	device["memory_available"] = strconv.FormatFloat(memoryAvailable, 'f', 1, 64)

	memoryBuffers, _ := strconv.ParseFloat(device["memory_buffers"], 64)
	memoryBuffers = memoryBuffers / 1024
	device["memory_buffers"] = strconv.FormatFloat(memoryBuffers, 'f', 1, 64)

	memoryTotal, _ := strconv.ParseFloat(device["memory_total"], 64)
	memoryTotal = memoryTotal / 1024
	device["memory_total"] = strconv.FormatFloat(memoryTotal, 'f', 1, 64)

	device["memory_used"] = strconv.FormatFloat(memoryTotal-memoryFree, 'f', 1, 64)
	device["memory_real_used"] = strconv.FormatFloat(memoryTotal-memoryAvailable, 'f', 1, 64)
	device["memory_percent"] = strconv.FormatFloat(100*(memoryTotal-memoryFree)/memoryTotal, 'f', 1, 64)
	device["memory_real_percent"] = strconv.FormatFloat(100*(memoryTotal-memoryAvailable)/memoryTotal, 'f', 1, 64)
	device["memory_cached_percent"] = strconv.FormatFloat(100*(memoryCached)/memoryTotal, 'f', 1, 64)
	swapFree, _ := strconv.ParseFloat(device["swap_free"], 64)
	swapTotal, _ := strconv.ParseFloat(device["swap_total"], 64)
	device["swap_free"] = strconv.FormatFloat(swapFree/1024, 'f', 1, 64)
	device["swap_total"] = strconv.FormatFloat(swapTotal/1024, 'f', 1, 64)
	device["swap_used"] = strconv.FormatFloat((swapTotal-swapFree)/1024, 'f', 1, 64)
	device["swap_used_percent"] = strconv.FormatFloat(100*(swapTotal-swapFree)/swapTotal, 'f', 1, 64)
	if swapFree == 0 && swapTotal == 0 {
		device["swap_used_percent"] = "0.0"
		device["swap_free"] = "0.0"
		device["swap_total"] = "0.0"
	}

	return device
}

func resolveTime(str string) string {
	var uptime string
	strF, _ := strconv.ParseFloat(str, 10)
	second := int64(strF)
	var min = second / 60
	var hours = min / 60
	var days = int64(math.Floor(float64(hours / 24)))
	hours = int64(math.Floor(float64(hours - (days * 24))))
	min = int64(math.Floor(float64(min - (days * 60 * 24) - (hours * 60))))

	if days != 0 {
		if days == 1 {
			uptime = strconv.FormatInt(days, 10) + " day "
		} else {
			uptime = strconv.FormatInt(days, 10) + " days "
		}
	}
	if hours != 0 {
		uptime = uptime + strconv.FormatInt(hours, 10) + ":"
	}

	return uptime + strconv.FormatInt(min, 10)
}

func bytesRound(number, round float64) string {
	var last string
	if number < 0 {
		last = "0" + "B"
	} else if number < 1024 {
		numberStr := strconv.FormatFloat(number, 'f', 1, 64)
		last = numberStr + "B"
	} else if number < 1048576 {
		number = number / 1024
		last = strconv.FormatFloat(math.Round(number*math.Pow(10, round))/math.Pow(10, round), 'f', 1, 64) + "KB"
	} else if number < 1048576000 {
		number = number / 1048576
		last = strconv.FormatFloat(math.Round(number*math.Pow(10, round))/math.Pow(10, round), 'f', 1, 64) + "MB"
	} else {
		number = number / 1048576000
		last = strconv.FormatFloat(math.Round(number*math.Pow(10, round))/math.Pow(10, round), 'f', 1, 64) + "GB"
	}
	return last
}
