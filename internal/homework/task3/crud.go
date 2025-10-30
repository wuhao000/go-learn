package task3

import (
  "database/sql"
  "fmt"

  "gorm.io/driver/sqlite"
  "gorm.io/gorm"

  _ "modernc.org/sqlite"
)

func createDb() *gorm.DB {
  // 使用database/sql连接
  sqlDB, err := sql.Open("sqlite", "file:test.db?cache=shared")
  if err != nil {
    panic(fmt.Sprintf("SQLite连接失败: %v", err))
  }

  db, err := gorm.Open(&sqlite.Dialector{
    Conn: sqlDB,
  })
  if err != nil {
    panic(fmt.Sprintf("GORM初始化失败: %v", err))
  }
  return db
}

//SQL语句练习
//题目1：基本CRUD操作
//假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
//要求 ：
//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

type User struct {
  Id    int
  Name  string
  Age   uint
  Grade string
}

func AddStudent(db *gorm.DB, u *User) error {
  if db == nil {
    return fmt.Errorf("database connection is nil")
  }

  err := db.AutoMigrate(User{})
  if err != nil {
    return fmt.Errorf("failed to migrate: %w", err)
  }

  err = db.Create(u).Error
  if err != nil {
    return fmt.Errorf("failed to create user: %w", err)
  }

  return nil
}

func QueryAgeGt18(db *gorm.DB) []User {
  var users []User
  db.Where("age > ?", 18).Find(&users)
  return users
}

// Update 姓名为 "张三" 的学生年级更新为 "四年级"。
func Update(db *gorm.DB, grade string) {
  db.Model(&User{}).Where("name = ?", "张三").Update("grade", grade)
}

// DeleteAgeLessThan15 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
func DeleteAgeLessThan15(db *gorm.DB) {
  db.Where("age < ?", 15).Delete(&User{})
}
