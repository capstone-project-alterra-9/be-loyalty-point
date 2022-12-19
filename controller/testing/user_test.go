package controller

import (
	"capstone-project/middlewares"
	_assignmentMock "capstone-project/controller/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	assignmentRepository _assignmentMock.Repository
	assignmentService    assignments.Usecase

	assignmentDomain assignments.Domain
)

func TestMain(m *testing.M) {
	assignmentService = assignments.NewAssignmentUsecase(&assignmentRepository, &middlewares.ConfigJWT{})

	assignmentDomain = assignments.Domain{
		IdAssignment        : "",  
		Name     			: "Tugas 1",         	
		Deadline			: "2022-11-05 13:45:20.930",     	
		Class				: "10",
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All Assignment | Valid", func(t *testing.T) {
		assignmentRepository.On("GetAssignments").Return([]assignments.Domain{assignmentDomain}).Once()

		result := assignmentService.GetAssignments()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All Assignment | InValid", func(t *testing.T) {
		assignmentRepository.On("GetAssignments").Return([]assignments.Domain{}).Once()

		result := assignmentService.GetAssignments()

		assert.Equal(t, 0, len(result))
	})
}

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
