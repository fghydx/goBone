package GLDBHelper

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	*sql.DB
	connectString string
}

func (m *Mysql) Open() (err error) {
	m.DB, err = sql.Open("mysql", m.connectString)
	if err != nil {
		return err
	}
	return nil
}

func GetDBByConnectString_Mysql(connectString string) *Mysql{
	db := &Mysql{
		connectString:connectString,
	}
	err := db.Open()
	if err != nil {
		fmt.Println("sql open:", err)
		return nil
	}
	return db
}

func GetDB_Mysql(server string,port int,DBName string,UName string,Pass string) *Mysql{
	connectString := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",UName,Pass,"tcp",server,port,DBName)
	db := &Mysql{
		connectString:connectString,
	}
	err := db.Open()
	if err != nil {
		fmt.Println("sql open:", err)
		return nil
	}
	return db
}
