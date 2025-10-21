package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Pendaftar struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	Matpel string `json:"matpel"`
	Durasi string `json:"durasi"`
	Jadwal string `json:"jadwal"`
	NoHP   string `json:"noHp"`
}

var registrations []Pendaftar
var db *sql.DB
var err error

func initDB() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/project_web?charset=utf8mb4&parseTime=True&loc=Local"
	// fmt.Println("dsn", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Gagal ambil koneksi SQL: %v", err)
	}

	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("❌ Database belum bisa diakses: %v", err)
	}

	// Auto migrate tabel dari struct
	if err = db.AutoMigrate(&Pendaftar{}); err != nil {
		log.Fatalf("❌ Gagal migrate: %v", err)
	}

	fmt.Println("✅ Berhasil terkoneksi ke database MySQL!")
	return db
}

func main() {
	// 🔗 Koneksi ke MySQL
	initDB()

	r := gin.Default()

	// Serve static frontend
	r.Static("/static", "./frontend")

	// API: menerima pendaftaran
	r.POST("/daftar", func(c *gin.Context) {
		var reg Pendaftar
		if err := c.BindJSON(&reg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		// reg.Nama = c.PostForm("name")
		// reg.Email = c.PostForm("email")
		// reg.Matpel = c.PostForm("matpel")
		// reg.Jadwal = c.PostForm("jadwal")
		// reg.NoHP = c.PostForm("noHp")

		fmt.Printf("📩 Data diterima: %+v\n", reg)

		registrations = append(registrations, reg)
		c.JSON(http.StatusOK, gin.H{
			"message": "Data pendaftar berhasil diregistrasi!",
			"data":    reg,
		})
	})

	// API: lihat semua pendaftar
	r.GET("/pendaftar", func(c *gin.Context) {
		c.JSON(http.StatusOK, registrations)
	})

	// Fallback ke index.html
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	r.Run(":8080")
}

// func RegisterHandler(w http.ResponseWriter, r *http.Request) {
// 	var data struct {
// 		Name   string `json:"name"`
// 		Email  string `json:"email"`
// 		Matpel string `json:"matpel"`
// 		Jadwal string `json:"jadwal"`
// 		NoHp   string `json:"noHp"`
// 	}
// 	json.NewDecoder(r.Body).Decode(&data)
// 	// Cek apakah email sudah ada
// 	var existing string
// 	err := db.QueryRow("SELECT email FROM users WHERE email = ?", data.Email).Scan(&existing)
// 	if err == nil {
// 		http.Error(w, `{"message":"Email sudah terdaftar"}`, http.StatusBadRequest)
// 		return
// 	}
// 	// Jika belum ada, baru insert
// 	_, err = db.Exec("INSERT INTO users (name,email,matpel,jadwal,no_hp) VALUES (?,?,?,?,?)",
// 		data.Name, data.Email, data.Matpel, data.Jadwal, data.NoHp)
// 	if err != nil {
// 		http.Error(w, `{"message":"Gagal menyimpan data"}`, http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(`{"message":"Berhasil mendaftar"}`))
// }
