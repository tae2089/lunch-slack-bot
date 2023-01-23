package config

import (
	"bc-labs-lunch-bot/ent/enttest"
	"context"
	"fmt"
	"github.com/go-playground/assert/v2"
	_ "github.com/mattn/go-sqlite3"
	"testing"
	"time"
)

func TestOpenDB(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	lunch, err := client.Lunch.Create().
		SetUserName("asd").
		SetCafeName("asd").
		SetRestaurantName("asd").
		SetPayer("asd").
		SetPaymentTime(time.Now()).
		Save(context.Background())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(lunch.CreatedAt)
	assert.Equal(t, lunch.CafeName, "asd")
}
