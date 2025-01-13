package mirrorMove

import (
    "github.com/jinzhu/gorm"
    Dto "mirrorMove/src/dtos"
)

type ActionRepository struct {
    db *gorm.DB
}


func NewActionRepository(db *gorm.DB) *ActionRepository {
    return &ActionRepository{db: db}
}

func (r *ActionRepository) GetAllActions() ([]Dto.Action, error) {
    var actions []Dto.Action

    if err := r.db.Select("id, name, created_at, token").Find(&actions).Error; err != nil {
        return nil, err
    }

    return actions, nil
}