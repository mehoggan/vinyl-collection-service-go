package endpoints

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mehoggan/vinyl-collection-service-go/types"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestNew(t *testing.T) {
	vc := NewVinylCollection()
	expected := []types.Album{
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
			Price:  39.99}}
	if !reflect.DeepEqual(vc.Albums, expected) {
		t.Errorf("Actual albums of (%v) was not equal to that of (%v)", vc.Albums,
			expected)
	}
}

func TestGetAlbumsHandler(t *testing.T) {
	var cleanString = func(str string) string {
		str = strings.ReplaceAll(str, " ", "")
		str = strings.ReplaceAll(str, "\n", "")
		str = strings.ReplaceAll(str, "\t", "")
		return str
	}

	expected := `[
		{"id": "1",
			"title":  "Blue Train",
			"artist": "John Coltrane",
			"price":  56.99},
		{"id": "2",
			"title":  "Jeru",
			"artist": "Gerry Mulligan",
			"price":  17.99},
		{"id": "3",
			"title":  "Sarah Vaughan and Clifford Brown",
			"artist": "Sarah Vaughan",
			"price":  39.99}]`
	t.Logf("Running TestGetAlbums...\n")
	vc := NewVinylCollection()
	router := SetUpRouter()
	router.GET("/albums", vc.GetAlbumsHandler)
	request, _ := http.NewRequest("GET", "/albums", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response, _ := ioutil.ReadAll(recorder.Body)
	assert.Equal(t, cleanString(string(expected)), cleanString(string(response)))
	assert.Equal(t, http.StatusOK, recorder.Code)
}
