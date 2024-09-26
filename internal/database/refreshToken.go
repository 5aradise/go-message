package database

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	Token     string
	ExpiresAt time.Time
	UserID    uint
}

func (sl *Database) UpdateRefreshTokenByUserName(name string, newToken string) error {
	res := sl.gormDB.
		Model(&RefreshToken{}).
		Where("user_id = (SELECT id FROM users WHERE name = ?)", name).
		Updates(map[string]interface{}{
			"token":      newToken,
			"expires_at": time.Now().Add(30 * 24 * time.Hour),
		})
	return res.Error
}
