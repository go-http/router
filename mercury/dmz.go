package router

import (
	"fmt"
	"net/url"
)

//设置DMZ主机(转发规则-DMZ主机)
func (cli *Client) SetDMZ(ip string, enable bool) error {
	q := url.Values{}
	q.Set("netmask", "255.255.255.0")
	if enable {
		q.Set("enable", "1")
	} else {
		q.Set("enable", "0")
	}
	q.Set("ipAddr", ip)
	q.Set("Save", "保存")

	bodyBytes, err := cli.request("/userRpm/DMZRpm.htm", q)
	if err != nil {
		return fmt.Errorf("请求失败:%s", err)
	}

	if err = getError(bodyBytes); err != nil {
		return err
	}

	slice := getJsArray(bodyBytes, "DMZInf")
	if len(slice) != 3 {
		return fmt.Errorf("错误的响应数据:%s", string(bodyBytes))
	}

	if slice[1] != ip {
		return fmt.Errorf("设置DMZ主机IP失败,%s!=%s", slice[1], ip)
	}

	if slice[0] != q.Get("enable") {
		return fmt.Errorf("设置DMZ状态失败")
	}

	return nil
}

//获取DMZ主机(转发规则-DMZ主机)
func (cli *Client) GetDMZ() (string, bool, error) {
	bodyBytes, err := cli.request("/userRpm/DMZRpm.htm", nil)
	if err != nil {
		return "", false, fmt.Errorf("请求失败:%s", err)
	}

	slice := getJsArray(bodyBytes, "DMZInf")
	if len(slice) != 3 {
		return "", false, fmt.Errorf("错误的响应数据:%s", string(bodyBytes))
	}

	return slice[1], slice[0] == "1", nil
}
