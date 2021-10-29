package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"service-app/auth"
	"strconv"
	"time"
)

var (
	// ErrNotFound is used when a specific User is requested but does not exist.
	ErrNotFound = errors.New("not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in its proper form")

	// ErrAuthenticationFailure occurs when a user attempts to authenticate but
	// anything goes wrong.
	ErrAuthenticationFailure = errors.New("authentication failed")

	// ErrForbidden occurs when a user tries to do something that is forbidden to them according to our access control policies.
	ErrForbidden = errors.New("attempted action is not allowed")
)

type DbService struct {
	*sql.DB
}

func NewDbService(db *sql.DB) *DbService {
	return &DbService{
		DB: db,
	}
}

func (db *DbService) Create(ctx context.Context, nu NewUser, now time.Time) (User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("generating password hash %w", err)
	}

	u := User{
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: hash,
		Roles:        nu.Roles,
		DateCreated:  now.UTC(),
		DateUpdated:  now.UTC(),
	}

	const q = `INSERT INTO users
		(name, email, password_hash, roles, date_created, date_updated)
		VALUES ( $1, $2, $3, $4, $5, $6)
		Returning id`

	var id int
	if err = db.QueryRowContext(ctx, q, u.Name, u.Email, u.PasswordHash, u.Roles, u.DateCreated, u.DateUpdated).Scan(&id); err != nil {
		return User{}, fmt.Errorf("inserting user %w", err)
	}

	u.ID = strconv.Itoa(id)
	return u, nil
}

// Authenticate finds a user by their email and verifies their password. On
// success it returns a Claims value representing this user. The claims can be
// used to generate a token for future authentication.
func (db *DbService) Authenticate(ctx context.Context, now time.Time, email, password string) (auth.Claims, error) {

	const q = `SELECT id,name,email,roles,password_hash FROM users WHERE email = $1`

	var u User

	err := db.QueryRowContext(ctx, q, email).Scan(&u.ID, &u.Name, &u.Email, &u.Roles, &u.PasswordHash)

	if err != nil {

		// Normally we would return ErrNotFound in this scenario but we do not want
		// to leak to an unauthenticated user which emails are in the system.
		if err == sql.ErrNoRows {
			return auth.Claims{}, ErrAuthenticationFailure
		}

		return auth.Claims{}, fmt.Errorf("selecting single user %w", err)
	}

	// Compare the provided password with the saved hash. Use the bcrypt
	if err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password)); err != nil {
		return auth.Claims{}, ErrAuthenticationFailure
	}

	// If we are this far the request is valid. Create some claims for the user
	// and generate their token.
	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "service project",
			Subject:   u.ID,
			Audience:  jwt.ClaimStrings{"students"},
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Roles: u.Roles,
	}

	return claims, nil
}

// Delete removes a user from the database.
func (db *DbService) Delete(ctx context.Context, email string) error { // id earlier in func param


	const q = `DELETE FROM users WHERE email = $1`

	if _, err := db.ExecContext(ctx, q, email); err != nil {
		return fmt.Errorf("deleting user %s %w", email, err)
	}

	return nil
}

// List retrieves a list of existing users from the database.
func (db *DbService) List(ctx context.Context) ([]User, error) {

	const q = `SELECT * FROM users`
	const q2 = `SELECT id,name,email,roles,password_hash FROM users`

	var users []User

	rows, err := db.QueryContext(ctx, q2)
	if err != nil {
		return nil, fmt.Errorf("selecting users %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var u User

		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Roles, &u.PasswordHash)

		if err != nil {
			return nil, fmt.Errorf("scanning users %w", err)
		}
		users = append(users, u)

	}

	return users, nil
}
