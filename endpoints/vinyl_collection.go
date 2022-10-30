package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mehoggan/vinyl-collection-service-go/types"
)

var collection *types.VinylCollection

func init() {
	collection = &types.VinylCollection{Albums: []types.Album{
		{ID: "1",
			Title:  "Blue Train",
			Artist: "John Coltrane",
			Price:  56.99},
		{ID: "2",
			Title:  "Jeru",
			Artist: "Gerry Mulligan",
			Price:  17.99},
		{ID: "3",
			Title:  "Sarah Vaughan and Clifford Brown",
			Artist: "Sarah Vaughan",
			Price:  39.99}}}
}

func GetAlbumsHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, collection.Albums)
}
