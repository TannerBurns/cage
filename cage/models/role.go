package models

import (
	"database/sql"
)

type Role struct {
	ID         int    `json:"id"`
	EmployeeID int    `json:"employee_id"`
	Role       string `json:"role"`
	RoleID     int    `json:"role_id"`
	Notes      string `json:"notes"`
}

type Roles struct {
	EmployeeID int    `json:"employee_id"`
	Roles      []Role `json:"roles"`
}

func (role *Role) AddRole(db *sql.DB) (err error) {
	query := `INSERT INTO roles (employee_id, role, role_id, notes) VALUES($1, $2, $3, $4) RETURNING id;`

	roles := Roles{EmployeeID: role.EmployeeID}
	err = roles.GetRoles(db)
	if err != nil {
		return
	}
	for _, r := range roles.Roles {
		if r.RoleID == role.RoleID {
			*role = r
			return
		}
	}

	err = db.QueryRow(query, role.EmployeeID, role.Role, role.RoleID, role.Notes).Scan(&role.ID)
	if err != nil {
		return
	}
	return
}

func (roles *Roles) GetRoles(db *sql.DB) (err error) {
	query := `SELECT * FROM roles WHERE employee_id=$1`

	rows, err := db.Query(query, roles.EmployeeID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		role := Role{}
		err = rows.Scan(&role.ID, &role.EmployeeID, &role.Role, &role.RoleID, &role.Notes)
		if err != nil {
			return
		}
		roles.Roles = append(roles.Roles, role)
	}
	return
}

func (role *Role) GetAccessLevel(db *sql.DB) (aLevel int, err error) {
	aLevel = -1
	permission := Permissions{RoleID: role.RoleID}
	err = permission.GetPermissions(db)
	if err != nil {
		return
	}
	aLevel = permission.AccessID
	return
}
