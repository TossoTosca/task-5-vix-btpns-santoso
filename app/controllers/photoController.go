package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/username/repo/models"
)

// GetPhotos mengembalikan semua data foto dari database
func GetPhotos(w http.ResponseWriter, r *http.Request) {
	photos, err := models.GetAllPhotos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(photos)
}

// GetPhoto mengembalikan data foto dengan ID tertentu dari database
func GetPhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	photo, err := models.GetPhotoByID(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(photo)
}

// CreatePhoto membuat data foto baru di database
func CreatePhoto(w http.ResponseWriter, r *http.Request) {
	var photo models.Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = models.CreatePhoto(&photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(photo)
}

// UpdatePhoto memperbarui data foto dengan ID tertentu di database
func UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	photo, err := models.GetPhotoByID(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = models.UpdatePhoto(&photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(photo)
}

// DeletePhoto menghapus data foto dengan ID tertentu dari database
func DeletePhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := models.DeletePhoto(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
