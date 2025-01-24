package mirrorMove

import (
    "log"
    "net/http"
    "encoding/json"
    Service "mirrorMove/src/service"
    Dto "mirrorMove/src/dtos"
)

type MoveController struct {
    service *Service.MoveService
}

func NewMoveController(service * Service.MoveService) * MoveController {
    return &MoveController{service : service}
}

func (this *MoveController) SearchMove(w http.ResponseWriter, r *http.Request){
    log.Println("GET /move/search")
    moveApiArgs := ManageMoveApiArguments(r)

    result, err := this.service.SearchMoves(moveApiArgs)
    ErrorResponseHandler(w, err)
    jsonResponse, err := json.Marshal(result)
    JSONResponseHandler(w, jsonResponse, err)
}

func (this *MoveController) GetMove(w http.ResponseWriter, r *http.Request){
    id := r.PathValue("id")
    log.Println("GET /move/"+id)
    result, err := this.service.GetMove(id)
    ErrorResponseHandler(w, err)
    jsonResponse, err := json.Marshal(result)
    JSONResponseHandler(w, jsonResponse, err)
}

func (this *MoveController) CreateMove(w http.ResponseWriter, r *http.Request){
    log.Println("POST /move")
    var moveCreate Dto.MoveCreate
    err := json.NewDecoder(r.Body).Decode(&moveCreate)
    if err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

    result, err := this.service.CreateMove(moveCreate)

    ErrorResponseHandler(w, err)
    jsonResponse, err := json.Marshal(result)
    JSONResponseHandler(w, jsonResponse, err)

}

func (this *MoveController) PatchMove(w http.ResponseWriter, r *http.Request) {
    log.Println("PATCH /move");
    var moveUpdate Dto.MoveUpdate
    err := json.NewDecoder(r.Body).Decode(&moveUpdate)

    if err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

    result, err := this.service.UpdateMove(moveUpdate)

    ErrorResponseHandler(w, err)
    jsonResponse, err := json.Marshal(result)
    JSONResponseHandler(w, jsonResponse, err)
   
}

func (this *MoveController) DeleteMove(w http.ResponseWriter, r *http.Request) {
    id :=r.PathValue("id")
    log.Println("DELETE /move/"+id)

    result, err := this.service.HideMove(id)
    ErrorResponseHandler(w, err)
    jsonResponse, err := json.Marshal(result)
    JSONResponseHandler(w, jsonResponse, err)
}

func ManageMoveApiArguments(r *http.Request) Dto.MoveApiArguments{
    queryValues := r.URL.Query()

    isHiddenStr := queryValues.Get("isHidden")
    var isHidden bool
    if isHiddenStr == "true" {
        isHidden = true
    } else {
        isHidden = false
    }

    moveApiArgs := Dto.MoveApiArguments {
        Name:      queryValues.Get("name"),
        IsHidden:  isHidden,
        Description: queryValues.Get("description"),
        SortOrder: Dto.SortOrder(queryValues.Get("sortOrder")),
        OrderBy:   Dto.OrderBy(queryValues.Get("orderBy")),
    }

    return moveApiArgs
}