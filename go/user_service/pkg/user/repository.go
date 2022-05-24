package user

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/coding-kiko/user_service/pkg/log"
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

type Repository interface {
	UpdateUser(user UpsertUserRequest) (User, error)
	InsertUser(user User) (User, error)
	UpdateAvatar(user User) (User, error)
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
	getUserQuery      = `SELECT id, username, email, created_at, country, birthdate, status, avatar FROM users WHERE id = $1`
	updateAvatarQuery = `UPDATE users SET avatar = $1 WHERE id = $2`
	baseUpdateQuery   = `UPDATE users SET `
)

// update user avatar by id
func (r *repo) UpdateAvatar(user User) (User, error) {
	_, err := r.db.Exec(updateAvatarQuery, user.Avatar, user.Id)
	if err != nil {
		r.logger.Error("repository.go", "InsertUser", err.Error())
		return User{}, err // probably user not found
	}

	r.logger.Info("repository.go", "InsertUser", "avatar updated successfully")
	return user, nil
}

// insert new user in the database
func (r *repo) InsertUser(user User) (User, error) {
	r.logger.Info("repository.go", "InsertUser", "inserting user")

	_, err := r.db.Exec(insertUserQuery, user.Id, user.Username, user.Email, user.Created_at, user.Country, user.Birthdate, user.Status, user.Avatar)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return User{}, errors.New("username or email already in use") // this should be validated by auth
		}
		r.logger.Error("repository.go", "InsertUser", err.Error())
		return User{}, err
	}

	r.logger.Info("repository.go", "InsertUser", "user inserted successfully")
	return user, nil
}

// Update one or many arbitrary values from user - not including avatar
func (r *repo) UpdateUser(req UpsertUserRequest) (User, error) {
	r.logger.Info("repository.go", "InsertUser", "updating user")

	updateQuery, args, err := PatchQueryConstructor(req)
	if err != nil {
		r.logger.Error("repository.go", "UpdateUser", err.Error())
		return User{}, err
	}
	_, err = r.db.Exec(updateQuery, args)
	if err != nil {
		r.logger.Error("repository.go", "UpdateUser", err.Error())
		if strings.Contains(err.Error(), "duplicate") {
			return User{}, errors.New("username or email already in use") // this should be validated by auth
		}
		return User{}, err // probably invalid id
	}

	updatedUser := User{}
	err = r.db.QueryRow(getUserQuery, req.Id).Scan(&updatedUser.Id, &updatedUser.Username, &updatedUser.Email, &updatedUser.Created_at, &updatedUser.Country, &updatedUser.Birthdate, &updatedUser.Status, &updatedUser.Avatar)
	if err != nil {
		r.logger.Error("repository.go", "UpdateUser", err.Error())
		return User{}, err
	}

	r.logger.Info("repository.go", "updateUser", "user updated successfully")
	return updatedUser, nil
}

// creates update query dynamically depending on the fields to be updated
func PatchQueryConstructor(req UpsertUserRequest) (string, interface{}, error) {
	i := 1 // increments accordingly eith the number of args
	queryParts := make([]string, 0)
	args := make([]interface{}, 0)

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
		return "", nil, errors.New("no valid field selected for update")
	}
	baseUpdateQuery += strings.Join(queryParts, ", ") + fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, *req.Id)

	return baseUpdateQuery, args, nil
}
