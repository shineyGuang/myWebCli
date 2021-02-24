package login

import (
	"database/sql"
	"myWebCli/db/mysql"
	"myWebCli/models"
	"myWebCli/utils/md5"
)

func Login(user *models.LoginForm) (err error) {
	originPwd := user.Pwd // 记录原始密码
	sqlStr := `select user_id, username, password from user where username=?`
	err = mysql.DB.Get(user, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	if md5.EncryptPassword(originPwd) == user.Pwd {
		return
	} else {
		return mysql.ErrorPasswordWrong
	}
}
