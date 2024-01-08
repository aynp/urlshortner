package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/aynp/urlshortner/db"
	"github.com/aynp/urlshortner/models"
	"github.com/aynp/urlshortner/utils"
	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	var err error

	path := c.Param("path")
	targetURL := db.Redis.Get(context.Background(), path)

	if targetURL.Err() != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
		})
		return
	}

	redirectURL, err := url.Parse(targetURL.Val())

	// Add URL Params
	err = addUrlParams(c, path, redirectURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	// Add Header Params
	err = addHeaderParams(c, path, redirectURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	// Add Auto Generated params
	err = addAutoGenParams(c, path, redirectURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	// Redirect
	c.Redirect(http.StatusPermanentRedirect, redirectURL.String())
}

// addUrlParams adds URL params to the redirect URL as set in DB
// It returns an error if a mandatory param is not present in the request
func addUrlParams(c *gin.Context, path string, redirectURL *url.URL) error {
	// Add URL Params
	urlParam := db.Redis.SMembers(context.Background(), fmt.Sprintf("%s:URLParams", path))

	if urlParam.Err() != nil {
		return urlParam.Err()
	}

	for _, urlParam := range urlParam.Val() {
		var unmarshalledURLParams models.URLParam
		err := json.Unmarshal([]byte(urlParam), &unmarshalledURLParams)
		if err != nil {
			return err
		}

		urlValue := c.Query(unmarshalledURLParams.SourceParam)

		if unmarshalledURLParams.IsMandatory && urlValue == "" {
			return fmt.Errorf("url param %s is mandatory", unmarshalledURLParams.SourceParam)
		}

		if urlValue != "" {
			redirectURL.Query().Add(unmarshalledURLParams.TargetParam, urlValue)
		}
	}

	return nil
}

// addHeaderParams adds header params to the redirect URL as set in DB
// It returns an error if a mandatory param is not present in the request
func addHeaderParams(c *gin.Context, path string, redirectURL *url.URL) error {
	headerParam := db.Redis.SMembers(context.Background(), fmt.Sprintf("%s:HeaderParams", path))

	if headerParam.Err() != nil {
		return headerParam.Err()
	}

	for _, urlParam := range headerParam.Val() {
		var unmarshalledHeaderParams models.HeaderParam
		err := json.Unmarshal([]byte(urlParam), &unmarshalledHeaderParams)
		if err != nil {
			return err
		}

		// get header value from request
		headerValue := c.GetHeader(unmarshalledHeaderParams.SourceParam)

		if unmarshalledHeaderParams.IsMandatory && headerValue == "" {
			return fmt.Errorf("header param %s is mandatory", unmarshalledHeaderParams.SourceParam)
		}

		if headerValue != "" {
			redirectURL.Query().Add(unmarshalledHeaderParams.TargetParam, headerValue)
		}
	}

	return nil
}

// addAutoGenParams adds either timestamp or uuid to the redirect URL
func addAutoGenParams(c *gin.Context, path string, redirectURL *url.URL) error {
	autoGenParam := db.Redis.SMembers(context.Background(), fmt.Sprintf("%s:AutoGenParams", path))

	if autoGenParam.Err() != nil {
		return autoGenParam.Err()
	}

	for _, urlParam := range autoGenParam.Val() {
		var unmarshalledAutoGenParams models.AutoGenParam
		err := json.Unmarshal([]byte(urlParam), &unmarshalledAutoGenParams)
		if err != nil {
			return err
		}

		if unmarshalledAutoGenParams.Type == "timestamp" {
			redirectURL.Query().Add(unmarshalledAutoGenParams.TargetKey, utils.TimestampString())
		}

		if unmarshalledAutoGenParams.Type == "uuid" {
			redirectURL.Query().Add(unmarshalledAutoGenParams.TargetKey, utils.UUIDString())
		}
	}

	return nil
}
