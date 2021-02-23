package SignUp

import (
	"database/sql"
	"myWebCli/db/mysql"
	"myWebCli/models"
	"myWebCli/utils/md5"
	"myWebCli/utils/snowflake"
)

func Register(user *models.RegisterForm) (err error) {
	sqlStr := `Select count(user_id) from user where username=?`
	var num int64
	err = mysql.DB.Get(&num, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if num > 0 {
		return mysql.ErrorUserExit
	}
	// 生成uid
	userId := snowflake.GenID()
	// 生成加密密码
	password := md5.EncryptPassword(user.Pwd)
	// 把用户插入数据库
	sqlStr = `INSERT INTO user(user_id, username, password, email, gender) values (?,?,?,?,?)`
	_, err = mysql.DB.Exec(sqlStr, userId, user.UserName, password, user.Email, user.Gender)
	return
}
