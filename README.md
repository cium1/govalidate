# govalidate
golang 数据验证包

```
	data := M{
		"username": "test",
		"password": 12345,
		"status":   false,
	}

	v := govalidate.New()

	v.AddColumn("username", "登录账户").Required("登录账户是必须的").AlphaNumeric("登录账户只能是字母和数值").Length(4, "登录账户字符长度应为4")
	v.AddColumn("password", "登录密码").Required("登录密码是必须的").AlphaDash("登录账户只能是字母数值和常规符号")
	v.AddColumn("status", "状态").Required("状态是必须的").Bool("状态是能是布尔型")
	v.AddColumn("other", "其他")

	if !v.Validate(data) {
	    fmt.Println(v.Error().field, v.Error().fieldAlias, v.Error().fieldData, v.Error().rule, v.Error().ruleArgs, v.Error().errorMessage)
	    return
	}

    fmt.Println(v.GetData())
```