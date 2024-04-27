package queriers

import "gorm.io/gen"

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE detail = @detail AND telephone_number = @telephoneNumber
	FindByDetailAndTelephoneNumber(detail, telephoneNumber string) (gen.T, error)
}
