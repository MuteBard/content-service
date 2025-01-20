package mirrorMove

import (
    "fmt"
    "strconv"
    "log"
    "time"
    "net/http"
    "encoding/json"

    Service "mirrorMove/src/service"
    Dto "mirrorMove/src/dtos"
)

func NewActionController(service *Service.ActionService)  {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /action/search", func(w http.ResponseWriter, r *http.Request) {
        log.Println("GET /action/search")
        actionApiArgs := ManageApiArguments(r)

        result, err := service.SearchActions(actionApiArgs)
        ErrorResponseHandler(w, err)
        jsonResponse, err := json.Marshal(result)
        JSONResponseHandler(w, jsonResponse, err)
    })

    mux.HandleFunc("GET /action/{id}", func(w http.ResponseWriter, r *http.Request) {
        id := r.PathValue("id")
        log.Println("GET /action/"+id)
        result, err := service.GetAction(id)
        ErrorResponseHandler(w, err)
        jsonResponse, err := json.Marshal(result)
        JSONResponseHandler(w, jsonResponse, err)
    })

    mux.HandleFunc("POST /action", func(w http.ResponseWriter, r *http.Request) {
        log.Println("POST /action")

        var data map[string]interface{}
        if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
            http.Error(w, "Bad Request", http.StatusBadRequest)
            return
        }

        action := Dto.Action {
            Name: data["name"].(string),
            CreatedAt: time.Now().Format(time.RFC3339),
            Description: data["description"].(string),
            Seconds:data["seconds"].(float64),
            Token: data["token"].(string),
        }

        result, err := service.CreateAction(action)
        ErrorResponseHandler(w, err)
        jsonResponse, err := json.Marshal(result)
        JSONResponseHandler(w, jsonResponse, err)
    })

    mux.HandleFunc("PATCH /action", func(w http.ResponseWriter, r *http.Request) {
        log.Println("PATCH /action")

        var data map[string]interface{}
        if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
            http.Error(w, "Bad Request", http.StatusBadRequest)
            return
        }

        id, err := strconv.Atoi(fmt.Sprintf("%v", data["id"]))
        if err != nil {
            fmt.Println(data["id"])
            http.Error(w, "Invalid ID", http.StatusBadRequest)
            return
        }

        action := Dto.Action {
            Id: uint(id),
            Name: data["name"].(string),
            CreatedAt: time.Now().Format(time.RFC3339),
            Description: data["description"].(string),
            Seconds:data["seconds"].(float64),
            Token: data["token"].(string),
        }

        result, err := service.UpdateAction(action)
        ErrorResponseHandler(w, err)
        jsonResponse, err := json.Marshal(result)
        JSONResponseHandler(w, jsonResponse, err)
    })

    mux.HandleFunc("DELETE /action/{id}", func(w http.ResponseWriter, r *http.Request) {
        id := r.PathValue("id")
        log.Println("DELETE /action/"+id)
        result, err := service.HideAction(id)
        ErrorResponseHandler(w, err)
        jsonResponse, err := json.Marshal(result)
        JSONResponseHandler(w, jsonResponse, err)
    })

    err := http.ListenAndServe("localhost:8080", mux)
    if err != nil {
        fmt.Println(err.Error())
    }
}

func ErrorResponseHandler(w http.ResponseWriter, err error){
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        fmt.Println("Error while getting actions:", err)
        return
    }
}

func JSONResponseHandler(w http.ResponseWriter, jsonResponse []byte, err error) {
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        fmt.Println("Error while encoding JSON:", err)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}
func ManageApiArguments(r *http.Request) Dto.ActionApiArguments{
    queryValues := r.URL.Query()

    isHiddenStr := queryValues.Get("isHidden")
    var isHidden bool
    if isHiddenStr == "true" {
        isHidden = true
    } else {
        isHidden = false
    }

    actionApiArgs := Dto.ActionApiArguments {
        Name:      queryValues.Get("name"),
        IsHidden:  isHidden,
        Description: queryValues.Get("description"),
        SortOrder: Dto.SortOrder(queryValues.Get("sortOrder")),
        OrderBy:   Dto.OrderBy(queryValues.Get("orderBy")),
    }

    return actionApiArgs
}