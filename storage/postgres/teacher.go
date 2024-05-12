package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type teacherRepo struct {
	db *pgxpool.Pool
}

func NewTeacher(db *pgxpool.Pool) teacherRepo {
	return teacherRepo{
		db: db,
	}
}

func (s *teacherRepo) Create(ctx context.Context, teacher models.AddTeacher) (string, error) {

	id := uuid.New()

	query := `
	INSERT INTO
		teachers (id, first_name, last_name, subject_id, start_working, phone, mail, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

	_, err := s.db.Exec(ctx, query, id, teacher.FirstName, teacher.LastName, teacher.SubjectId, teacher.StartWorking, teacher.Phone, teacher.Email, teacher.Password)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *teacherRepo) Update(ctx context.Context, teacher models.Teacher) (string, error) {
	query := `
	UPDATE
		teachers
	SET
		first_name = $2, last_name = $3, subject_id = $4, start_working = $5, phone = $6, mail = $7, updated_at = NOW()
	WHERE 
		id = $1 `

	_, err := s.db.Exec(ctx, query, teacher.Id, teacher.FirstName, teacher.LastName, teacher.SubjectId, teacher.StartWorking, teacher.Phone, teacher.Email)
	if err != nil {
		return "", err
	}
	return teacher.Id, nil
}

func (s *teacherRepo) Delete(ctx context.Context, id string) error {
	query := `
	DELETE
	FROM
		teachers
	WHERE 
		id = $1 `

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *teacherRepo) GetAll(ctx context.Context, req models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error) {
	resp := models.GetAllTeachersResponse{}
	filter := ""
	offest := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND first_name ILIKE '%` + req.Search + `%' `
	}

	query := `
	SELECT 
		id,
		first_name,
		last_name,
		subject_id,
		TO_CHAR(start_working,'YYYY-MM-DD HH:MM:SS'),
		phone,
		mail,
		TO_CHAR(created_at,'YYYY-MM-DD HH:MM:SS'),
		TO_CHAR(updated_at,'YYYY-MM-DD HH:MM:SS')
	FROM 
		teachers
	WHERE TRUE ` + filter + `
	OFFSET
		$1 
	LIMIT 
		$2;`

	rows, err := s.db.Query(ctx, query, offest, req.Limit)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			teacher   models.Teacher
			firstName, lastName, subjectId, startWorking, phone, mail, createdAt, updatedAt sql.NullString
		)

		if err := rows.Scan(
			&teacher.Id,
			&firstName,
			&lastName,
			&subjectId,
			&startWorking,
			&phone,
			&mail,
			&createdAt,
			&updatedAt,); err != nil {
			return resp, err
		}
		teacher.FirstName = pkg.NullStringToString(firstName)
		teacher.LastName = pkg.NullStringToString(lastName)
		teacher.SubjectId = pkg.NullStringToString(subjectId)
		teacher.StartWorking = pkg.NullStringToString(startWorking)
		teacher.Phone = pkg.NullStringToString(phone)
		teacher.Email = pkg.NullStringToString(mail)

		resp.Teachers = append(resp.Teachers, teacher)
	}

	err = s.db.QueryRow(ctx, `SELECT count(*) from teachers WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *teacherRepo) GetTeacher(ctx context.Context, id string) (models.Teacher, error) {

	query := `
	SELECT
		id,
		first_name,
		last_name,
		subject_id,
		TO_CHAR(start_working,'YYYY-MM-DD HH:MM:SS'),
		phone,
		mail,
		TO_CHAR(created_at,'YYYY-MM-DD HH:MM:SS'),
		TO_CHAR(updated_at,'YYYY-MM-DD HH:MM:SS')
	FROM
		teachers
	WHERE
		id = $1;
`
	row := s.db.QueryRow(ctx, query, id)

	var (
		teacher models.Teacher
		firstName, lastName, subjectId, startWorking, phone, mail, createdAt, updatedAt sql.NullString
)

	err := row.Scan(&teacher.Id, &firstName, &lastName, &subjectId, &startWorking, &phone, &mail, &createdAt, &updatedAt)

	teacher.FirstName = pkg.NullStringToString(firstName)
	teacher.LastName = pkg.NullStringToString(lastName)
	teacher.SubjectId = pkg.NullStringToString(subjectId)
	teacher.StartWorking = pkg.NullStringToString(startWorking)
	teacher.Phone = pkg.NullStringToString(phone)
	teacher.Email = pkg.NullStringToString(mail)

	if err != nil {
		return teacher, err
	}
	return teacher, nil
}
