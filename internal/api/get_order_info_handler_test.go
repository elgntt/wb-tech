package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"wb-tech/internal/model"
)

func Test_handler_GetOrderInfo(t *testing.T) {
	existsOrderUID := "test123"
	tests := []struct {
		name           string
		orderUID       string
		serviceBehave  func(service *MockService)
		wantStatusCode int
		wantErr        bool
	}{
		{
			name:     "Success",
			orderUID: existsOrderUID,
			serviceBehave: func(service *MockService) {
				service.EXPECT().GetFromCache(existsOrderUID).Return(model.OrderData{OrderUid: existsOrderUID}, nil)
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name:     "Error from service",
			orderUID: existsOrderUID,
			serviceBehave: func(service *MockService) {
				service.EXPECT().GetFromCache(existsOrderUID).Return(model.OrderData{OrderUid: existsOrderUID},
					errors.New("order not found"))
			},
			wantStatusCode: http.StatusInternalServerError,
			wantErr:        true,
		},
		{
			name:           "Invalid order uid",
			orderUID:       "",
			wantStatusCode: http.StatusNotFound,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := NewMockService(ctrl)
			if tt.serviceBehave != nil {
				tt.serviceBehave(mockService)
			}

			router := gin.Default()
			recorder := httptest.NewRecorder()

			h := handler{
				Service: mockService,
			}

			req, _ := http.NewRequest("GET", "/order/"+tt.orderUID, nil)
			router.LoadHTMLGlob("test_templates/*")
			// Perform the request and record the response.
			router.GET("/order/:orderUid", h.GetOrderInfo)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.wantStatusCode, recorder.Code)
		})
	}
}
