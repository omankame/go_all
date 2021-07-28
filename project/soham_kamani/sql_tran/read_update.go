package main

import (
         "log"
         "context"
         "database/sql"
         _ "github.com/lib/pq"
         "fmt"
)

const (
hostname = "localhost"
host_port = 5432
username = "postgres"
password = "postgres"
databasename = "bird_encyclopedia"
)

func main() {
     conStr := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", hostname, host_port, username, password, databasename)

     db, err := sql.Open("postgres", conStr)
     if err != nil {
        log.Fatal(err)
        return
     }

     defer db.Close()

    fmt.Println("Connection to Postgress sucesfully done")

    ctx := context.Background()

    Tx, err := db.BeginTx(ctx , nil) 
    if err != nil {
       log.Fatal(err)
    }


   _, err =  Tx.ExecContext(ctx, "INSERT INTO pets (name, species) VALUES ('jimmu', 'cat')")  
    if err != nil {
       Tx.Rollback()
       return
    }

    row := Tx.QueryRow("SELECT count(*) FROM pets WHERE species='cat'")
    fmt.Printf("%v", row)
    var countCat int
    err = row.Scan(&countCat)
    if err != nil {
       Tx.Rollback()
       return
    }

    _, err = Tx.ExecContext(ctx, "UPDATE food SET quantity=quantity+$1 WHERE name='Cat Food'", 5*countCat)
    if err != nil {
		Tx.Rollback()
		return
	}

    
        err = Tx.Commit()
	if err != nil {
		log.Fatal(err)
	} 

}      
