package main

import (
    "log"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    
    Controller "mirrorMove/src/controller"
    Service "mirrorMove/src/service"
    Repository "mirrorMove/src/repository"
    Env "mirrorMove/src/env"
    "net/http"
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

    movesRepo := Repository.NewMoveRepository(db)
    movesService := Service.NewMoveService(movesRepo)

    mux := http.NewServeMux()

    actionController :=  Controller.NewActionController(actionsService)
    moveController :=  Controller.NewMoveController(movesService)



    mux.HandleFunc("GET /action/search", actionController.SearchAction)
    mux.HandleFunc("GET /action/{id}", actionController.GetAction)
    mux.HandleFunc("POST /action", actionController.CreateAction)
    mux.HandleFunc("PATCH /action", actionController.PatchAction)
    mux.HandleFunc("DELETE /action/{id}", actionController.DeleteAction)

    mux.HandleFunc("GET /move/{id}", moveController.GetMove)
    mux.HandleFunc("POST /move", moveController.CreateMove)


    err = http.ListenAndServe("localhost:8080", mux)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Application started successfully")
}