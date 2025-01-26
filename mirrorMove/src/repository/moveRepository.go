package mirrorMove

import (
	"log"
	Dto "mirrorMove/src/dtos"
	"time"

	"github.com/jinzhu/gorm"
)

type MoveRepository struct {
	db *gorm.DB
}

func NewMoveRepository(db *gorm.DB) *MoveRepository {
	return &MoveRepository{db: db}
}

var joinFields = "moves.id as MId, moves.name as MName, moves.created_at as MCreatedAt, moves.updated_at as MUpdatedAt, moves.is_hidden as MIsHidden, moves.description_ as MDescription, moves.seconds as MSeconds, moves_actions.loops as Loops, actions.id as AId, actions.name as AName, actions.created_at as ACreatedAt, actions.updated_at as AUpdatedAt, actions.is_hidden as AIsHidden, actions.description_ as ADescription, actions.seconds as ASeconds, actions.token as AToken"

func (this *MoveRepository) buildSearchQuery(args Dto.MoveApiArguments) *gorm.DB {
	query := this.db

	query = query.Where("moves.is_hidden = ?", args.IsHidden)

	if args.Name != "" {
		query = query.Where("moves.name LIKE ?", "%"+args.Name+"%")
	}

	if args.Name != "" {
		query = query.Where("moves.description_ LIKE ?", "%"+args.Description+"%")
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
	case "UPDATEDAT":
		if args.SortOrder == Dto.ASC {
			query = query.Order("moves.updated_at ASC")
		} else if args.SortOrder == Dto.DESC {
			query = query.Order("moves.updated_at DESC")
		}
	default:
		query = query.Order("moves.updated_at DESC")
	}

	return query
}

func (this *MoveRepository) SearchMoves(args Dto.MoveApiArguments) ([]Dto.Move, error) {
	var joinMovesAction []Dto.JoinMovesAction
	var moves []Dto.Move

	query := this.buildSearchQuery(args)

	if err := query.Debug().
		Select(joinFields).
		Table("moves").
		Joins("LEFT OUTER JOIN moves_actions ON moves.id = moves_actions.move_id").
		Joins("LEFT OUTER JOIN actions ON moves_actions.action_id = actions.id").
		Scan(&joinMovesAction).Error; err != nil {
		return nil, err
	}

	moves = generateMoves(joinMovesAction)
	return moves, nil
}

func (this *MoveRepository) GetMove(id string) ([]Dto.Move, error) {
	var joinMovesAction []Dto.JoinMovesAction
	var moves []Dto.Move

	if err := this.db.Debug().
		Select(joinFields).
		Table("moves").
		Joins("LEFT OUTER JOIN moves_actions ON moves.id = moves_actions.move_id").
		Joins("LEFT OUTER JOIN actions ON moves_actions.action_id = actions.id").
		Where("moves.id = ?", id).
		Scan(&joinMovesAction).Error; err != nil {
		return nil, err
	}

	moves = generateMoves(joinMovesAction)

	return moves, nil
}

func (this *MoveRepository) CreateMove(moveCreate Dto.MoveCreate) ([]Dto.Move, error) {
	var newMove Dto.Move
	var moves []Dto.Move
	var loopableActions []Dto.LoopableAction

	newMove.Name = moveCreate.Name
	newMove.Description = moveCreate.Description
	newMove.CreatedAt = time.Now().Format(time.RFC3339)
	newMove.UpdatedAt = time.Now().Format(time.RFC3339)
	newMove.Seconds = moveCreate.Seconds

	//create move in db
	if err := this.db.Create(&newMove).Error; err != nil {
		return nil, err
	}

	//use the input's list of loopable actions
	for _, actionLoop := range moveCreate.ActionLoops {

		//create a new MovesAction object for the junction table move_action
		newMovesAction := Dto.MovesAction{
			MoveId:   newMove.Id,
			ActionId: actionLoop.ActionId,
			Loops:    actionLoop.Loops,
		}

		// create a new move_action
		if err := this.db.Create(&newMovesAction).Error; err != nil {
			return nil, err
		}

		existingAction := Dto.Action{}
		if err := this.db.Where("id = ?", actionLoop.ActionId).First(&existingAction).Error; err != nil {
			return nil, err
		}

		loopableAction := Dto.LoopableAction{
			Action: existingAction,
			Loops:  actionLoop.Loops,
		}

		loopableActions = append(loopableActions, loopableAction)

	}

	newMove.Actions = loopableActions
	moves = append(moves, newMove)

	return moves, nil
}

func (this *MoveRepository) UpdateMove(moveUpdate Dto.MoveUpdate) ([]Dto.Move, error) {
	var existingMove Dto.Move
	var moves []Dto.Move
	var loopableActions []Dto.LoopableAction

	if err := this.db.Where("id = ?", moveUpdate.Id).First(&existingMove).Error; err != nil {
		return nil, err
	}

	//create a move object with input name and time
	existingMove.Name = moveUpdate.Name
	existingMove.CreatedAt = moveUpdate.CreatedAt
	existingMove.UpdatedAt = time.Now().Format(time.RFC3339)
	existingMove.Description = moveUpdate.Description
	existingMove.Seconds = moveUpdate.Seconds

	if err := this.db.Save(&existingMove).Error; err != nil {
		return nil, err
	}

	// delete all pairings in joint table associated with the move id
	if err := this.db.Where("move_id = ?", existingMove.Id).Delete(&Dto.MovesAction{}).Error; err != nil {
		return nil, err
	}

	//use the input's list of loopable actions
	for _, actionLoop := range moveUpdate.ActionLoops {

		//create a new MovesAction object for the junction table move_action
		newMovesAction := Dto.MovesAction{
			MoveId:   existingMove.Id,
			ActionId: actionLoop.ActionId,
			Loops:    actionLoop.Loops,
		}

        existingAction := Dto.Action{}
        if err := this.db.Where("id = ?", actionLoop.ActionId).First(&existingAction).Error; err != nil {
            if err == gorm.ErrRecordNotFound {
                log.Fatalf("Action with id {%d} does not exist", actionLoop.ActionId)
                return nil, err
            }
            return nil, err
        }


		// create a new move_action
		if err := this.db.Create(&newMovesAction).Error; err != nil {
			return nil, err
		}



		loopableAction := Dto.LoopableAction{
			Action: existingAction,
			Loops:  actionLoop.Loops,
		}

		loopableActions = append(loopableActions, loopableAction)
	}

	existingMove.Actions = loopableActions
	moves = append(moves, existingMove)

	return moves, nil
}

func (this *MoveRepository) HideMove(id string) ([]Dto.Move, error) {
	var existingMove Dto.Move
	var moves []Dto.Move
	var joinMovesAction []Dto.JoinMovesAction

	if err := this.db.Where("id = ?", id).First(&existingMove).Error; err != nil {
		return nil, err
	}

	existingMove.IsHidden = true

	if err := this.db.Save(&existingMove).Error; err != nil {
		return nil, err
	}

	if err := this.db.Debug().
		Select(joinFields).
		Table("moves").
		Joins("LEFT OUTER JOIN moves_actions ON moves.id = moves_actions.move_id").
		Joins("LEFT OUTER JOIN actions ON moves_actions.action_id = actions.id").
		Where("moves.id = ?", id).
		Scan(&joinMovesAction).Error; err != nil {
		return nil, err
	}

	moves = generateMoves(joinMovesAction)
	return moves, nil

}

func generateMoves(joinMovesAction []Dto.JoinMovesAction) []Dto.Move {
	var moves []Dto.Move
	loopMap := make(map[uint][]Dto.LoopableAction)
	moveMap := make(map[uint]Dto.Move)
	for _, jma := range joinMovesAction {

		action := Dto.Action{
			Id:          jma.AId,
			Name:        jma.AName,
			CreatedAt:   jma.ACreatedAt,
			UpdatedAt:   jma.AUpdatedAt,
			IsHidden:    jma.AIsHidden,
			Description: jma.ADescription,
			Seconds:     jma.ASeconds,
			Token:       jma.AToken,
		}

		loopAction := Dto.LoopableAction{
			Loops:  jma.Loops,
			Action: action,
		}

		move := Dto.Move{
			Id:          jma.MId,
			Name:        jma.MName,
			CreatedAt:   jma.MCreatedAt,
			UpdatedAt:   jma.MUpdatedAt,
			IsHidden:    jma.MIsHidden,
			Description: jma.MDescription,
			Seconds:     jma.MSeconds,
		}

		loopMap[jma.MId] = append(loopMap[jma.MId], loopAction)
		moveMap[jma.MId] = move
	}

	for key, value := range moveMap {
		value.Actions = loopMap[key]
		moves = append(moves, value)
	}
	return moves
}
