package multichannel_timeout

import "C"
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 引入包，不使用，使其调用init函数注册mysql
)

type SqlInfo struct {
	MysqlHost string
	MysqlPort int
	MysqlUser string
	MysqlPassword string
	MysqlDB string
	CharSet string
}

type SqlConn struct {
	sqlInfo *SqlInfo
}

type SqlConnPool struct {
	ConnPool chan *sql.DB
}

func (SC *SqlConn) MysqlConn(sqlInfo *SqlInfo) (*sql.DB, error) {
	connectStr := sqlInfo.MysqlUser + "@" + sqlInfo.MysqlPassword + "@tcp(" + sqlInfo.MysqlHost + ":" + string(sqlInfo.MysqlPort) + ")" + "/" + sqlInfo.MysqlDB + "?charset=" + sqlInfo.CharSet

	conn, err := sql.Open("mysql", connectStr)
	if err != nil {
		Output("","Connect to mysql failed!")
		fmt.Print(err)
		return nil, err
	} else {
		Output("Sql connection success!", "")
		return conn,nil
	}
	//return nil,errors.New("connection failed")
}

func (CP *SqlConnPool) CreateConnPool(si *SqlInfo, poolWidth int) (chan *sql.DB, error) {
	ConnChan := make(chan *sql.DB, poolWidth)
	conPool := SqlConnPool{}
	for i := 0; i< poolWidth;i++ {
		sqlCon := SqlConn{}
		//sql_conn.sql_info = si
		conn,err := sqlCon.MysqlConn(si)
		if err != nil {
			fmt.Println("Connecting to mysql error")
			//fmt.Errorf("%e", err)
			return nil, err
		} else {
			ConnChan <- conn
		}

	}
	conPool.ConnPool = ConnChan
	return ConnChan,nil
}