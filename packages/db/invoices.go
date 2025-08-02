package db

import (
	"tina/packages/types"

	"gorm.io/gorm"
)

func InvoicesModel() *gorm.DB {
	return DB.Model(types.Invoice{})
}

func FindInvoice(query interface{}, args ...interface{}) (*types.Invoice, error) {
	var invoice types.Invoice

	result := DB.Model(types.Invoice{}).Where(query, args...).Find(&invoice)

	if result.RowsAffected == 0 {
		return nil, result.Error
	} else {
		return &invoice, nil
	}
}

func InsertInvoice(invoice *types.Invoice) error {
	result := DB.Model(types.Invoice{}).Create(invoice)
	return result.Error
}
func UpdateInvoice(invoice *types.Invoice) error {
	result := DB.Model(types.Invoice{}).Where("id = ?", invoice.Id).Updates(invoice)
	return result.Error
}

func FindInvoices(query interface{}, offset int, limit int, args ...interface{}) ([]types.Invoice, error) {
	var invoices []types.Invoice
	result := DB.Model(types.Invoice{}).Select("*").Where(query, args...).Offset(offset).Limit(limit).Find(invoices)
	return invoices, result.Error
}
