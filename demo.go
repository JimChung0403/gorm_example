package main

import (
    "github.com/BurntSushi/toml"
    "dbhelper"
    "fmt"
    "reflect"
    _ "github.com/jinzhu/gorm"
    "time"
    "model"
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

type Demo struct {
    ID        int64 `gorm:"primary_key"`
    CreatedAt time.Time `gorm:"column:creDate"`
    UpdatedAt time.Time `gorm:"column:updDate"`
    Name string
    City string `gorm:"default:'TPE'"`
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

    db := dbhelper.DB[CONFIG.DBName]

    var user model.User
    db.Joins("left join math on user.id = math.id").Where("user.id = ?", "36").Find(&user)


    fmt.Println(user)

}

