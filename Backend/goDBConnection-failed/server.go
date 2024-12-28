package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/* type User struct {
	ID   int    `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
} */

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("12 - Pong")
}

func ConnectDB() *gorm.DB {

	dsn := "host=localhost user=postgres password=sa dbname=learninggo port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("System not being ok %v", err)
	}
	log.Println("Database connection established")

	return db

}

func addUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User

		// Parse JSON body
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Insert user into the database
		if err := db.Create(&user).Error; err != nil {
			log.Printf("Failed to insert user: %v", err)
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		// Return success response
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("User %s created successfully", user.Name),
		})
	}
}

func main() {
	router := mux.NewRouter()

	//db connection

  db := ConnectDB()



  MigrateUser(db)


	router.HandleFunc("/ping", pingHandler).Methods("GET")


	router.HandleFunc("/addUser", addUser(db)).Methods("POST")


	log.Println("Server started on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}

}
