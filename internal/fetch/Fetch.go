package fetch

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	HTTPStatus int `json:"httpStatus"`
	Response   struct {
		Items []struct {
			Type         string  `json:"type"`
			HarvesterID  string  `json:"harvesterId"`
			CerebroScore float64 `json:"cerebro-score"`
			URL          string  `json:"url"`
			Title        string  `json:"title"`
			CleanImage   string  `json:"cleanImage"`
		} `json:"items"`
	} `json:"response"`
}

type ContentMarketing struct {
	HTTPStatus int `json:"httpStatus"`
	Response   struct {
		Items []struct {
			Type              string  `json:"type"`
			HarvesterID       string  `json:"harvesterId"`
			CommercialPartner string  `json:"commercialPartner"`
			LogoURL           string  `json:"logoURL"`
			CerebroScore      float64 `json:"cerebro-score"`
			URL               string  `json:"url"`
			Title             string  `json:"title"`
			CleanImage        string  `json:"cleanImage"`
		} `json:"items"`
	} `json:"response"`
}

func FetchArticle(c *gin.Context) (Article, error) {
	resp, err := http.Get("https://storage.googleapis.com/aller-structure-task/articles.json")
	if err != nil {
		log.Fatalln(err)
	}
	var body Article
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed reading body"})
		return body, err
	}

	return body, err
}

func FetchContentMarketing(c *gin.Context) (ContentMarketing, error) {
	resp, err := http.Get("https://storage.googleapis.com/aller-structure-task/contentmarketing.json")
	if err != nil {
		log.Fatalln(err)
	}

	var body ContentMarketing
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed reading body"})
		return body, err
	}

	return body, err
}
