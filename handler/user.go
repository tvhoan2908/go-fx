package handler

import (
	"log"

	"github.com/tvhoan2908/go-fx/db"
	"github.com/tvhoan2908/go-fx/model"
)

type UserHandler struct {
	DB *db.DB
}

func NewUserHandler(db *db.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) GetUsers() []model.User {
	var rows []model.User
	res, err := h.DB.Query("SELECT id, username FROM users")
	if err != nil {
		log.Fatal("Error ne", err)
		return nil
	}
	defer res.Close()
	for res.Next() {
		var element model.User
		err = res.Scan(&element.ID, &element.Username)
		log.Println("err", err)
		rows = append(rows, element)
	}
	return rows
}
