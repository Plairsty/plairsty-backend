package persistence

import (
	projectPb "awesomeProject/internal/proto/project"
	"database/sql"
	"github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ProjectRepository struct {
	DB *sql.DB
}

func (p ProjectRepository) Insert(username string, project *projectPb.ProjectFields) error {
	query := `
			INSERT INTO projects (username, title, description, leader, member_ids, guide_name, project_url, semester, start_date, end_date)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			`
	var startDate time.Time
	var endDate time.Time
	if project.GetStartDate() != nil {
		startDate = project.GetStartDate().AsTime()
		startDate = startDate.Add(time.Duration(project.StartDate.Seconds))
		startDate = startDate.AddDate(0, 0, 1)
	}
	if project.GetEndDate() != nil {
		endDate = project.GetEndDate().AsTime()
		endDate = endDate.Add(time.Duration(project.GetEndDate().GetSeconds()))
		endDate = endDate.AddDate(0, 0, 1)
	}
	_, err := p.DB.Exec(
		query,
		username,
		project.GetName(),
		project.GetDescription(),
		project.GetLeader(),
		pq.Array(project.GetMembersUsername()),
		project.GetGuideName(),
		project.GetProjectUrl(),
		project.GetSemester(),
		startDate,
		endDate,
	)
	return err
}

func (p ProjectRepository) Get(username string, projectId int64) (*projectPb.ProjectFields, error) {
	query := `
			SELECT title, description, leader, member_ids, guide_name, project_url, semester, start_date, end_date
			FROM projects
			WHERE id = $1 AND username = $2
			`
	var project projectPb.ProjectFields
	var startDate time.Time
	var endDate time.Time
	err := p.DB.QueryRow(query, projectId, username).Scan(
		&project.Name,
		&project.Description,
		&project.Leader,
		&project.MembersUsername,
		&project.GuideName,
		&project.ProjectUrl,
		&project.Semester,
		&startDate,
		&endDate,
	)
	if startDate.IsZero() {
		project.StartDate = nil
	} else {
		project.StartDate = &timestamppb.Timestamp{
			Seconds: startDate.Unix(),
			Nanos:   int32(startDate.Nanosecond()),
		}

		project.EndDate = &timestamppb.Timestamp{
			Seconds: endDate.Unix(),
			Nanos:   int32(endDate.Nanosecond()),
		}
	}
	return &project, err
}

func (p ProjectRepository) GetAll(username string) ([]*projectPb.ProjectFields, error) {
	query := `
			SELECT id, title, description, leader, member_ids, guide_name, project_url, semester, start_date, end_date
			FROM projects
			WHERE username = $1
			`
	rows, err := p.DB.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	var projects []*projectPb.ProjectFields
	for rows.Next() {
		var project projectPb.ProjectFields
		var startDate time.Time
		var endDate time.Time
		err := rows.Scan(
			&project.Id,
			&project.Name,
			&project.Description,
			&project.Leader,
			pq.Array(&project.MembersUsername),
			&project.GuideName,
			&project.ProjectUrl,
			&project.Semester,
			&startDate,
			&endDate,
		)

		if err != nil {
			return nil, err
		}
		if startDate.IsZero() {
			project.StartDate = nil
		} else {
			project.StartDate = &timestamppb.Timestamp{
				Seconds: startDate.Unix(),
				Nanos:   int32(startDate.Nanosecond()),
			}

			project.EndDate = &timestamppb.Timestamp{
				Seconds: endDate.Unix(),
				Nanos:   int32(endDate.Nanosecond()),
			}
		}
		projects = append(projects, &project)
	}
	return projects, nil
}

func (p ProjectRepository) Update(username string, project *projectPb.ProjectFields) error {
	query := `
			UPDATE projects
			SET title = $1, description = $2, leader = $3, member_ids = $4, guide_name = $5, project_url = $6, semester = $7, start_date = $8, end_date = $9
			WHERE id = $10 AND username = $11
			`
	var startDate time.Time
	var endDate time.Time
	if project.GetStartDate() != nil {
		startDate = project.GetStartDate().AsTime()
		startDate = startDate.Add(time.Duration(project.StartDate.Seconds))
		startDate = startDate.AddDate(0, 0, 1)
	}
	if project.GetEndDate() != nil {
		endDate = project.GetEndDate().AsTime()
		endDate = endDate.Add(time.Duration(project.GetEndDate().GetSeconds()))
		endDate = endDate.AddDate(0, 0, 1)
	}
	_, err := p.DB.Exec(
		query,
		project.GetName(),
		project.GetDescription(),
		project.GetLeader(),
		pq.Array(project.GetMembersUsername()),
		project.GetGuideName(),
		project.GetProjectUrl(),
		project.GetSemester(),
		startDate,
		endDate,
		project.GetId(),
		username,
	)
	return err
}

func (p ProjectRepository) Delete(username string, projectId int64) error {
	query := `
			DELETE FROM projects
			WHERE id = $1 AND username = $2
			`
	_, err := p.DB.Exec(query, projectId, username)
	return err
}

func (p ProjectRepository) GetProjectsBySemester(username string, semester int64) ([]*projectPb.ProjectFields, error) {
	query := `
			SELECT id, title, description, leader, member_ids, guide_name, project_url, semester, start_date, end_date
			FROM projects
			WHERE username = $1 AND semester = $2
			`
	rows, err := p.DB.Query(query, username, semester)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	var projects []*projectPb.ProjectFields
	for rows.Next() {
		var project projectPb.ProjectFields
		var startDate time.Time
		var endDate time.Time
		err := rows.Scan(
			&project.Id,
			&project.Name,
			&project.Description,
			&project.Leader,
			&project.MembersUsername,
			&project.GuideName,
			&project.ProjectUrl,
			&project.Semester,
			&startDate,
			&endDate,
		)
		if startDate.IsZero() {
			project.StartDate = nil
		} else {
			project.StartDate = &timestamppb.Timestamp{
				Seconds: startDate.Unix(),
				Nanos:   int32(startDate.Nanosecond()),
			}

			project.EndDate = &timestamppb.Timestamp{
				Seconds: endDate.Unix(),
				Nanos:   int32(endDate.Nanosecond()),
			}
		}
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}
	return projects, nil
}
