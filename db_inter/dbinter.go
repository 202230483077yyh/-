package db_inter

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func start_DB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// 获取菜谱表格某个字段的所有值
func Get_caipu_conlum(ziduan []string, m map[string]string) []string {
	str := make([]string, 0)
	db, err := start_DB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	_, err = db.Exec(`use caipu`)
	if err != nil {
		panic(err.Error())
	}
	ttmp := "select "
	for _, data := range ziduan {
		ttmp += data
		ttmp += " "
	}
	ttmp += "from dish "
	if len(m) != 0 {
		ttmp += "where "
	}
	for key, value := range m {
		ttmp += (key + " = " + value + ",")
	}
	ttmp = ttmp[:len(ttmp)-1]
	fmt.Println(ttmp)

	rows, err1 := db.Query(ttmp)
	if err1 != nil {
		panic(err1.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var tmp string
		err = rows.Scan(&tmp)
		if err != nil {
			panic(err.Error())
		}
		str = append(str, tmp)
	}
	return str
}
