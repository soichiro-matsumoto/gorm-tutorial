package services

import (
	"gorm-tutorial/src/config"

	"github.com/jinzhu/gorm"
)

// BaseService ...
type BaseService struct {
	config *config.DbConfig
}

// NewBaseService ...
func NewBaseService(path string) *BaseService {
	return &BaseService{
		config: config.NewDbConfig(path),
	}
}

// GormOpen ...
func (s *BaseService) GormOpen() (db *gorm.DB) {
	db, err := s.config.GormOpen()
	if err != nil {
		panic(err)
	}
	return db
}
