package router

import (
	"net/url"
	"os"
	"testing"
	"time"
)

func createClient(t *testing.T) *Client {
	address := os.Getenv("ADDR")
	if address == "" {
		t.Fatal("ADDR参数不能为空")
	}

	addr, _ := url.Parse(address)
	client, err := New(addr)
	if err != nil {
		t.Fatal(err)
	}

	return client
}

func TestGetWlanStation(t *testing.T) {
	hosts, err := createClient(t).GetWlanStation()
	if err != nil {
		t.Error(err)
		return
	}

	if hosts == nil {
		return
	}

	t.Log(" \t无线主机MAC地址  \t当前状态\t包下载\t包上传")
	for i, host := range hosts {
		t.Logf("%2d\t%s\t%s\t%s\t%s\n", i+1, host.Mac, host.Status, HumanInt(host.RecvPackage), HumanInt(host.SendPackage))
	}
}

func TestGetAllHostTraffic(t *testing.T) {
	hosts, err := createClient(t).GetAllHostTraffic()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(" \t当前IP地址\t流量(收/发)\t\t速度(收/发)")
	for i, host := range hosts {
		t.Logf("%2d\t%-13s\t%6s/%6s\t%6s/%6s\n", i+1, host.Ip, HumanInt(host.RecvBytes), HumanInt(host.SendBytes), HumanInt(host.RecvSpeed), HumanInt(host.SendSpeed))
	}
}

func TestGetAssignedIpList(t *testing.T) {
	hosts, err := createClient(t).GetAssignedIpList()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("   客户端名称               MAC地址            IP地址     过期时间")
	for i, host := range hosts {
		t.Logf("%2d %-20s %s   %-13s %s\n", i+1, host.Name, host.Mac, host.Ip, host.IpExpireAt.Format("15:04:05"))
	}
}

func TestGetStatus(t *testing.T) {
	t1 := time.Now()
	t1 = t1.AddDate(-1, 0, 0)
	t.Log(createClient(t).SetDateTimeCfg(t1, nil, true))
}
