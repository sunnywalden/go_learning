package db_demo

import (
	"errors"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func SqlRun() ([]string, error) {
	db, _ := sql.Open("mysql", "root:walden0517@tcp(127.0.0.1:3336)/easyst")
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	//连接数据库查询
	//for i := 0; i < 100; i++ {
	//	go func(i int) {
			mSql := "show tables"
			rows, _ := db.Query(mSql)
			res, err := rows.Columns()
			if err != nil {
				return nil, err
			} else {
				fmt.Printf("%s", res)
				return res, nil
			}
			defer rows.Close()
			return nil, errors.New("no res")
			//rows.Close() //这里如果不释放连接到池里，执行5次后其他并发就会阻塞
			//fmt.Println("第 ", i)
		//}(i)

	//}
	//
	//for {
	//	time.Sleep(time.Second)
	//}
}
