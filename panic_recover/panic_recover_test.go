package panic_recover

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"testing"
)

type SqlConn string

type Sql string

type QeuryResult string

var ConnError string = "Connection Error"

type GetConn interface {
	GetConnection() SqlConn
}

type GetSqlConn struct {
}

func (g *GetSqlConn) GetConnection() SqlConn {
	fmt.Print("Connecting now!\n")
	fmt.Print("Connection info!\n")
	return SqlConn("Fake Sql connection")
}

type CloseConn interface {
	CloseConnection(con string) (bool, error)
}

type CloseSqlConn struct {
}

func (c *CloseSqlConn) CloseConnection(con SqlConn) (res bool, err error) {
	fmt.Printf("Closing connection %s now!\n", con)
	return true,nil
}

type SqlQuery interface {
	SqlQuering(con SqlConn, sql string) (res string, err error)
}

type SqlQueryor struct {
}


func (q *SqlQueryor) SqlQuering(con SqlConn, sql string) (res string, err error) {
	defer func() {
		if err := recover();err != nil {
			fmt.Errorf("%s", err)
		}
	}()
	fmt.Print("Sql query starting!")
	fmt.Printf("Execute sql %s now!\n", sql)
	panic(errors.New("Sql Connection lost!"))

}

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
}