package router

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//获取NTP时间同步信息
func (cli *Client) GetDateTimeCfg() (time.Time, []string, error) {
	bodyBytes, err := cli.request("/userRpm/DateTimeCfgRpm.htm", nil)
	if err != nil {
		return time.Time{}, nil, fmt.Errorf("请求失败:%s", err)
	}

	if err = getError(bodyBytes); err != nil {
		return time.Time{}, nil, err
	}

	//依次是月日年时分秒时区NTPA及NTPB
	slice := getJsArray(bodyBytes, "timeInf")
	if len(slice) != 10 {
		return time.Time{}, nil, fmt.Errorf("错误的响应数据:%s", string(bodyBytes))
	}

	offsetMinutes, _ := strconv.Atoi(slice[6])
	offsetMinutes -= 12 * 60 //时差是相对于GMT-1200而不是GMT+0000计算的，所以需要纠正回来
	timezone := fmt.Sprintf("%+03d%02d", offsetMinutes/60, offsetMinutes%60)

	t, err := time.Parse("1-2-2006-15-4-5-0700", strings.Join(slice[:6], "-")+timezone)
	return t, slice[7:9], err
}

//触发NTP时间同步
func (cli *Client) SyncNtpTime(ntp []string) error {
	return cli.SetDateTimeCfg(time.Now(), ntp, false)
}

//手动设置时间
func (cli *Client) SetTime(t time.Time) error {
	return cli.SetDateTimeCfg(t, nil, true)
}

//设置时间同步配置
func (cli *Client) SetDateTimeCfg(t time.Time, ntp []string, manualSync bool) error {
	q := url.Values{}
	//选择要执行的操作
	if manualSync {
		q.Set("Save", "保存")
	} else {
		q.Set("GetGmtTime", "获取GMT时间")
	}

	//验证构造NTP服务器参数
	if ntp == nil || len(ntp) == 0 {
		ntp = []string{"0.0.0.0", "0.0.0.0"}
	} else if len(ntp) < 2 {
		ntp = append(ntp, ntp[0])
	}

	q.Set("ntpA", ntp[0])
	q.Set("ntpB", ntp[1])

	//时差是相对于GMT-1200而不是GMT+0000计算的
	_, offsetSeconds := t.Zone()
	offsetMinutes := offsetSeconds/60 + 12*60

	//构造时间参数
	q.Set("timezone", strconv.Itoa(offsetMinutes))
	q.Set("year", strconv.Itoa(t.Year()))
	q.Set("month", strconv.Itoa(int(t.Month())))
	q.Set("day", strconv.Itoa(t.Day()))
	q.Set("hour", strconv.Itoa(t.Hour()))
	q.Set("minute", strconv.Itoa(t.Minute()))
	q.Set("second", strconv.Itoa(t.Second()))

	bodyBytes, err := cli.request("/userRpm/DateTimeCfgRpm.htm", q)
	if err != nil {
		return fmt.Errorf("请求失败:%s", err)
	}

	return getError(bodyBytes)
}
