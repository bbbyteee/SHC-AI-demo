package session

import (
	"errors"
	"shc-ai-demo/common/mysql"
	"shc-ai-demo/model"
	"strings"
)

func GetSessionsByUserName(UserName int64) ([]model.Session, error) {
	var sessions []model.Session
	err := mysql.DB.Where("user_name = ?", UserName).Find(&sessions).Error
	return sessions, err
}

func CreateSession(session *model.Session) (*model.Session, error) {
	err := mysql.DB.Create(session).Error
	return session, err
}

func GetSessionByID(sessionID string) (*model.Session, error) {
	var session model.Session
	err := mysql.DB.Where("id = ?", sessionID).First(&session).Error
	return &session, err
}

func UpdateSessionTitle(id, title string) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return errors.New("title cannot be empty")
	}

	// 只更新 title
	err := mysql.DB.Model(&model.Session{}).
		Where("id = ?", id).
		Update("title", title).Error

	return err
}
