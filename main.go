// Package main adalah package utama yang menjalankan aplikasi web Health Tracker
package main

// Import package yang dibutuhkan:
// - gin: framework web untuk Go
// - controllers: package yang berisi handler HTTP
// - models: package yang berisi model data dan koneksi database
// - log: package untuk logging
import (
	"go-health-tracker/controllers"
	"go-health-tracker/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// fungsi main adalah entry point aplikasi yang melakukan:
// 1. Inisialisasi koneksi database MySQL menggunakan DSN (Data Source Name)
// 2. Setup router Gin dan template HTML
// 3. Pendefinisian endpoint API
// 4. Menjalankan server web pada port 8080
func main() {
	// Load variabel lingkungan dari file .env
	// Jika terjadi error, log error dan hentikan aplikasi
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// DSN untuk koneksi ke database MySQL:
	// - username: root
	// - password: kosong
	// - host: localhost
	// - port: 3306
	// - database: golang_digiup_ujikom
	// - charset: utf8mb4
	// - parseTime: true untuk parsing waktu otomatis
	// - timezone: Asia/Jakarta
	dsn := os.Getenv("DNS_MYSQL")
	models.InitDB(dsn)

	// Inisialisasi router Gin dengan konfigurasi default
	r := gin.Default()
	// Load semua file template HTML dari folder views
	r.LoadHTMLGlob("views/*")

	// Definisi route:
	// - GET /: menampilkan form input data
	// - POST /submit: memproses data yang dikirim dari form
	r.GET("/", controllers.Create)
	r.POST("/submit", controllers.Store)

	// Jalankan server pada port 8080
	// Jika terjadi error, log error dan hentikan aplikasi
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
