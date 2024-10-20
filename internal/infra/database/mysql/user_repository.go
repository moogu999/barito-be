package mysql

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/moogu999/barito-be/internal/user/domain/entity"
	"github.com/moogu999/barito-be/internal/user/domain/repository"
)

// @TODO refactor this
type UserModel struct {
	id        int64
	email     string
	password  string
	createdAt time.Time
	createdBy string
}

func (u UserModel) toUserEntity() entity.User {
	return entity.User{
		ID:        u.id,
		Email:     u.email,
		Password:  u.password,
		CreatedAt: u.createdAt,
		CreatedBy: u.createdBy,
	}
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	builder := sq.Select("id", "email", "password", "created_at", "created_by").
		From("users").
		Where(sq.Eq{"email": email})

	q, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]UserModel, 0)
	for rows.Next() {
		var user UserModel
		err = rows.Scan(
			&user.id,
			&user.email,
			&user.password,
			&user.createdAt,
			&user.createdBy,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, nil
	}

	user := users[0].toUserEntity()

	return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	builder := sq.Insert("users").
		Columns("email", "password", "created_at", "created_by").
		Values(user.Email, user.Password, user.CreatedAt, user.CreatedBy)
	q, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	res, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id

	return nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	builder := sq.Select("id", "email", "password", "created_at", "created_by").
		From("users").
		Where(sq.Eq{"id": id})

	q, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]UserModel, 0)
	for rows.Next() {
		var user UserModel
		err = rows.Scan(
			&user.id,
			&user.email,
			&user.password,
			&user.createdAt,
			&user.createdBy,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, nil
	}

	user := users[0].toUserEntity()

	return &user, nil
}
