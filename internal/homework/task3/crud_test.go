package task3

import (
  "fmt"
  "testing"
)

func TestCreate(t *testing.T) {
  db := createDb()
  u := User{
    Name:  "张三",
    Age:   20,
    Grade: "三年级",
  }
  err := AddStudent(db, &u)
  if err != nil {
    t.Fatalf("添加学生失败: %v", err)
  }
}

func TestQuery(t *testing.T) {
  db := createDb()

  // 先添加一个测试数据
  u := User{
    Name:  "李四",
    Age:   19,
    Grade: "二年级",
  }
  err := AddStudent(db, &u)
  if err != nil {
    t.Fatalf("添加测试数据失败: %v", err)
  }
  res := QueryAgeGt18(db)
  fmt.Printf("查询到年龄大于18的学生: %+v\n", res)
  if len(res) == 0 {
    t.Error("查询结果不应该为空")
  }
}

func TestUpdate(t *testing.T) {
  db := createDb()
  var u User
  Update(db, "三年级")
  db.Where("name = ?", "张三").First(&u)
  fmt.Println("grade of", u.Name, "is", u.Grade)
  Update(db, "四年级")
  db.Where("name = ?", "张三").First(&u)
  fmt.Println("grade of", u.Name, "is", u.Grade)
  if u.Grade != "四年级" {
    t.Error("grade of user should be", "四年级")
  }
}

func TestDelete(t *testing.T) {
  db := createDb()
  db.Create(&User{
    Name:  "王五",
    Age:   14,
    Grade: "一年级",
  })
  var count int64
  db.Model(&User{}).Where("age < ?", 15).Count(&count)
  if count < 1 {
    t.Error("count should not less than 1")
  }
  DeleteAgeLessThan15(db)
  db.Model(&User{}).Where("age < ?", 15).Count(&count)
  if count > 0 {
    t.Error("count should be 0")
  }
}
