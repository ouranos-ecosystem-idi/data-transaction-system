package common

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"data-spaces-backend/extension/logger"

	"github.com/google/uuid"
)

const (
	ErrorMessageCannotBeBlank = "cannot be blank"
	ErrorMessageInvalidUUID   = "invalid UUID"
	ErrorMessageNoLessThan    = "must be no less than 0.00001"
	ErrorMessage5thDecimal    = "must be a value up to the 5th decimal place"
)

// UUIDNotNil
// Summary: This is function which checks whether the UUID is not nil
// input: value(interface{}) value
// output: (error) error object
func UUIDNotNil(value interface{}) error {
	u, ok := value.(uuid.UUID)
	if !ok {
		logger.Set(nil).Warnf(ErrorMessageInvalidUUID)

		return fmt.Errorf(ErrorMessageInvalidUUID)
	}

	if u == uuid.Nil {
		logger.Set(nil).Warnf(ErrorMessageCannotBeBlank)

		return fmt.Errorf(ErrorMessageCannotBeBlank)
	}

	return nil
}

// StringUUIDValid
// Summary: This is function which checks whether the string is a valid UUID
// input: value(interface{}) value
// output: (error) error object
func StringUUIDValid(value interface{}) error {
	s, _ := value.(string)
	if s == "" {
		logger.Set(nil).Warnf(ErrorMessageCannotBeBlank)

		return fmt.Errorf(ErrorMessageCannotBeBlank)
	} else {
		if len(s) != 36 {
			logger.Set(nil).Warnf(ErrorMessageInvalidUUID)

			return fmt.Errorf(ErrorMessageInvalidUUID)
		}

		_, err := uuid.Parse(s)
		if err != nil {
			logger.Set(nil).Warnf(ErrorMessageInvalidUUID)

			return fmt.Errorf(ErrorMessageInvalidUUID)
		}
	}

	return nil
}

// StringPtrNilOrUUIDValid
// Summary: This is function which checks whether the string pointer is nil or a valid UUID
// input: value(interface{}) value
// output: (error) error object
func StringPtrNilOrUUIDValid(value interface{}) error {
	sp, _ := value.(*string)
	if sp == nil {
		return nil
	}
	s := *sp

	if len(s) != 36 {
		logger.Set(nil).Warnf(ErrorMessageInvalidUUID)

		return fmt.Errorf(ErrorMessageInvalidUUID)
	}

	_, err := uuid.Parse(s)
	if err != nil {
		logger.Set(nil).Warnf(ErrorMessageInvalidUUID)

		return fmt.Errorf(ErrorMessageInvalidUUID)
	}

	return nil
}

// BoolPtrNotNil
// Summary: This is function which checks whether the bool pointer is not nil
// input: value(interface{}) value
// output: (error) error object
func BoolPtrNotNil(value interface{}) error {
	bp, _ := value.(*bool)
	if bp == nil {
		logger.Set(nil).Warnf(ErrorMessageCannotBeBlank)

		return fmt.Errorf(ErrorMessageCannotBeBlank)
	}

	return nil
}

// FloatPtrPositive
// Summary: This is function which checks whether the float pointer not less than 0.00001
// input: value(interface{}) value
// output: (error) error object
func FloatPtrPositive(value interface{}) error {
	fp, _ := value.(*float64)
	if fp == nil || *fp < 0.00001 {
		logger.Set(nil).Warnf(ErrorMessageNoLessThan)

		return fmt.Errorf(ErrorMessageNoLessThan)
	}

	return nil
}

// FloatPtr5thDecimal
// Summary: This is function which checks whether the float pointer is a value up to the 5th decimal place
// input: value(interface{}) value
// output: (error) error object
func FloatPtr5thDecimal(value interface{}) error {
	fp, _ := value.(*float64)
	if fp == nil {
		return nil
	}

	if *fp == 0 {
		return nil
	}

	re, _ := regexp.Compile(`^[0-9]+\.?[0-9]{0,5}$`)
	stValue := strconv.FormatFloat(*fp, 'f', -1, 64)
	if !re.MatchString(stValue) {
		logger.Set(nil).Warnf(ErrorMessage5thDecimal)

		return fmt.Errorf(ErrorMessage5thDecimal)
	}

	return nil
}

// JoinErrors
// Summary: This is function which joins multiple error messages
// input: errors ([]error) error object
// output: (error) error object
func JoinErrors(errors []error) error {
	errDetails := ""
	errorMap := make(map[string]bool)

	filteredErrors := []error{}
	for _, v := range errors {
		if _, ok := errorMap[v.Error()]; ok {
			continue
		}
		errorMap[v.Error()] = true
		filteredErrors = append(filteredErrors, v)
	}

	for i, v := range filteredErrors {
		errDetails += v.Error()

		if i+1 < len(filteredErrors) {
			errDetails += "; "
		} else {
			if !strings.HasSuffix(errDetails, ".") {
				errDetails += "."
			}
		}
	}

	return fmt.Errorf(errDetails)
}
