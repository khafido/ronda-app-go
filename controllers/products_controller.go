package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"os"
	"io"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	_"github.com/khafido/ronda-app-api/auth"
	"github.com/khafido/ronda-app-api/models"
	"github.com/khafido/ronda-app-api/responses"
	"github.com/khafido/ronda-app-api/utils/formaterror"
)

// API
func (server *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {	
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    if err := r.ParseMultipartForm(2048); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	now := time.Now()	
	filename := fmt.Sprintf("%s%s", now.Format(time.RFC3339), filepath.Ext(handler.Filename))
	fmt.Println(filename);

	fileLocation := filepath.Join(dir, "files", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	harga,_ := strconv.ParseUint(r.FormValue("harga_produk"), 10, 64)
	product := models.Product{
		Nama_Produk: r.FormValue("nama_produk"),
		Kode_Produk: r.FormValue("kode_produk"),
		Harga_Produk: harga,
		Foto_Produk: filename,
	}
	
	// fmt.Println(product)
	product.Prepare()
	if errValidator := product.ValidateStruct(); errValidator != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, errValidator)
		return
	}


	productCreated, err := product.SaveProduct(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, productCreated.ID_Produk))
	responses.JSON(w, http.StatusCreated, map[string]string{"status":"Success Created Product"})	
}

func (server *Server) GetProducts(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}
	products, err := product.FindAllProducts(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, products)
}

func (server *Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id_produk"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	product := models.Product{}
	productGotten, err := product.FindProductByID(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, productGotten)
}

func (server *Server) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id_produk"], 10, 64)

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    if err := r.ParseMultipartForm(2048); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	now := time.Now()	
	filename := fmt.Sprintf("%s%s", now.Format(time.RFC3339), filepath.Ext(handler.Filename))
	fmt.Println(filename);

	fileLocation := filepath.Join(dir, "files", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	harga,_ := strconv.ParseUint(r.FormValue("harga_produk"), 10, 64)
	productUpdate := models.Product{
		Nama_Produk: r.FormValue("nama_produk"),
		Kode_Produk: r.FormValue("kode_produk"),
		Harga_Produk: harga,
		Foto_Produk: filename,
	}

	productUpdate.Prepare()
	if errValidator := productUpdate.ValidateStruct(); errValidator != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, errValidator)
		return
	}

	productUpdated, err := productUpdate.UpdateProduct(server.DB, uint32(pid))
	if err != nil && productUpdated == nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusCreated, map[string]string{"status":"Success Updated Product"})
}

func (server *Server) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	product := models.Product{}

	pid, err := strconv.ParseUint(vars["id_produk"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = product.DeleteProduct(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusCreated, map[string]string{"status":"Success Deleted Product"})
}
