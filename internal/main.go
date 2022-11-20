package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hsUlberg/GoAllerMedia/internal/fetch"
	"github.com/hsUlberg/GoAllerMedia/internal/utils"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		articles, err := fetch.FetchArticle(c)
		if err != nil {
			return
		}
		ContentMarketing, err := fetch.FetchContentMarketing(c)
		if err != nil {
			return
		}

		result := utils.SortArticles(articles, ContentMarketing)
		c.JSON(http.StatusOK, gin.H{"Response": result})

	})

	r.Run(":9090")
}
