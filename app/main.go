package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	tmpl = template.Must(template.ParseFiles("index.html"))
)

type Message struct {
	gorm.Model
	Content string `gorm:"not null"`
}

func initDB() {
	dsn := os.Getenv("DATABASE_URL")
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error", err)
	}
	db.AutoMigrate(&Message{})
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		msgContent := r.FormValue("message")
		if msgContent != "" {
			db.Create(&Message{Content: msgContent})
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var messages []Message
	db.Order("id desc").Find(&messages)
	tmpl.Execute(w, messages)
}

func main() {
	if _, err := os.Stat("index.html"); os.IsNotExist(err) {
		log.Fatal("error! there is no index.html was found!")
	}
	initDB()
	http.HandleFunc("/", handler)
	log.Println("server zapushyen na http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
