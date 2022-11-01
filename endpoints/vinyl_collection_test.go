package endpoints

import (
	"bytes"
	"encoding/json"
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

func CleanString(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\t", "")
	return str
}

func TestNew(t *testing.T) {
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
	if !reflect.DeepEqual(collection.Albums, expected) {
		t.Errorf("Actual albums of (%v) was not equal to that of (%v)",
			collection.Albums, expected)
	}
}

func TestGetAlbumsHandler(t *testing.T) {
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
	router := SetUpRouter()
	router.GET("/albums", GetAlbumsHandler)
	request, _ := http.NewRequest("GET", "/albums", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response, _ := ioutil.ReadAll(recorder.Body)
	assert.Equal(t, CleanString(string(expected)), CleanString(string(response)))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestPostAlbumHandler(t *testing.T) {
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
			"price":  39.99},
		{"id": "4",
			"title": "The Modern Sound of Betty Carter",
			"artist": "Betty Carter",
			"price": 49.99}]`

	t.Logf("Running TestPutAlbum...\n")
	router := SetUpRouter()
	router.POST("/albums", PostAlbumsHandler)
	body := types.Album{ID: "4",
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99}
	jsonValue, _ := json.Marshal(body)
	request, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonValue))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response, _ := ioutil.ReadAll(recorder.Body)
	assert.Equal(t, CleanString(string(expected)), CleanString(string(response)))
	assert.Equal(t, http.StatusCreated, recorder.Code)
}
