package cms

import (
        "database/sql"
         _ "github.com/lib/pq"
//         "log"
         "fmt"
)

const (
hostname = "localhost"
host_port = 5432
username = "postgres"
password = "postgres"
databasename = "goprojects"
)

var (
      store = newDB()
)

type PgStore struct {
        DB *sql.DB
}

func newDB() *PgStore {
     connStr := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", hostname, host_port, username, password, databasename)
     db, err := sql.Open("postgres", connStr)
     if err != nil {
        panic(err)
     }

     return &PgStore {
            DB: db,
     }
}

func GetPage(id string) (*Page, error) {
     var p Page
     err := store.DB.QueryRow("SELECT * FROM pages WHERE id = $1", id).Scan(&p.ID, &p.Title, &p.Content)
     fmt.Println(p)
     return &p, err
}

func GetPages() ([]*Page, error) {
//get all the pages
// run in loopand declare one slice of page and append all the page and return

        rows, err := store.DB.Query("SELECT * from pages")
        if err != nil {
              return nil, err 
        }
        pages := []*Page{}

       for rows.Next() {
           var p Page
           err = rows.Scan(&p.ID, &p.Title, &p.Content)
           if err != nil {
              return nil, err
           }
          pages = append(pages, &p)

       }

       return pages, nil
}

func GetPosts() ([]*Post, error) {
     rows, err := store.DB.Query("SELECT * from posts")
     if err != nil {
        return nil, err 
     }
     posts := []*Post{}
     for rows.Next() {
         var p Post 
            err = rows.Scan(&p.ID, &p.Title, &p.Content, &p.DatePublished)
            if err != nil {
               return nil, err
            }
        posts = append(posts, &p)
     }

     return posts, nil
} 


func GetPost(id string) (*Post, error) {
     var p Post
     err := store.DB.QueryRow("SELECT * FROM posts WHERE id = $1", id).Scan(&p.ID, &p.Title, &p.Content, &p.DatePublished)
     
     return &p, err 
}

func CreatePage(p *Page) (int, error) {
     var id int
     err := store.DB.QueryRow("INSERT INTO pages(title, content) VALUES($1, $2) RETURNING id", p.Title, p.Content).Scan(&id)
     return id, err
} 

func CreatePost(p *Post) (int, error) {
     var id int
     err := store.DB.QueryRow("INSERT INTO posts(title, content, date_created) VALUES($1, $2, $3) RETURNING id", p.Title, p.Content, p.DatePublished).Scan(&id)
     return id, err
}

func CreateComment(c *Comment) (int, error) {
     var id int
     err := store.DB.QueryRow("INSERT INTO comments(post_id, author, content,  date_created) VALUES($1, $2, $3, $4) RETURNING id", c.PostID, c.Author, c.Comment, c.DatePublished).Scan(&id)
     return id, err
}

func GetAll(postID int) ([]*Comment, error) {
     rows, err := store.DB.Query("SELECT * FROM comments WHERE id = $1", postID)
     if err != nil {
        return nil, err
     }
     comments := []*Comment{}
     for rows.Next() {
         var c Comment
            err = rows.Scan(&c.ID, &c.PostID, &c.Author, &c.Comment, &c.DatePublished)
            if err != nil {
               return nil, err
            }
        comments = append(comments, &c)
     }

     return comments, nil
}

