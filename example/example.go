package example

import (
    "model"
    "dbhelper"
    "strings"
    "errors"
    "fmt"
)

var DBName string

func T(dbname string) {
    DBName = dbname
}

func GetUserDataById(id int64) (model.User, error) {
    var user model.User
    db := dbhelper.DB[DBName]
    if err := db.Where("id = ? ", id).First(&user).Error; err != nil {
        return user, err
    }
    return user, nil
}

func GetUserDataById_v2(id int64) (model.User, error) {
    users, err := GetUserDataByIds([]int64{id})
    if err != nil {
        return model.User{}, err
    }

    if len(users) == 0 {
        return model.User{}, errors.New("not found data")
    } else {
        return users[0], nil
    }
}

func GetUserDataByFilter(params map[string]interface{}) ([]model.User, error) {
    var users []model.User
    db := dbhelper.DB[DBName]
    if err := db.Where(params).Find(&users).Error; err != nil {
        return users, err
    }
    return users, nil
}

func GetUserDataDefineField(fields []string) ([]model.User, error) {
    var users []model.User
    db := dbhelper.DB[DBName]
    if err := db.Select(strings.Join(fields, ",")).Find(&users).Error; err != nil {
        return users, err
    }
    return users, nil
}

func GetUserCountByFilter(params map[string]interface{}) (int64, error) {
    var cnt int64
    var user model.User
    db := dbhelper.DB[DBName]
    if err := db.Model(&user).Where(params).Count(&cnt).Error; err != nil {
        return cnt, err
    }
    return cnt, nil
}
func GetUserDataByIds(ids []int64) ([]model.User, error) {
    var users []model.User
    db := dbhelper.DB[DBName]
    if err := db.Where(ids).Find(&users).Error; err != nil {
        return users, err
    }
    return users, nil
}

func GetUserDataBySQL(id int64) ([]model.User, error) {
    var users []model.User
    db := dbhelper.DB[DBName]
    sqlStr := fmt.Sprintf("select * from user where id > %d", id)
    if err := db.Raw(sqlStr).Scan(&users).Error; err != nil {
        return users, err
    }
    return users, nil
}

func UseStatementQueryData() {
    db := dbhelper.DB[DBName]
    sqlStr := fmt.Sprintf("select name , students.id, math_score as s from students inner join math on students.id = math.id")
    rows, _ := db.Raw(sqlStr).Rows()
    defer rows.Close()
    for rows.Next() {
        var name string
        var id int64
        var ss int
        rows.Scan(&name, &id, &ss)
        fmt.Printf("name=%s, id=%d, score=%d\n", name, id, ss)
    }
}

func AddUser(user *model.User) (model.User, error) {
    db := dbhelper.DB[DBName]
    err := db.Create(&user).Error
    return *user, err
}

func UpdateUserByFilter(id int64, params map[string]interface{}) (int64) {
    var user model.User
    db := dbhelper.DB[DBName]
    return db.Model(&user).Where("id >= ? ", id).Updates(params).RowsAffected
}

func TransactionAndRollback() error {

    db := dbhelper.DB[DBName]
    tx := db.Begin()

    if tx.Error != nil {
        return tx.Error
    }

    if err := tx.Create(&model.User{Name: "Giraffe"}).Error; err != nil {
        tx.Rollback()
        return err
    }

    if err := tx.Create(&model.User{Old: 12313131312312312}).Error; err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit().Error
}


func Join1(id int64) (model.User, error) {
    var user model.User
    db := dbhelper.DB[DBName]
    if err := db.Joins("left join math on user.id = math.id").Where("user.id = ?", id).Find(&user).Error;
        err != nil {
            return user, err
    }
    return user, nil
}


func Join2(id int64)  {

    db := dbhelper.DB[DBName]
    db.Table("lottery_otp").Select("lottery_otp_secret_key.Key").Joins("inner join lottery_otp_secret_key on lottery_otp_secret_key.LotteryOtpId = lottery_otp.id").Scan(&cc)

}