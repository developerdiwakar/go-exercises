// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: contact.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createContact = `-- name: CreateContact :one
INSERT INTO contacts(
    first_name,
    last_name,
    phone_number,
    street,
    created_at,
    updated_at
) VALUES(
    $1,$2,$3,$4,$5,$6
) RETURNING id, first_name, last_name, phone_number, street, created_at, updated_at
`

type CreateContactParams struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Street      string    `json:"street"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (q *Queries) CreateContact(ctx context.Context, arg CreateContactParams) (Contact, error) {
	row := q.queryRow(ctx, q.createContactStmt, createContact,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.Street,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Contact
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.Street,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteContact = `-- name: DeleteContact :exec
DELETE FROM contacts
WHERE id = $1
`

func (q *Queries) DeleteContact(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteContactStmt, deleteContact, id)
	return err
}

const getContactById = `-- name: GetContactById :one
SELECT id, first_name, last_name, phone_number, street, created_at, updated_at 
FROM contacts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetContactById(ctx context.Context, id uuid.UUID) (Contact, error) {
	row := q.queryRow(ctx, q.getContactByIdStmt, getContactById, id)
	var i Contact
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.Street,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listContactsMany = `-- name: ListContactsMany :many
SELECT id, first_name, last_name, phone_number, street, created_at, updated_at 
FROM contacts
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListContactsManyParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListContactsMany(ctx context.Context, arg ListContactsManyParams) ([]Contact, error) {
	rows, err := q.query(ctx, q.listContactsManyStmt, listContactsMany, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Contact{}
	for rows.Next() {
		var i Contact
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.PhoneNumber,
			&i.Street,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateContact = `-- name: UpdateContact :one
UPDATE contacts
SET
    first_name  = coalesce($1, first_name),
    last_name   = coalesce($2, last_name),
    phone_number= coalesce($3, phone_number),
    street      = coalesce($4, street),
    updated_at  = coalesce($5, updated_at)
WHERE
    id = $6
RETURNING id, first_name, last_name, phone_number, street, created_at, updated_at
`

type UpdateContactParams struct {
	FirstName   sql.NullString `json:"first_name"`
	LastName    sql.NullString `json:"last_name"`
	PhoneNumber sql.NullString `json:"phone_number"`
	Street      sql.NullString `json:"street"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	ID          uuid.UUID      `json:"id"`
}

func (q *Queries) UpdateContact(ctx context.Context, arg UpdateContactParams) (Contact, error) {
	row := q.queryRow(ctx, q.updateContactStmt, updateContact,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.Street,
		arg.UpdatedAt,
		arg.ID,
	)
	var i Contact
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.Street,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}