package router

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

//获取客户端列表(DHCP服务器-客户端列表)
func (cli *Client) GetAssignedIpList() ([]Host, error) {
	bodyBytes, err := cli.request("/userRpm/AssignedIpAddrListRpm.htm", nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败:%s", err)
	}

	r := regexp.MustCompile(`(?s)var DHCPDynList=new Array\(([^;]*) \);`)
	match := r.FindSubmatch(bodyBytes)
	if len(match) < 2 {
		return nil, fmt.Errorf("非法的响应%s", string(bodyBytes))
	}

	str := string(match[1])
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, `"`, "", -1)
	slice := strings.Split(str, ",")

	var hosts []Host
	for i := 0; i < len(slice)-4; i += 4 {
		host := Host{
			Name: strings.TrimPrefix(slice[i], "android-"),
			Mac:  slice[i+1],
			Ip:   slice[i+2],
		}
		t, err := time.Parse("15:04:05", slice[i+3])
		if err == nil {
			d := time.Duration(t.Hour()) * time.Hour
			d += time.Duration(t.Minute()) * time.Minute
			d += time.Duration(t.Second()) * time.Second
			host.IpExpireAt = time.Now().Add(d)
		}
		hosts = append(hosts, host)
	}

	return hosts, nil
}
