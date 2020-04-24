package multichannel_timeout

import (
	"fmt"
	"testing"
)

func TestBufferChannel(t *testing.T) {
	rds := &SqlInfo{}
	rds.MysqlHost = "127.0.0.1"
	rds.MysqlPort = 3346
	rds.MysqlUser = "tezign"
	rds.MysqlPassword = "tezign0818"
	rds.MysqlDB = "ops"
	rds.CharSet = "utf8mb4"

	sqlConPool := SqlConnPool{}
	ConnChan, err := sqlConPool.CreateConnPool(rds, 5)
	if err != nil {
		fmt.Printf("Error")
	} else {
		chanLen := len(ConnChan)
		fmt.Printf("%d", chanLen)
		myConn := <- ConnChan
		res, err := myConn.Query("select * from aliyun_zones;")
		if err != nil {
			fmt.Print("Query failed!\n")

		} else {
			fmt.Print(res)
		}


	}
}
