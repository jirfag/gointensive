package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jirfag/gointensive/lec3/2_http_pool/workers"

	"github.com/golang/mock/gomock"
)

func TestHandlerOKCase(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `Hello, world!`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//go:generate mockgen -package server -source http.go -destination http_test_mock.go IPool
func TestHandlerErrorCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockedPool := NewMockIPool(ctrl)

	savedWp := wp
	defer func() {
		wp = savedWp
	}() // restore
	wp = mockedPool

	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)

	mockedPool.EXPECT().
		AddTaskSyncTimed(gomock.Any(), requestWaitInQueueTimeout).
		Return(nil, workers.ErrJobTimedOut)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
