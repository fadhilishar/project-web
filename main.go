package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Registration struct {
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	Matpel string `json:"matpel"`
	Durasi string `json:"durasi"`
	Jadwal string `json:"jadwal"`
	NoHP   string `json:"noHp"`
}

var registrations []Registration

func main() {
	r := gin.Default()

	// Serve static frontend
	r.Static("/static", "./frontend")

	// API: menerima pendaftaran
	r.POST("/daftar", func(c *gin.Context) {
		var reg Registration
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
			"message": "Registration received!",
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
