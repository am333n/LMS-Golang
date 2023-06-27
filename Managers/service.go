package managers

import (
	"context"
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
	Status string `json:"status"`
}
type Service interface {
	GetManagers(ctx context.Context) ([]Manager, error)
	PostManager(ctx context.Context,manager Manager) (string, error)
	GetManagerById(ctx context.Context,id int) (Manager, error)
	DeleteManager(ctx context.Context,id int) (string, error)
	UpdateManager(ctx context.Context,id int, manager Manager) (string, error)
}

func NewService(controller Controller) Service {
	return &RepoService{
		controller:controller,
	}

}

type RepoService struct {
	controller Controller
}

func (s RepoService) GetManagers(ctx context.Context) ([]Manager, error) {
	_, err := s.controller.CheckIfAdmin(ctx)
	if err!=nil{
		return nil,err
	}
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
func (s RepoService) PostManager(ctx context.Context,manager Manager) (string, error) {
	_, err := s.controller.CheckIfAdmin(ctx)
	if err!=nil{
		return "",err
	}
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
func (s RepoService) GetManagerById(ctx context.Context,id int) (Manager, error) {
	_, err := s.controller.CheckIfAdmin(ctx)
	if err!=nil{
		return Manager{},err
	}
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
func (s RepoService) DeleteManager(ctx context.Context,id int) (string, error) {
	_, err := s.controller.CheckIfAdmin(ctx)
	if err!=nil{
		return "",err
	}
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
func (s RepoService) UpdateManager(ctx context.Context,id int, manager Manager) (string, error) {
	_, err := s.controller.CheckIfAdmin(ctx)
	if err!=nil{
		return "",err
	}
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
