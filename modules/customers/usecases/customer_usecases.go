package usecases

import (
	"sort"

	"github.com/jinzhu/copier"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"github.com/sweetcandy273/go-teerawut/modules/entities"
	"github.com/sweetcandy273/go-teerawut/pkg/handlers/context"
	"gorm.io/gorm"
)

type customersUse struct {
	CustomersRepo entities.CustomersRepository
}

// Constructor
func NewCustomersUsecase(customersRepo entities.CustomersRepository) entities.CustomersUsecase {
	return &customersUse{
		CustomersRepo: customersRepo,
	}
}

// Create create
func (u *customersUse) Create(c *context.Context, req *entities.CreateCustomerRequest) error {
	err := req.Validate()
	if err != nil {
		logrus.Errorf("Validate error: %v", err)
		return err
	}
	actorID := c.GetUserID()
	customer := &entities.Customer{
		Actor: entities.Actor{
			CreatedByUserID: actorID,
			UpdatedByUserID: &actorID,
		},
	}
	_ = copier.CopyWithOption(&customer, req, copier.Option{IgnoreEmpty: true})
	for _, addr := range customer.Addresses {
		addr.Actor = entities.Actor{
			CreatedByUserID: actorID,
			UpdatedByUserID: &actorID,
		}
	}
	err = u.CustomersRepo.Create(customer)
	if err != nil {
		logrus.Errorf("Create customer error: %v", err)
		return err
	}

	return nil
}

// Update update
func (u *customersUse) Update(c *context.Context, req *entities.UpdateCustomerRequest) error {
	err := req.Validate()
	if err != nil {
		logrus.Errorf("Validate error: %v", err)
		return err
	}
	customer, err := u.CustomersRepo.GetByID(req.ID)
	if err != nil {
		logrus.Errorf("Get customer by id error: %v", err)
		return err
	}

	if customer == nil {
		logrus.Errorf("Customer not found")
		return gorm.ErrRecordNotFound
	}
	actorID := c.GetUserID()
	updateCustomer(actorID, req, customer)
	err = u.CustomersRepo.Update(customer)
	if err != nil {
		logrus.Errorf("Update customer error: %v", err)
		return err
	}

	// bulk update addresses
	if len(req.Addresses) > 0 {
		oldMap := make(map[uint]*entities.CustomerAddress)
		for _, addr := range customer.Addresses {
			oldMap[addr.ID] = addr
		}
		var newAddresses []*entities.CustomerAddress
		var updateAddresses []*entities.CustomerAddress
		newIDs := map[uint]bool{}
		for _, addr := range req.Addresses {
			if addr.ID == 0 {
				newAddresses = append(newAddresses,
					&entities.CustomerAddress{
						CustomerID: customer.ID,
						Address:    addr.Address,
						Village:    addr.Village,
						Detail:     addr.Detail,
						Actor: entities.Actor{
							CreatedByUserID: actorID,
							UpdatedByUserID: &actorID,
						},
					})
			} else {
				if updateAddr, exists := oldMap[addr.ID]; exists {
					updateAddresses = append(updateAddresses,
						&entities.CustomerAddress{
							Model:      updateAddr.Model,
							CustomerID: customer.ID,
							Address:    addr.Address,
							Village:    addr.Village,
							Detail:     addr.Detail,
							Actor: entities.Actor{
								UpdatedByUserID: &actorID,
							},
						})
				}
			}
			newIDs[addr.ID] = true
		}
		if len(newAddresses) > 0 {
			err = u.CustomersRepo.CreateAddress(newAddresses)
			if err != nil {
				logrus.Errorf("Create new customer addresses error: %v", err)
				return err
			}
		}
		if len(updateAddresses) > 0 {
			err = u.CustomersRepo.UpdateAddress(updateAddresses)
			if err != nil {
				logrus.Errorf("Update customer addresses error: %v", err)
				return err
			}
		}

		var deleteIDs []uint
		for _, addr := range customer.Addresses {
			if _, exists := newIDs[addr.ID]; !exists {
				deleteIDs = append(deleteIDs, addr.ID)
			}
		}
		if len(deleteIDs) > 0 {
			err = u.CustomersRepo.DeleteAddress(deleteIDs)
			if err != nil {
				logrus.Errorf("Delete customer addresses error: %v", err)
				return err
			}
		}
	}
	if len(req.Addresses) == 0 && len(customer.Addresses) > 0 {
		addressIDs := lo.Map(customer.Addresses, func(addr *entities.CustomerAddress, _ int) uint {
			return addr.ID
		})
		err = u.CustomersRepo.DeleteAddress(addressIDs)
		if err != nil {
			logrus.Errorf("Delete customer addresses error: %v", err)
			return err
		}
	}

	// bulk update air conditions
	if len(req.AirConditions) > 0 {
		oldMap := make(map[uint]*entities.CustomerAirCondition)
		for _, airCond := range customer.AirConditions {
			oldMap[airCond.ID] = airCond
		}
		var newAirConditions []*entities.CustomerAirCondition
		var updateAirConditions []*entities.CustomerAirCondition
		newIDs := map[uint]bool{}
		for _, airCond := range req.AirConditions {
			if airCond.ID == 0 {
				var airBrandID *uint
				if airCond.AirBrandID == 0 {
					airBrandID = nil
				} else {
					airBrandID = &airCond.AirBrandID
				}
				var airTypeID *uint
				if airCond.AirTypeID == 0 {
					airTypeID = nil
				} else {
					airTypeID = &airCond.AirTypeID
				}
				var btuID *uint
				if airCond.BtuID == 0 {
					btuID = nil
				} else {
					btuID = &airCond.BtuID
				}
				newAirConditions = append(newAirConditions,
					&entities.CustomerAirCondition{
						CustomerID: customer.ID,
						AirBrandID: airBrandID,
						AirTypeID:  airTypeID,
						BtuID:      btuID,
						RoomName:   airCond.RoomName,
						FromUs:     airCond.FromUs,
					})
			} else {
				if updateAirCond, exists := oldMap[airCond.ID]; exists {
					var airBrandID *uint
					if airCond.AirBrandID == 0 {
						airBrandID = nil
					} else {
						airBrandID = &airCond.AirBrandID
					}
					var airTypeID *uint
					if airCond.AirTypeID == 0 {
						airTypeID = nil
					} else {
						airTypeID = &airCond.AirTypeID
					}
					var btuID *uint
					if airCond.BtuID == 0 {
						btuID = nil
					} else {
						btuID = &airCond.BtuID
					}
					updateAirConditions = append(updateAirConditions,
						&entities.CustomerAirCondition{
							Model:      updateAirCond.Model,
							CustomerID: customer.ID,
							AirBrandID: airBrandID,
							AirTypeID:  airTypeID,
							BtuID:      btuID,
							RoomName:   airCond.RoomName,
							FromUs:     airCond.FromUs,
						})
				}
			}
			newIDs[airCond.ID] = true
		}
		if len(newAirConditions) > 0 {
			err = u.CustomersRepo.CreateAirConditions(newAirConditions)
			if err != nil {
				logrus.Errorf("Create new customer air conditions error: %v", err)
				return err
			}
		}
		if len(updateAirConditions) > 0 {
			err = u.CustomersRepo.UpdateAirConditions(updateAirConditions)
			if err != nil {
				logrus.Errorf("Update customer air conditions error: %v", err)
				return err
			}
		}
		var deleteIDs []uint
		for _, airCond := range customer.AirConditions {
			if _, exists := newIDs[airCond.ID]; !exists {
				deleteIDs = append(deleteIDs, airCond.ID)
			}
		}
		if len(deleteIDs) > 0 {
			err = u.CustomersRepo.DeleteAirConditions(deleteIDs)
			if err != nil {
				logrus.Errorf("Delete customer air conditions error: %v", err)
				return err
			}
		}
	}

	return nil
}

func updateCustomer(actorID uint, req *entities.UpdateCustomerRequest, customer *entities.Customer) {
	if req.Name != "" {
		customer.Name = req.Name
	}
	if req.PhoneNumber != "" {
		customer.PhoneNumber = req.PhoneNumber
	}
	if req.Detail != "" {
		customer.Detail = req.Detail
	}

	customer.Actor.UpdatedByUserID = lo.ToPtr(actorID)
}

// GetAll get all
func (u *customersUse) GetAll(c *context.Context, req *entities.GetAllCustomerRequest) ([]*entities.Customer, error) {
	customers, err := u.CustomersRepo.GetAll(req)
	if err != nil {
		logrus.Errorf("Get all customer error: %v", err)
		return nil, err
	}
	sort.Slice(customers, func(i, j int) bool {
		if customers[i].Name == customers[j].Name {
			return customers[i].ID < customers[j].ID
		}
		return customers[i].Name < customers[j].Name
	})
	return customers, nil
}

// Delete delete
func (u *customersUse) Delete(c *context.Context, req *entities.GetOne) error {
	customer, err := u.CustomersRepo.GetByID(req.ID)
	if err != nil {
		logrus.Errorf("Get customer by id error: %v", err)
		return err
	}

	if customer == nil {
		logrus.Errorf("Customer not found")
		return gorm.ErrRecordNotFound
	}

	err = u.CustomersRepo.Delete(customer.ID)
	if err != nil {
		logrus.Errorf("Delete customer error: %v", err)
		return err
	}
	return nil
}

// GetByID get by id
func (u *customersUse) GetByID(c *context.Context, req *entities.GetOne) (*entities.Customer, error) {
	customer, err := u.CustomersRepo.GetByID(req.ID)
	if err != nil {
		logrus.Errorf("Get customer by id %d error: %v", req, err)
		return nil, err
	}
	return customer, nil
}

// GetByDetailAndTelephoneNumber get by detail and telephone number
func (u *customersUse) GetByDetailAndTelephoneNumber(c *context.Context, req *entities.GetByDetailAndTelephoneNumberRequest) (any, error) {
	customer, err := u.CustomersRepo.FindByDetailAndTelephoneNumber(req.Detail, req.TelephoneNumber)
	if err != nil {
		logrus.Errorf("Get customer by detail %s and telephone number %s error: %v", req.Detail, req.TelephoneNumber, err)
		return nil, err
	}
	return customer, nil
}
