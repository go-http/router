package router

//所有的错误码定义
const (
	ERR_NO_ERROR                                     = 0
	ERR_FS_FILE_OPEN_FAILED                          = 10
	ERR_PPPOE_FIXED_IP                               = 1000
	ERR_PPPOE_TIMING_SET                             = 1001
	ERR_PPPOE_STRING_TOO_LONG                        = 1002
	ERR_PPPOE_USERNAME_TOO_LONG                      = 1003
	ERR_PPPOE_PWD_TOO_LONG                           = 1004
	ERR_PPPOE_AUTO_OFF_WAITE_TIME                    = 1005
	ERR_PPPOE_LCP_MRU                                = 1006
	ERR_PPPOE_ECHO_REQ_INTERVAL                      = 1007
	ERR_DHCP_SERVER_ADDR_POOL_ERROR                  = 2000
	ERR_DHCP_SERVER_GATEWAY_ERROR                    = 2001
	ERR_DHCP_SERVER_DNS_ERROR                        = 2002
	ERR_DHCP_SERVER_BAK_DNS_ERROR                    = 2003
	ERR_DHCP_SERVER_LEASE                            = 2004
	ERR_DHCP_SERVER_START_IP_ADDR                    = 2005
	ERR_DHCP_SERVER_END_IP_ADDR                      = 2006
	ERR_DHCP_SERVER_START_BIGGER_END                 = 2007
	ERR_DHCP_SERVER_ADD_RANGE                        = 2008
	ERR_FIX_MAP_MAC_ADDR_ERROR                       = 3000
	ERR_FIX_MAP_IP_ADDR_ERROR                        = 3001
	ERR_FIX_MAP_REC_EXIST                            = 3002
	ERR_FIX_MAP_PAGE_NUM_ERROR                       = 3003
	ERR_FIX_MAP_RECORD_ALREADY_FULL                  = 3004
	ERR_FIX_MAP_RECORD_MAC_ALREADY_EXIST             = 3005
	ERR_FIX_MAP_RECORD_IP_ALREADY_EXIST              = 3006
	ERR_FIX_MAP_IP_EQUAL_LANIP                       = 3007
	ERR_STATIC_ROUTR_ENABLE                          = 4000
	ERR_STATIC_ROUTR_DESTINATION_IP                  = 4001
	ERR_STATIC_ROUTR_SUBNETMASK_IP                   = 4002
	ERR_STATIC_ROUTR_SUBNETMASK_DISMATCH_IP          = 4003
	ERR_STATIC_ROUTR_GATEWAY_IP                      = 4004
	ERR_STATIC_ROUTR_NOEMPTY                         = 4005
	ERR_STATIC_ROUTR_DUPLICATION                     = 4006
	ERR_STATIC_ROUTR_DEFAULT_GATEWAY                 = 4007
	ERR_STATIC_ROUTR_NOT_SAME_NETWORK                = 4008
	ERR_STATIC_ROUTR_CONFLICT_LAN_WAN                = 4009
	ERR_STATIC_DEST_CONFLICT_LAN                     = 4010
	ERR_STATIC_DEST_CONFLICT_WAN                     = 4011
	ERR_STATIC_ROUTR_ALREADY_FULL                    = 4012
	ERR_STATIC_ROUTR_SAVE                            = 4013
	ERR_STATIC_ROUTR_OTHER                           = 4014
	ERR_WAN_DOWN_BANDWIDTH                           = 4015
	ERR_WAN_UP_BANDWIDTH                             = 4016
	ERR_NETWORK_MTU                                  = 5000
	ERR_LAN_IP_ERROR                                 = 5001
	ERR_LAN_MASK_ERROR                               = 5002
	ERR_WAN_IP_ERROR                                 = 5003
	ERR_WAN_MASK_ERROR                               = 5004
	ERR_WAN_DNS_ERROR                                = 5005
	ERR_WAN_BACKDNS_ERROR                            = 5006
	ERR_WAN_GATE_ERROR                               = 5007
	ERR_WAN_LAN_CONFLICT                             = 5008
	ERR_WAN_TYPE                                     = 5009
	ERR_LAN_IP_SET                                   = 5010
	ERR_LAN_MASK_SET                                 = 5011
	ERR_WAN_IP_SERVER                                = 5012
	ERR_WAN_IP_SET                                   = 5013
	ERR_WAN_MASK_SET                                 = 5014
	ERR_WAN_DNS_SET                                  = 5015
	ERR_WAN_GATE_SET                                 = 5016
	ERR_WAN_MAC_ADDR                                 = 5017
	ERR_WAN_MAC_DUPLICATE                            = 5018
	ERR_WAN_MAC_EQ_LAN_MAC                           = 5019
	ERR_SNTP_MONTH                                   = 5020
	ERR_SNTP_DAY                                     = 5021
	ERR_SNTP_YEAR                                    = 5022
	ERR_SNTP_HOUR                                    = 5023
	ERR_SNTP_MINUTE                                  = 5024
	ERR_SNTP_SECOND                                  = 5025
	ERR_SNTP_TIME_SET                                = 5026
	ERR_SNTP_TIMEZONE                                = 5027
	ERR_SNTP_GET_GMT_FAILED                          = 5028
	ERR_SAME_WAN_IP                                  = 5029
	ERR_SNTP_SERVER_A                                = 5030
	ERR_SNTP_SERVER_B                                = 5031
	ERR_SERVER_IP_ERROR                              = 5032
	ERR_LAN_WAN_CONFLICT                             = 5033
	ERR_WAN_LINK_MODE_ERROR                          = 5034
	ERR_TFTP_OVER_FILE_LEN                           = 7000
	ERR_TFTP_IP_ERROR                                = 7001
	ERR_TFTP_FILE_NOTFOUND                           = 7002
	ERR_TFTP_SERVER_NOTEXIST                         = 7003
	ERR_FIREWALL_START_TIME_FORMAT_ERROR             = 8000
	ERR_FIREWALL_END_TIME_FORMAT_ERROR               = 8001
	ERR_FIREWALL_TIME_START_BIGGER_END               = 8002
	ERR_FIREWALL_LAN_IP_FORMAT_ERROR                 = 8003
	ERR_FIREWALL_LAN_PORT_FORMAT_ERROR               = 8004
	ERR_FIREWALL_WAN_IP_FORMAT_ERROR                 = 8005
	ERR_FIREWALL_WAN_PORT_FORMAT_ERROR               = 8006
	ERR_FIREWALL_PROTOCOL_TYPE_ERROR                 = 8007
	ERR_FIREWALL_RECORD_ALREADY_EXIST                = 8008
	ERR_FIREWALL_IP_RECORD_ALREADY_FULL              = 8009
	ERR_FIREWALL_DOMAIN_NAME_LEN_OVER                = 9000
	ERR_FIREWALL_DOMAIN_NAME_ERROR                   = 9001
	ERR_FIREWALL_DOMAIN_IS_SUBSET                    = 9002
	ERR_FIREWALL_DOMAIN_RECORD_ALREADY_FULL          = 9003
	ERR_FIREWALL_TIME_NOT_FULL                       = 10000
	ERR_FIREWALL_TIME_FORMAT_ERROR                   = 10001
	ERR_FIREWALL_WZD_TIME_ALREADY_EXIST              = 10002
	ERR_FIREWALL_WZD_TIME_IS_SUBSET                  = 10003
	ERR_FIREWALL_WZD_IP_FORMAT_ERROR                 = 10004
	ERR_FIREWALL_WZD_ADDR_ALREADY_EXIST              = 10005
	ERR_FIREWALL_WZD_PORT_FORMAT_ERROR               = 10006
	ERR_FIREWALL_WZD_PORT_IS_SUBSET                  = 10007
	ERR_MAC_FILTER_PAGE_NUM_ERROR                    = 11000
	ERR_MAC_FILTER_RECORD_ALREADY_EXIST              = 11001
	ERR_MAC_FILTER_RECORD_ALREADY_FULL               = 11002
	ERR_MAC_FILTER_FORMAT_ERROR                      = 11003
	ERR_REMOTE_MANAGE_IP_FORMAT_ERROR                = 12000
	ERR_REMOTE_MANAGE_PORT_FORMAT_ERROR              = 12001
	ERR_REMOTE_MANAGE_PORT_OUT_OF_RANGE              = 12002
	ERR_REMOTE_MANAGE_UPDATE_FORBIDDEN               = 12003
	ERR_DMZ_HOST_IP_ADDR                             = 13000
	ERR_VS_PAGE_NUM_ERROR                            = 14000
	ERR_VS_PORT_OUT_RANGE                            = 14001
	ERR_VS_PORT_FORMAT_ERROR                         = 14002
	ERR_VS_IP_ADDRESS                                = 14003
	ERR_VS_RECORD_ALREADY_EXIST                      = 14004
	ERR_VS_PROTOCOL_TYPE_ERROR                       = 14005
	ERR_VS_RECORD_ALREADY_FULL                       = 14006
	ERR_VS_PORT_DUPLICATE                            = 14007
	ERR_VS_PUBLIC_PORT_TOO_MANY                      = 14008
	ERR_SPECIAL_APP_PUBLIC_PORT                      = 15000
	ERR_SPECIAL_APP_DUPLICATE_PUBLIC_PORT            = 15001
	ERR_SPECIAL_APP_DUPLICATE_TAG_PORT               = 15002
	ERR_SPECIAL_APP_RECORD_ALREADY_FULL              = 15003
	ERR_SPECIAL_APP_PUBLIC_PORT_RESERVE              = 15004
	ERR_SPECIAL_APP_PUBLIC_PORT_TOO_MANY             = 15005
	ERR_DDNS_USER_NAME_EMPTY                         = 16000
	ERR_DDNS_PWD_EMPTY                               = 16001
	ERR_DDNS_USER_HAS_SPACE                          = 16002
	ERR_DDNS_PWD_HAS_SPACE                           = 16003
	ERR_DDNS_LIST_FULL                               = 16004
	ERR_DDNS_LIST_INDEX_OUT_RANGE                    = 16005
	ERR_DDNS_ENTRY_BE_DELETE                         = 16006
	ERR_USER_NAME_LENGTH                             = 17000
	ERR_USER_PWD_LENGTH                              = 17001
	ERR_USER_NAME_ERROR                              = 17002
	ERR_USER_PWD_ERROR                               = 17003
	ERR_USER_PWD_INVALID_CHAR                        = 17004
	ERR_FIRMWARE_FAIL                                = 18000
	ERR_FIRMWARE_VALID_LENGTH                        = 18001
	ERR_FIRMWARE_VALID_FORMAT                        = 18002
	ERR_FIRMWARE_VALID_PRODUCTID                     = 18003
	ERR_FIRMWARE_VALID_MDCHECK                       = 18004
	ERR_FIRMWARE_VALID_VERSION                       = 18005
	ERR_FIRMWARE_VALID_LANGUAGE                      = 18006
	ERR_FIRMWARE_VALID_CONTRY                        = 18007
	ERR_FIRMWARE_VALID_OEM                           = 18008
	ERR_ARP_REC_IP_EXIST                             = 20000
	ERR_ARP_FIXMAP_FULL                              = 20001
	ERR_ARP_REC_IP_EXIST_ADD_SUCC                    = 20002
	ERR_ARP_REC_IP_EXIST_ADD_FAIL                    = 20003
	ERR_ARP_IP_EXIST_AND_FIXMAP_FULL                 = 20004
	ERR_ARP_FIXMAP_FULL_IGNORE_OTHER_ENTRYS          = 20005
	ERR_ARP_FIXMAP_MAC_ERR                           = 20006
	ERR_ARP_IP_SAME_AS_LANIP                         = 20007
	ERR_SYS_LOG_SYS_STATUS                           = 21000
	ERR_SYS_LOG_SRV_ID                               = 21001
	ERR_SYS_LOG_SRV_STATUS                           = 21002
	ERR_SYS_LOG_SRV_ADDRESS                          = 21003
	ERR_SYS_LOG_SRV_ADDR_EXIST                       = 21004
	ERR_SYS_LOG_SRV_PORT                             = 21005
	ERR_SYS_LOG_SETTING_EMERGENCY                    = 21006
	ERR_SYS_LOG_SETTING_ALERT                        = 21007
	ERR_SYS_LOG_SETTING_CRITICAL                     = 21008
	ERR_SYS_LOG_SETTING_ERROR                        = 21009
	ERR_SYS_LOG_SETTING_WARNING                      = 21010
	ERR_SYS_LOG_SETTING_NOTICE                       = 21011
	ERR_SYS_LOG_SETTING_INFORMATIONAL                = 21012
	ERR_SYS_LOG_SETTING_DEBUG                        = 21013
	ERR_SYS_LOG_SETTING_EMPTY                        = 21014
	ERR_FIREWALL_SYSLOG_SERVER_INVALID_ID            = 22000
	ERR_FIREWALL_SYSLOG_SERVER_NOT_DEFINED           = 22001
	ERR_FIREWALL_SCREEN_UNKNOWN_DEFENCE              = 22002
	ERR_FIREWALL_SCREEN_SCAN_THRESHOLD               = 22003
	ERR_FIREWALL_SCREEN_DOS_THRESHOLD                = 22004
	ERR_TDDP_UPLOAD_FILE_TOO_LONG                    = 23000
	ERR_TDDP_UPLOAD_FILE_FORMAT_ERR                  = 23001
	ERR_TDDP_UPLOAD_FILE_NAME_ERR                    = 23002
	ERR_COMMON_ERROR                                 = 25000
	ERR_TDDP_DOWNLOAD_FILE_TOO_LONG                  = 25001
	ERR_VS_RECORD_CONFLICT_REMOTE_WEB_PORT           = 25002
	ERR_DST_HOUR                                     = 26000
	ERR_DST_DAY                                      = 26001
	ERR_DST_MONTH                                    = 26002
	ERR_DST_BEGIN_END                                = 26003
	ERR_WLAN_CONFIG_BASE                             = 26100
	ERR_WLAN_CONFIG_SECURITY                         = 26101
	ERR_WLAN_CONFIG_KEY                              = 26102
	ERR_WLAN_MAC_FILTER_PAGE_NUM_ERROR               = 26103
	ERR_WLAN_MAC_FILTER_RECORD_ALREADY_EXIST         = 26104
	ERR_WLAN_MAC_FILTER_RECORD_ALREADY_FULL          = 26105
	ERR_IP_NOT_IN_THE_SAME_SUBNET                    = 26106
	ERR_WLAN_SSID_LEN                                = 26107
	ERR_WLAN_REGION                                  = 26108
	ERR_WLAN_CHANNEL_WIDTH                           = 26109
	ERR_WLAN_STATIC_RATE                             = 26110
	ERR_WLAN_MODE                                    = 26111
	ERR_WLAN_BROADCAST                               = 26112
	ERR_WLAN_MAC_ADDR_INVALID                        = 26113
	ERR_WLAN_RADIUS_IP_INVALID                       = 26114
	ERR_WLAN_CHANNEL                                 = 26115
	ERR_QOS_TOTAL_EGRESS_100M                        = 27000 /* ZJin 090903: make the err msg more detailedly */
	ERR_QOS_TOTAL_INGRESS_100M                       = 27001
	ERR_QOS_TOTAL_EGRESS_1000M                       = 27002
	ERR_QOS_TOTAL_INGRESS_1000M                      = 27003
	ERR_QOS_NOBUF                                    = 27004 /* 已经没有空间*/
	ERR_QOS_NOENT                                    = 27005 /* 不存在在信息*/
	ERR_QOS_EXIST                                    = 27006 /* 该信息已经存在*/
	ERR_QOS_USEDBW                                   = 27007 /* 新的QoS 带宽小于已使用的带宽*/
	ERR_QOS_NOBW                                     = 27008 /* 系统不能满足所要求的带宽*/
	ERR_QOS_BADRULE                                  = 27009 /* 规则有交集*/
	ERR_QOS_TYPE                                     = 27010 /* 错误的类型*/
	ERR_QOS_INGRESS_BANDWIDTH                        = 27011 /*下行带宽总算大于系统提供总数*/
	ERR_QOS_EGRESS_BANDWIDTH                         = 27012 /*上行带宽总数大于系统提供总算*/
	ERR_QOS_MAX                                      = 27013 /*the max error code */
	ERR_PARENT_CTRL_FULL                             = 28000 /* 家长控制列表满 */
	ERR_PARENT_CTRL_URLDESC                          = 28001 /* URL 描述重复 */
	ERR_PARENT_CTRL_SAME_MAC_WITH_PARENT             = 28002 /* 小孩MAC与家长MAC相同*/
	ERR_ACC_CTRL_HOST_FULL                           = 29000
	ERR_ACC_CTRL_TARGET_FULL                         = 29001
	ERR_ACC_CTRL_SCHEDULE_FULL                       = 29002
	ERR_ACC_CTRL_RULE_FULL                           = 29003
	ERR_ACC_CTRL_SAME_NAME                           = 29004
	ERR_ACC_CTRL_REFERED                             = 29005
	ERR_ACC_CTRL_RULE_CONFLICT                       = 29006
	ERR_ACC_PARTIAL_DEL                              = 29007
	ERR_ACC_DEL_NONE                                 = 29008
	ERR_FILTER_MAC                                   = 29009
	ERR_ACC_CTRL_HOST_IPSTART                        = 29010
	ERR_ACC_CTRL_HOST_IPEND                          = 29011
	ERR_ACC_CTRL_TARGET_IPSTART                      = 29012
	ERR_ACC_CTRL_TARGET_IPEND                        = 29013
	ERR_ACC_CTRL_HOST_IPSTART_NOT_IN_THE_SAME_SUBNET = 29014
	ERR_ACC_CTRL_HOST_IPEND_NOT_IN_THE_SAME_SUBNET   = 29015
	ERR_LAN_GATEWAY_ERROR                            = 500 //for 501
	ERR_LAN_GATEWAY_IP                               = 501
	ERR_LAN_SAME_IP                                  = 510
	ERR_LAN_SAME_MAC                                 = 511
	ERR_DMZ_IP                                       = 512
	ERR_PORT_HAS_BEEN_USED                           = 520
)

//错误码对应的错误提示
var errorStringMap = map[int]string{
	ERR_NO_ERROR:                                     "有错误发生，请重试！",
	ERR_PPPOE_FIXED_IP:                               "IP地址错误，请重新输入。",
	ERR_PPPOE_TIMING_SET:                             "定时连接设置错误。",
	ERR_PPPOE_STRING_TOO_LONG:                        "PPPOE密码或用户名字长度出错。",
	ERR_PPPOE_USERNAME_TOO_LONG:                      "上网帐号长度不得超过119个字符，请重新输入。",
	ERR_PPPOE_PWD_TOO_LONG:                           "上网密码长度不得超过119个字符，请重新输入。",
	ERR_PPPOE_AUTO_OFF_WAITE_TIME:                    "自动断线等待时间超出合法范围（10－99），请重新输入。",
	ERR_PPPOE_LCP_MRU:                                "MTU设置超出合法范围（576－1500），请重新输入。",
	ERR_PPPOE_ECHO_REQ_INTERVAL:                      "在线检测间隔时间错误。",
	ERR_DHCP_SERVER_ADDR_POOL_ERROR:                  "IP地址池错误，请重新输入。",
	ERR_DHCP_SERVER_GATEWAY_ERROR:                    "网关错误，请重新输入。",
	ERR_DHCP_SERVER_DNS_ERROR:                        "首选DNS服务器地址错误，请重新输入。",
	ERR_DHCP_SERVER_BAK_DNS_ERROR:                    "备用DNS服务器地址错误，请重新输入。",
	ERR_DHCP_SERVER_LEASE:                            "地址租期超出范围（1-2880）。请重新输入。",
	ERR_DHCP_SERVER_START_IP_ADDR:                    "地址池开始地址错误，请重新输入。",
	ERR_DHCP_SERVER_END_IP_ADDR:                      "地址池结束地址错误，请重新输入。",
	ERR_DHCP_SERVER_START_BIGGER_END:                 "地址池开始地址大于结束地址。",
	ERR_DHCP_SERVER_ADD_RANGE:                        "地址池范围不得大于256，请重新输入。",
	ERR_FIX_MAP_MAC_ADDR_ERROR:                       "MAC地址错误，请重新输入。",
	ERR_FIX_MAP_IP_ADDR_ERROR:                        "IP地址错误，请重新输入。",
	ERR_FIX_MAP_REC_EXIST:                            "静态地址分配条目已存在，请重新输入。",
	ERR_FIX_MAP_PAGE_NUM_ERROR:                       "静态地址分配页面号错误，请重新输入。",
	ERR_FIX_MAP_RECORD_ALREADY_FULL:                  "静态地址分配列表已满。",
	ERR_FIX_MAP_RECORD_MAC_ALREADY_EXIST:             "MAC地址已被包含在其它静态地址分配条目中。",
	ERR_FIX_MAP_RECORD_IP_ALREADY_EXIST:              "IP地址已被包含在其它静态地址分配条目中。",
	ERR_FIX_MAP_IP_EQUAL_LANIP:                       "IP地址不能与LAN口IP地址相同。",
	ERR_STATIC_ROUTR_ENABLE:                          "有错误发生，请重试！",
	ERR_STATIC_ROUTR_DESTINATION_IP:                  "目的IP地址错误，请重新输入。",
	ERR_STATIC_ROUTR_SUBNETMASK_IP:                   "子网掩码错误，请重新输入。",
	ERR_STATIC_ROUTR_SUBNETMASK_DISMATCH_IP:          "子网掩码和目的网络地址不匹配，请重新输入。",
	ERR_STATIC_ROUTR_GATEWAY_IP:                      "网关错误，请重新输入。",
	ERR_STATIC_ROUTR_NOEMPTY:                         "有错误发生，请重试！",
	ERR_STATIC_ROUTR_DUPLICATION:                     "静态路由表条目已存在，请重新输入。",
	ERR_STATIC_ROUTR_DEFAULT_GATEWAY:                 "默认网关，请在网络参数菜单下进行设置。",
	ERR_STATIC_ROUTR_NOT_SAME_NETWORK:                "网关地址必须与LAN或WAN在同一子网。",
	ERR_STATIC_ROUTR_CONFLICT_LAN_WAN:                "静态路由条目与LAN口设置或WAN口设置冲突，请重新输入。",
	ERR_STATIC_DEST_CONFLICT_LAN:                     "目的网络地址不能与LAN口地址处于同一子网，请重新输入。",
	ERR_STATIC_DEST_CONFLICT_WAN:                     "目的网络地址不能处于WAN口地址所在的子网内部，请重新输入。",
	ERR_STATIC_ROUTR_ALREADY_FULL:                    "静态路由表已满！",
	ERR_STATIC_ROUTR_SAVE:                            "有错误发生，请重试！",
	ERR_STATIC_ROUTR_OTHER:                           "有错误发生，请重试！",
	ERR_WAN_DOWN_BANDWIDTH:                           "下行带宽超出允许范围，请重新输入（1-100000）。",
	ERR_WAN_UP_BANDWIDTH:                             "上行带宽超出允许范围，请重新输入（1-100000）。",
	ERR_QOS_TOTAL_EGRESS_100M:                        "上行总带宽超出允许范围，请重新输入（1-100000）。建议您设置为ISP分配带宽。",
	ERR_QOS_TOTAL_INGRESS_100M:                       "下行总带宽超出允许范围，请重新输入（1-100000）。建议您设置为ISP分配带宽。",
	ERR_QOS_TOTAL_EGRESS_1000M:                       "上行总带宽超出允许范围，请重新输入（1-1000000）。建议您设置为ISP分配带宽。",
	ERR_QOS_TOTAL_INGRESS_1000M:                      "下行总带宽超出允许范围，请重新输入（1-1000000）。建议您设置为ISP分配带宽。",
	ERR_QOS_NOBUF:                                    "控制条目已满，已经没有空间。",
	ERR_QOS_NOENT:                                    "不存在该条目信息。",
	ERR_QOS_EXIST:                                    "该条目信息已经存在。",
	ERR_QOS_USEDBW:                                   "新配置的总带宽小于已分配的带宽之和。",
	ERR_QOS_NOBW:                                     "系统不能满足所要求的带宽，请重新输入。",
	ERR_QOS_BADRULE:                                  "设定的规则有交集，请重新输入。",
	ERR_QOS_TYPE:                                     "错误的类型。",
	ERR_QOS_INGRESS_BANDWIDTH:                        "下行带宽总算大于系统提供总数。",
	ERR_QOS_EGRESS_BANDWIDTH:                         "上行带宽总数大于系统提供总算。",
	ERR_ARP_IP_SAME_AS_LANIP:                         "禁止将LAN口IP与其它MAC地址绑定。",
	ERR_NETWORK_MTU:                                  "MTU输入错误，请重新输入（576－1500）。",
	ERR_LAN_IP_ERROR:                                 "LAN口IP地址错误，请重新输入。",
	ERR_LAN_MASK_ERROR:                               "LAN口子网掩码错误，请重新输入。",
	ERR_WAN_IP_ERROR:                                 "WAN口IP地址错误，请重新输入。",
	ERR_WAN_MASK_ERROR:                               "WAN口子网掩码错误，请重新输入。",
	ERR_WAN_DNS_ERROR:                                "首选DNS服务器地址错误，请重新输入。",
	ERR_WAN_BACKDNS_ERROR:                            "备用DNS服务器地址错误，请重新输入。",
	ERR_WAN_GATE_ERROR:                               "WAN口网关错误，请重新输入。",
	ERR_WAN_LAN_CONFLICT:                             "WAN口IP地址和LAN口IP地址不能处于同一子网，请重新输入。",
	ERR_WAN_TYPE:                                     "错误的网络连接类型，请选择正确的连接类型。",
	ERR_LAN_IP_SET:                                   "有错误发生，请重试！",
	ERR_LAN_MASK_SET:                                 "有错误发生，请重试！",
	ERR_WAN_IP_SERVER:                                "WAN口IP错误。",
	ERR_WAN_IP_SET:                                   "有错误发生，请重试！",
	ERR_WAN_MASK_SET:                                 "有错误发生，请重试！",
	ERR_WAN_DNS_SET:                                  "有错误发生，请重试！",
	ERR_WAN_GATE_SET:                                 "有错误发生，请重试！",
	ERR_WAN_MAC_ADDR:                                 "MAC地址错误，请重新输入。",
	ERR_WAN_MAC_DUPLICATE:                            "重复的MAC地址。",
	ERR_WAN_MAC_EQ_LAN_MAC:                           "WAN口MAC地址和LAN口MAC地址冲突。",
	ERR_SNTP_MONTH:                                   "时间输入错误，请重新输入（月的范围为从01到12）。",
	ERR_SNTP_DAY:                                     "时间输入错误，请重新输入（日的范围为从01到31）。",
	ERR_SNTP_YEAR:                                    "时间输入错误，请重新输入（年的范围为从1970到2037）。",
	ERR_SNTP_HOUR:                                    "时间输入错误，请重新输入（时的范围为从00到23）。",
	ERR_SNTP_MINUTE:                                  "时间输入错误，请重新输入（分的范围为从00到59）。",
	ERR_SNTP_SECOND:                                  "时间输入错误，请重新输入（秒的范围为从00到59）。",
	ERR_SNTP_TIME_SET:                                "时间设置失败，请重试。",
	ERR_SNTP_TIMEZONE:                                "时区错误，请选择正确的时区。",
	ERR_SNTP_GET_GMT_FAILED:                          "获取网络时间错误，请检查是否正确连接到网络。",
	ERR_TFTP_OVER_FILE_LEN:                           "文件名超出长度，允许的文件名最大长度64字符。",
	ERR_TFTP_IP_ERROR:                                "TFTP服务器错误。",
	ERR_TFTP_FILE_NOTFOUND:                           "找不到指定的文件。",
	ERR_TFTP_SERVER_NOTEXIST:                         "探测不到服务器。",
	ERR_FIREWALL_START_TIME_FORMAT_ERROR:             "开始时间错误，请以格式hhmm输入24小时制时间。",
	ERR_FIREWALL_END_TIME_FORMAT_ERROR:               "结束时间错误，请以格式hhmm输入24小时制时间。",
	ERR_FIREWALL_TIME_START_BIGGER_END:               "开始时间不能晚于结束时间，请重新输入。",
	ERR_FIREWALL_LAN_IP_FORMAT_ERROR:                 "LAN口IP地址错误，请重新输入。",
	ERR_FIREWALL_LAN_PORT_FORMAT_ERROR:               "LAN口端口号超出允许范围（1-65535），请重新输入。",
	ERR_FIREWALL_WAN_IP_FORMAT_ERROR:                 "WAN口IP地址错误，请重新输入。",
	ERR_FIREWALL_WAN_PORT_FORMAT_ERROR:               "WAN口端口号超出允许范围（1-65535），请重新输入。",
	ERR_FIREWALL_PROTOCOL_TYPE_ERROR:                 "协议选择错误，请重试。",
	ERR_FIREWALL_RECORD_ALREADY_EXIST:                "条目已存在。",
	ERR_FIREWALL_IP_RECORD_ALREADY_FULL:              "IP地址过滤表已满。",
	ERR_FIREWALL_DOMAIN_NAME_LEN_OVER:                "域名长度超长，请重新输入。",
	ERR_FIREWALL_DOMAIN_NAME_ERROR:                   "域名不合法，请重新输入。",
	ERR_FIREWALL_DOMAIN_IS_SUBSET:                    "域名存在包含关系，已包含另一条目或已被另一条目包含，请检查并重新输入。",
	ERR_FIREWALL_DOMAIN_RECORD_ALREADY_FULL:          "域名过滤表已满。",
	ERR_FIREWALL_TIME_NOT_FULL:                       "请同时输入开始时间和结束时间。",
	ERR_FIREWALL_TIME_FORMAT_ERROR:                   "时间格式错误，请以hhmm格式输入24小时制时间。",
	ERR_FIREWALL_WZD_TIME_ALREADY_EXIST:              "时间范围与其它条目冲突。",
	ERR_FIREWALL_WZD_TIME_IS_SUBSET:                  "时间范围已被其它条目包含。",
	ERR_FIREWALL_WZD_IP_FORMAT_ERROR:                 "地址错误，请重新输入IP地址或IP地址段或域名。<br>如果输入的是IP地址段，请确保起始IP和结束IP在同一IP地址段，并且可用。",
	ERR_FIREWALL_WZD_ADDR_ALREADY_EXIST:              "IP地址或域名冲突。",
	ERR_FIREWALL_WZD_PORT_FORMAT_ERROR:               "端口号错误，<br>请输入1-65535之间的数字，或以“-”连接的端口段，端口或端口段之间以“,”分隔。<br><br>",
	ERR_FIREWALL_WZD_PORT_IS_SUBSET:                  "端口号已被其它条目使用。",
	ERR_MAC_FILTER_PAGE_NUM_ERROR:                    "MAC过滤表页错误。",
	ERR_MAC_FILTER_RECORD_ALREADY_EXIST:              "条目已存在，请重新输入。",
	ERR_MAC_FILTER_RECORD_ALREADY_FULL:               "MAC地址过滤表已满。",
	ERR_MAC_FILTER_FORMAT_ERROR:                      "MAC地址错误，请重新输入。",
	ERR_REMOTE_MANAGE_IP_FORMAT_ERROR:                "IP地址错误，请重新输入。",
	ERR_REMOTE_MANAGE_PORT_FORMAT_ERROR:              "端口号错误，请重新输入。",
	ERR_REMOTE_MANAGE_PORT_OUT_OF_RANGE:              "远程管理端口号超出范围（1-65535）或为浏览器不支持的端口（21、25、110等），请重新输入。",
	ERR_REMOTE_MANAGE_UPDATE_FORBIDDEN:               "远程管理禁止页面升级。",
	ERR_DMZ_HOST_IP_ADDR:                             "DMZ主机错误，请重新输入。",
	ERR_VS_PAGE_NUM_ERROR:                            "虚拟服务器表页错误。",
	ERR_VS_PORT_OUT_RANGE:                            "端口号超出范围（1-65535），请重新输入。",
	ERR_VS_PORT_FORMAT_ERROR:                         "端口号可以是数字或者数字范围（以“-”连接），请重新输入。",
	ERR_VS_IP_ADDRESS:                                "IP地址错误，请重新输入。",
	ERR_VS_RECORD_ALREADY_EXIST:                      "条目已存在或端口号已被其它条目使用，请重新输入。",
	ERR_VS_PROTOCOL_TYPE_ERROR:                       "协议类型错误，请重新选择。",
	ERR_VS_RECORD_ALREADY_FULL:                       "虚拟服务器表已满。",
	ERR_VS_PORT_DUPLICATE:                            "端口号已经被占用，请重新输入。",
	ERR_VS_PUBLIC_PORT_TOO_MANY:                      "开放端口数量过多，请重新输入。",
	ERR_SPECIAL_APP_PUBLIC_PORT:                      "开放端口错误，请重新输入。",
	ERR_SPECIAL_APP_DUPLICATE_PUBLIC_PORT:            "开放端口与已存在条目冲突，请重新输入。<br>两个条目不能设置相同的开放端口。",
	ERR_SPECIAL_APP_DUPLICATE_TAG_PORT:               "触发条件与已存在条目冲突，请重新输入。<br>两个条目不能将触发协议和触发端口都设为相同。",
	ERR_SPECIAL_APP_RECORD_ALREADY_FULL:              "特殊应用程序表已满。",
	ERR_SPECIAL_APP_PUBLIC_PORT_RESERVE:              "开放端口号不能包含路由器的内部保留端口(2049-2240)，请重新输入。",
	ERR_SPECIAL_APP_PUBLIC_PORT_TOO_MANY:             "开放端口数量过多，请重新输入。",
	ERR_DDNS_USER_NAME_EMPTY:                         "用户名不能为空，请重新输入。",
	ERR_DDNS_PWD_EMPTY:                               "密码不能为空，请重新输入。",
	ERR_DDNS_USER_HAS_SPACE:                          "用户名不能包含空格，请重新输入。",
	ERR_DDNS_PWD_HAS_SPACE:                           "密码不能包含空格，请重新输入。",
	ERR_DDNS_LIST_FULL:                               "DDNS条目已满。",
	ERR_DDNS_LIST_INDEX_OUT_RANGE:                    "条目索引超出范围。",
	ERR_DDNS_ENTRY_BE_DELETE:                         "条目已被删除。",
	ERR_USER_NAME_LENGTH:                             "用户名长度超过15个字符，请重新输入。",
	ERR_USER_PWD_LENGTH:                              "密码长度超过15个字符，请重新输入。",
	ERR_USER_NAME_ERROR:                              "旧用户名错误，请重新输入。",
	ERR_USER_PWD_ERROR:                               "旧密码错误，请重新输入。",
	ERR_USER_PWD_INVALID_CHAR:                        "用户名或密码包含非法字符，请重新输入。",
	ERR_FIRMWARE_FAIL:                                "升级不成功，请检查您上传的文件是否正确。",
	ERR_FIRMWARE_VALID_LENGTH:                        "升级文件长度错误，请检查文件并重试。",
	ERR_FIRMWARE_VALID_FORMAT:                        "升级文件格式错误，请检查文件并重试。",
	ERR_FIRMWARE_VALID_PRODUCTID:                     "升级不成功，请检查升级文件是否和产品类型相匹配。",
	ERR_FIRMWARE_VALID_MDCHECK:                       "升级文件校验失败，请确认文件没有被修改。",
	ERR_FIRMWARE_VALID_VERSION:                       "不能升级到当前的软件版本。",
	ERR_FIRMWARE_VALID_LANGUAGE:                      "不支持的语言选项。",
	ERR_FIRMWARE_VALID_CONTRY:                        "不支持的国家或区域。",
	ERR_FIRMWARE_VALID_OEM:                           "未被允许的OEM厂家。",
	ERR_ARP_REC_IP_EXIST:                             "条目与已分配的静态条目冲突。",
	ERR_ARP_FIXMAP_FULL:                              "静态列表已满，添加失败。",
	ERR_ARP_REC_IP_EXIST_ADD_SUCC:                    "忽略与静态条目冲突条目，成功添加其他条目。",
	ERR_ARP_REC_IP_EXIST_ADD_FAIL:                    "所有欲添加条目均与静态条目冲突，添加失败。",
	ERR_ARP_IP_EXIST_AND_FIXMAP_FULL:                 "静态列表已满。",
	ERR_ARP_FIXMAP_FULL_IGNORE_OTHER_ENTRYS:          "静态列表已添加满，忽略多余条目。",
	ERR_ARP_FIXMAP_MAC_ERR:                           "MAC地址格式不对，请按照“00-00-00-00-00-00”格式输入。",
	ERR_SYS_LOG_SYS_STATUS:                           "有错误发生，请重试！",
	ERR_SYS_LOG_SRV_ID:                               "有错误发生，请重试！",
	ERR_SYS_LOG_SRV_STATUS:                           "有错误发生，请重试！",
	ERR_SYS_LOG_SRV_ADDRESS:                          "有错误发生，请重试！",
	ERR_SYS_LOG_SRV_ADDR_EXIST:                       "有错误发生，请重试！",
	ERR_SYS_LOG_SRV_PORT:                             "有错误发生，请重试！",
	ERR_SYS_LOG_SETTING_EMERGENCY:                    "有错误发生，请重试！",
	ERR_SYS_LOG_SETTING_ALERT:                        "有错误发生，请重试！",
	ERR_SYS_LOG_SETTING_CRITICAL:                     "有错误发生，请重试！",
	ERR_SYS_LOG_SETTING_ERROR:                        "有错误发生，请重试！",
	ERR_SYS_LOG_SETTING_WARNING:                      "有错误发生，请重试！",
	ERR_SYS_LOG_SETTING_NOTICE:                       "有错误发生，请重试！",
	ERR_SYS_LOG_SETTING_INFORMATIONAL:                "有错误发生，请重试！",
	ERR_SYS_LOG_SETTING_DEBUG:                        "有错误发生，请重试！",
	ERR_SYS_LOG_SETTING_EMPTY:                        "有错误发生，请重试！",
	ERR_FIREWALL_SYSLOG_SERVER_INVALID_ID:            "有错误发生，请重试！",
	ERR_FIREWALL_SYSLOG_SERVER_NOT_DEFINED:           "有错误发生，请重试！",
	ERR_FIREWALL_SCREEN_UNKNOWN_DEFENCE:              "有错误发生，请重试！",
	ERR_FIREWALL_SCREEN_SCAN_THRESHOLD:               "有错误发生，请重试！",
	ERR_FIREWALL_SCREEN_DOS_THRESHOLD:                "有错误发生，请重试！",
	ERR_TDDP_UPLOAD_FILE_TOO_LONG:                    "上传文件大小太大。 ",
	ERR_TDDP_UPLOAD_FILE_FORMAT_ERR:                  "文件格式错误。",
	ERR_TDDP_UPLOAD_FILE_NAME_ERR:                    "文件名太长，文件名长度必须小于64字节。",
	ERR_COMMON_ERROR:                                 "有错误发生，请重试！",
	ERR_DST_HOUR:                                     "DST时间错误，请重新输入。",
	ERR_DST_DAY:                                      "DST时间错误，请重新输入。",
	ERR_DST_MONTH:                                    "DST时间错误，请重新输入。",
	ERR_DST_BEGIN_END:                                "开始时间与结束时间不能相同。",
	ERR_TDDP_DOWNLOAD_FILE_TOO_LONG:                  "下载的文件太大。",
	ERR_VS_RECORD_CONFLICT_REMOTE_WEB_PORT:           "虚拟服务器端口与远程Web管理端口冲突。",
	ERR_WLAN_CONFIG_BASE:                             "无线设置错误。",
	ERR_WLAN_CONFIG_SECURITY:                         "无线安全设置错误。",
	ERR_WLAN_CONFIG_KEY:                              "WEP密钥错误。",
	ERR_WLAN_MAC_FILTER_PAGE_NUM_ERROR:               "无线MAC过滤表页不存在，请重试。",
	ERR_WLAN_MAC_FILTER_RECORD_ALREADY_EXIST:         "条目已存在，请重新输入。",
	ERR_WLAN_MAC_FILTER_RECORD_ALREADY_FULL:          "无线MAC地址过滤表已满。",
	ERR_IP_NOT_IN_THE_SAME_SUBNET:                    "IP地址和当前LAN不在同一子网。",
	ERR_WLAN_SSID_LEN:                                "SSID长度错误。",
	ERR_WLAN_REGION:                                  "地区码错误。",
	ERR_WLAN_CHANNEL_WIDTH:                           "频带宽度错误。",
	ERR_WLAN_STATIC_RATE:                             "静态速率错误。",
	ERR_WLAN_MODE:                                    "无线模式错误。",
	ERR_WLAN_BROADCAST:                               "广播错误。",
	ERR_WLAN_MAC_ADDR_INVALID:                        "MAC地址错误，请重新输入。",
	ERR_WLAN_RADIUS_IP_INVALID:                       "Radius服务器IP非法。",
	ERR_WLAN_CHANNEL:                                 "信道错误。",
	ERR_SNTP_SERVER_A:                                "NTP服务器1IP地址错误，请重新输入。",
	ERR_SNTP_SERVER_B:                                "NTP服务器2IP地址错误，请重新输入。",
	ERR_SERVER_IP_ERROR:                              "服务器IP地址错误，请重新输入。",
	ERR_LAN_WAN_CONFLICT:                             "LAN口IP地址和WAN口IP地址不能处于同一子网，请重新输入。",
	ERR_WAN_LINK_MODE_ERROR:                          "WAN口速率/双工模式设置错误。",
	ERR_PARENT_CTRL_FULL:                             "家长控制列表满。",
	ERR_PARENT_CTRL_URLDESC:                          "URL 描述重复。",
	ERR_PARENT_CTRL_SAME_MAC_WITH_PARENT:             "小孩PC的MAC地址和家长控制PC的MAC地址相同。",
	ERR_ACC_CTRL_HOST_FULL:                           "主机列表已满。",
	ERR_ACC_CTRL_TARGET_FULL:                         "访问目标列表已满。",
	ERR_ACC_CTRL_SCHEDULE_FULL:                       "日程计划列表已满。",
	ERR_ACC_CTRL_RULE_FULL:                           "访问控制规则管理列表已满。",
	ERR_ACC_CTRL_SAME_NAME:                           "有重复的列表命名。",
	ERR_ACC_CTRL_REFERED:                             "条目被引用，不能删除。",
	ERR_ACC_CTRL_RULE_CONFLICT:                       "存在重复的条目。",
	ERR_ACC_PARTIAL_DEL:                              "部分条目被引用，不能删除，已删除未被引用的条目。",
	ERR_ACC_DEL_NONE:                                 "所有条目均被引用，不能删除。",
	ERR_FILTER_MAC:                                   "MAC地址输入错误。",
	ERR_ACC_CTRL_HOST_IPSTART:                        "主机起始IP地址输入错误。",
	ERR_ACC_CTRL_HOST_IPEND:                          "主机结束IP地址输入错误。",
	ERR_ACC_CTRL_TARGET_IPSTART:                      "访问目标起始IP地址输入错误。",
	ERR_ACC_CTRL_TARGET_IPEND:                        "访问目标结束IP地址输入错误。",
	ERR_ACC_CTRL_HOST_IPSTART_NOT_IN_THE_SAME_SUBNET: "主机起始IP地址必须与LAN在同一子网。",
	ERR_ACC_CTRL_HOST_IPEND_NOT_IN_THE_SAME_SUBNET:   "主机结束IP地址必须与LAN在同一子网。",
	ERR_LAN_GATEWAY_ERROR:                            "网关地址必须和IP地址在同一个子网内,请重新输入。 ",
	ERR_LAN_GATEWAY_IP:                               "网关错误。",
	ERR_LAN_SAME_IP:                                  "不可以输入LAN IP地址，请重新输入。",
	ERR_LAN_SAME_MAC:                                 "不可以输入LAN MAC地址，请重新输入。",
	ERR_DMZ_IP:                                       "IP地址错误，请重新输入。",
	ERR_PORT_HAS_BEEN_USED:                           "端口号已被占用，请重新输入。",
}
