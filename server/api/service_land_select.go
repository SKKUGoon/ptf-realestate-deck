package api

import (
	"fmt"
	"log"
	"net/http"

	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
)

type requestCoordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Select a point, give back a polygon information
// ex curl:
// ```
//
//	curl -X POST http://localhost:10500/api/v1/land/pnu \
//	 -H 'Content-Type: application/json' \
//	 -d '{"latitude": 37.5665, "longitude": 126.9780}'
//
// ```
func servicePnuSelect(c *gin.Context, ctrl *Controller) {
	// Client request parsing
	request := requestCoordinate{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Printf("failed to parse json body: %v", err)
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	// Controller Retrieves data
	tx, err := ctrl.Data.Map.Tx(ctx)
	if err != nil {
		log.Printf("failed to create a transaction: %v", err)
		c.JSON(500, gin.H{"error": "failed to create transaction"})
		return
	}
	defer tx.Client().Close()

	data, err := tx.Pnu.Query().
		Where(func(s *sql.Selector) {
			s.Where(
				sql.ExprP(
					"public.ST_Contains(geometry, public.ST_SetSRID(public.ST_Point($1, $2), 4326))",
					request.Longitude,
					request.Latitude,
				),
			)
		}).
		All(ctx)
	if err != nil {
		log.Printf("failed to query: %v", err)
		c.JSON(500, gin.H{"error": "failed to query database"})
		return
	}

	fmt.Println(data, "query success")
	c.JSON(http.StatusOK, data)
}
