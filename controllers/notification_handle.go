package controllers

import (
	"CeltinelaBack/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type NotificationController struct {
	NotificationService *models.NotificationService
}

func (nc NotificationController) Notification(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	alias, err := nc.NotificationService.Notification(token)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonAlias, err := json.Marshal(alias)

	if err != nil {
		fmt.Println(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonAlias)
}
