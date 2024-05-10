package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type studentRepo struct {
	db *pgxpool.Pool
}

func NewStudent(db *pgxpool.Pool) studentRepo {
	return studentRepo{
		db: db,
	}
}

func (s *studentRepo) Create(ctx context.Context, student models.Student) (string, error) {

	id := uuid.New()

	query := `
	INSERT INTO
		students (id, first_name, last_name, age, external_id, phone, mail) VALUES ($1, $2, $3, $4, $5, $6, $7);`

	_, err := s.db.Exec(ctx, query, id, student.FirstName, student.LastName, student.Age, student.ExternalId, student.Phone, student.Email)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *studentRepo) Update(ctx context.Context, student models.Student) (string, error) {
	query := `
	UPDATE
		students
	SET
		first_name = $2, last_name = $3, age = $4, external_id = $5, phone = $6, mail = $7, updated_at = NOW()
	WHERE 
		id = $1; `

	_, err := s.db.Exec(ctx, query, student.Id, student.FirstName, student.LastName, student.Age, student.ExternalId, student.Phone, student.Email)
	if err != nil {
		return "", err
	}
	return student.Id, nil
}

func (s *studentRepo) UpdateStatus(ctx context.Context, student models.Student) (string, error) {
	fmt.Println(student.IsActive)
	if student.IsActive {
		student.IsActive = false
	} else {
		student.IsActive = true
	}
	fmt.Println(student.IsActive)
	query := `
	UPDATE
		students
	SET
	is_active = $2
	WHERE 
		id = $1;`

	_, err := s.db.Exec(ctx, query, student.Id, student.IsActive)
	if err != nil {
		return "", err
	}
	return student.Id, nil
}

func (s *studentRepo) Delete(ctx context.Context, id string) error {
	query := `
	DELETE
	FROM
		students
	WHERE 
		id = $1 `

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *studentRepo) GetAll(ctx context.Context, req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error) {
	resp := models.GetAllStudentsResponse{}
	filter := ""
	offest := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND first_name ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT id,
					first_name,
					last_name
				FROM students
				WHERE TRUE ` + filter + `
				OFFSET $1 LIMIT $2
					`
	rows, err := s.db.Query(ctx, query, offest, req.Limit)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			student  models.GetStudent
			lastName sql.NullString
		)
		if err := rows.Scan(
			&student.Id,
			&student.FirstName,
			&lastName); err != nil {
			return resp, err
		}

		student.LastName = pkg.NullStringToString(lastName)
		resp.Students = append(resp.Students, student)
	}

	err = s.db.QueryRow(ctx, `SELECT count(*) from students WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *studentRepo) GetStudent(ctx context.Context, id string) (models.GetStudent, error) {

	query := `
	SELECT
		id,
		first_name,
		last_name,
		age,
		external_id,
		phone,
		mail,
		TO_CHAR(created_at,'YYYY-MM-DD HH:MM:SS'),
		TO_CHAR(updated_at,'YYYY-MM-DD HH:MM:SS')
	FROM
		students
	WHERE
		id = $1;
`
	row := s.db.QueryRow(ctx, query, id)

	var (
		student models.GetStudent
		lastName, updatedAt sql.NullString
	)

	err := row.Scan(&student.Id, &student.FirstName, &lastName, &student.Age, &student.ExternalId, &student.Phone, &student.Email, &student.CreatedAt, &updatedAt)

	student.LastName = pkg.NullStringToString(lastName)
	student.UpdatedAt = pkg.NullStringToString(updatedAt)


	if err != nil {
		return student, err
	}

	return student, nil
}
