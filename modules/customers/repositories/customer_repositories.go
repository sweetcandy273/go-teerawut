package repositories

import (
	"github.com/sirupsen/logrus"
	"github.com/sweetcandy273/go-teerawut/modules/entities"

	"gorm.io/gorm"
)

type customersRepo struct {
	DB *gorm.DB
}

// NewCustomersRepository new customers repository
func NewCustomersRepository(db *gorm.DB) entities.CustomersRepository {
	return &customersRepo{
		DB: db,
	}
}

// Create create
func (r *customersRepo) Create(c *entities.Customer) error {
	err := r.DB.Create(&c).Error
	if err != nil {
		logrus.Errorf("Create customer error: %v", err)
		return err
	}

	return nil
}

// GetByID get by id
func (r *customersRepo) GetByID(id uint) (*entities.Customer, error) {
	var customer entities.Customer
	err := r.DB.Where("id = ?", id).First(&customer).Error
	if err != nil {
		logrus.Errorf("Get customer by id error: %v", err)
		return nil, err
	}

	return &customer, nil
}

// Update update
func (r *customersRepo) Update(c *entities.Customer) error {
	err := r.DB.Save(&c).Error
	if err != nil {
		logrus.Errorf("Update customer error: %v", err)
		return err
	}

	return nil
}

// GetAll get all
func (r *customersRepo) GetAll(req *entities.GetAllCustomerRequest) ([]*entities.Customer, error) {
	var customers []*entities.Customer
	err := query(r.DB, req).Find(&customers).Error
	if err != nil {
		logrus.Errorf("Query :: Get all customers error: %v", err)
		return nil, err
	}
	return customers, nil
}

func query(db *gorm.DB, req *entities.GetAllCustomerRequest) *gorm.DB {
	if req.ID != nil {
		db = db.Where("id = ?", *req.ID)
	}
	if req.Name != nil {
		db = db.Where("name = ?", *req.Name)
	}
	if req.Surname != nil {
		db = db.Where("surname = ?", *req.Surname)
	}
	if req.Nickname != nil {
		db = db.Where("nickname = ?", *req.Nickname)
	}
	if req.TelephoneNumber != nil {
		db = db.Where("telephone_number = ?", *req.TelephoneNumber)
	}
	if req.PhoneNumber != nil {
		db = db.Where("phone_number = ?", *req.PhoneNumber)
	}
	if req.Detail != nil {
		db = db.Where("detail = ?", *req.Detail)
	}
	if req.Query != nil {
		db = db.Where("name LIKE ? OR surname LIKE ? OR nickname LIKE ? OR telephone_number LIKE ? OR phone_number LIKE ? OR detail LIKE ?",
			"%"+*req.Query+"%", "%"+*req.Query+"%", "%"+*req.Query+"%", "%"+*req.Query+"%", "%"+*req.Query+"%", "%"+*req.Query+"%")
	}
	return db
}
