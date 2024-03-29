package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aynp/urlshortner/db"
	"github.com/aynp/urlshortner/models"
	"github.com/aynp/urlshortner/utils"
	"github.com/gin-gonic/gin"
)

func CreateShortURL(c *gin.Context) {
	var createRequest models.CreateRequest
	c.BindJSON(&createRequest)

	shortURL := utils.GenerateRandomString()

	db.Redis.Set(context.Background(), shortURL, createRequest.OriginalURL, 0)

	// set URL params
	err := setURLParams(shortURL, &createRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// set Header params
	err = setHeaderParams(shortURL, &createRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	// set Auto Generated params
	err = setAutoGenParams(shortURL, &createRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// return the shortened URL
	c.JSON(http.StatusOK, gin.H{
		"short_url": shortURL,
	})
}

// setUrlParams sets the URL params for the short URL to the DB
func setURLParams(shortURL string, createRequest *models.CreateRequest) error {
	for _, urlParam := range createRequest.URLParams {
		x := db.Redis.SAdd(context.Background(), fmt.Sprintf("%s:URLParams", shortURL), models.URLParam{
			SourceParam: urlParam.SourceParam,
			TargetParam: urlParam.TargetParam,
			IsMandatory: urlParam.IsMandatory,
		})

		if x.Err() != nil {
			return x.Err()
		}
	}

	return nil
}

// setHeaderParams sets the Header params for the short URL to the DB
func setHeaderParams(shortURL string, createRequest *models.CreateRequest) error {
	for _, headerParam := range createRequest.HeaderParams {
		res := db.Redis.SAdd(context.Background(), fmt.Sprintf("%s:HeaderParams", shortURL), models.HeaderParam{
			SourceParam: headerParam.SourceParam,
			TargetParam: headerParam.TargetParam,
			IsMandatory: headerParam.IsMandatory,
		})

		if res.Err() != nil {
			return res.Err()
		}
	}

	return nil
}

// setAutoGenParams sets the Auto Generated params for the short URL to the DB
func setAutoGenParams(shortURL string, createRequest *models.CreateRequest) error {
	for _, autoGenParam := range createRequest.AutoGenParams {
		res := db.Redis.SAdd(context.Background(), fmt.Sprintf("%s:AutoGenParams", shortURL), models.AutoGenParam{
			Type:      autoGenParam.Type,
			TargetKey: autoGenParam.TargetKey,
		})

		if res.Err() != nil {
			return res.Err()
		}
	}

	return nil
}
