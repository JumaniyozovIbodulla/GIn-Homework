package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type subjectsRepo struct {
	db *pgxpool.Pool
}

func NewSubject(db *pgxpool.Pool) subjectsRepo {
	return subjectsRepo{
		db: db,
	}
}

func (s *subjectsRepo) Create(ctx context.Context, subject models.AddSubject) (string, error) {
	id := uuid.New()

	query := `
	INSERT INTO
		subjects (id, name, type) VALUES ($1, $2, $3);`

	_, err := s.db.Exec(ctx, query, id, subject.Name, subject.Type)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *subjectsRepo) Update(ctx context.Context, subject models.Subjects) (string, error) {
	query := `
	UPDATE
		subjects
	SET
		name = $2, type = $3, created_at = $4, updated_at = NOW()
	WHERE 
		id = $1;`

	_, err := s.db.Exec(ctx, query, subject.Id, subject.Name, subject.Type, subject.CreatedAt)
	if err != nil {
		return "", err
	}
	return subject.Id, nil
}

func (s *subjectsRepo) Delete(ctx context.Context, id string) error {
	query := `
	DELETE
	FROM
		subjects
	WHERE 
		id = $1;`

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *subjectsRepo) GetAll(ctx context.Context, req models.GetAllSubjectsRequest) (models.GetAllSubjectsResponse, error) {
	resp := models.GetAllSubjectsResponse{}
	filter := ""
	offest := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND first_name ILIKE '%` + req.Search + `%' `
	}

	query := `
	SELECT
		id,
		name,
		type,
		TO_CHAR(created_at,'YYYY-MM-DD HH:MM:SS'),
		TO_CHAR(updated_at,'YYYY-MM-DD HH:MM:SS')
	FROM
		subjects
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
			subject models.Subjects
			name, typeSubject, updatedAt sql.NullString
		)
		if err := rows.Scan(
			&subject.Id,
			&name,
			&typeSubject,
			&subject.CreatedAt,
			&updatedAt); err != nil {
			return resp, err
		}
		subject.Name = pkg.NullStringToString(name)
		subject.Type = pkg.NullStringToString(typeSubject)
		subject.UpdatedAt = pkg.NullStringToString(updatedAt)

		resp.Subjects = append(resp.Subjects, subject)
	}

	err = s.db.QueryRow(ctx, `SELECT COUNT(*) from subjects WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *subjectsRepo) GetSubject(ctx context.Context, id string) (models.Subjects, error) {

	query := `
	SELECT
		id,
		name,
		type,
		TO_CHAR(created_at,'YYYY-MM-DD HH:MM:SS'),
		TO_CHAR(updated_at,'YYYY-MM-DD HH:MM:SS')
	FROM
		subjects
	WHERE
		id = $1;`

	row := s.db.QueryRow(ctx, query, id)

	var (
		subject models.Subjects
		name, typeSubject, updatedAt sql.NullString
)

	err := row.Scan(&subject.Id, &name, &typeSubject, &subject.CreatedAt, &updatedAt)

	subject.Name = pkg.NullStringToString(name)
	subject.Type = pkg.NullStringToString(typeSubject)
	subject.UpdatedAt = pkg.NullStringToString(updatedAt)

	if err != nil {
		return subject, err
	}

	return subject, nil
}
