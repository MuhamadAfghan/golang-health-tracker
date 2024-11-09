// Package controllers berisi handler untuk mengelola request HTTP pada aplikasi Health Tracker
// File ini menangani:
// 1. Menampilkan form input data kesehatan
// 2. Memproses dan validasi data form
// 3. Menghitung BMI dan kategori kesehatan
// 4. Menyimpan data ke database
// 5. Mengirim response ke client

// Import package yang dibutuhkan:
// - net/http: untuk keperluan HTTP request/response
// - strconv: untuk konversi string ke tipe numerik
// - time: untuk manipulasi waktu dan tanggal
// - gin: framework web
// - models: package lokal berisi model data
package controllers

import (
	"net/http"
	"strconv"
	"time"

	"go-health-tracker/models"

	"github.com/gin-gonic/gin"
)

// Create adalah handler untuk menampilkan halaman form HTML
// Parameter:
// - c *gin.Context: konteks Gin yang berisi informasi request/response
// Response:
// - Menampilkan template create.html dengan status 200 OK
func Create(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", nil)
}

// SubmitForm adalah handler untuk memproses data form yang dikirim
// Alur proses:
// 1. Menerima data form (email, nama, tinggi, berat, tanggal lahir)
// 2. Validasi dan konversi tipe data
// 3. Hitung BMI dan tentukan kategori kesehatan
// 4. Simpan ke database
// 5. Kirim response JSON dengan status dan kategori kesehatan
//
// Parameter:
// - c *gin.Context: konteks Gin yang berisi informasi request/response
//
// Response:
// - Sukses: JSON dengan pesan sukses dan kategori kesehatan (status 200)
// - Error: JSON dengan pesan error (status 400/500)
func Store(c *gin.Context) {
	var user models.User

	// Ambil data dari form
	user.Email = c.PostForm("email")
	user.FullName = c.PostForm("full_name")

	heightStr := c.PostForm("height")

	// Parsing tinggi badan dari string ke float64
	height, err := strconv.ParseFloat(heightStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid height format"})
		return
	}
	user.Height = height

	// Parsing berat badan dari string ke float64
	weightStr := c.PostForm("weight")

	weight, err := strconv.ParseFloat(weightStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid weight format"})
		return
	}
	user.Weight = weight

	// Parsing tanggal lahir dari string ke time.Time
	birthDateStr := c.PostForm("birth_date")

	birthDate, err := time.Parse("2006-01-02", birthDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}
	user.BirthDate = birthDate.Format("2006-01-02")

	// Hitung kategori kesehatan berdasarkan BMI
	bmi := user.Weight / ((user.Height / 100) * (user.Height / 100))
	switch {
	case bmi < 17:
		user.HealthCategory = "Sangat Kurus"
	case bmi >= 17 && bmi < 18.5:
		user.HealthCategory = "Kurus"
	case bmi >= 18.5 && bmi < 25:
		user.HealthCategory = "Normal"
	case bmi >= 25 && bmi < 30:
		user.HealthCategory = "Gemuk"
	case bmi >= 30 && bmi < 35:
		user.HealthCategory = "Sangat Gemuk"
	default:
		user.HealthCategory = "Obesitas"
	}

	// Simpan data user ke database
	db := models.GetDB()
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kirim respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "Data saved successfully", "health_category": user.HealthCategory})
}
