package websuite

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/krakenlab/gspec"
)

// WebSuite mock controller requests
type WebSuite struct {
	gspec.Suite
	Server http.Handler
}

// MockRequest on Server
func (suite *WebSuite) MockRequest(method, path, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))

	w := httptest.NewRecorder()
	suite.Server.ServeHTTP(w, req)
	return w
}

// MockRequestWithSession on Server
func (suite *WebSuite) MockRequestWithSession(lastResponse *httptest.ResponseRecorder, method, path, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))

	for _, cookie := range lastResponse.Result().Cookies() {
		req.AddCookie(cookie)
	}

	w := httptest.NewRecorder()
	suite.Server.ServeHTTP(w, req)
	return w
}

// GET on Server
func (suite *WebSuite) GET(path, body string) *httptest.ResponseRecorder {
	return suite.MockRequest("GET", path, body)
}

// POST on Server
func (suite *WebSuite) POST(path, body string) *httptest.ResponseRecorder {
	return suite.MockRequest("POST", path, body)
}

// PUT on Server
func (suite *WebSuite) PUT(path, body string) *httptest.ResponseRecorder {
	return suite.MockRequest("PUT", path, body)
}

// PATCH on Server
func (suite *WebSuite) PATCH(path, body string) *httptest.ResponseRecorder {
	return suite.MockRequest("PATCH", path, body)
}

// DELETE on Server
func (suite *WebSuite) DELETE(path, body string) *httptest.ResponseRecorder {
	return suite.MockRequest("DELETE", path, body)
}

// CookiedGET on Server
func (suite *WebSuite) CookiedGET(lastReq *httptest.ResponseRecorder, path, body string) *httptest.ResponseRecorder {
	return suite.MockRequestWithSession(lastReq, "GET", path, body)
}

// CookiedPOST on Server
func (suite *WebSuite) CookiedPOST(lastReq *httptest.ResponseRecorder, path, body string) *httptest.ResponseRecorder {
	return suite.MockRequestWithSession(lastReq, "POST", path, body)
}

// CookiedPUT on Server
func (suite *WebSuite) CookiedPUT(lastReq *httptest.ResponseRecorder, path, body string) *httptest.ResponseRecorder {
	return suite.MockRequestWithSession(lastReq, "PUT", path, body)
}

// CookiedPATCH on Server
func (suite *WebSuite) CookiedPATCH(lastReq *httptest.ResponseRecorder, path, body string) *httptest.ResponseRecorder {
	return suite.MockRequestWithSession(lastReq, "PATCH", path, body)
}

// CookiedDELETE on Server
func (suite *WebSuite) CookiedDELETE(lastReq *httptest.ResponseRecorder, path, body string) *httptest.ResponseRecorder {
	return suite.MockRequestWithSession(lastReq, "DELETE", path, body)
}
