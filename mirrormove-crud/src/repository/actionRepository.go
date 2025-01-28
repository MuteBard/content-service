package mirrorMove

import (
    "time"
    "github.com/jinzhu/gorm"
    Dto "mirrorMove/src/dtos"
)

type ActionRepository struct {
    db *gorm.DB
}


func NewActionRepository(db *gorm.DB) *ActionRepository {
    return &ActionRepository{db: db}
}

var fields = "id, name, created_at, updated_at, description_, seconds, is_hidden, token"

func (this *ActionRepository) buildSearchQuery(args Dto.ActionApiArguments) *gorm.DB {
    query := this.db

    query = query.Where("is_hidden = ?", args.IsHidden)

    if (args.Name != "") {
        query = query.Where("name LIKE ?", "%" + args.Name + "%")
    }

    if (args.Name != "") {
        query = query.Where("description_ LIKE ?", "%" + args.Description + "%")
    }

    switch args.OrderBy {
    case "NAME":
        if args.SortOrder == Dto.ASC {
            query = query.Order("name ASC")
        } else if args.SortOrder == Dto.DESC {
            query = query.Order("name DESC")
        }
    case "SECONDS":
        if args.SortOrder == Dto.ASC {
            query = query.Order("seconds ASC")
        } else if args.SortOrder == Dto.DESC {
            query = query.Order("seconds DESC")
        }
    case "CREATEDAT":
        if args.SortOrder == Dto.ASC {
            query = query.Order("created_at ASC")
        } else if args.SortOrder == Dto.DESC {
            query = query.Order("created_at DESC")
        }
    default:
        query = query.Order("created_at DESC")
    }

    return query
}

func (this *ActionRepository) SearchActions(args Dto.ActionApiArguments) ([]Dto.Action, error) {
    var actions []Dto.Action

    query := this.buildSearchQuery(args)

    if err := query.Select(fields).Find(&actions).Error; err != nil {
        return nil, err
    }

    return actions, nil
}


func (this *ActionRepository) GetAction(id string) ([]Dto.Action, error) {
    var actions []Dto.Action

    if err := this.db.Where("id = ?", id).Find(&actions).Error; err != nil {
        return nil, err
    }

    return actions, nil
}

func (this *ActionRepository) CreateAction(action Dto.Action) ([]Dto.Action, error) {
    var actions []Dto.Action

    err := this.db.Create(&action).Error
    if err != nil {
        return nil, err 
    }
   
    actions = append(actions,action)
    
    return actions, nil
}

func (this *ActionRepository) UpdateAction(action Dto.Action) ([]Dto.Action, error) {
    var existingAction Dto.Action
    var actions []Dto.Action

    
    if err := this.db.Where("id = ?", action.Id).First(&existingAction).Error; err != nil {
        return nil, err
    }

    // Update the existing action with the new values
    existingAction.Name = action.Name
    existingAction.CreatedAt = action.CreatedAt
    existingAction.UpdatedAt = time.Now().Format(time.RFC3339)
    existingAction.Description = action.Description
    existingAction.Seconds = action.Seconds
    existingAction.Token = action.Token

    // Save the updated action back to the database
    if err := this.db.Save(&existingAction).Error; err != nil {
        return nil, err
    }

    actions = append(actions, existingAction)
    
    return actions, nil
}

func (this *ActionRepository) HideAction(id string) ([]Dto.Action, error) {
    var action Dto.Action
    var actions []Dto.Action

    if err := this.db.Where("id = ?", id).First(&action).Error; err != nil {
        return nil, err
    }

    action.IsHidden = true
    hiddenAction := action
    actions = append(actions, hiddenAction)

    if err := this.db.Save(&action).Error; err != nil {
        return nil, err
    }

    return actions, nil
}

