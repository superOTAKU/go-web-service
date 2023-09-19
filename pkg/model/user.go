package model

type User struct {
	Id   int64 `gorm:"primaryKey;autoIncrement"`
	Name string
}

func (u *User) TableName() string {
	return "user"
}
