package models

import (
	"database/sql"
)

type Permissions struct {
	RoleID   int
	AccessID int
}

func (perm *Permissions) CreatePermissions(db *sql.DB) (err error) {
	var permID int
	query := `INSERT INTO permissions (role_id, access_id) VALUES($1, $2) RETURNING id;`

	err = db.QueryRow(query, perm.RoleID, perm.AccessID).Scan(&permID)
	if err != nil {
		return
	}
	return
}

func (perm *Permissions) GetPermissions(db *sql.DB) (err error) {
	query := `SELECT access_id FROM permissions WHERE role_id=$1`

	err = db.QueryRow(query, perm.RoleID).Scan(&perm.AccessID)
	if err != nil {
		return
	}
	return
}
