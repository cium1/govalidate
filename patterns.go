package govalidate

import "regexp"

const (
	// Username 用户名
	Username string = "^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~.-]+$"
	// Host Host地址
	Host string = "^[^\\s]+\\.[^\\s]+$"
	// Email 邮件地址
	Email string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	// CreditCard 银行卡号
	CreditCard string = "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$"
	// Alpha 字母
	Alpha string = "^[a-zA-Z]+$"
	// AlphaNumeric 字母+数值
	AlphaNumeric string = "^[a-zA-Z0-9]+$"
	// AlphaDash 字母+数值+常规符号
	AlphaDash string = "^[a-zA-Z0-9\\-\\_]+$"
	// Numeric 数值
	Numeric string = "^[0-9]+$"
	// Int Int类型
	Int string = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
	// Float 浮点数类型
	Float string = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
	// Money 货币金额
	Money string = "^(0|[1-9]+[0-9]*)(.[0-9]{1,2})?$/"
	// HexColor HEX 颜色
	HexColor string = "^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	// RgbColor RGB 颜色
	RgbColor string = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	// ASCII ascii值
	ASCII string = "^[\x00-\x7F]+$"
	// Base64 值
	Base64 string = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	// DNSName dns名称
	DNSName string = `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`
	// IP 地址
	IP string = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	// URLSchema Url 头
	URLSchema string = `((ftp|tcp|udp|wss?|https?):\/\/)`
	// URLUsername Url 名称
	URLUsername string = `(\S+(:\S*)?@)`
	//URLPath Url路径
	URLPath string = `((\/|\?|#)[^\s]*)`
	// URLPort Url端口
	URLPort string = `(:(\d{1,5}))`
	// URLIP url ip
	URLIP string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))`
	// URLSubDomain url子域名
	URLSubDomain string = `((www\.)|([a-zA-Z0-9]+([-_\.]?[a-zA-Z0-9])*[a-zA-Z0-9]\.[a-zA-Z0-9]+))`
	// URL 地址
	URL string = `^` + URLSchema + `?` + URLUsername + `?` + `((` + URLIP + `|(\[` + IP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + URLSubDomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + URLPort + `?` + URLPath + `?$`
)

var (
	rxpUsername     = regexp.MustCompile(Username)
	rxpHost         = regexp.MustCompile(Host)
	rxpEmail        = regexp.MustCompile(Email)
	rxpCreditCard   = regexp.MustCompile(CreditCard)
	rxpAlpha        = regexp.MustCompile(Alpha)
	rxpAlphaNumeric = regexp.MustCompile(AlphaNumeric)
	rxpAlphaDash    = regexp.MustCompile(AlphaDash)
	rxpNumeric      = regexp.MustCompile(Numeric)
	rxpInt          = regexp.MustCompile(Int)
	rxpFloat        = regexp.MustCompile(Float)
	rxpMoney        = regexp.MustCompile(Money)
	rxpHexColor     = regexp.MustCompile(HexColor)
	rxpRgbColor     = regexp.MustCompile(RgbColor)
	rxpASCII        = regexp.MustCompile(ASCII)
	rxpBase64       = regexp.MustCompile(Base64)
	rxpDNSName      = regexp.MustCompile(DNSName)
	rxpIP           = regexp.MustCompile(IP)
	rxpURL          = regexp.MustCompile(URL)
)
