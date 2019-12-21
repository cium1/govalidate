package govalidate

import (
	"regexp"
)

// M is data map
type M map[string]interface{}

// Validate struct
type Validate struct {
	columns []column
	data    M
	error   *Error
}

type column struct {
	name  string
	alias string
	rule  *Rule
}

// New start run
func New() *Validate {
	return new(Validate)
}

// AddColumn add column
func (v *Validate) AddColumn(name string, alias string) *Rule {

	if !v.insetColumn(name) {
		v.columns = append(v.columns, column{
			name:  name,
			alias: alias,
			rule:  &Rule{},
		})
	}

	return v.getColumn(name).rule
}

// Validate is map data validate
func (v *Validate) Validate(data map[string]interface{}) bool {

	v.data = make(map[string]interface{})

	for _, column := range v.columns {

		rules := column.rule

		for _, item := range rules.item {

			if !item.verifyFunc(data, column.name, item.args...) {
				v.error = new(Error)
				v.error.field = column.name
				v.error.fieldAlias = column.alias
				v.error.fieldData = data[column.name]
				v.error.rule = item.name
				v.error.ruleArgs = item.args
				v.error.errorMessage = item.message
				return false
			}
		}

		v.data[column.name] = data[column.name]
	}
	return true
}

func (v *Validate) Error() *Error {
	return v.error
}

// GetData get validate data
func (v *Validate) GetData() map[string]interface{} {
	return v.data
}

func (v *Validate) insetColumn(name string) bool {
	for _, column := range v.columns {
		if column.name == name {
			return true
		}
	}
	return false
}

func (v *Validate) getColumn(name string) column {
	for _, column := range v.columns {
		if column.name == name {
			return column
		}
	}
	return column{}
}

func (v *Validate) required(data map[string]interface{}, column string, args ...interface{}) bool {
	_, ok := data[column]
	return ok
}

func (v *Validate) alpha(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpAlpha.MatchString(ToString(value))
}

func (v *Validate) alphaNumeric(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpAlphaNumeric.MatchString(ToString(value))
}

func (v *Validate) alphaDash(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpAlphaDash.MatchString(ToString(value))
}

func (v *Validate) between(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	if len(args) < 2 {
		return false
	}

	var (
		val float64
		min float64
		max float64
		err error
	)

	if val, err = ToFloat(value); err != nil {
		return false
	}

	if min, err = ToFloat(args[0]); err != nil {
		return false
	}

	if max, err = ToFloat(args[1]); err != nil {
		return false
	}

	return val >= min && val <= max
}

func (v *Validate) bool(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	_, err := ToBoolean(ToString(value))
	if err != nil {
		return false
	}

	return true
}

func (v *Validate) float(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpFloat.MatchString(ToString(value))
}

func (v *Validate) timeBefore(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	this, err := ToTime(value)
	if err != nil {
		return false
	}

	if len(args) < 1 {
		return false
	}

	refer, err := ToTime(args[0])
	if err != nil {
		return false
	}

	return this.Before(refer)
}

func (v *Validate) timeAfter(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	this, err := ToTime(value)
	if err != nil {
		return false
	}

	if len(args) < 1 {
		return false
	}

	refer, err := ToTime(args[0])
	if err != nil {
		return false
	}

	return this.After(refer)
}

func (v *Validate) equal(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	if len(args) < 1 {
		return false
	}

	return ToString(value) == ToString(args[0])
}

func (v *Validate) different(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	if len(args) < 1 {
		return false
	}

	return ToString(value) != ToString(args[0])
}

func (v *Validate) equalWithColumn(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	if len(args) < 3 {
		return false
	}

	return ToString(value) == ToString(args[2])
}

func (v *Validate) differentWithColumn(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	if len(args) < 3 {
		return false
	}

	return ToString(value) != ToString(args[2])
}

func (v *Validate) in(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	if len(args) < 1 {
		return false
	}

	val := ToString(value)

	for _, v := range args {
		if val == ToString(v) {
			return true
		}
	}

	return false
}

func (v *Validate) integer(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpInt.MatchString(ToString(value))
}

func (v *Validate) ip(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpIP.MatchString(ToString(value))
}

func (v *Validate) notIn(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	if len(args) < 1 {
		return true
	}

	val := ToString(value)

	for _, v := range args {
		if val == ToString(v) {
			return false
		}
	}

	return true
}

func (v *Validate) length(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	if len(args) < 1 {
		return false
	}

	length, err := ToInt(args[0])
	if err != nil {
		return false
	}

	return int64(len([]rune(ToString(value)))) == length
}

func (v *Validate) lengthMax(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	if len(args) < 1 {
		return false
	}

	length, err := ToInt(args[0])
	if err != nil {
		return false
	}

	return int64(len([]rune(ToString(value)))) <= length
}

func (v *Validate) lengthMin(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	if len(args) < 1 {
		return false
	}

	length, err := ToInt(args[0])
	if err != nil {
		return false
	}

	return int64(len([]rune(ToString(value)))) >= length
}

func (v *Validate) betweenLen(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	if len(args) < 2 {
		return false
	}

	startLen, err := ToInt(args[0])
	if err != nil {
		return false
	}

	endLen, err := ToInt(args[1])
	if err != nil {
		return false
	}

	valLen := int64(len([]rune(ToString(value))))

	return valLen >= startLen && valLen <= endLen
}

func (v *Validate) max(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	if len(args) < 1 {
		return false
	}

	val, err := ToFloat(value)
	if err != nil {
		return false
	}

	max, err := ToFloat(args[0])
	if err != nil {
		return false
	}

	return val <= max
}

func (v *Validate) min(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	if len(args) < 1 {
		return false
	}

	val, err := ToFloat(value)
	if err != nil {
		return false
	}

	min, err := ToFloat(args[0])
	if err != nil {
		return false
	}

	return val >= min
}

func (v *Validate) money(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpMoney.MatchString(ToString(value))
}

func (v *Validate) regexp(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	if len(args) < 1 {
		return false
	}

	rxp := ToString(args[1])

	return regexp.MustCompile(rxp).MatchString(ToString(value))
}

func (v *Validate) username(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpUsername.MatchString(ToString(value))
}

func (v *Validate) host(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpHost.MatchString(ToString(value))
}

func (v *Validate) email(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpEmail.MatchString(ToString(value))
}

func (v *Validate) creditCard(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpCreditCard.MatchString(ToString(value))
}

func (v *Validate) numeric(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpNumeric.MatchString(ToString(value))
}

func (v *Validate) hexColor(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpHexColor.MatchString(ToString(value))
}

func (v *Validate) rgbColor(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpRgbColor.MatchString(ToString(value))
}

func (v *Validate) ascii(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpASCII.MatchString(ToString(value))
}

func (v *Validate) base64(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpBase64.MatchString(ToString(value))
}

func (v *Validate) dnsName(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpDNSName.MatchString(ToString(value))
}

func (v *Validate) url(data map[string]interface{}, column string, args ...interface{}) bool {

	value, ok := data[column]
	if !ok {
		return true
	}

	return rxpURL.MatchString(ToString(value))
}
