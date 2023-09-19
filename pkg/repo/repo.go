package repo

import "gorm.io/gorm"

func InitRepos(db *gorm.DB) {
	InitUserRepo(db)
}
