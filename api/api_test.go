package api_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	"teststaticwongnai2/api"
	mockapi_test "teststaticwongnai2/api/mock_api"
	"teststaticwongnai2/client"
)

func TestApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := mockapi_test.NewMockHandlerMock(ctrl)
	h := api.Handler{
		Serv: mock,
	}
	mock.EXPECT().GetHttp().Return(nil, client.Age{})
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)
	h.GetRespone(ginContext)
}
