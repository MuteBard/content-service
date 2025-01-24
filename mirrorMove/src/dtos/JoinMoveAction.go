package mirrorMove

type JoinMovesAction struct {
	MId          uint    `gorm:"column:id"`
	MName        string  `gorm:"column:name"`
	MCreatedAt   string  `gorm:"column:created_at"`
	MUpdatedAt   string  `gorm:"column:updated_at"`
	MIsHidden    bool    `gorm:"column:is_hidden"`
	MDescription string  `gorm:"column:description_"`
	MSeconds     float64 `gorm:"column:seconds"`
	AId          uint    `gorm:"column:id"`
	AName        string  `gorm:"column:name"`
	ACreatedAt   string  `gorm:"column:created_at"`
	AUpdatedAt   string  `gorm:"column:updated_at"`
	AIsHidden    bool    `gorm:"column:is_hidden"`
	ADescription string  `gorm:"column:description_"`
	ASeconds     float64 `gorm:"column:seconds"`
	AToken       string  `gorm:"column:token"`
	Loops        int     `gorm:"column:loops"`
}
