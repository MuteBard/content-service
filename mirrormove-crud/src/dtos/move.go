package mirrorMove

type Move struct {
    Id              uint     `gorm:"column:id"`
    Name            string   `gorm:"column:name"`
    CreatedAt       string   `gorm:"column:created_at"`
    UpdatedAt       string   `gorm:"column:updated_at"`
    IsHidden        bool     `gorm:"column:is_hidden"`
    Description     string   `gorm:"column:description_"`
    Seconds         float64  `gorm:"column:seconds"`
    Actions         []LoopableAction
}

