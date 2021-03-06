package courts

import (
	"../helpers"
	"github.com/gin-gonic/gin"
)

// SearchAction - query data from the instances and format them
func SearchAction(context *gin.Context) {
	var results []helpers.CourtItem
	responseCode := 200
	responses := helpers.GetAPIData(context.Param("searchTerm"))
	results = helpers.ParseResponses(responses)

	context.JSON(responseCode, results)
}

// IndexAction - application start page
func IndexAction(context *gin.Context) {
	context.HTML(200, "index.html", gin.H{
		"title": "Meta Search",
	})
}
