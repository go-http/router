package router

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var WlanBandWidthSlice = []string{"", "20MHz", "自动", "40MHz"}
var DHCPStatusSlice = []string{"", "正在获取IP地址", "WAN口未连接"}
var WanStatusSlice = []string{"", "已禁用", "等待应答超时", "网线没有插好", "线路正常"}
var WlanTypeSlice = []string{"", "11b only", "11g only", "11n only", "11bg mixed", "11bgn mixed"}
var WanTypeSlice = []string{"", "动态IP", "静态IP", "PPPoE拨号", "802.1X", "BPA", "L2TP", "PPTP", "DHCP+"}
var WDSStatusSlice = []string{"未开启", "初始化...", "扫描...", "认证...", "关联...", "成功", "检测状态..."}

//获取运行状态
func (cli *Client) GetStatus() error {
	bodyBytes, err := cli.request("/userRpm/StatusRpm.htm", nil)
	if err != nil {
		return fmt.Errorf("请求失败:%s", err)
	}

	var idx int
	var slices []string

	//解析路由器状态
	slices = getJsArray(bodyBytes, "statusPara")
	if slices != nil && len(slices) > 3 {
		//slices[0]//WLAN是否启用
		//slices[1]//WAN个数
		//slices[2]//wan参数数量
		//slices[3]//页面刷新间隔ms

		idx, _ = strconv.Atoi(slices[4])
		cli.UpTime = time.Duration(idx) * time.Second
		cli.SoftwareVersion = slices[5]
		cli.HardwareVersion = slices[6]
	}

	//解析LAN口状态
	slices = getJsArray(bodyBytes, "lanPara")
	if slices != nil && len(slices) > 3 {
		cli.Lan.Mac = slices[0]
		cli.Lan.Ip = slices[1]
		cli.Lan.IpMask = slices[2]
	}

	//解析WAN口信息
	slices = getJsArray(bodyBytes, "wanPara")
	if slices != nil && len(slices) > 4 {
		idx, _ = strconv.Atoi(slices[0])
		cli.Wan.Status = WanStatusSlice[idx]

		cli.Wan.Mac = slices[1]
		cli.Wan.Ip = slices[2]

		idx, _ = strconv.Atoi(slices[3])
		cli.Wan.Status = WanTypeSlice[idx]

		cli.Wan.IpMask = slices[4]

		//slices[5]: IEEE8021x WAN是否已登陆    {"", "已登陆","未登陆"}
		//slices[6]: IEEE8021x登陆状态， {"","未登录","登录中","登录成功"}

		cli.Wan.GatewayIp = slices[7]

		//slices[8:9]: DHCP相关按钮状态

		idx, _ = strconv.Atoi(slices[10])
		cli.Wan.DHCPStatus = DHCPStatusSlice[idx]

		cli.Wan.Dns = strings.Split(strings.Replace(slices[11], " ", "", -1), ",")
		cli.Wan.ConnectedDuration = slices[12]
	}

	//解析WAN口流量统计数据
	slices = getJsArray(bodyBytes, "statistList")
	if slices != nil && len(slices) > 4 {
		cli.Wan.UploadBytes, _ = strconv.Atoi(slices[0])
		cli.Wan.DownloadBytes, _ = strconv.Atoi(slices[1])
		cli.Wan.UploadPackages, _ = strconv.Atoi(slices[2])
		cli.Wan.DownloadPackages, _ = strconv.Atoi(slices[3])
	}

	//解析WLAN信息
	slices = getJsArray(bodyBytes, "wlanPara")
	if slices != nil && len(slices) > 11 {
		cli.Wlan.Enable = (slices[0] != "0")
		cli.Wlan.Ssid = slices[1]
		cli.Wlan.Channel = slices[2]

		idx, _ = strconv.Atoi(slices[3])
		cli.Wlan.Mode = WlanTypeSlice[idx]

		cli.Wlan.Mac = slices[4]
		cli.Wlan.Ip = slices[5]

		idx, _ = strconv.Atoi(slices[6])
		cli.Wlan.BandWidth = WlanBandWidthSlice[idx]

		//slices[7]:??
		//slices[8]:发送速率
		//slices[9]:信道号

		idx, _ = strconv.Atoi(slices[10])
		cli.Wlan.WdsStatus = WDSStatusSlice[idx]
	}

	return nil
}
