package repositories

import (
	"github.com/sirupsen/logrus"
	"github.com/sweetcandy273/go-teerawut/modules/entities"
	"github.com/sweetcandy273/go-teerawut/query"

	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	query.SetDefault(r.DB)
	q := query.Customer
	customer, err := q.Where(q.ID.Eq(id)).Preload(q.Addresses).First()
	if err != nil {
		logrus.Errorf("Get customer by id %d error: %v", id, err)
		return nil, err
	}

	return customer, nil
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
	err := queryCustomer(r.DB, req).Preload("Addresses").Find(&customers).Error
	if err != nil {
		logrus.Errorf("Query :: Get all customers error: %v", err)
		return nil, err
	}
	return customers, nil
}

func queryCustomer(db *gorm.DB, req *entities.GetAllCustomerRequest) *gorm.DB {
	if req.ID != nil {
		db = db.Where("id = ?", *req.ID)
	}
	if req.Name != nil {
		db = db.Where("name = ?", *req.Name)
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
		db = db.Where("fname LIKE ? OR lname LIKE ? OR nickname LIKE ? OR phone_number LIKE ? OR detail LIKE ?",
			"%"+*req.Query+"%", "%"+*req.Query+"%", "%"+*req.Query+"%", "%"+*req.Query+"%", "%"+*req.Query+"%", "%"+*req.Query+"%")
	}
	return db
}

// Delete delete
func (r *customersRepo) Delete(id uint) error {
	err := r.DB.Where("id = ?", id).Delete(&entities.Customer{}).Error
	if err != nil {
		logrus.Errorf("Delete customer id %v error: %v", id, err)
		return err
	}
	return nil
}

// FindByDetailAndTelephoneNumber find by detail and telephone number
func (r *customersRepo) FindByDetailAndTelephoneNumber(detail, telephoneNumber string) (gen.T, error) {
	query.SetDefault(r.DB)
	q := query.Customer
	customer, err := q.FindByDetailAndTelephoneNumber(detail, telephoneNumber)
	if err != nil {
		logrus.Errorf("Get customer by detail %s and telephone number %s error: %v", detail, telephoneNumber, err)
		return nil, err
	}

	return customer, nil
}

// CreateAddress create address
func (r *customersRepo) CreateAddress(addresses []*entities.CustomerAddress) error {
	query.SetDefault(r.DB)
	q := query.CustomerAddress
	err := q.CreateInBatches(addresses, 100)
	if err != nil {
		logrus.Errorf("Create customer address error: %v", err)
		return err
	}
	return nil
}

// UpdateAddress update address
func (r *customersRepo) UpdateAddress(addresses []*entities.CustomerAddress) error {
	query.SetDefault(r.DB)
	q := query.CustomerAddress
	err := q.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"address", "village", "detail"}),
	}).CreateInBatches(addresses, 100)
	if err != nil {
		logrus.Errorf("Update customer address error: %v", err)
		return err
	}
	return nil
}

// DeleteAddress delete address
func (r *customersRepo) DeleteAddress(ids []uint) error {
	query.SetDefault(r.DB)
	q := query.CustomerAddress
	_, err := q.Where(q.ID.In(ids...)).Delete()
	if err != nil {
		logrus.Errorf("Delete customer address error: %v", err)
		return err
	}
	return nil
}
