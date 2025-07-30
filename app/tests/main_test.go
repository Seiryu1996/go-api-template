package main

import (
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"gin-app/infra"
	"gin-app/models"
	"gin-app/router"

	"encoding/json"
	"net/http"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"bytes"
	"gin-app/dto"
	"gin-app/services"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalln("Error loading .env file")
	}
	code := m.Run()
	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	db.Exec("DELETE FROM items")
	db.Exec("DELETE FROM users")
	db.Exec("ALTER TABLE items AUTO_INCREMENT = 1")
	db.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
	items := []models.Item{
		{Name: "テスト1", Price: 1000, Description: "test1", SoldOut: false, UserID: 1},
		{Name: "テスト2", Price: 2000, Description: "test2", SoldOut: true, UserID: 1},
		{Name: "テスト3", Price: 3000, Description: "", SoldOut: false, UserID: 2},
	}
	users := []models.User{
		{Name: "User1", Email: "test1@example.com", Password: "user123"},
		{Name: "User2", Email: "test2@example.com", Password: "user456"},
	}
	for _, user := range users {
		db.Create(&user)
	}
	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{})
	setupTestData(db)
	router := router.SetupRouter(db)
	return router
}

func TestFindAll(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)

	router.ServeHTTP(w, req)

	var res map[string][]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["data"]))
}

func TestCreate(t *testing.T) {
	router := setup()
	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)

	createItemInput := dto.CreateItemInput{
		Name:        "test999",
		Price:       9999,
		Description: "ZZZZZ",
	}
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+*token)

	router.ServeHTTP(w, req)

	var res map[string]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "test999", res["data"].Name)
}

func TestCreateUnauthorized(t *testing.T) {
	router := setup()

	createItemInput := dto.CreateItemInput{
		Name:        "test999",
		Price:       9999,
		Description: "ZZZZZ",
	}
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))

	router.ServeHTTP(w, req)

	var res map[string]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
