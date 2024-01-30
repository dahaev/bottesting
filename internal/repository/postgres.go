package repository

import (
	"context"
	"time"

	"github.com/dahaev/bottesting/pkg/models"
)

func (repo *Repository) CreateLadyAccount(ctx context.Context, account *models.Account) error {
	registrationTime := time.Now()
	stmt, err := repo.db.PrepareContext(ctx, "INSERT INTO ladies(account_name, created_at, location, region, description) VALUES($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, account.UserName, registrationTime, account.Location, account.Region, account.Description)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetAccountLady(ctx context.Context, userName string) (*models.Account, error) {
	stmt, err := repo.db.PrepareContext(ctx, "SELECT user_id, account_name, created_at, rating, location, region, description from ladies where account_name=$1")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, userName)
	var user models.Account

	err = row.Scan(&user.ID, &user.UserName, &user.Created, &user.Rating, &user.Location, &user.Region, &user.Description)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *Repository) CreateDonAccount(ctx context.Context, username string) error {
	now := time.Now()
	query := "INSERT INTO don(username, created_at) VALUES($1,$2)"
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, username, now)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetDonAccount(ctx context.Context, userName string) (*models.DonAccount, error) {
	stmt, err := repo.db.PrepareContext(ctx, "SELECT user_id, username, rating, created_at from don where username=$1")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, userName)
	var user models.DonAccount
	err = row.Scan(&user.ID, &user.UserName, &user.Rating, &user.Created)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *Repository) CreateReview(ctx context.Context, ladyName string, donName string, review string, rating int) error {
	reviewTime := time.Now()
	query := "INSERT INTO reviews(lady_id, don_id, description,rating, created_at) VALUES ($1,$2,$3,$4, $5)"
	don, err := repo.GetDonAccount(ctx, donName)
	if err != nil {
		return err
	}
	lady, err := repo.GetAccountLady(ctx, ladyName)
	if err != nil {
		return err
	}
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, lady.ID, don.ID, review, rating, reviewTime)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetLadyReviews(ctx context.Context, ladyName string) ([]models.Review, error) {
	query := "SELECT reviews.id,don.username,reviews.description, reviews.rating,reviews.created_at from reviews inner join ladies ON reviews.lady_id = ladies.user_id inner join don on reviews.don_id = don.user_id\nwhere ladies.user_id = $1"
	lady, err := repo.GetAccountLady(ctx, ladyName)
	if err != nil {
		return nil, err
	}
	reviews := make([]models.Review, 0)
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx, lady.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		review := models.Review{}
		err = rows.Scan(&review.ID, &review.DonName, &review.Description, &review.Rating, &review.Date)
		if err != nil {
			continue
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}
