package dao

import (
	_ "github.com/mysql"
	"database/sql"
	"Main/model"
)


func init(){
	db, e := sql.Open("mysql", "root:123abcd@tcp(47.104.244.69:3306)/zhenhun?charset=utf8")
	checkerr(e)
	ping := db.Ping()
	checkerr(ping)
	stmt,err:=db.Prepare("INSERT INTO people VALUES (?,?,?,?,?,?,?,?,?,?)")
	checkerr(err)
	Stmt=stmt
}

func InsertUser(user model.UserMes) (bool,error){
	result,err:=Stmt.Exec(user.Name,user.Gender,user.Age,
				user.Height,user.Income,user.Marriage,
					user.Education,user.Location,user.Infor,
					user.Require)

	if err!=nil {
		return false,err
	}

	effect,err:=result.RowsAffected()
	if err!=nil {
		return false,err
	}

	if effect<1 {
		return false,nil
	}

	return true,nil

}

