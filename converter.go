package govalidate

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// ToString convert the input to a string.
func ToString(obj interface{}) string {
	res := fmt.Sprintf("%v", obj)
	return string(res)
}

// ToJSON convert the input to a valid JSON string
func ToJSON(obj interface{}) (string, error) {
	res, err := json.Marshal(obj)
	if err != nil {
		res = []byte("")
	}
	return string(res), err
}

//ToFloat2 convert the input string to a float, or 0.0 if the input is not a float.
func ToFloat2(str string) (float64, error) {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		res = 0.0
	}
	return res, err
}

// ToFloat convert the input string to a float, or 0.0 if the input is not a float.
func ToFloat(value interface{}) (res float64, err error) {
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		res = float64(val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		res = float64(val.Uint())
	case reflect.Float32, reflect.Float64:
		res = val.Float()
	case reflect.String:
		res, err = strconv.ParseFloat(val.String(), 64)
		if err != nil {
			res = 0.0
		}
	default:
		err = fmt.Errorf("conversion failed, type is %T", value)
		res = 0.0
	}

	return
}

// ToInt convert the input string or any int type to an integer type 64, or 0 if the input is not an integer.
func ToInt(value interface{}) (res int64, err error) {
	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		res = val.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		res = int64(val.Uint())
	case reflect.String:
		v, err := strconv.Atoi(val.String())
		if err != nil {
			res = 0
		}
		res = int64(v)

	default:
		err = fmt.Errorf("conversion failed, type is %T", value)
		res = 0
	}

	return
}

// ToBoolean convert the input string to a boolean.
func ToBoolean(str string) (bool, error) {
	return strconv.ParseBool(str)
}

// ToTime convert the input interface to a time.
func ToTime(value interface{}) (t time.Time, err error) {

	switch value.(type) {
	case time.Time:
		t = value.(time.Time)
		err = nil
	case string:
		t, err = time.Parse(time.RFC3339, value.(string))

	default:
		val, err := ToInt(value)
		if err != nil {
			t = time.Now()
		}
		t = time.Unix(val, 0)
	}

	return
}
