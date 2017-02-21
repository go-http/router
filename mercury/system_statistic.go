package router

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const (
	TrafficType = iota
	TrafficTypeAll
	TrafficTypeTCP
	TrafficTypeUDP
	TrafficTypeICMP
)

const (
	SortType = iota
	SortTypeIP
	SortTypeDownloadBytes
	SortTypeUploadBytes
	SortTypeDownloadBps
	SortTypeUploadBps
)

const DefaultPageSize = 1000

func (cli *Client) GetAllHostTraffic() ([]Host, error) {
	return cli.GetHostStatistic(TrafficTypeAll, SortTypeIP, DefaultPageSize, 1)
}

//获取流量统计数据(系统工具-流量统计)
func (cli *Client) GetHostStatistic(trafficType, sortType, pageSize, pageIndex int) ([]Host, error) {
	q := url.Values{}
	q.Set("contType", strconv.Itoa(trafficType))
	q.Set("sortType", strconv.Itoa(sortType))
	q.Set("Goto_page", strconv.Itoa(pageIndex))
	q.Set("Num_per_page", strconv.Itoa(pageSize))

	bodyBytes, err := cli.request("/userRpm/SystemStatisticRpm.htm", q)
	if err != nil {
		return nil, fmt.Errorf("请求失败:%s", err)
	}

	r := regexp.MustCompile(`(?s)var statList=new Array\(([^;]*) \);`)
	match := r.FindSubmatch(bodyBytes)
	if len(match) < 2 {
		return nil, fmt.Errorf("非法的响应%s", string(bodyBytes))
	}

	var hosts []Host
	for _, str := range strings.Split(string(match[1]), "\n") {
		slices := strings.Split(strings.Replace(str, `"`, "", -1), ",")
		if len(slices) < 7 {
			continue
		}
		var host Host
		host.Ip = slices[1]
		host.Mac = slices[2]
		host.SendBytes, _ = strconv.Atoi(slices[3])
		host.RecvBytes, _ = strconv.Atoi(slices[4])
		host.SendSpeed, _ = strconv.Atoi(slices[5])
		host.RecvSpeed, _ = strconv.Atoi(slices[6])

		hosts = append(hosts, host)
	}

	return hosts, nil
}
