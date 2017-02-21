package router

import "time"

type Host struct {
	Ip          string
	Name        string
	Mac         string
	Status      string
	RecvPackage int
	SendPackage int
	SendBytes   int
	RecvBytes   int
	SendSpeed   int
	RecvSpeed   int
	IpExpireAt  time.Time
}
