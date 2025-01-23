package mirrorMove

import (
    "time"
    "github.com/jinzhu/gorm"
	Dto "mirrorMove/src/dtos"
)

type MoveRepository struct {
    db *gorm.DB
}


func NewMoveRepository(db *gorm.DB) *MoveRepository {
    return &MoveRepository{db: db}
}

func (this *MoveRepository) GetMove(id string) ([]Dto.Move, error) {
    var moves []Dto.Move

    if err := this.db.Where("id = ?", id).Find(&moves).Error; err != nil {
        return nil, err
    }

    return moves, nil
}

func (this *MoveRepository) CreateMove(moveCreate Dto.MoveCreate)  ([]Dto.Move, error){
    var moves []Dto.Move
    var loopableActions []Dto.LoopableAction

    //create a move object with input name and time
    newMove := Dto.Move {
        Name: moveCreate.Name,
        CreatedAt: time.Now().Format(time.RFC3339),
    }

    //create move in db
    if err := this.db.Create(&newMove).Error; err != nil {
        return nil, err
    }

    //use the input's list of loopable actions
    for _, actionLoop := range moveCreate.ActionLoops {

        //create a new moveAction object for the junction table move_action
        newMoveAction := Dto.MoveAction {
            MoveId: newMove.Id,
            ActionId: actionLoop.ActionId,
            Loops: actionLoop.Loops,
        }

        // create a new move_action
        if err := this.db.Create(&newMoveAction).Error; err != nil {
            return nil, err
        }

        existingAction := Dto.Action{}
        if err := this.db.Where("id = ?", actionLoop.ActionId).First(&existingAction).Error; err != nil {
            return nil, err
        }
            
        loopableAction := Dto.LoopableAction {
            Action: existingAction,
            Loops: actionLoop.Loops,
        }

        loopableActions = append(loopableActions, loopableAction)

    }

    newMove.Actions = loopableActions

    moves = append(moves, newMove)

    return moves, nil;

    }


