package authcontroller

import (
	"time"
	"net/http"
	"encoding/json"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/Rezapahlevi3108/task-5-pbi-btpns-reza-pahlevi-kurniawan/helper"
	"github.com/Rezapahlevi3108/task-5-pbi-btpns-reza-pahlevi-kurniawan/models"
	"github.com/Rezapahlevi3108/task-5-pbi-btpns-reza-pahlevi-kurniawan/config"
)

type User struct {
	models.User
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	var user User
	if err := models.DB.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Email atau password salah"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Email atau password salah"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	createdTime, err := time.Parse("2006-01-02 15:04:05.000", user.CreatedAt)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	user.User.CreatedAt = createdTime

	updatedTime, err := time.Parse("2006-01-02 15:04:05.000", user.UpdatedAt)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	user.User.UpdatedAt = updatedTime

	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim {
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims {
			Issuer: "task5-pbi-btpns-reza-pahlevi-kurniawan",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	http.SetCookie(w, &http.Cookie {
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "Login berhasil"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie {
		Name: "token",
		Path: "/",
		Value: "",
		HttpOnly: true,
		MaxAge: -1,
	})

	response := map[string]string{"message": "Logout berhasil"}
	helper.ResponseJSON(w, http.StatusOK, response)
}