package mirrorMove

type Action struct {
    ID        uint   `gorm:"column:id"`
    Name      string `gorm:"column:name"`
    CreatedAt string `gorm:"column:created_at"`
    IsHidden  bool   `gorm:"column:is_hidden"`
    Token     string `gorm:"column:token"`
}