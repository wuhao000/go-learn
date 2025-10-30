package task3

// 使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func queryPosts() []Post {
  db := createDb()
  var posts []Post
  db.Preload("Comments").Where("user_id = ?", 5).
    Find(&posts)
  return posts
}

//使用Gorm查询评论数量最多的文章信息
func queryMostComments() Post {
  db := createDb()
  db.Create(&BlogUser{
    Name: "王五",
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
          {
            Content: "博客1 评论3",
          },
          {
            Content: "博客1 评论4",
          },
          {
            Content: "博客1 评论5",
          },
        },
      },
    },
  })
  var post Post
  subQuery := db.Model(&Comment{}).Select("post_id").
    Group("post_id").
    Order("count(*) desc").
    Limit(1)
  db.Preload("Comments").Where("id in (?)", subQuery).
    Find(&post)
  return post
}
