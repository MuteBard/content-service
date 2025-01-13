package mirrorMove

import (
    "fmt"
    "net/http"
    "encoding/json"
    Service "mirrorMove/src/service"
)

func NewActionController(service *Service.ActionService)  {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /actions", func(w http.ResponseWriter, r *http.Request) {
        result, err := service.GetAllActions()
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            fmt.Println("Error while getting actions:", err)
            return
        }

        jsonResponse, err := json.Marshal(result)
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            fmt.Println("Error while encoding JSON:", err)
            return
        }
        
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(jsonResponse)
    })

    err := http.ListenAndServe("localhost:8080", mux)
    if err != nil {
        fmt.Println(err.Error())
    }
}