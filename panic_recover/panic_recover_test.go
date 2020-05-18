package panic_recover

import (
	"fmt"
	sh "github.com/sunnywalden/go_learning/say_hello"
	sql "github.com/sunnywalden/go_learning/panic_recover"
	"os"
	"testing"
)

func TestDeferPanic(t *testing.T) {
	sqlConn := new(sql.GetSqlConn)
	conn := sqlConn.GetConnection()
	sqlQuery := new(sql.SqlQueryor)
	queryRes,err := sqlQuery.SqlQuering(conn, "use demo;select * from user")
	if err != nil {
		os.Exit(-1)
	} else {
		fmt.Print(queryRes)
		closeConn := new(sql.CloseSqlConn)
		closeConn.CloseConnection(conn)
	}
	sh.SayHello()
}
