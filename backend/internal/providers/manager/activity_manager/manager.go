package activity_manager

import (
	"database/sql"

	"He110/PersonalWebSite/internal/generated"
	"go.uber.org/zap"
)

type ActivityManager struct {
	db *sql.DB
	logger *zap.Logger
}

func NewActivityManager(db *sql.DB, l *zap.Logger) *ActivityManager {
	return &ActivityManager{db: db, logger: l}
}

func (m *ActivityManager) FindAll() ([]generated.ActivityItem, error) {
	query := "SELECT `title`, `description`, `image_url`, `type`, `link` FROM `activity_item` ORDER BY `sort`"
	// TODO add labels support
	rows, err := m.db.Query(query)
	if err != nil {
		m.logger.Error(`cannot select items from an activity_item table`, zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var items []generated.ActivityItem

	for rows.Next() {
		item := generated.ActivityItem{}
		if err = rows.Scan(&item.Title, &item.Description, &item.ImageUrl, &item.Type, &item.Link); err != nil {
			m.logger.Error(`error on scanning row for the activity item`, zap.Error(err))
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}