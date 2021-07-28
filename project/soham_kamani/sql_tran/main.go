package main

import (
         "log"
         "database/sql"
         _ "github.com/lib/pq"
         "context"
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
     connStr := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", hostname, host_port, username, password, databasename)
     db, err := sql.Open("postgres", connStr) 
     if err != nil {
        log.Fatal(err)
     }

     ctx := context.Background()

     Tx, err := db. BeginTx(ctx, nil )
     if err != nil {
        log.Fatal(err)
     }

     _, err = Tx.ExecContext(ctx, "INSERT INTO pets (name, species) VALUES ('Fido', 'dog'), ('Albert', 'cat')") 
     if err != nil {
        Tx.Rollback()
     }

     _, err = Tx.ExecContext(ctx, "INSERT INTO food (name, quantity) VALUES ('Dog Biscuit', 3), ('Cat Food', 5)")
     if err != nil {
        Tx.Rollback()
     }

     err = Tx.Commit() 
     if err != nil {
        Tx.Rollback()
     }

     fmt.Println("So finally transaction done sucessfully")

}
      
