// @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
// @Description: Golang implementation of pi-dashboard
// @Author: github.com/plutobell
// @Creation: 2020-08-01
// @Last modification: 2021-09-02
// @Version: 1.6.0

package device

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/plutobell/pi-dashboard-go/config"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

//piCpuModelInfo Raspberry Pi CPU型号信息
var piCpuModelInfo = map[string]string{
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
var command map[string]string = map[string]string{
	"ip":               "ip a | grep -w inet | grep -v inet6 | grep -v 127 | awk '{ print $2 }'",
	"login_user_count": "who -q | awk 'NR==2{print $2}'",
	"system":           "cat /etc/os-release | grep PRETTY_NAME=",
	"uname":            "uname -a",
	"model":            "cat /proc/cpuinfo | grep -i Model |sort -u |head -1",
	"cpu_status":       "top -bn1 | grep -w '%Cpu(s):' | awk '{ print $2,$4,$6,$8,$10,$12,$14,$16}'",
	"cpu_model_name":   "lscpu | grep 'Model name' | awk '{ print $3}'",
	"cpu_freq":         "cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq",
}

var (
	arch  string
	os    string
	model string
)

type Host struct {
	Model          string `json:"model"`
	HostName       string `json:"hostname"`
	Uptime         string `json:"uptime_raw"`
	UptimeFormat   string `json:"uptime"`
	IP             string `json:"ip"`
	System         string `json:"system"`
	Uname          string `json:"uname"`
	CurrentUser    string `json:"current_user"`
	OS             string `json:"os"`
	LoginUserCount string `json:"login_user_count"`
	NowTimeHMS     string `json:"now_time_hms"`
	NowTimeYMD     string `json:"now_time_ymd"`
}

type CPU struct {
	Arch        string `json:"arch"`
	Cores       string `json:"cpu_cores"`
	Mhz         string `json:"cpu_freq"`
	ModelName   string `json:"cpu_model_name"`
	Revision    string `json:"cpu_revision"`
	Idle        string `json:"cpu_status_idle"`
	Iowait      string `json:"cpu_status_iowait"`
	Irq         string `json:"cpu_status_irq"`
	Nice        string `json:"cpu_status_nice"`
	Softirg     string `json:"cpu_status_softirq"`
	System      string `json:"cpu_status_system"`
	User        string `json:"cpu_status_user"`
	Temperature string `json:"cpu_temperature"`
	UsedPercent string `json:"cpu_used"`
}

type Disk struct {
	Path        string `json:"disk_name"`
	Free        string `json:"disk_free"`
	Total       string `json:"disk_total"`
	Used        string `json:"disk_used"`
	UsedPercent string `json:"disk_used_percent"`
}

type Memory struct {
	Total         string `json:"memory_total"`
	Used          string `json:"memory_used"`
	Free          string `json:"memory_free"`
	Percent       string `json:"memory_percent"`
	Available     string `json:"memory_available"`
	Buffers       string `json:"memory_buffers"`
	Cached        string `json:"memory_cached"`
	CachedPercent string `json:"memory_cached_percent"`
	RealPercent   string `json:"memory_real_percent"`
	RealUsed      string `json:"memory_real_used"`

	SwapTotal       string `json:"swap_total"`
	SwapFree        string `json:"swap_free"`
	SwapUsed        string `json:"swap_used"`
	SwapUsedPercent string `json:"swap_used_percent"`
}

type Net struct {
	Name           string `json:"net_dev_name"`
	DataIn         string `json:"net_status_in_data"`
	DataOut        string `json:"net_status_out_data"`
	DataInFormat   string `json:"net_status_in_data_format"`
	DatatOutFormat string `json:"net_status_out_data_format"`
	PackageIn      string `json:"net_status_in_package"`
	PackageOut     string `json:"net_status_out_package"`
	LoadAvg        string `json:"net_status_load_average"`
}

type NetLo struct {
	Name           string `json:"net_lo_dev_name"`
	DataIn         string `json:"net_status_lo_in_data"`
	DataOut        string `json:"net_status_lo_out_data"`
	DataInFormat   string `json:"net_status_lo_in_data_format"`
	DatatOutFormat string `json:"net_status_lo_out_data_format"`
	PackageIn      string `json:"net_status_lo_in_package"`
	PackageOut     string `json:"net_status_lo_out_package"`
	LoadAvg        string `json:"net_status_lo_load_average"`
}

type Load struct {
	Load1  string `json:"load_average_1m"`
	Load5  string `json:"load_average_5m"`
	Load15 string `json:"load_average_15m"`
}

type Process struct {
	Running string `json:"load_average_process_running"`
	Total   string `json:"load_average_process_total"`
}

func init() {
	arch = runtime.GOARCH
	os = runtime.GOOS

}

func (hh *Host) Get() {
	h, err := host.Info()
	if err != nil {
		panic(err)
	}

	nowTime := strings.Split(time.Now().Format("2006-01-02 15:04:05"), " ")
	hh.NowTimeYMD = nowTime[0]
	hh.NowTimeHMS = nowTime[1]

	hh.HostName = h.Hostname
	hh.Uptime = strconv.Itoa(int(h.Uptime))
	hh.UptimeFormat = resolveTime(hh.Uptime)
	hh.OS = h.OS

	if hh.OS == "linux" {
		res, err := Popen(command["model"])
		if err != nil {
			panic(err)
		}
		res = strings.Replace(res, "\n", "", -1)
		if strings.Contains(arch, "arm") {
			if strings.Contains(res, ":") {
				hh.Model = strings.TrimSpace(strings.Split(res, ":")[1])
			} else {
				hh.Model = "NaN"
			}
		} else {
			hh.Model = "Linux Computer"
		}
		model = hh.Model
	}

	if hh.OS == "windows" {
		hh.System = h.Platform + " " + h.PlatformVersion
	} else {
		res, err := Popen(command["system"])
		if err != nil {
			panic(err)
		}
		res = strings.Replace(res, "\n", "", -1)
		if strings.Contains(res, "\"") {
			hh.System = strings.Replace(strings.Replace(strings.Split(res, "\"")[1], " GNU/Linux ", " ", -1), "\"", "", -1)
		} else {
			hh.System = "NaN"
		}
	}

	if hh.OS == "linux" {
		res, err := Popen(command["ip"])
		if err != nil {
			panic(err)
		}
		res = strings.Replace(res, "\n", "", -1)
		hh.IP = strings.Split(res, "/")[0]

		res, err = Popen(command["uname"])
		if err != nil {
			panic(err)
		}
		res = strings.Replace(res, "\n", "", -1)
		hh.Uname = res

		res, err = Popen(command["login_user_count"])
		if err != nil {
			panic(err)
		}
		res = strings.Replace(res, "\n", "", -1)
		hh.LoginUserCount = strings.Split(res, "=")[1]

		hh.CurrentUser = config.LinuxUserInfo.Username
	}
}

func (cc *CPU) Get() {
	cpuInfo, err := cpu.Info()
	if err != nil {
		panic(err)
	}

	cpuCounts, err := cpu.Counts(false)
	if err != nil {
		panic(err)
	}
	cc.Cores = strconv.Itoa(cpuCounts)
	cc.Arch = arch
	cc.Revision = ""

	if strings.Contains(arch, "arm") {
		res, err := Popen(command["cpu_model_name"])
		if err != nil {
			panic(err)
		}
		res = strings.Replace(res, "\n", "", -1)
		cc.ModelName = res
		for key, value := range piCpuModelInfo {
			if strings.Contains(model, key) {
				cc.Revision = cc.ModelName
				cc.ModelName = value
				break
			}
		}
	} else if cpuInfo[0].ModelName != "" {
		cc.ModelName = cpuInfo[0].ModelName
	} else {
		cc.ModelName = "NaN"
	}

	cpuPercent, err := cpu.Percent(0, false)
	cc.UsedPercent = strconv.FormatFloat(cpuPercent[0], 'f', 1, 64)

	if strings.Contains(arch, "arm") {
		res, err := Popen(command["cpu_freq"])
		if err != nil {
			panic(err)
		}
		res = strings.Replace(res, "\n", "", -1)
		cpuFreq, _ := strconv.ParseInt(res, 10, 64)
		cc.Mhz = strconv.FormatInt(cpuFreq/1000, 10)
	} else {
		cc.Mhz = strconv.FormatInt(int64(cpuInfo[0].Mhz), 10)
	}

	if os == "linux" {
		res, err := Popen(command["cpu_status"])
		if err != nil {
			panic(err)
		}
		res = strings.Replace(res, "\n", "", -1)
		cpuStatusRaw := res
		exceptionSituation := []string{"id,", "wa,", "hi,", "si,"}
		for _, exception := range exceptionSituation {
			if exception == "id" {
				cpuStatusRaw = strings.Replace(cpuStatusRaw, exception, "100.0", -1)
			} else {
				cpuStatusRaw = strings.Replace(cpuStatusRaw, exception, "0.0", -1)
			}
		}
		cpuStatus := strings.Split(cpuStatusRaw, " ")
		cc.User = cpuStatus[0]
		cc.Nice = cpuStatus[1]
		cc.System = cpuStatus[2]
		cc.Idle = cpuStatus[3]
		cc.Iowait = cpuStatus[4]
		cc.Irq = cpuStatus[5]
		cc.Softirg = cpuStatus[6]
	} else {
		cpuTimes, _ := cpu.Times(false)
		if err != nil {
			panic(err)
		}
		timesTotal := cpuTimes[0].Total()
		cc.Idle = strconv.FormatFloat(cpuTimes[0].Idle/timesTotal*100, 'f', 1, 64)
		cc.Iowait = strconv.FormatFloat(cpuTimes[0].Iowait/timesTotal*100, 'f', 1, 64)
		cc.Irq = strconv.FormatFloat(cpuTimes[0].Irq/timesTotal*100, 'f', 1, 64)
		cc.Nice = strconv.FormatFloat(cpuTimes[0].Nice/timesTotal*100, 'f', 1, 64)
		cc.Softirg = strconv.FormatFloat(cpuTimes[0].Softirq/timesTotal*100, 'f', 1, 64)
		cc.System = strconv.FormatFloat((cpuTimes[0].System / timesTotal * 100), 'f', 1, 64)
		cc.User = strconv.FormatFloat(cpuTimes[0].User/timesTotal*100, 'f', 1, 64)
	}

	temperature, err := host.SensorsTemperatures()
	if err != nil {
		panic(err)
	}
	if len(temperature) != 0 {
		cc.Temperature = strconv.FormatFloat(temperature[0].Temperature, 'f', 1, 64)
	} else {
		cc.Temperature = "NaN"
	}
}

func (dd *Disk) Get(path string) {
	d, err := disk.Usage(path)
	if err != nil {
		panic(err)
	}

	dd.Path = strings.ToUpper(d.Path)
	dd.Total = strconv.FormatFloat(float64(d.Total)/1024/1024/1024, 'f', 1, 64)
	dd.Free = strconv.FormatFloat(float64(d.Free)/1024/1024/1024, 'f', 1, 64)
	dd.Used = strconv.FormatFloat(float64(d.Used)/1024/1024/1024, 'f', 1, 64)
	dd.UsedPercent = strconv.FormatFloat(d.UsedPercent, 'f', 1, 64)
}

func (mm *Memory) Get() {
	m, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}

	mm.Total = strconv.FormatFloat(float64(m.Total)/1024/1024, 'f', 1, 64)
	mm.Free = strconv.FormatFloat(float64(m.Free)/1024/1024, 'f', 1, 64)
	mm.Buffers = strconv.FormatFloat(float64(m.Buffers)/1024/1024, 'f', 1, 64)
	mm.Cached = strconv.FormatFloat(float64(m.Cached)/1024/1024, 'f', 1, 64)
	mm.Available = strconv.FormatFloat(float64(m.Available)/1024/1024, 'f', 1, 64)
	mm.Used = strconv.FormatFloat(float64((m.Used+m.Buffers+m.Cached))/1024/1024, 'f', 1, 64)
	mm.Percent = strconv.FormatFloat(float64((m.Used+m.Buffers+m.Cached))/float64(m.Total)*100, 'f', 1, 64)
	mm.CachedPercent = strconv.FormatFloat(float64(m.Cached)/float64(m.Total)*100, 'f', 1, 64)
	mm.RealUsed = strconv.FormatFloat(float64((m.Total-m.Available)/1024/1024), 'f', 1, 64)
	mm.RealPercent = strconv.FormatFloat((float64(m.Total)-float64(m.Available))/float64(m.Total)*100, 'f', 1, 64)

	mm.SwapTotal = strconv.FormatFloat(float64(m.SwapTotal)/1024/1024, 'f', 1, 64)
	mm.SwapFree = strconv.FormatFloat(float64(m.SwapFree)/1024/1024, 'f', 1, 64)
	mm.SwapUsed = strconv.FormatFloat((float64(m.SwapTotal)-float64(m.SwapFree))/1024/1024, 'f', 1, 64)
	if m.SwapTotal != 0 && m.SwapTotal-m.SwapFree != 0 {
		swapUsedFloat, _ := strconv.ParseFloat(mm.SwapUsed, 64)
		swapTotalFloat, _ := strconv.ParseFloat(mm.SwapTotal, 64)
		mm.SwapUsedPercent = strconv.FormatFloat(swapUsedFloat/swapTotalFloat*100, 'f', 1, 64)
	} else {
		mm.SwapUsedPercent = "0.0"
	}
	if m.SwapFree == 0 && m.SwapTotal == 0 {
		mm.SwapUsedPercent = "0.0"
		mm.SwapFree = "0.0"
		mm.SwapTotal = "0.0"
		mm.SwapUsed = "0.0"
	}

}

func (nn *Net) Get(name string) {
	n, err := net.IOCounters(true)
	if err != nil {
		panic(err)
	}

	var existed bool
	for _, item := range n {
		if item.Name == name {
			nn.Name = item.Name
			nn.DataIn = strconv.Itoa(int(item.BytesRecv))
			nn.DataOut = strconv.Itoa(int(item.BytesSent))
			nn.PackageIn = strconv.Itoa(int(item.PacketsRecv))
			nn.PackageOut = strconv.Itoa(int(item.PacketsSent))

			existed = true
		}
	}
	if !existed {
		fmt.Println("Not found " + name)
	} else {
		packageInFloat, _ := strconv.ParseFloat(nn.PackageIn, 64)
		packageOutFloat, _ := strconv.ParseFloat(nn.PackageOut, 64)
		nn.LoadAvg = strconv.FormatFloat((packageInFloat+packageOutFloat)/2, 'f', 1, 64)

		dataInFloat, _ := strconv.ParseFloat(nn.DataIn, 64)
		dataOutFloat, _ := strconv.ParseFloat(nn.DataOut, 64)
		nn.DataInFormat = bytesRound(float64(dataInFloat), 2)
		nn.DatatOutFormat = bytesRound(float64(dataOutFloat), 2)
	}
}

func (nn *NetLo) Get(name string) {
	n, err := net.IOCounters(true)
	if err != nil {
		panic(err)
	}

	var existed bool
	for _, item := range n {
		if item.Name == name {
			nn.Name = item.Name
			nn.DataIn = strconv.Itoa(int(item.BytesRecv))
			nn.DataOut = strconv.Itoa(int(item.BytesSent))
			nn.PackageIn = strconv.Itoa(int(item.PacketsRecv))
			nn.PackageOut = strconv.Itoa(int(item.PacketsSent))

			existed = true
		}
	}
	if !existed {
		fmt.Println("Not found " + name)
	} else {
		packageInFloat, _ := strconv.ParseFloat(nn.PackageIn, 64)
		packageOutFloat, _ := strconv.ParseFloat(nn.PackageOut, 64)
		nn.LoadAvg = strconv.FormatFloat((packageInFloat+packageOutFloat)/2, 'f', 1, 64)

		dataInFloat, _ := strconv.ParseFloat(nn.DataIn, 64)
		dataOutFloat, _ := strconv.ParseFloat(nn.DataOut, 64)
		nn.DataInFormat = bytesRound(float64(dataInFloat), 2)
		nn.DatatOutFormat = bytesRound(float64(dataOutFloat), 2)
	}
}

func (ll *Load) Get() {
	l, err := load.Avg()
	if err != nil {
		panic(err)
	}

	ll.Load1 = strconv.FormatFloat(l.Load1, 'f', 2, 64)
	ll.Load5 = strconv.FormatFloat(l.Load5, 'f', 2, 64)
	ll.Load15 = strconv.FormatFloat(l.Load15, 'f', 2, 64)
}

func (pp *Process) Get() {
	p, err := load.Misc()
	if err != nil {
		panic(err)
	}

	pp.Running = strconv.Itoa(p.ProcsRunning)
	pp.Total = strconv.Itoa(p.ProcsTotal)
}

func (hh Host) String() string {
	s, _ := json.Marshal(hh)
	return string(s)
}

func (cc CPU) String() string {
	s, _ := json.Marshal(cc)
	return string(s)
}

func (dd Disk) String() string {
	s, _ := json.Marshal(dd)
	return string(s)
}

func (mm Memory) String() string {
	s, _ := json.Marshal(mm)
	return string(s)
}

func (nn Net) String() string {
	s, _ := json.Marshal(nn)
	return string(s)
}

func (ll Load) String() string {
	s, _ := json.Marshal(ll)
	return string(s)
}

func (pp Process) String() string {
	s, _ := json.Marshal(pp)
	return string(s)
}

func Info() map[string]interface{} {
	device := make(map[string]interface{})

	host := new(Host)
	host.Get()
	hostMap, _ := struct2Map(host, "json")
	device = mergeMap(device, hostMap)
	// fmt.Println("host:", host)

	cpu := new(CPU)
	cpu.Get()
	cpuMap, _ := struct2Map(cpu, "json")
	device = mergeMap(device, cpuMap)
	// fmt.Println("cpu:", cpu)

	disk := new(Disk)
	disk.Get(config.Disk)
	diskMap, _ := struct2Map(disk, "json")
	device = mergeMap(device, diskMap)
	// fmt.Println("disk:", disk)

	mem := new(Memory)
	mem.Get()
	memMap, _ := struct2Map(mem, "json")
	device = mergeMap(device, memMap)
	// fmt.Println("memory:", mem)

	netLo := new(NetLo)
	netLo.Get("lo")
	netLoMap, _ := struct2Map(netLo, "json")
	device = mergeMap(device, netLoMap)
	// fmt.Println("net:", netLo)

	net := new(Net)
	net.Get(config.Net)
	netMap, _ := struct2Map(net, "json")
	device = mergeMap(device, netMap)
	// fmt.Println("net:", net)

	load := new(Load)
	load.Get()
	loadMap, _ := struct2Map(load, "json")
	device = mergeMap(device, loadMap)
	// fmt.Println("load:", load)

	process := new(Process)
	process.Get()
	processMap, _ := struct2Map(process, "json")
	device = mergeMap(device, processMap)
	// fmt.Println("process:", process)

	if strings.Contains(strings.ToLower(device["model"].(string)), "raspberry") {
		device["device_photo"] = "raspberrypi.png"
		device["favicon"] = "raspberrypi.ico"
	} else {
		device["device_photo"] = "linux.png"
		device["favicon"] = "linux.ico"
	}

	return device
}

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

func struct2Map(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}

func mergeMap(map1, map2 map[string]interface{}) map[string]interface{} {
	for k, v := range map2 {
		map1[k] = v
	}

	return map1
}
