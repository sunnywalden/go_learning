package panic_recover

import (
	"fmt"
	sh "github.com/sunnywalden/go_learning/say_hello"
	"os"
	"testing"
)

func TestDeferPanic(t *testing.T) {
	sqlConn := new(GetSqlConn)
	conn := sqlConn.GetConnection()
	sqlQuery := new(SqlQueryor)
	queryRes,err := sqlQuery.SqlQuering(conn, "use demo;select * from user")
	if err != nil {
		os.Exit(-1)
	} else {
		fmt.Print(queryRes)
		closeConn := new(CloseSqlConn)
		closeConn.CloseConnection(conn)
	}
	sh.SayHello()
}
