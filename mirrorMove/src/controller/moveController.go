package mirrorMove

import (
    "log"
    // "fmt"
    // "time"
    // "strconv"
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

func (mc *MoveController) GetMove(w http.ResponseWriter, r *http.Request){
    id := r.PathValue("id")
    log.Println("GET /move/"+id)
    result, err := mc.service.GetMove(id)
    ErrorResponseHandler(w, err)
    jsonResponse, err := json.Marshal(result)
    JSONResponseHandler(w, jsonResponse, err)
}

func (mc *MoveController) CreateMove(w http.ResponseWriter, r *http.Request){
    var moveCreate Dto.MoveCreate
    err := json.NewDecoder(r.Body).Decode(&moveCreate)
    if err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

    result, err := mc.service.CreateMove(moveCreate)

    ErrorResponseHandler(w, err)
    jsonResponse, err := json.Marshal(result)
    JSONResponseHandler(w, jsonResponse, err)

}



