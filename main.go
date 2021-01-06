package main

import (
    "github.com/BurntSushi/toml"
    "dbhelper"
    "fmt"
    "reflect"
    "./example"

    "model"
    "time"
)


type Config struct {
    DBConfigFile string
    DBName string
    DbServers map[string]dbhelper.DBServer
}

var CONFIG Config

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



func main() {

    err := LoadConfig()
    if err != nil{
        fmt.Println("Init Config Fail: ", err.Error())
        return
    }

    err = dbhelper.InitDB(CONFIG.DbServers)
    if err != nil{
        fmt.Println("Init DB Fail: ", err.Error())
        return
    }

    defer func() {
        fmt.Println("Close DB")
        for _, k := range reflect.ValueOf(dbhelper.DB).MapKeys(){
            dbhelper.DB[fmt.Sprintf("%v", k)].Close()
        }
    }()
    example.T(CONFIG.DBName)


    fmt.Println("* input: id, output: struct")
    fmt.Println(example.GetUserDataById(1))
    fmt.Println(example.GetUserDataById_v2(2))

    fmt.Println("\n* input: ids, output: struct of list")
    fmt.Println(example.GetUserDataByIds([]int64{1, 2, 3}))

    fmt.Println("\n* input: filter condition , output: struct of list")
    var parmas map[string]interface{}
    parmas = make(map[string]interface{})
    fmt.Sprintf("%v", parmas)
    parmas["name"] = "CC"
    parmas["tag"] = 2
    fmt.Println(example.GetUserDataByFilter(parmas))

    fmt.Println("\n* input: filter condition , output: count() ")
    fmt.Println(example.GetUserCountByFilter(parmas))

    fmt.Println("\n* input: select column , output: struct of list ")
    fmt.Println(example.GetUserDataDefineField([]string{"name", "old"}))


    fmt.Println("\n* use sql query data , output: struct of list ")
    fmt.Println(example.GetUserDataBySQL(2))


    fmt.Println("\n* demo join by sql , output: define data ")
    example.UseStatementQueryData()


    fmt.Println("\n* insert row ")
    user := &model.User{
       Name:"DD",
       Old: 34,
       Tag: 1,
    }
    fmt.Println(example.AddUser(user))


    fmt.Println("\n* update rows, output: update count ")
    p2 := make(map[string]interface{})
    p2["tag"] = time.Now().Unix()
    fmt.Println(example.UpdateUserByFilter(3, p2))


    fmt.Println("\n* transaction & rollback ")
    example.TransactionAndRollback()


    fmt.Println("\n* Join 1")
    fmt.Println(example.Join1(36))

    return
}

