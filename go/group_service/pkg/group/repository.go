package group

import (
	// std lib
	"database/sql"
	"fmt"
	"strings"
	"time"

	// Internal
	"github.com/coding-kiko/group_service/pkg/errors"
	"github.com/coding-kiko/group_service/pkg/log"
)

var (
	createGroupQuery = `INSERT INTO groups(id, name, size, country, admin_id, access_code, access_code_expiration_time, avatar, created_at, description) 
						Values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	joinUsersGroupsQuery    = `INSERT INTO users_groups(user_id, group_id) Values($1, $2)`
	getGroupQuery           = `SELECT id, name, size, country, admin_id, avatar, created_at, description FROM groups WHERE id = $1`
	baseUpdateQuery         = `UPDATE groups SET `
	updateAvatarQuery       = `UPDATE groups SET avatar = $1 WHERE id = $2`
	accessCodeQuery         = `UPDATE groups SET access_code = $1, access_code_expiration_time = $2 WHERE id = $3`
	checkGroupAdminQuery    = `SELECT admin_id FROM groups WHERE id = $1`
	validateAccessCodeQuery = `SELECT id, access_code_expiration_time FROM groups WHERE access_code = $1`
	increaseGroupSize       = `UPDATE groups SET size = size + 1 WHERE id = $1`
	getGroupMembersQuery    = `SELECT users.id, username, email, created_at, country, birthdate, status, avatar FROM users_groups INNER JOIN users ON users.id = users_groups.user_id WHERE users_groups.group_id = $1`
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

type Repository interface {
	CreateGroup(group Group) (Group, error)
	GetGroup(id string) (Group, error)
	UpdateGroup(req UpdateGroupRequest) (Group, error)
	UpdateAvatar(req UpdateAvatarRequest) (Group, error)
	StoreAccessCode(req AccessCode) (AccessCode, error)
	JoinGroup(req AccessCode) (Group, error)
	GetGroupMembers(req GetGroupMembersRequest) ([]User, error)
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

func (r *repo) GetGroupMembers(req GetGroupMembersRequest) ([]User, error) {
	i := 1
	members := []User{}

	rows, err := r.db.Query(getGroupMembersQuery, req.Id)
	if err != nil {
		return []User{}, errors.NewNotFound()
	}
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Created_at, &user.Country, &user.Birthdate, &user.Status, &user.Avatar)
		if err != nil {
			return []User{}, errors.NewNotFound()
		}
		members = append(members, user)
		if i == req.Amount {
			break
		}
		i++
	}
	r.logger.Info("repository.go", "GetGroupMembers", "group members retrieved successfully")
	return members, nil
}

func (r *repo) JoinGroup(req AccessCode) (Group, error) {
	var expiration int64

	// look for any groups with the requested access code
	err := r.db.QueryRow(validateAccessCodeQuery, req.AccessCode).Scan(&req.GroupId, &expiration)
	if err != nil {
		return Group{}, errors.NewInvalidAccessCode("invalid access code")
	}
	// if acces code matches but already expires
	if time.Now().Unix() > expiration {
		return Group{}, errors.NewInvalidAccessCode("access code expired")
	}

	// finally inner join user and group
	_, err = r.db.Exec(joinUsersGroupsQuery, req.UserId, req.GroupId)
	if err != nil {
		if strings.Contains(err.Error(), "unique") {
			return Group{}, errors.NewPermissionForbidden("you already belong to that group")
		}
		return Group{}, errors.NewNotFound("group not found")
	}
	// increase group size
	_, err = r.db.Exec(increaseGroupSize, req.GroupId)
	if err != nil {
		return Group{}, errors.NewNotFound("group not found")
	}

	updatedGroup, err := r.GetGroup(req.GroupId)
	if err != nil {
		return Group{}, err
	}

	r.logger.Info("repository.go", "JoinGroup", "joined group successfully")
	return updatedGroup, nil
}

func (r *repo) StoreAccessCode(code AccessCode) (AccessCode, error) {
	var admin_id string

	// check if user requesting access code is the group admin
	err := r.db.QueryRow(checkGroupAdminQuery, code.GroupId).Scan(&admin_id)
	if err != nil {
		return AccessCode{}, errors.NewNotFound("group not found")
	}

	if admin_id != code.UserId {
		return AccessCode{}, errors.NewPermissionForbidden()
	}

	_, err = r.db.Exec(accessCodeQuery, code.AccessCode, code.Expiration, code.GroupId)
	if err != nil {
		return AccessCode{}, errors.NewNotFound("group not found")
	}

	// chanchada: to make them not show on json response
	code.UserId = ""
	code.Expiration = 0
	code.GroupId = ""

	return code, nil
}

func (r *repo) UpdateAvatar(req UpdateAvatarRequest) (Group, error) {
	_, err := r.db.Exec(updateAvatarQuery, req.Avatar, req.Id)
	if err != nil {
		return Group{}, errors.NewNotFound("group not found")
	}

	updatedGroup, err := r.GetGroup(req.Id)
	if err != nil {
		return Group{}, err
	}

	r.logger.Info("repository.go", "UpdateAvatar", "avatar updated successfully")
	return updatedGroup, nil
}

// update group dynaically
func (r *repo) UpdateGroup(req UpdateGroupRequest) (Group, error) {
	updateQuery, args, err := PatchQueryConstructor(req)
	if err != nil {
		return Group{}, err
	}

	_, err = r.db.Exec(updateQuery, args...)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return Group{}, errors.NewInvalidCredentials("group name already in use")
		}
		return Group{}, err
	}

	updatedGroup, err := r.GetGroup(*req.Id)
	if err != nil {
		return Group{}, err
	}

	r.logger.Info("repository.go", "UpdateUser", "group updated successfully")
	return updatedGroup, nil
}

// Get group by Id
func (r *repo) GetGroup(id string) (Group, error) {
	group := Group{}
	err := r.db.QueryRow(getGroupQuery, id).Scan(&group.Id, &group.Name, &group.Size, &group.Country, &group.Admin_id, &group.Avatar, &group.Created_at, &group.Description)
	if err != nil {
		return Group{}, errors.NewNotFound()
	}
	return group, nil
}

func (r *repo) CreateGroup(group Group) (Group, error) {
	_, err := r.db.Exec(createGroupQuery, group.Id, group.Name, group.Size, group.Country, group.Admin_id, group.Access_code, group.Access_code_expiration_time, group.Avatar, group.Created_at, group.Description)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return Group{}, errors.NewInvalidCredentials()
		}
		return Group{}, err
	}

	_, err = r.db.Exec(joinUsersGroupsQuery, group.Admin_id, group.Id)
	if err != nil {
		return Group{}, err
	}

	r.logger.Info("repository.go", "CreateGroup", "group created successfully")
	return Group{
		Id:          group.Id,
		Name:        group.Name,
		Admin_id:    group.Admin_id,
		Access_code: group.Access_code,
		Size:        group.Size,
		Created_at:  group.Created_at,
		Avatar:      group.Avatar,
		Country:     group.Country,
		Description: group.Description,
	}, nil
}

// func (r *repo) JoinGroup(req JoinGroupRequest) (JoinGroupResponse, error) {

// }

// creates group update query dynamically depending on the fields to be updated
func PatchQueryConstructor(req UpdateGroupRequest) (string, []interface{}, error) {
	i := 1 // increments accordingly eith the number of args
	queryParts := make([]string, 0)
	args := make([]interface{}, 0)
	query := baseUpdateQuery

	if req.Name != nil {
		field := fmt.Sprintf("name = $%d", i)
		queryParts = append(queryParts, field)
		args = append(args, strings.ToLower(*req.Name))
		i++
	}

	if req.Admin_id != nil {
		queryParts = append(queryParts, fmt.Sprintf("email = $%d", i))
		args = append(args, strings.ToLower(*req.Admin_id))
		i++
	}

	if req.Country != nil {
		queryParts = append(queryParts, fmt.Sprintf("country = $%d", i))
		args = append(args, strings.ToUpper(*req.Country))
		i++
	}

	if req.Description != nil {
		queryParts = append(queryParts, fmt.Sprintf("description = $%d", i))
		args = append(args, strings.TrimSpace(*req.Description))
		i++
	}

	if len(queryParts) == 0 {
		return "", nil, errors.NewInvalidUpdate()
	}
	query += strings.Join(queryParts, ", ") + fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, *req.Id)

	return query, args, nil
}
