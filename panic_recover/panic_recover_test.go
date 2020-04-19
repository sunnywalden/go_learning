package panic_recover

import (
	"fmt"
	sh "go_learning/say_hello"
	"os"
	"testing"
)

func TestDeferPanic(t *testing.T) {
	sql_conn := new(GetSqlConn)
	conn := sql_conn.GetConnection()
	sql_query := new(SqlQueryor)
	query_res,err := sql_query.SqlQuering(conn, "use demo;select * from user")
	if err != nil {
		os.Exit(-1)
	} else {
		fmt.Print(query_res)
		close_conn := new(CloseSqlConn)
		close_conn.CloseConnection(conn)
	}
	sh.SayHello()
}
