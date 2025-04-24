package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/AugustSerenity/service_auth/internal/model"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) SaveToken(ctx context.Context, refresh *model.RefreshToken) error {
	fmt.Println("Сохраняем токен")

	query := `
		INSERT INTO refresh_tokens (user_id, access_jti, hashed_token, ip, used, created_at)
		VALUES ($1, $2, $3, $4, false, NOW())
	`

	_, err := s.db.ExecContext(ctx, query, refresh.UserID, refresh.AccessJTI, refresh.HashedToken, refresh.IP)
	if err != nil {
		log.Printf("Ошибка при создании токена: %v", err)
		return err
	}

	fmt.Println("Токен успешно сохранен")
	return nil
}

func (s *Storage) FindRefreshToken(ctx context.Context, userID, accessJTI string) (*model.RefreshToken, error) {
	query := `
		SELECT id, user_id, access_jti, hashed_token, ip, used, created_at
		FROM refresh_tokens
		WHERE user_id = $1 AND access_jti = $2
		LIMIT 1
	`

	row := s.db.QueryRowContext(ctx, query, userID, accessJTI)

	var token model.RefreshToken
	err := row.Scan(
		&token.ID,
		&token.UserID,
		&token.AccessJTI,
		&token.HashedToken,
		&token.IP,
		&token.Used,
		&token.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("refresh token not found")
		}
		return nil, fmt.Errorf("failed to query refresh token: %w", err)
	}

	return &token, nil
}

func (s *Storage) MarkTokenUsed(ctx context.Context, tokenID uint) error {
	query := `UPDATE refresh_tokens 
	          SET used = true 
	          WHERE id = $1 AND used = false`

	_, err := s.db.ExecContext(ctx, query, tokenID)
	if err != nil {
		log.Printf("Error marking token as used: %v", err)
		return fmt.Errorf("could not mark token as used: %w", err)
	}

	return nil
}
