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
	//pending or accepted status
	Status     string `json:"status"`
}
type Service interface {
	GetManagers() ([]Manager, error)
	PostManager(manager Manager) (string, error)
	GetManagerById(id int) (Manager, error)
	DeleteManager(id int) (string, error)
	UpdateManager(id int, manager Manager) (string, error)
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
func (RepoService) GetManagerById(id int) (Manager, error) {
	db, err := dc.GetDB()
	var manager Manager
	if err != nil {
		return Manager{}, err
	}
	err = db.Where("Id=?", id).First(&manager).Error
	if err != nil {
		return Manager{}, err
	}
	return manager, nil
}
func (RepoService) DeleteManager(id int) (string, error) {
	db, err := dc.GetDB()
	var manager Manager
	if err != nil {
		return "", err
	}
	err = db.Where("Id=?", id).Delete(&manager).Error
	if err != nil {
		return "", err
	}
	return "Manager Deleted", nil

}
func (RepoService) UpdateManager(id int, manager Manager) (string, error) {
	db, err := dc.GetDB()
	if err != nil {
		return "", err
	}
	err = db.Where("id=?", id).Updates(&manager).Error
	if err != nil {
		return "", err
	}
	return "Manager Successfully Updated", nil
}
