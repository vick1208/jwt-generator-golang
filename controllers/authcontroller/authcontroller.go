package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt/v4"
	"github.com/vick1208/jwt-go/config"
	"github.com/vick1208/jwt-go/helper"
	"github.com/vick1208/jwt-go/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userReg models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userReg); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	var user models.User
	if err := models.DB.Where("username = ?", userReg.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Username atau Password salah"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReg.Password)); err != nil {
		response := map[string]string{"message": "Username atau Password salah"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jwt-go",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokAlg := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokAlg.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})
	response := map[string]string{"message": "login berhasil"}
	helper.ResponseJSON(w, http.StatusOK, response)

}
func Register(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan json
	var userReg models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userReg); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(userReg.Password), bcrypt.DefaultCost)
	userReg.Password = string(hashPass)

	if err := models.DB.Create(&userReg).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
func Logout(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "logout berhasil"}
	helper.ResponseJSON(w, http.StatusOK, response)

}
