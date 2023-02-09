package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emilgibi/products-microservices/models"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (handler *Handler) Connect(host, user, pass, dbName, port string) {
	var err error
	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"
	handler.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	handler.DB.AutoMigrate(models.Product{})
	if err != nil {
		panic(err)
	}

}

func (handler *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	handler.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}

func (handler *Handler) Getuser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var products models.Product
	handler.DB.First(&products, params["id"])
	json.NewEncoder(w).Encode(products)
}
