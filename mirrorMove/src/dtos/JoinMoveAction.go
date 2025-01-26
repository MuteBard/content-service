package mirrorMove

type JoinMovesAction struct {
    MId          uint    `gorm:"column:MId"`
    MName        string  `gorm:"column:MName"`
    MCreatedAt   string  `gorm:"column:MCreatedAt"`
    MUpdatedAt   string  `gorm:"column:MUpdatedAt"`
    MIsHidden    bool    `gorm:"column:MIsHidden"`
    MDescription string  `gorm:"column:MDescription"`
    MSeconds     float64 `gorm:"column:MSeconds"`
    AId          uint    `gorm:"column:AId"`
    AName        string  `gorm:"column:AName"`
    ACreatedAt   string  `gorm:"column:ACreatedAt"`
    AUpdatedAt   string  `gorm:"column:AUpdatedAt"`
    AIsHidden    bool    `gorm:"column:AIsHidden"`
    ADescription string  `gorm:"column:ADescription"`
    ASeconds     float64 `gorm:"column:ASeconds"`
    AToken       string  `gorm:"column:AToken"`
    Loops        int     `gorm:"column:Loops"`
}
