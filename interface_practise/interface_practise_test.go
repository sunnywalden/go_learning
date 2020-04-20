package interface_practise

import (
	"fmt"
	"testing"
)

type SqlConn string

type LogOut interface {
	LogOutput(err error,res ...interface{})
}

type Logger struct {
}

func (l *Logger) LogOutput(err error,res ...interface{}) {
	if err == nil {
		fmt.Printf("Result: %s", res)
	} else {
		fmt.Printf("Error occurred %s", err)
	}
}

type GetConn interface {
	GetConnection() SqlConn
}

type CloseConn interface {
	CloseConnection(con string) (bool, error)
}



type SqlCon interface {
	GetConn
	CloseConn
}


type GetSqlConn struct {
}

func (g *GetSqlConn) GetConnection() SqlConn {
	fmt.Print("Connecting now!\n")
	fmt.Print("Connection info!\n")
	return "Fake Sql connection"
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
	fmt.Printf("Execute sql %s now!\n", sql)
	return "fake result",nil
}

type SqlUpdator struct {
}

func (u *SqlUpdator) SqlUpdating(con SqlConn, sql string) (bool, error) {
	fmt.Printf("Update sql %s now!\n", sql)
	return true, nil
}

func TestPolymorphism(t *testing.T) {
	con := &GetSqlConn{}
	conn := con.GetConnection()
	l := &Logger{}
	q := &SqlQueryor{}
	query_res,query_err := q.SqlQuering(conn, "fake sql")
	l.LogOutput(query_err, query_res)
	u := &SqlUpdator{}
	update_res,update_err := u.SqlUpdating(conn, "fake updating sql")
	l.LogOutput(update_err, update_res)
	c := &CloseSqlConn{}
	close_res, close_err := c.CloseConnection(conn)
	l.LogOutput(close_err, close_res)
}
