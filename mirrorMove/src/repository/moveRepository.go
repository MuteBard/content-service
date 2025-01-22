package mirrorMove

import (
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

