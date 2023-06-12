package managers

import (
	dc "lms/Database"

	"gorm.io/gorm"
)

type Manager struct {
	gorm.Model
	Name       string `json:"name"`
	Department string `json:"department"`
	Address    string `json:"address"`
	DOB        string `json:"dob"`
	Email      string `json:"email"`
	Phone      int    `json:"phone"`
}
type Service interface {
	GetManagers() ([]Manager, error)
	PostManager(manager Manager) (string, error)
}
type RepoService struct{}

func (RepoService) GetManagers() ([]Manager, error) {
	var managers []Manager
	db, err := dc.GetDB()
	if err != nil {
		return managers, err
	}
	err = db.Find(&managers).Error
	if err != nil {
		return managers, err
	}
	return managers, nil

}
func (RepoService) PostManager(manager Manager) (string, error) {
	db, err := dc.GetDB()
	if err != nil {
		return "", err
	}
	err = db.Create(&manager).Error
	if err != nil {
		return "", err
	}
	return "New Manager Added", nil

}
