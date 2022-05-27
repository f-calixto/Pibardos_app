package group

import (
	// std lib
	"database/sql"
	"fmt"
	"strings"

	// Internal
	"github.com/coding-kiko/group_service/pkg/log"
)

var (
	createGroupQuery = `INSERT INTO groups(id, name, size, country, admin_id, access_code, access_code_issue_time, avatar, created_at) 
						Values($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	joinUsersGroupsQuery = `INSERT INTO users_groups(user_id, group_id) Values($1, $2)`
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

type Repository interface {
	// JoinGroup(req JoinGroupRequest) (JoinGroupResponse, error)
	CreateGroup(group Group) (CreateGroupResponse, error)
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

func (r *repo) CreateGroup(group Group) (CreateGroupResponse, error) {
	r.logger.Info("repository.go", "CreateGroup", "creating group")
	_, err := r.db.Exec(createGroupQuery, group.Id, group.Name, group.Size, group.Country, group.Admin_id, group.Access_code, group.Access_code_issue_time, group.Avatar_route, group.Created_at)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			fmt.Println("UNIQUE constraint violated")
		}
		r.logger.Error("repository.go", "CreateGroup", err.Error())
		return CreateGroupResponse{}, err
	}
	// _, err = r.db.Exec(joinUsersGroupsQuery, group.Admin_id, group.Id)
	// if err != nil {
	// 	r.logger.Error("repository.go", "CreateGroup", err.Error())
	// 	return CreateGroupResponse{}, err
	// }
	r.logger.Info("repository.go", "CreateGroup", "group created successfully")
	return CreateGroupResponse{
		Id:           group.Id,
		Name:         group.Name,
		Admin_id:     group.Admin_id,
		Access_code:  group.Access_code,
		Size:         group.Size,
		Created_at:   group.Created_at,
		Avatar_route: group.Avatar_route,
		Country:      group.Country,
	}, nil
}

// func (r *repo) JoinGroup(req JoinGroupRequest) (JoinGroupResponse, error) {

// }
