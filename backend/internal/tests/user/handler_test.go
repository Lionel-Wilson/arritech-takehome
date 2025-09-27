package user

import (
	"encoding/json"
	"github.com/Lionel-Wilson/arritech-takehome/internal/http/router"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/domain"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetUsersHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := user.NewMockService(ctrl)
	logger := logrus.New()

	testTime := time.Date(2023, time.March, 1, 10, 0, 0, 0, time.UTC)

	tests := []struct {
		name       string
		mock       func(mockService *user.MockService)
		wantStatus int
		wantBody   map[string]any
	}{
		{
			name: "success",
			mock: func(mockService *user.MockService) {
				mockService.EXPECT().GetUsers(gomock.Any(), domain.GetUsersParams{
					Query: "",
					Page:  1,
					Size:  10,
				}).Return([]domain.User{
					{
						ID:        1,
						Firstname: "Todo",
						Lastname:  "Aoi",
						Age:       20,
						Email:     "BoogieWoogie@gmail.com",
						CreatedAt: testTime,
						UpdatedAt: testTime,
					},
					{
						ID:        2,
						Firstname: "Itadori",
						Lastname:  "Yuji",
						Age:       20,
						Email:     "StraightHands@gmail.com",
						CreatedAt: testTime,
						UpdatedAt: testTime,
					},
				},
					int64(2), nil,
				)
			},
			wantStatus: http.StatusOK,
			wantBody: map[string]interface{}{
				"users": []interface{}{
					map[string]interface{}{
						"id":         float64(1),
						"firstname":  "Todo",
						"lastname":   "Aoi",
						"age":        float64(20),
						"email":      "BoogieWoogie@gmail.com",
						"created_at": testTime.String(),
						"updated_at": testTime.String(),
					},
					map[string]interface{}{
						"id":         float64(2),
						"firstname":  "Itadori",
						"lastname":   "Yuji",
						"age":        float64(20),
						"email":      "StraightHands@gmail.com",
						"created_at": testTime.String(),
						"updated_at": testTime.String(),
					},
				},
				"page":  float64(1),
				"size":  float64(10),
				"total": float64(2),
			},
		},
		{
			name: "internal server error",
			mock: func(mockService *user.MockService) {
				mockService.EXPECT().GetUsers(gomock.Any(), domain.GetUsersParams{
					Query: "",
					Page:  1,
					Size:  10,
				}).Return(nil,
					int64(0), assert.AnError,
				)
			},
			wantStatus: http.StatusInternalServerError,
			wantBody: map[string]interface{}{
				"error": "Something went wrong",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mux := router.New(
				logger,
				mockService,
			)
			tc.mock(mockService)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/users/", nil)
			mux.ServeHTTP(w, req)

			require.Equal(t, tc.wantStatus, w.Code)
			var actual map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &actual)
			require.NoError(t, err)

			assert.Equal(t, tc.wantBody, actual)
		})
	}
}
