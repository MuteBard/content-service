package mirrorMove

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    Service "mirrorMove/src/service"
    Dto "mirrorMove/src/dtos"
)

type ActionController struct {
    service *Service.ActionService
}

func NewActionController(service * Service.ActionService) * ActionController {
    return &ActionController{service : service}
}

func (this *ActionController) SearchAction(w http.ResponseWriter, r *http.Request){
    log.Println("GET /action/search")
    actionApiArgs := ManageActionApiArguments(r)

    result, err := this.service.SearchActions(actionApiArgs)
    ErrorResponseHandler(w, err)
    jsonResponse, err := json.Marshal(result)
    JSONResponseHandler(w, jsonResponse, err)
}

func (this *ActionController) GetAction(w http.ResponseWriter, r *http.Request){
    id := r.PathValue("id")
    log.Println("GET /action/"+id)
    result, err := this.service.GetAction(id)
    ErrorResponseHandler(w, err)
    jsonResponse, err := json.Marshal(result)
    JSONResponseHandler(w, jsonResponse, err)
}

func (this *ActionController) CreateAction(w http.ResponseWriter, r *http.Request){
    log.Println("POST /action")
    var actionCreate Dto.ActionCreate
    if err := json.NewDecoder(r.Body).Decode(&actionCreate); err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    result, err := this.service.CreateAction(actionCreate)

    ErrorResponseHandler(w, err)
    jsonResponse, err := json.Marshal(result)
    JSONResponseHandler(w, jsonResponse, err)
}

func (this *ActionController) PatchAction(w http.ResponseWriter, r *http.Request){
    log.Println("PATCH /action")
    var actionUpdate Dto.ActionUpdate

    if err := json.NewDecoder(r.Body).Decode(&actionUpdate); err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    result, err := this.service.UpdateAction(actionUpdate)
    ErrorResponseHandler(w, err)
    jsonResponse, err := json.Marshal(result)
    JSONResponseHandler(w, jsonResponse, err)
}

func (this *ActionController) DeleteAction(w http.ResponseWriter, r *http.Request){
    id := r.PathValue("id")
    log.Println("DELETE /action/"+id)
    result, err := this.service.HideAction(id)
    ErrorResponseHandler(w, err)
    jsonResponse, err := json.Marshal(result)
    JSONResponseHandler(w, jsonResponse, err)
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

func ManageActionApiArguments(r *http.Request) Dto.ActionApiArguments{
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