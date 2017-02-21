# go-router-api
golang版的路由器操作API接口。

主要用于通过API自动控制路由器，而不是通过WEB管理界面手动操作。
目前支持[Mercury](http://www.mercurycom.com.cn)的部分操作，并在MW300R中测试。

# 原理
通过捕获WEB管理界面的请求并分析参数，得到操作路由器的方法。

# 支持的API列表
- [x] 获取DHCP客户端列表
- [x] 获取WiFi连接设备列表
- [x] 获取流量统计列表
- [ ] 修改WiFi密码及SSID
- [ ] 修改IP及Mac地址绑定
- [ ] 其他API

# TODO
* 支持其他更多品牌的路由器，比如TPLink、DLink、Netgear等
