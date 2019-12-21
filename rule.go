package govalidate

// Rule struct
type Rule struct {
	item []item
}

type item struct {
	name       string
	message    string
	args       []interface{}
	verifyFunc Func
}

// Func validate func
type Func func(data map[string]interface{}, column string, args ...interface{}) bool

// Required 必须存在值
func (r *Rule) Required(message string) *Rule {

	r.item = append(r.item, item{
		name:       "required",
		message:    message,
		verifyFunc: (&Validate{}).required,
	})

	return r
}

// Bool 布尔型
func (r *Rule) Bool(message string) *Rule {

	r.item = append(r.item, item{
		name:       "bool",
		message:    message,
		verifyFunc: (&Validate{}).bool,
	})

	return r
}

// Alpha 是否是字母 即[a-zA-Z]
func (r *Rule) Alpha(message string) *Rule {

	r.item = append(r.item, item{
		name:       "alpha",
		message:    message,
		verifyFunc: (&Validate{}).alpha,
	})

	return r
}

// AlphaNumeric 是否是字母和数字组成 即[a-zA-Z0-9]
func (r *Rule) AlphaNumeric(message string) *Rule {

	r.item = append(r.item, item{
		name:       "alphaNumeric",
		message:    message,
		verifyFunc: (&Validate{}).alphaNumeric,
	})

	return r
}

// AlphaDash 是否是字母和数字下划线破折号组成 即[a-zA-Z0-9\-\_]
func (r *Rule) AlphaDash(message string) *Rule {

	r.item = append(r.item, item{
		name:       "alphaDash",
		message:    message,
		verifyFunc: (&Validate{}).alphaDash,
	})

	return r
}

// Between 是否在 min max 之间
func (r *Rule) Between(min int64, max int64, message string) *Rule {

	r.item = append(r.item, item{
		name:       "between",
		message:    message,
		args:       []interface{}{min, max},
		verifyFunc: (&Validate{}).between,
	})

	return r
}

// Float 是否为小数
func (r *Rule) Float(message string) *Rule {

	r.item = append(r.item, item{
		name:       "float",
		message:    message,
		verifyFunc: (&Validate{}).float,
	})

	return r
}

// TimeBefore 是否在某时间之前
func (r *Rule) TimeBefore(t interface{}, message string) *Rule {

	r.item = append(r.item, item{
		name:       "dateBefore",
		message:    message,
		args:       []interface{}{t},
		verifyFunc: (&Validate{}).timeBefore,
	})

	return r
}

// TimeAfter 是否在某时间之后
func (r *Rule) TimeAfter(t interface{}, message string) *Rule {

	r.item = append(r.item, item{
		name:       "dateAfter",
		message:    message,
		args:       []interface{}{t},
		verifyFunc: (&Validate{}).timeAfter,
	})

	return r
}

// Equal 验证值是否相等
func (r *Rule) Equal(val interface{}, message string) *Rule {

	r.item = append(r.item, item{
		name:       "equal",
		message:    message,
		args:       []interface{}{val},
		verifyFunc: (&Validate{}).equal,
	})

	return r
}

// Different 验证值是否不相等
func (r *Rule) Different(val interface{}, message string) *Rule {

	r.item = append(r.item, item{
		name:       "different",
		message:    message,
		args:       []interface{}{val},
		verifyFunc: (&Validate{}).different,
	})

	return r
}

// EqualWithColumn 验证节点值是否相等
func (r *Rule) EqualWithColumn(node string, message string) *Rule {

	r.item = append(r.item, item{
		name:       "equalWithColumn",
		message:    message,
		args:       []interface{}{node},
		verifyFunc: (&Validate{}).equalWithColumn,
	})

	return r
}

// DifferentWithColumn 验证节点值是否不相等
func (r *Rule) DifferentWithColumn(node string, message string) *Rule {

	r.item = append(r.item, item{
		name:       "differentWithColumn",
		message:    message,
		args:       []interface{}{node},
		verifyFunc: (&Validate{}).differentWithColumn,
	})

	return r
}

// In 是否包含
func (r *Rule) In(s []interface{}, message string) *Rule {

	r.item = append(r.item, item{
		name:       "in",
		message:    message,
		args:       s,
		verifyFunc: (&Validate{}).in,
	})

	return r
}

// Integer integer 数据
func (r *Rule) Integer(message string) *Rule {

	r.item = append(r.item, item{
		name:       "integer",
		message:    message,
		args:       nil,
		verifyFunc: (&Validate{}).integer,
	})

	return r
}

// IP 有效
func (r *Rule) IP(message string) *Rule {

	r.item = append(r.item, item{
		name:       "ip",
		message:    message,
		args:       nil,
		verifyFunc: (&Validate{}).ip,
	})

	return r
}

// NotIn 不包含
func (r *Rule) NotIn(s []interface{}, message string) *Rule {

	r.item = append(r.item, item{
		name:       "notIn",
		message:    message,
		args:       s,
		verifyFunc: (&Validate{}).notIn,
	})

	return r
}

// Length 验证长度
func (r *Rule) Length(len int64, message string) *Rule {

	r.item = append(r.item, item{
		name:       "length",
		message:    message,
		args:       []interface{}{len},
		verifyFunc: (&Validate{}).length,
	})

	return r
}

// LengthMax 最大长度
func (r *Rule) LengthMax(len int64, message string) *Rule {

	r.item = append(r.item, item{
		name:       "lengthMax",
		message:    message,
		args:       []interface{}{len},
		verifyFunc: (&Validate{}).lengthMax,
	})

	return r
}

// LengthMin 最小长度
func (r *Rule) LengthMin(len int64, message string) *Rule {

	r.item = append(r.item, item{
		name:       "lengthMin",
		message:    message,
		args:       []interface{}{len},
		verifyFunc: (&Validate{}).lengthMin,
	})

	return r
}

// BetweenLen 长度范围
func (r *Rule) BetweenLen(minLen int64, maxLen int64, message string) *Rule {

	r.item = append(r.item, item{
		name:       "betweenLen",
		message:    message,
		args:       []interface{}{minLen, maxLen},
		verifyFunc: (&Validate{}).betweenLen,
	})

	return r
}

// Max 最大值
func (r *Rule) Max(max int64, message string) *Rule {

	r.item = append(r.item, item{
		name:       "max",
		message:    message,
		args:       []interface{}{max},
		verifyFunc: (&Validate{}).max,
	})

	return r
}

// Min 最小值
func (r *Rule) Min(max int64, message string) *Rule {

	r.item = append(r.item, item{
		name:       "min",
		message:    message,
		args:       []interface{}{max},
		verifyFunc: (&Validate{}).min,
	})

	return r
}

// Money 有效货币金额
func (r *Rule) Money(message string) *Rule {

	r.item = append(r.item, item{
		name:       "money",
		message:    message,
		args:       nil,
		verifyFunc: (&Validate{}).money,
	})

	return r
}

// Regexp 自定义正则
func (r *Rule) Regexp(regexp string, message string) *Rule {

	r.item = append(r.item, item{
		name:       "regexp",
		message:    message,
		args:       []interface{}{regexp},
		verifyFunc: (&Validate{}).regexp,
	})

	return r
}

// Username 合法用户名
func (r *Rule) Username(message string) *Rule {

	r.item = append(r.item, item{
		name:       "username",
		message:    message,
		args:       nil,
		verifyFunc: (&Validate{}).username,
	})

	return r
}

// Host 有效Host地址
func (r *Rule) Host(message string) *Rule {

	r.item = append(r.item, item{
		name:       "host",
		message:    message,
		args:       nil,
		verifyFunc: (&Validate{}).host,
	})

	return r
}

// Email 电子邮箱地址
func (r *Rule) Email(message string) *Rule {

	r.item = append(r.item, item{
		name:       "email",
		message:    message,
		args:       nil,
		verifyFunc: (&Validate{}).email,
	})

	return r
}

// CreditCard 银行卡号
func (r *Rule) CreditCard(message string) *Rule {

	r.item = append(r.item, item{
		name:       "creditCard",
		message:    message,
		args:       nil,
		verifyFunc: (&Validate{}).creditCard,
	})

	return r
}

// Numeric 数值
func (r *Rule) Numeric(message string) *Rule {

	r.item = append(r.item, item{
		name:       "numeric",
		message:    message,
		args:       nil,
		verifyFunc: (&Validate{}).numeric,
	})

	return r
}

// HexColor Hex颜色
func (r *Rule) HexColor(message string) *Rule {

	r.item = append(r.item, item{
		name:       "hexColor",
		message:    message,
		args:       nil,
		verifyFunc: (&Validate{}).hexColor,
	})

	return r
}

// RgbColor RGB颜色
func (r *Rule) RgbColor(message string) *Rule {

	r.item = append(r.item, item{
		name:       "rgbColor",
		message:    message,
		args:       nil,
		verifyFunc: (&Validate{}).rgbColor,
	})

	return r
}

// ASCII Ascii值
func (r *Rule) ASCII(message string) *Rule {

	r.item = append(r.item, item{
		name:       "ascii",
		message:    message,
		args:       nil,
		verifyFunc: (&Validate{}).ascii,
	})

	return r
}

// Base64 base64值
func (r *Rule) Base64(message string) *Rule {

	r.item = append(r.item, item{
		name:       "base64",
		message:    message,
		args:       nil,
		verifyFunc: (&Validate{}).base64,
	})

	return r
}

// DNSName dns名称
func (r *Rule) DNSName(message string) *Rule {

	r.item = append(r.item, item{
		name:       "dnsName",
		message:    message,
		args:       nil,
		verifyFunc: (&Validate{}).dnsName,
	})

	return r
}

// URL url地址
func (r *Rule) URL(message string) *Rule {

	r.item = append(r.item, item{
		name:       "url",
		message:    message,
		args:       nil,
		verifyFunc: (&Validate{}).url,
	})

	return r
}
