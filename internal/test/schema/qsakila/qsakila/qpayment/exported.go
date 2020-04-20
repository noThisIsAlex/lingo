// Code generated by Lingo for table sakila.payment - DO NOT EDIT

package qpayment

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QPayment {
	return instance
}

func PaymentId() path.Int16Path {
	return instance.paymentId
}

func CustomerId() path.Int16Path {
	return instance.customerId
}

func StaffId() path.BoolPath {
	return instance.staffId
}

func RentalId() path.IntPath {
	return instance.rentalId
}

func Amount() path.BinaryPath {
	return instance.amount
}

func PaymentDate() path.TimePath {
	return instance.paymentDate
}

func LastUpdate() path.TimePath {
	return instance.lastUpdate
}
