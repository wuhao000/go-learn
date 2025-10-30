package task3

import (
  "fmt"
  "testing"
)

func TestCreateTables(t *testing.T) {
  db := createDb()
  err := db.AutoMigrate(&BlogUser{})
  if err != nil {
    return
  }
  err = db.AutoMigrate(&Post{})
  if err != nil {
    return
  }
  err = db.AutoMigrate(&Comment{})
  if err != nil {
    return
  }
}

func TestInsertData(t *testing.T) {
  db := createDb()
  //db.Where("1=1").Delete(&BlogUser{})
  //db.Where("1=1").Delete(&Comment{})
  //db.Where("1=1").Delete(&Post{})
  db.Create(&BlogUser{
    Name: "张三",
    Posts: []Post{
      {
        Content: "博客1",
        Comments: []Comment{
          {
            Content: "博客1 评论1",
          },
          {
            Content: "博客1 评论2",
          },
        },
      },
      {
        Content: "博客2",
        Comments: []Comment{
          {
            Content: "博客2 评论1",
          },
          {
            Content: "博客2 评论2",
          },
        },
      },
    },
  })
  var users []BlogUser
  var comments []Comment
  var posts []Post
  db.Preload("Posts.Comments").Find(&users)
  db.Preload("Comments").Find(&posts)
  db.Find(&comments)
  fmt.Println(users)
  fmt.Println(comments)
  fmt.Println(posts)
  if len(users) != 1 {
    t.Error("should have 1 user")
  }
  if len(posts) != 2 {
    t.Error("should have 2 posts")
  }
  if len(comments) != 4 {
    t.Error("should have 4 comments")
  }
}
