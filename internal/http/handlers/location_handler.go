package handlers

import (
	"go-transjakarta/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLastLocation(c *gin.Context) {
	vehicleID := c.Param("vehicle_id")

	loc, err := database.GetLastLocation(vehicleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB Error " + err.Error()})
		return
	}

	if loc == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, loc)
}

func GetLocationHistory(c *gin.Context) {
	vehicleID := c.Param("vehicle_id")
	start := c.Query("start")
	end := c.Query("end")

	locations, err := database.GetHistory(vehicleID, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vehicle_id": vehicleID,
		"history":    locations,
	})
}
