package router

import (
	"fmt"
	"net"
	"net/url"
	"time"
)

type Client struct {
	Username string
	Password string
	baseUri  string

	UpTime time.Duration

	SoftwareVersion string
	HardwareVersion string

	Lan struct {
		Mac    string
		Ip     string
		IpMask string
	}
	Wlan struct {
		Enable    bool
		Ssid      string
		Channel   string
		Mode      string
		BandWidth string
		Ip        string
		Mac       string
		WdsStatus string
	}
	Wan struct {
		Status            string
		Mac               string
		Ip                string
		IpMask            string
		GatewayIp         string
		DHCPStatus        string
		Dns               []string
		ConnectedDuration string
		UploadBytes       int
		DownloadBytes     int
		UploadPackages    int
		DownloadPackages  int
	}

	InsecureSkipVerify bool
}

const (
	DefaultUsername = "admin"
	DefaultPassword = "admin"
)

func New(addr *url.URL) (*Client, error) {
	client := Client{Username: DefaultUsername, Password: DefaultPassword}
	if addr.User != nil {
		client.Username = addr.User.Username()
		client.Password, _ = addr.User.Password()
	}

	ips, err := net.LookupIP(addr.Hostname())
	if err != nil || len(ips) == 0 {
		return nil, fmt.Errorf("主机解析失败:%s", err)
	}

	port := addr.Port()
	if port == "" {
		port = "80"
	}

	client.baseUri = fmt.Sprintf("%s://%s:%s", addr.Scheme, ips[0], port)

	return &client, nil
}
