package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Hirogava/ParkingDealer/internal/models/routresponse"
	"github.com/Hirogava/ParkingDealer/internal/repository/postgres/api"
	"github.com/Hirogava/ParkingDealer/internal/service/funcgraf"

	"github.com/gin-gonic/gin"
)

func InitParams(r *gin.Engine, manager *api.Manager) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/critical", func(ctx *gin.Context) {
			GetCritical(ctx, manager)
		})
	}
}

func GetCritical(ctx *gin.Context, manager *api.Manager) {
	var response map[string]interface{}

	err := ctx.BindJSON(&response)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error"})
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Fatal error"})
	}

	reader := bytes.NewReader(jsonData)

	resp, err := http.Post("http://routing.api.2gis.com/routing/7.0.0/global?key=52e0e0ae-e911-4c14-9359-e53ab161888a", "application/json", reader)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch from external API"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read API response"})
		return
	}

	var apiResponse *routresponse.RouteResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmparshal API response"})
		return
	}

	weatherResp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%g&lon=%g&appid=%s", apiResponse.Query.Points[0].Lat, apiResponse.Query.Points[0].Lon, os.Getenv("WEATHER_KEY")))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch from external API"})
		return
	}
	defer weatherResp.Body.Close()

	weatherBody, err := io.ReadAll(weatherResp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read API response"})
		return
	}

	var weatherApiResponse *routresponse.WeatherResponse
	if err := json.Unmarshal(weatherBody, &weatherApiResponse); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmparshal API response"})
		return
	}

	coordinatesResp, err := http.Get(fmt.Sprintf("https://catalog.api.2gis.com/3.0/items/geocode?lat=%g&lon=%g&fields=items.point&key=%s", apiResponse.Query.Points[0].Lat, apiResponse.Query.Points[0].Lon, os.Getenv("2GIS_KEY")))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch from external API"})
		return
	}
	defer coordinatesResp.Body.Close()

	coordinatesBody, err := io.ReadAll(coordinatesResp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read API response"})
		return
	}

	var coordinatesApiResponse *routresponse.LatLon
	if err := json.Unmarshal(coordinatesBody, &coordinatesApiResponse); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmparshal API response"})
		return
	}

	result, gl, err := manager.GetCriticalManeuvers(apiResponse, weatherApiResponse)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error"})
	}

	finalResult, err := funcgraf.ProcessRoutesFromAPI(result, gl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  finalResult.Status,
		"query":   finalResult.Query,
		"type":    finalResult.Type,
		"message": finalResult.Message,
		"result":  finalResult.Result,
	})
}
