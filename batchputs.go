package main

import (
    "fmt"
    "time"
    "github.com/theplant/batchputs"
    "dbhelper"
    //"model"
    "reflect"
    "github.com/BurntSushi/toml"
)

type Config struct {
    DBConfigFile string
    DBName       string
    DbServers    map[string]dbhelper.DBServer
}

func LoadConfig() error {
    file := "./config.toml"
    _, err := toml.DecodeFile(file, &CONFIG)
    if err != nil {
        return err
    }
    _, err = toml.DecodeFile(CONFIG.DBConfigFile, &CONFIG)
    if err != nil {
        return err
    }

    return nil
}

var CONFIG Config

func main() {
    err := LoadConfig()
    if err != nil {
        fmt.Println("Init Config Fail: ", err.Error())
        return
    }

    err = dbhelper.InitDB(CONFIG.DbServers)
    if err != nil {
        fmt.Println("Init DB Fail: ", err.Error())
        return
    }

    defer func() {
        fmt.Println("Close DB")
        for _, k := range reflect.ValueOf(dbhelper.DB).MapKeys() {
            dbhelper.DB[fmt.Sprintf("%v", k)].Close()
        }
    }()

    db := dbhelper.DB[CONFIG.DBName]

    rows := [][]interface{}{}
    for i := 0; i < 10; i++ {
        rows = append(rows, []interface{}{
            fmt.Sprintf("CODE_%d", i),
            fmt.Sprintf("short name %d", i),
            i+1,
        })
    }
    columns := []string{"name", "city", "id"}

    start := time.Now()
    err = batchputs.Put(db.DB(), "mysql", "demo", "id", columns, rows)
    if err != nil {
        panic(err)
    }
    duration := time.Since(start)
    fmt.Println("Inserts 30000 records using less than 3 seconds:", duration.Seconds() < 3)
    //
    //rows = [][]interface{}{}
    //for i := 0; i < 20000; i++ {
    //    rows = append(rows, []interface{}{
    //        fmt.Sprintf("CODE_%d", i),
    //        fmt.Sprintf("short name %d", i),
    //        i + 1,
    //    })
    //}
    //start = time.Now()
    //err = batchputs.Put(db.DB(), dialect, "countries", "code", columns, rows)
    //if err != nil {
    //    panic(err)
    //}
    //duration = time.Since(start)
    //fmt.Println("Updates 20000 records using less than 3 seconds:", duration.Seconds() < 3)
}
