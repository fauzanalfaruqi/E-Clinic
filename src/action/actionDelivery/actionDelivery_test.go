package actionDelivery

import (
	"avengers-clinic/model/dto/actionDto"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockActionUsecase struct {
	mock.Mock
}

func (mock *mockActionUsecase) GetAll() ([]actionDto.Action, error) {
	args := mock.Called()
	return args.Get(0).([]actionDto.Action), args.Error(1)
}

func (mock *mockActionUsecase) GetByID(actionID string) (actionDto.Action, error) {
	args := mock.Called(actionID)
	return args.Get(0).(actionDto.Action), args.Error(1)
}

func (mock *mockActionUsecase) Create(req actionDto.CreateRequest) (actionDto.Action, error) {
	args := mock.Called(req)
	return args.Get(0).(actionDto.Action), args.Error(1)
}

func (mock *mockActionUsecase) Update(req actionDto.UpdateRequest) (actionDto.Action, error) {
	args := mock.Called(req)
	return args.Get(0).(actionDto.Action), args.Error(1)
}

func (mock *mockActionUsecase) Delete(actionID string) error {
	args := mock.Called(actionID)
	return args.Error(0)
}

func (mock *mockActionUsecase) SoftDelete(actionID string) error {
	args := mock.Called(actionID)
	return args.Error(0)
}

func (mock *mockActionUsecase) Restore(actionID string) error {
	args := mock.Called(actionID)
	return args.Error(0)
}

type actionDeliveryTestSuite struct {
	suite.Suite
	router *gin.Engine
	actionUC *mockActionUsecase
}

func (suite *actionDeliveryTestSuite) SetupTest() {
	suite.router = gin.New()
	suite.actionUC = new(mockActionUsecase)
	
	apiGroup := suite.router.Group("/api")
	v1Group := apiGroup.Group("/v1")
	NewActionDelivery(v1Group, suite.actionUC)
}

// Start Get All
// func (suite *actionDeliveryTestSuite) TestGetAllSuccess() {
// 	actions := []actionDto.Action{{ID: "1", Name: "Konsultasi", Price: 20000}}

// 	suite.actionUC.On("GetAll").Return(actions, nil)
	
// 	res := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodGet, "/api/v1/actions", nil)

// 	token, _ := utils.GenerateJWT()
// }
// End Get All

func (suite *actionDeliveryTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
}

func TestActionDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(actionDeliveryTestSuite))
}