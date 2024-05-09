package postgres

import (
	"backend_course/lms/api/models"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type timeRepo struct {
	db *pgxpool.Pool
}

func NewTime(db *pgxpool.Pool) timeRepo {
	return timeRepo{
		db: db,
	}
}

func (s *timeRepo) Create(ctx context.Context, time models.Time) (string, error) {
	id := uuid.New()

	query := `
	INSERT INTO
		time_table (id, teacher_id, student_id, subject_id, from_date, to_date) VALUES ($1, $2, $3, $4, $5, $6);`

	_, err := s.db.Exec(ctx, query, id, time.TeacherId, time.StudentId, time.SubjectId, time.FromDate, time.ToDate)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *timeRepo) Update(ctx context.Context, time models.Time) (string, error) {
	query := `
	UPDATE
		time_table
	SET
		teacher_id = $2, student_id = $3, subject_id = $4, from_date = $5, to_date = $6
	WHERE 
		id = $1; `

	_, err := s.db.Exec(ctx, query, time.Id, time.TeacherId, time.StudentId, time.SubjectId, time.FromDate, time.ToDate)
	if err != nil {
		return "", err
	}
	return time.Id, nil
}

func (s *timeRepo) Delete(ctx context.Context, id string) error {
	query := `
	DELETE
	FROM
		time_table
	WHERE 
		id = $1 `

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *timeRepo) GetAll(ctx context.Context, req models.GetAllTimeRequest) (models.GetAllTimeResponse, error) {
	resp := models.GetAllTimeResponse{}
	filter := ""
	offest := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND first_name ILIKE '%` + req.Search + `%' `
	}

	query := `
	SELECT
		*
	FROM 
		time_table
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
			time models.Time
		)
		if err := rows.Scan(
			&time.Id,
			&time.TeacherId,
			&time.StudentId,
			&time.SubjectId,
			&time.FromDate,
			&time.ToDate); err != nil {
			return resp, err
		}

		resp.Time = append(resp.Time, time)
	}

	err = s.db.QueryRow(ctx, `SELECT COUNT(*) from time_table WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *timeRepo) GetTime(ctx context.Context, id string) (models.Time, error) {

	query := `
	SELECT
		*
	FROM
		time_table
	WHERE
		id = $1;
`
	row := s.db.QueryRow(ctx, query, id)

	var time models.Time

	err := row.Scan(&time.Id, &time.TeacherId, &time.StudentId, &time.SubjectId, &time.FromDate, &time.ToDate)

	if err != nil {
		return time, err
	}

	return time, nil
}
