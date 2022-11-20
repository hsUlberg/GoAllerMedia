package utils

import (
	"github.com/hsUlberg/GoAllerMedia/internal/fetch"
)

type result []interface{}

func SortArticles(articles fetch.Article, contentmarketing fetch.ContentMarketing) result {
	cmLength := len(contentmarketing.Response.Items)
	var result result
	cmIndex := 0
	for index, item := range articles.Response.Items {
		result = append(result, item)
		modulo := (index + 1) % 5
		if modulo == 0 {
			if cmIndex < cmLength {
				result = append(result, contentmarketing.Response.Items[cmIndex])
			} else {
				m := make(map[string]interface{})
				m["type"] = "Ad"
				result = append(result, m)
			}
			cmIndex += 1
		}
	}
	return result
}
