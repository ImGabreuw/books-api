package controllers

import (
	"books-api/database"
	"books-api/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ShowBook(ctx *gin.Context) {
	id := ctx.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(
			400,
			gin.H{"error": "ID has to be a integer"},
		)
		return
	}

	db := database.GetDataBase()

	var book model.Book

	err = db.First(&book, newId).Error

	if err != nil {
		ctx.JSON(
			400,
			gin.H{"error": "Cannot find book: " + err.Error()},
		)
		return
	}

	ctx.JSON(200, book)
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(
			400,
			gin.H{"error": "ID has to be a integer"},
		)
		return
	}

	db := database.GetDataBase()

	err = db.Delete(&model.Book{}, newId).Error

	if err != nil {
		ctx.JSON(
			400,
			gin.H{"error": "Cannot delete book: " + err.Error()},
		)
		return
	}

	ctx.Status(204)
}

func CreateBook(ctx *gin.Context) {
	db := database.GetDataBase()

	var book model.Book

	err := ctx.ShouldBindJSON(&book)

	if err != nil {
		ctx.JSON(
			400,
			gin.H{"error": "Cannot bind JSON: " + err.Error()},
		)
		return
	}

	err = db.Create(&book).Error

	if err != nil {
		ctx.JSON(
			400,
			gin.H{"error": "Cannot create book: " + err.Error()},
		)
		return
	}

	ctx.JSON(201, book)
}

func UpdateBook(ctx *gin.Context) {
	db := database.GetDataBase()

	var book model.Book

	err := ctx.ShouldBindJSON(&book)

	if err != nil {
		ctx.JSON(
			400,
			gin.H{"error": "Cannot bind JSON: " + err.Error()},
		)
		return
	}

	err = db.Save(&book).Error

	if err != nil {
		ctx.JSON(
			400,
			gin.H{"error": "Cannot update book: " + err.Error()},
		)
		return
	}

	ctx.JSON(200, book)
}

func ShowBooks(ctx *gin.Context) {
	db := database.GetDataBase()

	var books []model.Book

	err := db.Find(&books).Error

	if err != nil {
		ctx.JSON(
			400,
			gin.H{"error": "Cannot list books: " + err.Error()},
		)
		return
	}

	ctx.JSON(200, books)
}
