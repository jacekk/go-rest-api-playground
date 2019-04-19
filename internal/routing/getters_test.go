package routing // import "github.com/jacekk/go-rest-api-playground/internal/routing"

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	router := SetupRouter()

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping/non/existing", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, recorder.Code, http.StatusNotFound)
	assert.Equal(t, recorder.Body.String(), "404 page not found")
}

func TestGetPlainText(t *testing.T) {
	router := SetupRouter()

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/simple/plain", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, recorder.Code, http.StatusOK)
	assert.Contains(t, recorder.Body.String(), "Works now --")
}

func TestGetQuery(t *testing.T) {
	router := SetupRouter()

	query := url.Values{}
	query.Set("uno", "foo uno")
	query.Add("tres", "tres 1")
	query.Add("tres", "treS-2")

	expectedQuery := map[string]interface{}{
		"uno":  []string{"foo uno"},
		"tres": []string{"tres 1", "treS-2"},
	}
	expectedBody := map[string]interface{}{
		"dos":   "default-dos",
		"now":   time.Now().Format(TIME_FORMAT),
		"query": expectedQuery,
		"uno":   "foo uno",
	}

	reqUrl := &url.URL{
		Path:     "/simple/json",
		RawQuery: query.Encode(),
	}
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", reqUrl.String(), nil)
	router.ServeHTTP(recorder, req)
	resp := recorder.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	respJson, _ := simplejson.NewJson(body)
	unoValue := respJson.GetPath("uno").MustString()
	expectedBodyBytes, _ := json.Marshal(expectedBody)

	if unoValue != "foo uno" {
		t.Errorf("Expected 'foo uno' as 'uno' value, got '%s'.", unoValue) // less verbose and less readable way
	}

	assert.Equal(t, recorder.Code, http.StatusOK)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.Contains(t, resp.Header.Get("Content-Type"), "application/json")
	assert.JSONEq(t, recorder.Body.String(), string(expectedBodyBytes)) // or ...
	// assert.JSONEq(t, string(body), expectedBody)
}

// func GetQuery(ctx *gin.Context) {
// 	response := gin.H{
// 		"now":   time.Now(),
// 		"uno":   ctx.Query("uno"),
// 		"dos":   ctx.DefaultQuery("dos", "default-dos"),
// 		"query": ctx.Request.URL.Query(),
// 	}
// 	ctx.JSON(http.StatusOK, response)
// }

// func GetParams(ctx *gin.Context) {
// 	response := gin.H{
// 		"now":    time.Now(),
// 		"dos":    ctx.Param("dos"),
// 		"tres":   ctx.Param("tres"),
// 		"params": ctx.Params,
// 	}
// 	ctx.JSON(http.StatusOK, response)
// }
