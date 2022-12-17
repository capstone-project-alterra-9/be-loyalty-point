package mocking

import (
	// jwtAuth "capstone-project/middleware"
	mocksService "capstone-project/service/mocks"
	// "github.com/stretchr/testify/mock"
	"testing"

	// "github.com/stretchr/testify/assert"
	"capstone-project/entity"
	// "capstone-project/service"
	"github.com/labstack/echo/v4"
	gomock "github.com/golang/mock/gomock"
	// "time"
	// "errors"
)

var (
	Service mocksService.MockSvc
)

func TestUseReturnsErrorFromDo(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()
	var c echo.Context

	var userViewEntity entity.UsersView

    // dummyError := errors.New("dummy error")
    mockDoer := mocksService.NewMockSvc(mockCtrl)
    // testUser := []entity.UsersView{entity.UsersView}

    // Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return dummyError from the mocked call.
    mockDoer.EXPECT().GetUsersPagination(c).Return([]entity.UsersView{userViewEntity}, nil).Times(1)

	_, err := Service.GetUsersPagination(mockDoer)
	if err != nil {
		t.Fatal(err)
	}
}

// func TestMain(m *testing.M) {
// 	assignmentService = assignments.NewAssignmentUsecase(&assignmentRepository, &middlewares.ConfigJWT{})

// 	userViewEntity = entity.UsersView{
// 		ID: "07757383-d7df-46cd-83b0-177018060afc",
// 		Role: "user",
// 		CreatedAt: "2022-12-07T16:20:24.409423Z",
// 		UpdatedAt: "2022-12-07T16:20:24.409423Z",
// 		Username: "test 2",
// 		Email: "test2@gmail.com",
// 		Password: "$2a$10$9BGXMP4hsZPhcXsGFxKwruI131OWwQ6AmYbXcC6i84NSQEbpbPT0G",
// 		Points: 20000,
// 		CostPoints: 0,
// 	}

// 	m.Run()
// }


// func TestGetAll(t *testing.T) {
// 	controllerMocks := mocks.NewController(t)
// 	var c echo.Context
// 	t.Run("Get users pagination | Valid", func(t *testing.T) {
// 		controllerMocks.On("GetUsersPagination", mock.AnythingOfType("echo.Context")).
// 		Return([]entity.UsersView{userViewEntity}, nil,).Once()

// 		result,_ := Service.GetUsersPagination(c)

// 		assert.Equal(t, 1, len(result))
// 	})

// 	t.Run("Get users pagination | InValid", func(t *testing.T) {
// 		controllerMocks.On("GetUsersPagination", mock.AnythingOfType("echo.Context")).
// 		Return([]entity.UsersView{userViewEntity}, nil,).Once()

// 		result,_ := Service.GetUsersPagination(c)

// 		assert.Equal(t, 0, len(result))
// 	})
// }

// func TestGetByID(t *testing.T) {
// 	t.Run("Get By ID | Valid", func(t *testing.T) {
// 		assignmentRepository.On("GetByID", "7111a840-3099-4c62-85b6-00d665d42cba").Return(assignmentDomain).Once()

// 		result := assignmentService.GetByID("7111a840-3099-4c62-85b6-00d665d42cba")

// 		assert.NotNil(t, result)
// 	})

// 	t.Run("Get By ID | InValid", func(t *testing.T) {
// 		assignmentRepository.On("GetByID", "-7111a840-3099-4c62-85b6-00d665d42cba").Return(assignments.Domain{}).Once()

// 		result := assignmentService.GetByID("-7111a840-3099-4c62-85b6-00d665d42cba")

// 		assert.NotNil(t, result)
// 	})
// }

// func TestCreate(t *testing.T) {
// 	t.Run("Create | Valid", func(t *testing.T) {
// 		assignmentRepository.On("Create", &assignmentDomain).Return(assignmentDomain).Once()

// 		result := assignmentService.Create(&assignmentDomain)

// 		assert.NotNil(t, result)
// 	})

// 	t.Run("Create | InValid", func(t *testing.T) {
// 		assignmentRepository.On("Create", &assignments.Domain{}).Return(assignments.Domain{}).Once()

// 		result := assignmentService.Create(&assignments.Domain{})

// 		assert.NotNil(t, result)
// 	})
// }

// func TestUpdate(t *testing.T) {
// 	t.Run("Update | Valid", func(t *testing.T) {
// 		assignmentRepository.On("Update", "7111a840-3099-4c62-85b6-00d665d42cba", &assignmentDomain).Return(assignmentDomain).Once()

// 		result := assignmentService.Update("7111a840-3099-4c62-85b6-00d665d42cba", &assignmentDomain)

// 		assert.NotNil(t, result)
// 	})

// 	t.Run("Update | InValid", func(t *testing.T) {
// 		assignmentRepository.On("Update", "7111a840-3099-4c62-85b6-00d665d42cba", &assignments.Domain{}).Return(assignments.Domain{}).Once()

// 		result := assignmentService.Update("7111a840-3099-4c62-85b6-00d665d42cba", &assignments.Domain{})

// 		assert.NotNil(t, result)
// 	})
// }

// func TestDelete(t *testing.T) {
// 	t.Run("Delete | Valid", func(t *testing.T) {
// 		assignmentRepository.On("Delete", "7111a840-3099-4c62-85b6-00d665d42cba").Return(true).Once()

// 		result := assignmentService.Delete("7111a840-3099-4c62-85b6-00d665d42cba")

// 		assert.True(t, result)
// 	})

// 	t.Run("Delete | InValid", func(t *testing.T) {
// 		assignmentRepository.On("Delete", "-7111a840-3099-4c62-85b6-00d665d42cba").Return(false).Once()

// 		result := assignmentService.Delete("-7111a840-3099-4c62-85b6-00d665d42cba")

// 		assert.False(t, result)
// 	})
// }
