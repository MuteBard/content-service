package mirrorMove

type Action struct {
    ID        int    `gorm:"column:id"`
    Name      string `gorm:"column:name"`
    CreatedAt string `gorm:"column:created_at"`
    Token     string `gorm:"column:token"`
}