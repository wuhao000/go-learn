package task3

import (
  "fmt"
  "testing"
)

// 使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func TestQueryPosts(t *testing.T) {
  var posts []Post = queryPosts()
  fmt.Println(posts)
  if len(posts) == 0 {
    t.Error("posts should not be empty")
  }
  for _, p := range posts {
    if len(p.Comments) == 0 {
      t.Error("comments of post should not be empty")
    }
  }
}

//使用Gorm查询评论数量最多的文章信息
func TestQueryMostComments(t *testing.T) {
  var post Post = queryMostComments()
  fmt.Println(post)
  if len(post.Comments) < 3 {
    t.Error("comments should be more than 2")
  }
}
