package services_test

import (
	"fmt"
	"gorm-tutorial/src/services"
	"testing"
)

var config string = "../db_config.json"

func TestGetActor(t *testing.T) {
	as := services.NewActorService(config)
	a := as.GetActor(1)
	fmt.Println(a)
}

func TestGetActors(t *testing.T) {
	as := services.NewActorService(config)
	a := as.GetActors()
	fmt.Println(a)
}

func TestUpdateActor(t *testing.T) {
	as := services.NewActorService(config)
	a := as.UpdateActor(1, "埼玉のｲﾙﾏﾆｱ", "ｱｰｲ")
	fmt.Println(a)
}
