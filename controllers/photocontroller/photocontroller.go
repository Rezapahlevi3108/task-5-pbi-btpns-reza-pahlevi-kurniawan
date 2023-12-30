package photocontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Rezapahlevi3108/task-5-pbi-btpns-reza-pahlevi-kurniawan/helper"
	"github.com/Rezapahlevi3108/task-5-pbi-btpns-reza-pahlevi-kurniawan/models"
)

type Photo struct {
	models.Photo
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func CreatePhoto(w http.ResponseWriter, r *http.Request) {
	var photoInput models.Photo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&photoInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if err := models.DB.Create(&photoInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func GetPhotos(w http.ResponseWriter, r *http.Request) {
	var photos []Photo
	if err := models.DB.Find(&photos).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	helper.ResponseJSON(w, http.StatusOK, photos)
}

func UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	photoID, err := strconv.ParseInt(params["photoId"], 10, 64)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var updatedPhoto models.Photo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedPhoto); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if err := models.DB.Model(&models.Photo{}).Where("id = ?", photoID).Updates(&updatedPhoto).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func DeletePhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	photoID, err := strconv.ParseInt(params["photoId"], 10, 64)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	if err := models.DB.Delete(&models.Photo{}, photoID).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}