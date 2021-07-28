package cms3

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

    err = db.Ping()

    return &PgStore{
		DB: db,
	}
}

 
// CreatePage creates a new page in our DB
func CreatePage(p *Page) (int, error) {
	var id int
	err := store.DB.QueryRow("INSERT INTO pages(title, content) VALUES($1, $2) RETURNING id", p.Title, p.Content).Scan(&id)
	return id, err
}

// CreatePost creates a new post in our DB
func CreatePost(p *Post) (int, error) {
	var id int
	err := store.DB.QueryRow("INSERT INTO posts(title, content, date_created) VALUES($1, $2, $3) RETURNING id", p.Title, p.Content, p.DatePublished).Scan(&id)
	return id, err
}

// GetPage gets the page by it's ID.
func GetPage(id string) (*Page, error) {
	var p Page
	err := store.DB.QueryRow("SELECT * FROM pages WHERE id = $1", id).Scan(&p.ID, &p.Title, &p.Content)
        fmt.Println(p)
	return &p, err
}

// GetPages returns every page from our DB.
func GetPages() ([]*Page, error) {
	rows, err := store.DB.Query("SELECT * FROM pages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
 
