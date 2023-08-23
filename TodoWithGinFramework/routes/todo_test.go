package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
)

func TestGetTodos(t *testing.T) {
	// Routing
	// SetupAppForTest()
	gin.SetMode(gin.TestMode)
	SetupRouter()
	req, _ := http.NewRequest(http.MethodGet, "/todos", nil)

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 but got %d", rr.Code)
	}

	t.Log("everything working fine!")
}



