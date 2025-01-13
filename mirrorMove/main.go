package main

import (
    "log"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    
    Controller "mirrorMove/src/controller"
    Service "mirrorMove/src/service"
    Repository "mirrorMove/src/repository"
    Env "mirrorMove/src/env"
)


func main() {

    connectionString := Env.GetENV().DB_CONNECTION

    db, err := gorm.Open("mysql", connectionString)
    if err != nil {
        log.Fatal("Error connecting to the database: ", err)
    }
    defer db.Close()

    actionsRepo := Repository.NewActionRepository(db)
    actionsService := Service.NewActionService(actionsRepo)
    Controller.NewActionController(actionsService)

    log.Println("Application started successfully")
}