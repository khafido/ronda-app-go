package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/khafido/ronda-app-api/models"
)

var users = []models.User{
	models.User{
		Nama: "User 1",
		Password: "password1",
    	Username: "user1",
		Foto: "https://blog.golang.org/lib/godoc/images/go-logo-blue.svg",
	},
	models.User{
		Nama: "User 2",
		Password: "password2",
    	Username: "user2",
		Foto: "https://blog.golang.org/lib/godoc/images/go-logo-blue.svg",
	},
	models.User{
		Nama: "khafido",
		Password: "khaf123",
    	Username: "khaf",
		Foto: "https://blog.golang.org/lib/godoc/images/go-logo-blue.svg",
	},
}

var products = []models.Product{
	models.Product{
		ID_Produk: 1,
		Nama_Produk: "Produk 1",
		Kode_Produk: "P1",
		Harga_Produk: 10000,
    	Foto_Produk: "https://blog.golang.org/lib/godoc/images/go-logo-blue.svg",
	},
	models.Product{
		ID_Produk: 2,
		Nama_Produk: "Produk 2",
		Kode_Produk: "P2",
		Harga_Produk: 20000,
    	Foto_Produk: "https://blog.golang.org/lib/godoc/images/go-logo-blue.svg",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.User{}, &models.Product{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Product{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	// USERS
	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	// PRODUCTS
	for i, _ := range products {
		err = db.Debug().Model(&models.User{}).Create(&products[i]).Error
		if err != nil {
			log.Fatalf("cannot seed table: %v", err)
		}		
	}
}
