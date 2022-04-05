package rest

import (
	"github.com/ashtishad/banking/pkg/lib"
	"strconv"
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
