package interface_practise

import (
	"fmt"
	"testing"
)

type SqlConn string

type LogOut interface {
		LogOutput(err error,res ...interface{})
}

func LogOutput(err error,res ...interface{}) {
	if err == nil {
		fmt.Printf("Result: %s", res)
	} else {
		fmt.Errorf("Error occured %s", err)
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

type SqlQuery interface {
	SqlQuering(con SqlConn, sql string) (string, error)
}

type SqlUpdate interface {
	SqlUpdating(con SqlConn, sql string) (bool, error)
}

func GetConnection() SqlConn {
	fmt.Print("Connecting now!\n")
	fmt.Print("Connection info!\n")
	return "Fake Sql connection"
}

func CloseConnection(con SqlConn) (res bool, err error) {
	fmt.Printf("Closing connection %s now!\n", con)
	return true,nil
}

func SqlQuering(con SqlConn, sql string) (res string, err error) {
	fmt.Printf("Execute sql %s now!\n", sql)
	return "fake result",nil
}

func SqlUpdating(con SqlConn, sql string) (bool, error) {
	fmt.Printf("Update sql %s now!\n", sql)
	return true, nil
}

func TestPolymorphism(t *testing.T) {
	conn := GetConnection()
	query_res,query_err := SqlQuering(conn, "fake sql")
	LogOutput(query_err, query_res)
	update_res,update_err := SqlUpdating(conn, "fake updating sql")
	LogOutput(update_err, update_res)
	close_res, close_err := CloseConnection(conn)
	LogOutput(close_err, close_res)
}
