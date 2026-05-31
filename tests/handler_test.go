package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	contracts "project/internal/application/contracts"
	linkhandlers "project/internal/controllers/link_handlers"
	"project/tests/mocks"
)

func newRouter(handler *linkhandlers.LinkHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/link", handler.PostLink)
	r.GET("/link/:shortUrl", handler.GetLink)
	return r
}

func TestPostLink_Success(t *testing.T) {
	svc := new(mocks.MockLinkService)
	svc.On("AddNewLink", contracts.LinkData{Url: "https://example.com"}).
		Return(&contracts.LinkAddResult{ShortedUrl: "abc1234"}, http.StatusCreated)

	r := newRouter(linkhandlers.NewLinkHandler(svc))

	req := httptest.NewRequest(http.MethodPost, "/link", strings.NewReader(`{"full_url":"https://example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp contracts.LinkAddResult
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "abc1234", resp.ShortedUrl)
	svc.AssertExpectations(t)
}

func TestPostLink_BadJSON(t *testing.T) {
	svc := new(mocks.MockLinkService)
	r := newRouter(linkhandlers.NewLinkHandler(svc))

	req := httptest.NewRequest(http.MethodPost, "/link", strings.NewReader(`not json`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	svc.AssertNotCalled(t, "AddNewLink")
}

func TestPostLink_ServiceInternalError(t *testing.T) {
	svc := new(mocks.MockLinkService)
	svc.On("AddNewLink", contracts.LinkData{Url: "https://example.com"}).
		Return(nil, http.StatusInternalServerError)

	r := newRouter(linkhandlers.NewLinkHandler(svc))

	req := httptest.NewRequest(http.MethodPost, "/link", strings.NewReader(`{"full_url":"https://example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}

func TestPostLink_Conflict(t *testing.T) {
	svc := new(mocks.MockLinkService)
	svc.On("AddNewLink", contracts.LinkData{Url: "https://example.com"}).
		Return(nil, http.StatusConflict)

	r := newRouter(linkhandlers.NewLinkHandler(svc))

	req := httptest.NewRequest(http.MethodPost, "/link", strings.NewReader(`{"full_url":"https://example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	svc.AssertExpectations(t)
}

func TestGetLink_Success(t *testing.T) {
	svc := new(mocks.MockLinkService)
	svc.On("ExtractFullLink", contracts.ShortLinkData{ShortLink: "abc1234"}).
		Return(&contracts.LinkExtractResult{FullUrl: "https://example.com"}, http.StatusFound)

	r := newRouter(linkhandlers.NewLinkHandler(svc))

	req := httptest.NewRequest(http.MethodGet, "/link/abc1234", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "https://example.com", w.Header().Get("Location"))
	svc.AssertExpectations(t)
}

func TestGetLink_NotFound(t *testing.T) {
	svc := new(mocks.MockLinkService)
	svc.On("ExtractFullLink", contracts.ShortLinkData{ShortLink: "missing"}).
		Return(nil, http.StatusInternalServerError)

	r := newRouter(linkhandlers.NewLinkHandler(svc))

	req := httptest.NewRequest(http.MethodGet, "/link/missing", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}

func TestGetLink_ServiceInternalError(t *testing.T) {
	svc := new(mocks.MockLinkService)
	svc.On("ExtractFullLink", contracts.ShortLinkData{ShortLink: "abc1234"}).
		Return(nil, http.StatusInternalServerError)

	r := newRouter(linkhandlers.NewLinkHandler(svc))

	req := httptest.NewRequest(http.MethodGet, "/link/abc1234", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}
