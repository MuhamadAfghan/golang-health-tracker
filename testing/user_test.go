package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"go-health-tracker/controllers"
	"go-health-tracker/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// Test_Store menguji fungsi Store
func Test_Store(t *testing.T) {
	// Cetak direktori kerja saat ini untuk debugging
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}
	log.Printf("Current working directory: %s", dir)

	// Muat variabel lingkungan dari file .env yang berada di direktori parent
	err = godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Mengatur Gin ke mode testing
	gin.SetMode(gin.TestMode)
	// Membuat instance router Gin
	router := gin.Default()
	// Mendefinisikan route POST /submit yang akan memanggil fungsi Store
	router.POST("/submit", controllers.Store)
	// Mendefinisikan Data Source Name (DSN) untuk koneksi database
	dsn := os.Getenv("DNS_MYSQL")
	// Menginisialisasi koneksi ke database menggunakan DSN
	models.InitDB(dsn)

	// Membuat instance url.Values untuk menyimpan data form
	form := url.Values{}
	form.Add("email", "test@example.com")
	form.Add("full_name", "John Doe")
	form.Add("height", "170")
	form.Add("weight", "70")
	form.Add("birth_date", "2000-01-01")

	// Membuat HTTP request baru dengan metode POST ke endpoint /submit dan mengirim data form yang telah di-encode
	req, err := http.NewRequest(http.MethodPost, "/submit", strings.NewReader(form.Encode()))
	// Menangani error jika request tidak dapat dibuat
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	// Mengatur header Content-Type menjadi application/x-www-form-urlencoded
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// Membuat response recorder untuk merekam respons dari server
	w := httptest.NewRecorder()
	// Menjalankan request menggunakan router Gin
	router.ServeHTTP(w, req)
	// Memastikan bahwa status code respons adalah 200 OK
	assert.Equal(t, http.StatusOK, w.Code)
	// Memastikan bahwa respons berisi pesan "Data saved successfully"
	assert.Contains(t, w.Body.String(), "Data saved successfully")
	// Memastikan bahwa respons berisi kategori kesehatan
	assert.Contains(t, w.Body.String(), "health_category")
}
