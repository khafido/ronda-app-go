package models

import (
	"errors"
	"html"
	"strings"
	_"log"
	_"fmt"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator/v10"
	"github.com/khafido/ronda-app-api/utils"
)

type Product struct {
	ID_Produk uint32 `gorm:"primary_key;auto_increment" json:"id_produk"`
	Nama_Produk string `gorm:"size:255;not null;unique" json:"nama_produk" validate:"required"`
	Kode_Produk string `gorm:"size:255;not null;unique" json:"kode_produk" validate:"required"`
	Harga_Produk uint64 `gorm:"size:11;not null;" json:"harga_produk" validate:"required"`
	Foto_Produk string `gorm:"size:150;not null;" json:"foto_produk"`	
}

func (p *Product) Prepare() {	
	p.Nama_Produk = html.EscapeString(strings.TrimSpace(p.Nama_Produk))	
	p.Kode_Produk = html.EscapeString(strings.TrimSpace(p.Kode_Produk))	
	p.Foto_Produk = html.EscapeString(strings.TrimSpace(p.Foto_Produk))
}

func (p *Product) ValidateStruct() map[string]string {
    validate := validator.New()
    err := validate.Struct(p)
    if err != nil {
    	return utils.ParseValidator(err.(validator.ValidationErrors))
    }
    return nil
}

func (p *Product) SaveProduct(db *gorm.DB) (*Product, error) {
	var err error
	err = db.Debug().Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) FindAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	return &products, err
}

func (p *Product) FindProductByID(db *gorm.DB, pid uint32) (*Product, error) {	
	var err error
	err = db.Debug().Model(Product{}).Where("id_produk = ?", pid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Product{}, errors.New("Product Not Found")
	}

	return p, err
}

func (p *Product) UpdateProduct(db *gorm.DB, pid uint32) (*Product, error) {
	db = db.Debug().Model(&Product{}).Where("id_produk = ?", pid).Take(&Product{}).UpdateColumns(
		map[string]interface{}{
			"nama_produk": p.Nama_Produk,
			"kode_produk": p.Kode_Produk,
			"harga_produk": p.Harga_Produk,
			"foto_produk": p.Foto_Produk,
		},
	)

	if db.Error != nil {
		return &Product{}, db.Error
	}

	err := db.Debug().Model(&Product{}).Where("id_produk = ?", pid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}

	return p, nil
}

func (p *Product) DeleteProduct(db *gorm.DB, pid uint32) (int64, error) {
	db = db.Debug().Model(&Product{}).Where("id_produk = ?", pid).Take(&Product{}).Delete(&Product{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
