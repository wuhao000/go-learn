package task3

import (
  "fmt"
  "testing"
)

//续使用博客系统的模型。
//要求 ：
//为 Post 模型添加一个钩子函数，
// 在文章创建时自动更新用户的文章数量统计字段。
//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

func TestCreateHook(t *testing.T) {
  db := createDb()
  err := db.AutoMigrate(&BlogUser{})
  if err != nil {
    return
  }
  db.Delete(&BlogUser{}, "id = 1")
  db.Create(
    &BlogUser{
      Id:   1,
      Name: "李义",
    },
  )
  var u BlogUser
  db.First(&u, "id = 1")
  fmt.Println(u)
  p1 := u.PostCount
  db.Create(&Post{
    UserId:  1,
    Content: "测试文章内容",
  })
  db.Find(&u, "id = 1")
  p2 := u.PostCount
  fmt.Println(u)
  if p2 <= p1 {
    t.Errorf("p2应该比p1大1，但是实际p1 %v p2 %v", p1, p2)
  }
}

func TestDeleteHook(t *testing.T) {
  db := createDb()
  err := db.AutoMigrate(&Post{})
  if err != nil {
    return
  }
  postId := 1000
  commentId := 10000
  db.Delete(&Post{}, "id = ?", postId)
  db.Delete(&Comment{}, "post_id = ?", postId)
  db.Create(&Post{
    Id:            postId,
    Content:       "这是测试",
    CommentStatus: "有评论",
  })
  var post Post
  db.First(&post, "id = ?", postId)
  fmt.Println(post.CommentStatus)
  comment := Comment{
    Id:      commentId,
    PostId:  postId,
    Content: "评论",
  }
  db.Create(&comment)
  db.Delete(&comment)
  db.First(&post, "id = ?", postId)
  fmt.Println(post.CommentStatus)
}
