package services

import (
	"gorm-tutorial/src/models"
)

// ActorService ...
type ActorService struct {
	*BaseService
}

// NewActorService ...
func NewActorService(path string) *ActorService {
	return &ActorService{NewBaseService(path)}
}

// GetActors ...
func (s *ActorService) GetActors() *[]models.Actor {

	db := s.GormOpen()
	defer db.Close()

	actors := []models.Actor{}
	db.Find(&actors)

	return &actors
}

// GetActor ...
func (s *ActorService) GetActor(actorID int) *models.Actor {

	db := s.GormOpen()
	defer db.Close()

	actor := models.Actor{}
	actor.ActorID = actorID

	db.First(&actor)

	return &actor
}

// CreateActor ...
func (s *ActorService) CreateActor(firstName string, lastName string) *models.Actor {

	db := s.GormOpen()
	defer db.Close()

	actor := models.Actor{}
	actor.FirstName = firstName
	actor.LastName = lastName

	db.Create(&actor)

	return &actor
}

// UpdateActor ...
func (s *ActorService) UpdateActor(actorID int, firstName string, lastName string) models.Actor {

	db := s.GormOpen()
	defer db.Close()

	actor := models.Actor{}
	actor.ActorID = actorID
	db.Find(&actor)

	actor.FirstName = firstName
	actor.LastName = lastName

	db.Save(&actor)

	return actor
}

// DeleteActor ...
func (s *ActorService) DeleteActor(actorID int) {

	db := s.GormOpen()
	defer db.Close()

	actor := models.Actor{}
	actor.ActorID = actorID
	db.Find(&actor)

	db.Delete(&actor)
}
