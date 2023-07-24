package controllers

import (
	"encoding/json"
	"golang-rest-api/domain"
	"golang-rest-api/infra/database"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []domain.Personality
	database.DB.Find(&personalities)

	json.NewEncoder(w).Encode(personalities)
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err.Error())
	}

	var personality domain.Personality
	database.DB.First(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality domain.Personality
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Create(&personality)
	json.NewEncoder(w).Encode(personality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err.Error())
	}

	var personality domain.Personality
	database.DB.Delete(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err.Error())
	}

	var personality domain.Personality
	database.DB.First(&personality, id)

	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Save(&personality)

	json.NewEncoder(w).Encode(personality)
}
