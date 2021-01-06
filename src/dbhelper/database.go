package dbhelper

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

type DBServer struct {
    Host      string
    Port      int
    User      string
    Pwd       string
    Database  string
    IdleConn int
    MaxConn  int
}

var DB map[string]*gorm.DB

func InitDB(dbServers map[string]DBServer) error {
    DB = make(map[string]*gorm.DB)

    for k, v := range dbServers {
        connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", v.User, v.Pwd, v.Host, v.Port, v.Database)
        gdb, err := gorm.Open("mysql", connectString)
        if err != nil {
            return err
        }
        gdb.SingularTable(true)
        gdb.DB().SetMaxIdleConns(v.IdleConn)
        gdb.DB().SetMaxOpenConns(v.MaxConn)

        // if want to loop at sql statements
        gdb.LogMode(true)
        DB[string(k)] = gdb
    }

    return nil
}
