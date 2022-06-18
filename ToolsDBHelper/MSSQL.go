package ToolsDBHelper

import (
	"database/sql"
	"fmt"
	//_ "github.com/mattn/go-adodb"
	_ "github.com/denisenkom/go-mssqldb"
)

type Mssql struct {
	*sql.DB
	connectString string
}

func (m *Mssql) Open() (err error) {
	m.DB, err = sql.Open("mssql", m.connectString)
	if err != nil {
		return err
	}
	return nil
}

func GetDBByConnectString_Mssql(connectString string) *Mssql {
	db := &Mssql{
		connectString: connectString,
	}
	err := db.Open()
	if err != nil {
		fmt.Println("sql open:", err)
		return nil
	}
	return db
}

// GetDB_Mssql encrypt = true,false,disable
func GetDB_Mssql(server string, port int, DBName string, UName string, Pass string, encrypt string) *Mssql {
	connectString := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s;encrypt=%s", server, port, DBName, UName, Pass, encrypt)
	db := &Mssql{
		connectString: connectString,
	}
	err := db.Open()
	if err != nil {
		fmt.Println("sql open:", err)
		return nil
	}
	return db
}
