package handler

import (
	"encoding/json"
	"fmt"
	"go-run/model"
	"go-run/storage"
	"go-run/utils"
	"log"
	"net/http"
)

func (a *application) createUser(w http.ResponseWriter, r *http.Request) {
	var usr model.User
	if err := json.NewDecoder(r.Body).Decode(&usr); err != nil {
		utils.SendMessageAsJsonResponse(w, "Malformed JSON in body", http.StatusBadRequest)
		return
	}

	if err := usr.Validate(); err != nil {
		msg := fmt.Sprintf("Invalid user in body: %+v", err)
		utils.SendMessageAsJsonResponse(w, msg, http.StatusBadRequest)
		return
	}

	if err := a.DB.PutUser(usr); err != nil {
		if err != storage.ErrUserExistsAlready {
			utils.SendMessageAsJsonResponse(w, "User already exists with this email", http.StatusBadRequest)
		} else {
			utils.SendMessageAsJsonResponse(w, "Some error occurred while creating user", http.StatusInternalServerError)
			log.Printf("ERROR: Failed to create user in DB: %+v", err)
		}
		return
	}

	utils.SendMessageAsJsonResponse(w, "User successfully created", http.StatusCreated)
}
