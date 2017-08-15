//路由器API应用范例之记录WiFi设备在线时长
//
//    本例的主要功能，就是每分钟刷新一次路由器上的WiFi设备清单
//    并累加记录设备的在线时长
package main

import (
	"flag"
	"log"
	"net/url"
	"time"

	"../../mercury"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//设备在线时长信息
type Duration struct {
	Mac     string
	Date    time.Time
	Minutes int
	gorm.Model
}

func main() {
	var dsn, address string
	flag.StringVar(&dsn, "dsn", "", "数据库服务器DSN")
	flag.StringVar(&address, "addr", "", "路由器地址http://user:pass@192.168.1.1")
	flag.Parse()

	dbConn, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("联接数据库失败", err)
	}

	dbConn.AutoMigrate(&Duration{})

	addr, err := url.Parse(address)
	if err != nil {
		log.Fatal("不合法的路由器地址:%s", err)
	}

	routerClient, err := router.New(addr)
	if err != nil {
		log.Fatal("创建路由器客户端失败:", err)
	}
	routerClient.InsecureSkipVerify = true

	refresh(dbConn, routerClient)
	for range time.Tick(time.Minute) {
		log.Println("tick")

		err = refresh(dbConn, routerClient)
		if err != nil {
			log.Println("刷新失败", err)
		}
	}
}

//刷新无线设备在线记录
func refresh(db *gorm.DB, routerClient *router.Client) error {
	wlanStations, err := routerClient.GetWlanStation()
	if err != nil {
		return err
	}

	now := time.Now()

	//计算本地时区的零点值，作为日期
	_, zoneOffset := now.Zone()
	date := now.Truncate(24 * time.Hour).Add(time.Duration(-zoneOffset) * time.Second)

	for _, wlanStation := range wlanStations {
		//更新在线时长
		var duration Duration
		db.FirstOrInit(&duration, Duration{Mac: wlanStation.Mac, Date: date})
		duration.Minutes += 1
		if err := db.Save(&duration).Error; err != nil {
			log.Println("保存设备", wlanStation.Mac, "在线时长信息错误", err)
		}
	}

	return nil
}
