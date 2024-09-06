package main

import (
	"golang/bmysql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   string `json:"data"`
}

func main() {
	ginApp := gin.Default()
	bmysql.InitializeConnection()
	ginApp.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})

	ginApp.GET("/albums/all", getAllAlbums)
	ginApp.DELETE("/albums/:id", deleteAlbum)
	ginApp.POST("/album", createAlbum)
	ginApp.PUT("/album/:id", updateAlbum)
	// listen on port:8080
	ginApp.Run()
}

func getAllAlbums(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, bmysql.GetAllAlbums())
}

func deleteAlbum(ctx *gin.Context) {
	var resp response
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		resp = response{
			Code:   http.StatusInternalServerError,
			Status: "FAILED",
			Data:   err.Error()}
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	rows := bmysql.DeleteAlbum(int(id))
	if rows == 0 {
		resp = response{
			Code:   http.StatusNotFound,
			Status: "FAILED",
			Data:   "No record found"}
		ctx.JSON(http.StatusNotFound, resp)
		return
	}

	resp = response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Deleted Successfully"}
	ctx.JSON(http.StatusOK, resp)
}

func createAlbum(ctx *gin.Context) {
	var resp response
	var album bmysql.Album

	err := ctx.ShouldBindJSON(&album)
	if err != nil {
		resp = response{
			Code:   http.StatusInternalServerError,
			Status: "FAILED",
			Data:   err.Error()}
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	rows, err := bmysql.InsertAlbum(album)
	if err != nil && rows == 0 {
		resp = response{
			Code:   http.StatusInternalServerError,
			Status: "FAILED",
			Data:   err.Error()}
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Created Successfully"}
	ctx.JSON(http.StatusOK, resp)
}

func updateAlbum(ctx *gin.Context) {
	var resp response
	var album bmysql.Album
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		resp = response{
			Code:   http.StatusInternalServerError,
			Status: "FAILED",
			Data:   err.Error()}
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	err = ctx.ShouldBindJSON(&album)
	if err != nil {
		resp = response{
			Code:   http.StatusInternalServerError,
			Status: "FAILED",
			Data:   err.Error()}
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	rows, err := bmysql.UpdateAlbum(id,album)
	if err != nil && rows == 0 {
		resp = response{
			Code:   http.StatusInternalServerError,
			Status: "FAILED",
			Data:   err.Error()}
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp = response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Update Successfully"}
	ctx.JSON(http.StatusOK, resp)
}
