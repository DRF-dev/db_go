package main

import (
	"context"
	"crud_golang/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

//User est la structure de nos utilisateurs
type User struct {
	Nom    string `json:"FirstName"`
	Prenom string `json:"LastName"`
}

func main() {
	client, cancel, err := models.ConnectDB()
	if err != nil {
		log.Fatalf("Erreur ligne 14 : %v\n", err)
	}
	defer cancel()
	collection := client.Database("dbgo").Collection("db_go")

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.File("./view/index.html")
	})
	api := router.Group("/api")
	{
		api.POST("/add", func(c *gin.Context) {
			elm := &User{
				Nom:    c.Request.PostFormValue("nom"),
				Prenom: c.Request.PostFormValue("prenom"),
			}
			j, err := bson.Marshal(elm)
			if err != nil {
				log.Fatalf("Erreur ligne 38 %v\n", err)
			}
			collection.InsertOne(context.Background(), j)
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
				"msg":    "Message enrengistré avec succès",
			})
		})
	}

	router.Run(":4000")
}
