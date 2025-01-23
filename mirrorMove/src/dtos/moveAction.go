package mirrorMove

type MoveAction struct {
    Id         uint   `gorm:"column:id"`
    MoveId     uint   `gorm:"column:move_id"`
    ActionId   uint   `gorm:"column:action_id"`
    Loops      int    `gorm:"column:loops"`
}