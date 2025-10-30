package task3

import (
  "fmt"

  "gorm.io/gorm"
)

//续使用博客系统的模型。
//要求 ：
//为 Post 模型添加一个钩子函数，
// 在文章创建时自动更新用户的文章数量统计字段。

func (u *Post) AfterCreate(tx *gorm.DB) (err error) {
  var count int64
  tx.Model(&Post{}).Where("user_id = ?", u.UserId).Count(&count)
  tx.Model(&BlogUser{}).Where("id = ?", u.UserId).Update("post_count", count)
  return
}

// AfterDelete 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
  var count int64
  tx.Model(&Comment{}).Where("post_id = ?", c.PostId).Count(&count)
  if count == 0 {
    tx.Model(&Post{}).Where("id = ?", c.PostId).Update("comment_status", "无评论")
  }
  fmt.Println("执行了钩子函数，count = ", count, " post id is ", c.PostId)
  return
}
