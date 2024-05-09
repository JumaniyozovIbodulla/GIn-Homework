package postgres

import (
	"backend_course/lms/api/models"
	"context"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateStudent(t *testing.T) {
	studentRepo := NewStudent(db)

	reqStudent := models.Student{
		FirstName: faker.Name(),
		Age:       10,
		LastName:  faker.Word(),
	}

	id, err := studentRepo.Create(context.Background(), reqStudent)
	if assert.NoError(t, err) {
		createdStudent, err := studentRepo.GetStudent(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqStudent.FirstName, createdStudent.FirstName)
			assert.Equal(t, reqStudent.Age, createdStudent.Age)
			assert.Equal(t, reqStudent.LastName, createdStudent.LastName)
		} else {
			return
		}
		fmt.Println("Created student", createdStudent)
	}
}

func TestUpdateStudent(t *testing.T) {
	studentRepo := NewStudent(db)

	reqStudent := models.Student{
		Id:         "3eccd9c5-c3ed-460d-9b2e-e10b600c3a0e",
		FirstName:  faker.Name(),
		LastName:   faker.Word(),
		Age:        12,
		ExternalId: faker.ID,
		Phone:      faker.Phonenumber(),
		Email:      faker.Email(),
	}

	id, err := studentRepo.Update(context.Background(), reqStudent)
	if assert.NoError(t, err) {
		createdStudent, err := studentRepo.GetStudent(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqStudent.FirstName, createdStudent.FirstName)
			assert.Equal(t, reqStudent.Age, createdStudent.Age)
			assert.Equal(t, reqStudent.LastName, createdStudent.LastName)
		} else {
			return
		}
	}
}

func TestDeleteStudent(t *testing.T) {
	studentRepo := NewStudent(db)

	reqStudent := models.Student{
		FirstName: faker.Name(),
		Age:       10,
		LastName:  faker.Word(),
	}

	id, err := studentRepo.Create(context.Background(), reqStudent)
	if assert.NoError(t, err) {
		err := studentRepo.Delete(context.Background(), id)
		if assert.NoError(t, err) {
			return
		}
	}
}

func TestGetAllStudent(t *testing.T) {
	studentRepo := NewStudent(db)
	reqStudent := models.Student{
		FirstName: faker.Name(),
		Age:       10,
		LastName:  faker.Word(),
	}

	response, err := studentRepo.GetAll(context.Background(), models.GetAllStudentsRequest{})
	if assert.NoError(t, err) {
		count := response.Count

		_, err := studentRepo.Create(context.Background(), reqStudent)

		if assert.NoError(t, err) {
			testResponse, err := studentRepo.GetAll(context.Background(), models.GetAllStudentsRequest{})
			if assert.NoError(t, err) {
				testCount := testResponse.Count
				assert.Equal(t, count + 1, testCount)
			} else {
				return
			}
		}
	}
}