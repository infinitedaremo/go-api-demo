-- name: CreatePerson :one
INSERT INTO person (first_name, last_name)
VALUES (?, ?)
RETURNING *;

-- name: GetPerson :one
SELECT first_name, last_name FROM person
WHERE id = ? LIMIT 1;

-- name: GetPortfolio :many
SELECT sqlc.embed(person), sqlc.embed(work_experience)
FROM person
JOIN work_experience ON work_experience.person_id = person.id
WHERE work_experience.person_id = ?
ORDER BY work_experience.start_date DESC
LIMIT 25;

-- name: CreateWorkExperience :exec
INSERT INTO work_experience (person_id, company_name, job_title, start_date, end_date)
VALUES (?, ?, ?, ?, ?);