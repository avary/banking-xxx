package rest

import (
	"github.com/ashtishad/banking/pkg/lib"
	"strconv"
	"strings"
)

// getCustomerIdFromPath returns the account id from the path
// as database id type is big int, we need to convert it to int
func getCustomerIdFromPath(customerIdParam string) (int64, lib.RestErr) {
	id, err := strconv.ParseInt(customerIdParam, 10, 64)
	if err != nil {
		return 0, lib.NewBadRequestError("user id should be a number")
	}
	return id, nil
}

// validateCustomerStatus takes status and validates it
func validateCustomerStatus(status string) (int8, lib.RestErr) {
	status = strings.ToLower(status)
	switch status {
	case "0":
		return 0, nil
	case "1":
		return 1, nil
	case "active":
		return 1, nil
	case "inactive":
		return 0, nil
	default:
		return 0, lib.NewBadRequestError("status is invalid")
	}
}
