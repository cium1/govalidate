package govalidate

import (
	"testing"
)

var (
	v *Validate
)

func TestValidate(t *testing.T) {

	t.Parallel()

	data := M{
		"username": "test",
		"password": 12345,
		"status":   false,
	}

	v = New()
	v.AddColumn("username", "").Required("").AlphaNumeric("").Length(4, "")
	v.AddColumn("password", "").Required("").AlphaDash("")
	v.AddColumn("status", "").Required("").Bool("")
	v.AddColumn("other", "")

	if v.Validate(data) != true {
		t.Error(v.Error().field, v.Error().fieldAlias, v.Error().fieldData, v.Error().rule, v.Error().ruleArgs, v.Error().errorMessage)
	}

}

func TestRequired(t *testing.T) {

	t.Parallel()

	data := M{"user": "test"}
	v = New()
	v.AddColumn("username", "username").Required("")
	result := v.Validate(data)

	if result != false {
		t.Errorf("Expected Required(%q) to be %v, got %v,", "username", false, result)
	}
}

func TestAlpha(t *testing.T) {
	t.Parallel()

	var tests = []*struct {
		value    M
		expected bool
	}{
		{M{"t1": "\n"}, false},
		{M{"t1": "\r"}, false},
		{M{"t1": "Ⅸ"}, false},
		{M{"t1": ""}, false},
		{M{"t1": "   fooo   "}, false},
		{M{"t1": "abc!!!"}, false},
		{M{"t1": "abc1"}, false},
		{M{"t1": "abc〩"}, false},
		{M{"t1": "abc"}, true},
		{M{"t1": "소주"}, false},
		{M{"t1": "ABC"}, true},
		{M{"t1": "FoObAr"}, true},
		{M{"t1": "소aBC"}, false},
		{M{"t1": "소"}, false},
		{M{"t1": "달기&Co."}, false},
		{M{"t1": "〩Hours"}, false},
		{M{"t1": "\ufff0"}, false},
		{M{"t1": "\u0070"}, true},
		{M{"t1": "\u0026"}, false},
		{M{"t1": "\u0030"}, false},
		{M{"t1": "123"}, false},
		{M{"t1": "0123"}, false},
		{M{"t1": "-00123"}, false},
		{M{"t1": "0"}, false},
		{M{"t1": "-0"}, false},
		{M{"t1": "123.123"}, false},
		{M{"t1": 123}, false},
		{M{"t1": 0123}, false},
		{M{"t1": -00123}, false},
		{M{"t1": 0}, false},
		{M{"t1": -0}, false},
		{M{"t1": 123.123}, false},
		{M{"t1": " "}, false},
		{M{"t1": "."}, false},
		{M{"t1": "-1¾"}, false},
		{M{"t1": "1¾"}, false},
		{M{"t1": "〥〩"}, false},
		{M{"t1": "모자"}, false},
		{M{"t1": "ix"}, true},
		{M{"t1": "۳۵۶۰"}, false},
		{M{"t1": "1--"}, false},
		{M{"t1": "1-1"}, false},
		{M{"t1": "-"}, false},
		{M{"t1": "--"}, false},
		{M{"t1": "1++"}, false},
		{M{"t1": "1+1"}, false},
		{M{"t1": "+"}, false},
		{M{"t1": "++"}, false},
		{M{"t1": "+1"}, false},
	}
	v = New()
	for _, test := range tests {
		v.AddColumn("t1", "").Alpha("")
		result := v.Validate(test.value)
		if result != test.expected {
			t.Error(test.value, test.expected, result)
		}
	}
}

func TestAlphaNumeric(t *testing.T) {

	t.Parallel()

	var tests = []*struct {
		value    M
		expected bool
	}{
		{M{"t1": "\n"}, false},
		{M{"t1": "\r"}, false},
		{M{"t1": "Ⅸ"}, false},
		{M{"t1": ""}, false},
		{M{"t1": "   fooo   "}, false},
		{M{"t1": "abc!!!"}, false},
		{M{"t1": "abc1"}, true},
		{M{"t1": "abc〩"}, false},
		{M{"t1": "abc"}, true},
		{M{"t1": "소주"}, false},
		{M{"t1": "ABC"}, true},
		{M{"t1": "FoObAr"}, true},
		{M{"t1": "소aBC"}, false},
		{M{"t1": "소"}, false},
		{M{"t1": "달기&Co."}, false},
		{M{"t1": "〩Hours"}, false},
		{M{"t1": "\ufff0"}, false},
		{M{"t1": "\u0070"}, true},
		{M{"t1": "\u0026"}, false},
		{M{"t1": "\u0030"}, true},
		{M{"t1": "123"}, true},
		{M{"t1": "0123"}, true},
		{M{"t1": "-00123"}, false},
		{M{"t1": "0"}, true},
		{M{"t1": "-0"}, false},
		{M{"t1": "123.123"}, false},
		{M{"t1": 123}, true},
		{M{"t1": 0123}, true},
		{M{"t1": -00123}, false},
		{M{"t1": 0}, true},
		{M{"t1": -0}, true},
		{M{"t1": 123.123}, false},
		{M{"t1": " "}, false},
		{M{"t1": "."}, false},
		{M{"t1": "-1¾"}, false},
		{M{"t1": "1¾"}, false},
		{M{"t1": "〥〩"}, false},
		{M{"t1": "모자"}, false},
		{M{"t1": "ix"}, true},
		{M{"t1": "۳۵۶۰"}, false},
		{M{"t1": "1--"}, false},
		{M{"t1": "1-1"}, false},
		{M{"t1": "-"}, false},
		{M{"t1": "--"}, false},
		{M{"t1": "1++"}, false},
		{M{"t1": "1+1"}, false},
		{M{"t1": "+"}, false},
		{M{"t1": "++"}, false},
		{M{"t1": "+1"}, false},
	}
	v = New()
	for _, test := range tests {
		v.AddColumn("t1", "").AlphaNumeric("")
		result := v.Validate(test.value)
		if result != test.expected {
			t.Error(test.value, test.expected, result)
		}
	}
}

func TestAlphaDash(t *testing.T) {
	t.Parallel()

	var tests = []*struct {
		value    M
		expected bool
	}{
		{M{"t1": "\n"}, false},
		{M{"t1": "\r"}, false},
		{M{"t1": "Ⅸ"}, false},
		{M{"t1": ""}, false},
		{M{"t1": "   fooo   "}, false},
		{M{"t1": "abc!!!"}, false},
		{M{"t1": "abc1"}, true},
		{M{"t1": "abc〩"}, false},
		{M{"t1": "abc"}, true},
		{M{"t1": "소주"}, false},
		{M{"t1": "ABC"}, true},
		{M{"t1": "FoObAr"}, true},
		{M{"t1": "소aBC"}, false},
		{M{"t1": "소"}, false},
		{M{"t1": "달기&Co."}, false},
		{M{"t1": "〩Hours"}, false},
		{M{"t1": "\ufff0"}, false},
		{M{"t1": "\u0070"}, true},
		{M{"t1": "\u0026"}, false},
		{M{"t1": "\u0030"}, true},
		{M{"t1": "123"}, true},
		{M{"t1": "0123"}, true},
		{M{"t1": "-00123"}, true},
		{M{"t1": "0"}, true},
		{M{"t1": "-0"}, true},
		{M{"t1": "123.123"}, false},
		{M{"t1": 123}, true},
		{M{"t1": 0123}, true},
		{M{"t1": -00123}, true},
		{M{"t1": 0}, true},
		{M{"t1": -0}, true},
		{M{"t1": 123.123}, false},
		{M{"t1": " "}, false},
		{M{"t1": "."}, false},
		{M{"t1": "-1¾"}, false},
		{M{"t1": "1¾"}, false},
		{M{"t1": "〥〩"}, false},
		{M{"t1": "모자"}, false},
		{M{"t1": "ix"}, true},
		{M{"t1": "۳۵۶۰"}, false},
		{M{"t1": "1--"}, true},
		{M{"t1": "1-1"}, true},
		{M{"t1": "-"}, true},
		{M{"t1": "--"}, true},
		{M{"t1": "1++"}, false},
		{M{"t1": "1+1"}, false},
		{M{"t1": "+"}, false},
		{M{"t1": "++"}, false},
		{M{"t1": "+1"}, false},
	}
	v = New()
	for _, test := range tests {
		v.AddColumn("t1", "").AlphaDash("")
		result := v.Validate(test.value)
		if result != test.expected {
			t.Error(test.value, test.expected, result)
		}
	}
}

func TestBetween(t *testing.T) {

	t.Parallel()

	var tests = []*struct {
		value    M
		min      int64
		max      int64
		expected bool
	}{
		{M{"t1": 2}, -1, 1, false},
		{M{"t1": -2}, -1, 1, false},
		{M{"t1": 0}, -1, 1, true},
		{M{"t1": 1}, -1, 1, true},
		{M{"t1": -1}, -1, 1, true},
		{M{"t1": "2"}, -1, 1, false},
		{M{"t1": "-2"}, -1, 1, false},
		{M{"t1": "0"}, -1, 1, true},
		{M{"t1": "1"}, -1, 1, true},
		{M{"t1": "-1"}, -1, 1, true},
		{M{"t1": "aa"}, -1, 1, false},
	}
	v = New()
	for _, test := range tests {
		v.AddColumn("t1", "").Between(test.min, test.max, "")
		result := v.Validate(test.value)
		if result != test.expected {
			t.Error(test.value, test.min, test.max, test.expected, result)
		}
	}
}
