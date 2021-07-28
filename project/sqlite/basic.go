package main

import (
         "github.com/jinzhu/gorm"
//         "github.com/jinzhu/gorm/dialects/sqlite"
         "fmt"
         _ "github.com/jinzhu/gorm/dialects/sqlite"
)
//var db *gorm.DB

type User struct {
    gorm.Model
    Name  string
    Email string
}

func main() {

     fmt.Println("Go ORM Tutorial")
     initialMigration()
     createUser()
}

func initialMigration() {

     db, err := gorm.Open("sqlite3", "test.db")
     if err != nil {
           panic(err)
     }

     defer db.Close()

     db.AutoMigrate(&User{})

}

func createUser() {


    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
       panic(err)
    }
    defer db.Close()
    db.Create(&User{Name: "Onkar", Email: "o@test.com", })

    var user []User

    db.Find(&user)
    fmt.Println("{}", user)

}
    
      
