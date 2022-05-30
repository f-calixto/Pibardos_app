package user

import (
	// std lib
	"database/sql"
	"fmt"
	"strings"

	// Internal
	"github.com/coding-kiko/user_service/pkg/errors"
	"github.com/coding-kiko/user_service/pkg/log"
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

type Repository interface {
	UpdateUser(req UpsertUserRequest) (User, error)
	InsertUser(user User) (User, error)
	UpdateAvatar(req UpdateAvatarRequest) (User, error)
	GetUser(id string) (User, error)
	GetUserGroups(id string) ([]Group, error)
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

var (
	insertUserQuery = `INSERT INTO users(id, username, email, created_at, country, birthdate, status, avatar) 
						Values($1, $2, $3, $4, $5, $6, $7, $8)`
	getUserQuery       = `SELECT id, username, email, created_at, country, birthdate, status, avatar FROM users WHERE id = $1`
	updateAvatarQuery  = `UPDATE users SET avatar = $1 WHERE id = $2`
	baseUpdateQuery    = `UPDATE users SET `
	getUserGroupsQuery = `SELECT groups.id, name, size, admin_id, country, avatar, created_at, description FROM users_groups INNER JOIN groups ON groups.id = users_groups.group_id WHERE users_groups.user_id = $1`
)

// Get all of one user's groups
func (r *repo) GetUserGroups(id string) ([]Group, error) {
	groups := []Group{}

	rows, err := r.db.Query(getUserGroupsQuery, id)
	if err != nil {
		return []Group{}, err
	}
	defer rows.Close()

	for rows.Next() {
		group := Group{}
		err := rows.Scan(&group.Id, &group.Name, &group.Size, &group.Admin_id, &group.Country, &group.Avatar, &group.Created_at, &group.Description)
		if err != nil {
			return []Group{}, err
		}
		groups = append(groups, group)
	}
	r.logger.Info("repository.go", "GetUserGroups", "User groups retrieved successfully")
	return groups, nil
}

// Get user by id
func (r *repo) GetUser(id string) (User, error) {
	user := User{}
	err := r.db.QueryRow(getUserQuery, id).Scan(&user.Id, &user.Username, &user.Email, &user.Created_at, &user.Country, &user.Birthdate, &user.Status, &user.Avatar)
	if err != nil {
		return User{}, errors.NewNotFound()
	}
	return user, nil
}

// update user avatar by id
func (r *repo) UpdateAvatar(req UpdateAvatarRequest) (User, error) {
	_, err := r.db.Exec(updateAvatarQuery, req.Avatar, req.Id)
	if err != nil {
		return User{}, errors.NewNotFound()
	}

	updatedUser, err := r.GetUser(req.Id)
	if err != nil {
		return User{}, err
	}
	r.logger.Info("repository.go", "UpdateAvatar", "avatar updated successfully")
	return updatedUser, nil
}

// insert new user in the database
func (r *repo) InsertUser(user User) (User, error) {
	r.logger.Info("repository.go", "InsertUser", "inserting user")

	_, err := r.db.Exec(insertUserQuery, user.Id, user.Username, user.Email, user.Created_at, user.Country, user.Birthdate, user.Status, user.Avatar)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return User{}, errors.NewInvalidCredentials()
		}
		r.logger.Error("repository.go", "InsertUser", err.Error())
		return User{}, err
	}

	r.logger.Info("repository.go", "InsertUser", "user inserted successfully")
	return user, nil
}

// Update one or many arbitrary values from user - not including avatar
func (r *repo) UpdateUser(req UpsertUserRequest) (User, error) {
	updateQuery, args, err := PatchQueryConstructor(req)
	if err != nil {
		return User{}, err
	}

	_, err = r.db.Exec(updateQuery, args...)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return User{}, errors.NewInvalidCredentials()
		}
		return User{}, errors.NewNotFound()
	}

	updatedUser, err := r.GetUser(*req.Id)
	if err != nil {
		return User{}, err
	}

	r.logger.Info("repository.go", "UpdateUser", "user updated successfully")
	return updatedUser, nil
}

// creates user update query dynamically depending on the fields to be updated
func PatchQueryConstructor(req UpsertUserRequest) (string, []interface{}, error) {
	i := 1 // increments accordingly eith the number of args
	queryParts := make([]string, 0)
	args := make([]interface{}, 0)
	query := baseUpdateQuery

	if req.Username != nil {
		field := fmt.Sprintf("username = $%d", i)
		queryParts = append(queryParts, field)
		args = append(args, strings.ToLower(*req.Username))
		i++
	}

	if req.Email != nil {
		queryParts = append(queryParts, fmt.Sprintf("email = $%d", i))
		args = append(args, strings.ToLower(*req.Email))
		i++
	}

	if req.Country != nil {
		queryParts = append(queryParts, fmt.Sprintf("country = $%d", i))
		args = append(args, strings.ToUpper(*req.Country))
		i++
	}

	if req.Birthdate != nil {
		queryParts = append(queryParts, fmt.Sprintf("birthdate = $%d", i))
		args = append(args, *req.Birthdate)
		i++
	}

	if req.Status != nil {
		queryParts = append(queryParts, fmt.Sprintf("status = $%d", i))
		args = append(args, strings.TrimSpace(*req.Status))
		i++
	}

	if len(queryParts) == 0 {
		return "", nil, errors.NewInvalidUpdate()
	}
	query += strings.Join(queryParts, ", ") + fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, *req.Id)

	return query, args, nil
}
