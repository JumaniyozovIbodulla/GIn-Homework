package postgres

// import (
// 	"backend_course/lms/api/models"
// 	"context"
// 	"testing"

// 	"github.com/go-faker/faker/v4"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateSubject(t *testing.T) {
// 	subjectRepo := NewSubject(db)

// 	reqSubject := models.Subjects{
// 		Name:      faker.Name(),
// 		Type:      faker.Name(),
// 		UpdatedAt: "2024-05-09 12:34:56",
// 	}

// 	id, err := subjectRepo.Create(context.Background(), reqSubject)
// 	if assert.NoError(t, err) {
// 		createdSubject, err := subjectRepo.GetSubject(context.Background(), id)
// 		if assert.NoError(t, err) {
// 			assert.Equal(t, reqSubject.Name, createdSubject.Name)
// 			assert.Equal(t, reqSubject.Type, createdSubject.Type)
// 			assert.Equal(t, reqSubject.UpdatedAt, createdSubject.UpdatedAt)
// 		} else {
// 			return
// 		}
// 	}
// }
