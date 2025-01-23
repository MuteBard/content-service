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

func (this *MoveRepository) buildSearchQuery(args Dto.MoveApiArguments) *gorm.DB {
    query := this.db

    query = query.Where("moves.is_hidden = ?", args.IsHidden)

    if (args.Name != "") {
        query = query.Where("moves.name LIKE ?", "%" + args.Name + "%")
    }

    if (args.Name != "") {
        query = query.Where("moves.description_ LIKE ?", "%" + args.Description + "%")
    }

    switch args.OrderBy {
    case "NAME":
        if args.SortOrder == Dto.ASC {
            query = query.Order("moves.name ASC")
        } else if args.SortOrder == Dto.DESC {
            query = query.Order("moves.name DESC")
        }
    case "SECONDS":
        if args.SortOrder == Dto.ASC {
            query = query.Order("moves.seconds ASC")
        } else if args.SortOrder == Dto.DESC {
            query = query.Order("moves.seconds DESC")
        }
    case "CREATEDAT":
        if args.SortOrder == Dto.ASC {
            query = query.Order("moves.created_at ASC")
        } else if args.SortOrder == Dto.DESC {
            query = query.Order("moves.created_at DESC")
        }
    default:
        query = query.Order("moves.created_at DESC")
    }

    return query
}



func (this *MoveRepository) SearchMoves(args Dto.MoveApiArguments) ([]Dto.Move, error) {
    var joinMoveAction []Dto.JoinMoveAction 
    var moves []Dto.Move

    query := this.buildSearchQuery(args)

    if err := query.Debug().
        Select("moves.id, moves.name, moves.created_at, moves.is_hidden, moves.description_, moves.seconds, move_actions.loops, actions.id, actions.name, actions.created_at, actions.is_hidden, actions.description_, actions.seconds, actions.token").
        Table("moves").
        Joins("LEFT OUTER JOIN move_actions ON moves.id = move_actions.move_id").
        Joins("LEFT OUTER JOIN actions ON move_actions.action_id = actions.id").
        Scan(&joinMoveAction).Error; err != nil {
        return nil, err
    }

    moves = generateMoves(joinMoveAction)
    return moves, nil
}





func (this *MoveRepository) GetMove(id string) ([]Dto.Move, error) {
    var joinMoveAction []Dto.JoinMoveAction 
    var moves []Dto.Move

    if err := this.db.Debug().
        Select("moves.id, moves.name, moves.created_at, moves.is_hidden, moves.description_, moves.seconds, move_actions.loops, actions.id, actions.name, actions.created_at, actions.is_hidden, actions.description_, actions.seconds, actions.token").
        Table("moves").
        Joins("LEFT OUTER JOIN move_actions ON moves.id = move_actions.move_id").
        Joins("LEFT OUTER JOIN actions ON move_actions.action_id = actions.id").
        Where("moves.id = ?", id).
        Scan(&joinMoveAction).Error; err != nil {
        return nil, err
    }

    moves = generateMoves(joinMoveAction)

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


func generateMoves(joinMoveAction []Dto.JoinMoveAction) []Dto.Move {
    var moves []Dto.Move
    loopMap := make(map[uint][]Dto.LoopableAction)
    moveMap := make(map[uint]Dto.Move)
    for _, jma := range joinMoveAction {
        
        action := Dto.Action{
            Id: jma.AId,
            Name: jma.AName,
            CreatedAt: jma.ACreatedAt,
            IsHidden: jma.AIsHidden,
            Description: jma.ADescription,
            Seconds: jma.ASeconds,
            Token: jma.AToken,
        }

        loopAction := Dto.LoopableAction{
            Loops: jma.Loops,
            Action: action,
        }

        move := Dto.Move{
            Id: jma.MId,
            Name: jma.MName,
            CreatedAt: jma.MCreatedAt,
            IsHidden: jma.MIsHidden,
            Description: jma.MDescription,
            Seconds: jma.MSeconds,
        }

        loopMap[jma.MId] = append(loopMap[jma.MId], loopAction)
        moveMap[jma.MId] = move
    }

    for key, value := range moveMap {
        value.Actions = loopMap[key]
        moves = append(moves, value)
    }
    return moves;
}
