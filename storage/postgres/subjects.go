package postgres

import (
	"backend_course/lms/api/models"
	"context"

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

func (s *subjectsRepo) Create(ctx context.Context, subject models.Subjects) (string, error) {
	id := uuid.New()

	query := `
	INSERT INTO
		subjects (id, name, type, updated_at) VALUES ($1, $2, $3, $4);`

	_, err := s.db.Exec(ctx, query, id, subject.Name, subject.Type, subject.UpdatedAt)
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
		name = $2, type = $3, created_at = $4, updated_at = $5
	WHERE 
		id = $1; `

	_, err := s.db.Exec(ctx, query, subject.Id, subject.Name, subject.Type, subject.CreatedAt, subject.UpdatedAt)
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
		id = $1 `

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
		*
	FROM 
		subjects
	WHERE 
		TRUE ` + filter + `
	OFFSET $1 LIMIT $2
					`
	rows, err := s.db.Query(ctx, query, offest, req.Limit)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			subject models.Subjects
		)
		if err := rows.Scan(
			&subject.Id,
			&subject.Name,
			&subject.Type,
			&subject.CreatedAt,
			&subject.UpdatedAt); err != nil {
			return resp, err
		}

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
		*
	FROM
		subjects
	WHERE
		id = $1;
`
	row := s.db.QueryRow(ctx, query, id)

	var subject models.Subjects

	err := row.Scan(&subject.Id, &subject.Name, &subject.Type, &subject.CreatedAt, &subject.UpdatedAt)

	if err != nil {
		return subject, err
	}

	return subject, nil
}
