package router

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

//WiFi连接状态
var wlanStatus = [...]string{
	"认证",
	"连接",
	"连接(WPA)",
	"连接(WPA-PSK)",
	"连接(WPA2)",
	"连接(WPA2-PSK)",
	"802-1X",
	"加入",
	"启用",
	"关闭",
	"断开",
}

//获取无线设备清单(无线设置-主机状态)
func (cli *Client) GetWlanStation() ([]Host, error) {
	hosts := make([]Host, 0)

	page := 1
	for {
		pageHosts, more, err := cli.getWlanStationPage(page)
		if err != nil {
			return nil, err
		}

		page += 1
		hosts = append(hosts, pageHosts...)

		if !more {
			break
		}

	}

	return hosts, nil
}

func (cli *Client) getWlanStationPage(page int) ([]Host, bool, error) {
	q := url.Values{"Page": []string{strconv.Itoa(page)}}

	bodyBytes, err := cli.request("/userRpm/WlanStationRpm.htm", q)
	if err != nil {
		return nil, false, fmt.Errorf("请求错误: %s", err)
	}

	r := regexp.MustCompile(`(?s)new Array\(([^\)]*)\)`)
	submatchs := r.FindAllSubmatch(bodyBytes, -1)
	if len(submatchs) < 3 {
		return nil, false, fmt.Errorf("错误的响应%s", string(bodyBytes))
	}

	//第一个匹配数组是主机统计信息
	var total, pageIndex, perPage, refresh, itemCount int
	hostParaString := strings.Replace(string(submatchs[0][1]), "\n", "", -1)
	fmt.Sscanf(hostParaString, "%d,%d,%d,%d,%d", &total, &pageIndex, &perPage, &refresh, &itemCount)

	hostListString := strings.Replace(string(submatchs[1][1]), "\n", "", -1)
	hostListSlice := strings.Split(hostListString, ",")

	hosts := make([]Host, 0)
	for i := 0; i < len(hostListSlice)-itemCount; i += itemCount {
		rawMacStr := hostListSlice[i]
		statusIdx, _ := strconv.Atoi(hostListSlice[i+1])

		host := Host{}
		host.Mac = rawMacStr[1 : len(rawMacStr)-1]
		host.Status = wlanStatus[statusIdx]
		host.RecvPackage, _ = strconv.Atoi(hostListSlice[i+4])
		host.SendPackage, _ = strconv.Atoi(hostListSlice[i+5])

		hosts = append(hosts, host)
	}

	var more bool
	if total > perPage && pageIndex*perPage < total {
		more = true
	}

	return hosts, more, nil
}
