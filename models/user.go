package models

import (
	"encoding/json"
	"errors"
)

type LoginForm struct {
	UserId   int64  `json:"user_id" db:"user_id"`
	UserName string `json:"userName" db:"username"`
	Pwd      string `json:"pwd" db:"password"`
}

type RegisterForm struct {
	UserId   int64  `json:"user_id" db:"user_id"`
	UserName string `json:"userName" db:"username"`
	Pwd      string `json:"pwd" db:"password"`
	RePwd    string `json:"re_pwd"`
	Email    string `json:"email" db:"email"`
	Gender   int    `json:"gender" db:"gender"`
}

func (s *RegisterForm) UnmarshalJson(data []byte) (err error) {
	required := RegisterForm{}
	if err = json.Unmarshal(data, &required); err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段userName")
	} else if len(required.Pwd) == 0 {
		err = errors.New("缺少必填字段pwd")
	} else if required.Pwd != required.RePwd {
		err = errors.New("两次密码输入不一致")
	} else {
		s.UserId = required.UserId
		s.UserName = required.UserName
		s.Pwd = required.Pwd
		s.Email = required.Email
		s.Gender = required.Gender
	}
	return
}
