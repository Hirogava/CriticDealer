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
	"github.com/Hirogava/ParkingDealer/internal/config/logger" 

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

// GetCritical обрабатывает запрос на безопасный маршрут
// @Summary Получить безопасный маршрут с анализом аварийности
// @Description Запрашивает маршрут у 2GIS, получает погоду и рассчитывает опасность участков
// @Tags routes
// @Accept json
// @Produce json
// @Param request body map[string]interface{} true "JSON-запрос для API 2GIS"
// @Success 200 {object} routresponse.RouteResponse "Успешный анализ маршрута"
// @Failure 400 {string} string "Ошибка в данных запроса"
// @Failure 500 {string} string "Ошибка обработки маршрута или внешнего API"
// @Router /api/v1/critical [get]
func GetCritical(ctx *gin.Context, manager *api.Manager) {
	var response map[string]interface{}

	err := ctx.BindJSON(&response)
	if err != nil {
		logger.Logger.Info("Received request for critical route analysis")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error"})
	}

	logger.Logger.Debug("JSON request parsed successfully")
	jsonData, err := json.Marshal(response)
	if err != nil {
		logger.Logger.Error("Failed to marshal JSON data", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Fatal error"})
	}

	reader := bytes.NewReader(jsonData)

	logger.Logger.Info("Sending request to 2GIS Routing API")
	resp, err := http.Post("http://routing.api.2gis.com/routing/7.0.0/global?key=" + os.Getenv("2GIS_KEY"), "application/json", reader)
	if err != nil {
		logger.Logger.Error("Failed to fetch from 2GIS API", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch from external API"})
		return
	}
	defer resp.Body.Close()

	logger.Logger.Debug("2GIS API response status", "status_code", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Logger.Error("Failed to read 2GIS API response", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read API response"})
		return
	}

	var apiResponse *routresponse.RouteResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		logger.Logger.Error("Failed to unmarshal 2GIS API response", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmparshal API response"})
		return
	}
	logger.Logger.Info("2GIS route data received successfully")

	logger.Logger.Info("Fetching weather data")
	weatherResp, err := http.Get(fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?lat=%g&lon=%g&appid=%s&units=metric&lang=ru",
		apiResponse.Query.Points[0].Lat,
		apiResponse.Query.Points[0].Lon,
		os.Getenv("WEATHER_KEY"),
	))
	if err != nil {
		logger.Logger.Error("Failed to fetch weather data", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch from external API"})
		return
	}
	defer weatherResp.Body.Close()

	weatherBody, err := io.ReadAll(weatherResp.Body)
	if err != nil {
		logger.Logger.Error("Failed to read weather API response", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read API response"})
		return
	}

	var weatherApiResponse *routresponse.WeatherResponse
	if err := json.Unmarshal(weatherBody, &weatherApiResponse); err != nil {
		logger.Logger.Error("Failed to unmarshal weather response", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmparshal API response"})
		return
	}
	logger.Logger.Info("Weather data received successfully")

	logger.Logger.Info("Fetching coordinates data")
	coordinatesResp, err := http.Get(fmt.Sprintf("https://catalog.api.2gis.com/3.0/items/geocode?lat=%g&lon=%g&fields=items.point&key=%s", apiResponse.Query.Points[0].Lat, apiResponse.Query.Points[0].Lon, os.Getenv("2GIS_KEY")))
	if err != nil {
		logger.Logger.Error("Failed to fetch coordinates data", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch from external API"})
		return
	}
	defer coordinatesResp.Body.Close()

	coordinatesBody, err := io.ReadAll(coordinatesResp.Body)
	if err != nil {
		logger.Logger.Error("Failed to read coordinates API response", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read API response"})
		return
	}

	var coordinatesApiResponse *routresponse.LatLon
	if err := json.Unmarshal(coordinatesBody, &coordinatesApiResponse); err != nil {
		logger.Logger.Error("Failed to unmarshal coordinates response", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmparshal API response"})
		return
	}
	logger.Logger.Info("Coordinates data received successfully")

	logger.Logger.Info("Starting critical maneuvers analysis")
	result, gl, err := manager.GetCriticalManeuvers(apiResponse, weatherApiResponse)
	if err != nil {
		logger.Logger.Error("Critical maneuvers analysis failed", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error"})
	}

	logger.Logger.Info("Processing routes with graph algorithms")
	finalResult, err := funcgraf.ProcessRoutesFromAPI(result, gl)
	if err != nil {
		logger.Logger.Error("Route processing failed", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error"})
		return
	}
	logger.Logger.Info("Route analysis completed successfully", 
		"status", finalResult.Status,
		"route_type", finalResult.Type)

	ctx.JSON(http.StatusOK, finalResult)
}
