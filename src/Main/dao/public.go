package dao

import "database/sql"

var(
	Stmt *sql.Stmt
)

const(
	INSERT="DbDao.Insert"
)

func checkerr(err error){

	if err != nil {
		panic(err)
	}
}
