package model

type MyModel struct {
    ID uint `gorm:"primary_key"`
}

type User struct {
    Id   int64 `gorm:"primary_key"`
    Name string
    Old  int64
    Tag int64

    //// 如果 struct 有CreatedAt/UpdatedAt,
    //// 新增資料時, 會自動塞CreatedAt
    //// 更新資料時, 會自動update UpdatedAt
    //// 兩個欄位資料庫 Data type 應為TIMESTAMP
    //CreatedAt time.Time `gorm:"column:xxxxxx"`
    //UpdatedAt time.Time `gorm:"column:xxxxxx"`
}



// if you want to rest table name
//func (User) TableName() string {
//    return "user"
//}


type Math struct {
    Id   int64 `gorm:"primary_key"`
    MathScore int64 `gorm:"column:math_score"`
}