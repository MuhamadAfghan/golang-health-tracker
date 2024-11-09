// Package models menyediakan model data dan interaksi database untuk aplikasi
package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// Variabel global untuk koneksi database
var db *gorm.DB

// User merepresentasikan model pengguna dalam database
// Meng-extend gorm.Model yang menyediakan field ID, CreatedAt, UpdatedAt, dan DeletedAt
type User struct {
	gorm.Model
	// Email adalah alamat email pengguna (wajib diisi, maksimal 100 karakter)
	Email string `gorm:"type:varchar(100);not null"`
	// FullName adalah nama lengkap pengguna (wajib diisi, maksimal 100 karakter)
	FullName string `gorm:"type:varchar(100);not null"`
	// Height adalah tinggi badan pengguna dalam sentimeter (wajib diisi)
	Height float64 `gorm:"not null"`
	// Weight adalah berat badan pengguna dalam kilogram (wajib diisi)
	Weight float64 `gorm:"not null"`
	// BirthDate adalah tanggal lahir pengguna dalam format YYYY-MM-DD (wajib diisi)
	BirthDate string `gorm:"type:date;not null"`
	// HealthCategory merepresentasikan kategori BMI pengguna (wajib diisi, maksimal 50 karakter)
	HealthCategory string `gorm:"type:varchar(50);not null"`
}

// InitDB menginisialisasi koneksi database dan melakukan auto-migrasi
// Parameter:
//   - dsn: String koneksi database dalam format "user:pass@tcp(host:port)/dbname"
//
// Fungsi ini akan menampilkan log.Fatal jika terjadi kesalahan koneksi atau migrasi
func InitDB(dsn string) {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate model User
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
}

// GetDB mengembalikan instance koneksi database global
// Mengembalikan *gorm.DB yang dapat digunakan untuk operasi database
func GetDB() *gorm.DB {
	return db
}
