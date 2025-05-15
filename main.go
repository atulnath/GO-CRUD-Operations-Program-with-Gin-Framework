package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title Gin API
// @version 1.0

// @description This is a sample server.
func main() {
	myRouter := gin.Default()
	myRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Can you hear me?",
			"status":  "ok",
		})
	})

	// GET request with query parameters
	// Get used to get the data
	myRouter.GET("/me/:second_id/:first_id", func(c *gin.Context) {
		firstID := c.Param("first_id")
		secondID := c.Param("second_id")
		c.JSON(http.StatusOK, gin.H{
			"first_id":  firstID,
			"second_id": secondID,
		})
	})

	// POST request with JSON body
	// Post used to create a new user
	myRouter.POST("/userlogin", func(c *gin.Context) {

		//Saving the data: email and password to the database

		type UserData struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		//Let us create a new user
		var user UserData

		//error checking
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"email":    user.Email,
			"password": user.Password,
		})
	})

	// PUT request with query parameters
	// PUT used to update the data
	myRouter.PUT("/userlogin", func(c *gin.Context) {

		//Saving the data: email and password to the database

		type UserData struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		//Let us create a new user
		var user UserData

		//error checking
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"email":    user.Email,
			"password": user.Password,
		})
	})

	//PATCH request with query parameters
	// PATCH used to update the data, small part of the data
	// differecne between PUT and PATCH
	// PUT used to update the data
	// PATCH used to update the data, small part of the data
	myRouter.PATCH("/userlogin", func(c *gin.Context) {

		//Saving the data: email and password to the database

		type UserData struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		//Let us create a new user
		var user UserData

		//error checking
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"email":    user.Email,
			"password": user.Password,
		})
	})

	// DELETE request with query parameters
	// DELETE used to delete the data
	// in real world, we detete the data from the database by using the id
	myRouter.DELETE("/userlogin/:id", func(c *gin.Context) {
		var user_id = c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"id":    user_id,
			"Message": "User Id deleted",
		})
	})
	myRouter.Run(":8080") // listen and serve on
}
