package traceability

import (
	"fmt"

	"data-spaces-backend/domain/common"
	"data-spaces-backend/extension/logger"
)

// EnumRequestTypeValid
// Summary: This is function which checks whether the RequestType is valid.
// input: value (interface{}) value
// output: (error) error object
func EnumRequestTypeValid(value interface{}) error {
	v, _ := value.(RequestType)
	switch v {
	case RequestTypeCFP:
		return nil
	default:
		logger.Set(nil).Warnf(common.InvalidEnumError(v.ToString()))

		return fmt.Errorf(common.InvalidEnumError(v.ToString()))
	}
}

// EnumCfpTypeValid
// Summary: This is function which checks whether the CfpType is valid.
// input: value (interface{}) value
// output: (error) error object
func EnumCfpTypeValid(value interface{}) error {
	v, _ := value.(CfpType)
	switch v {
	case CfpTypePreProduction, CfpTypeMainProduction, CfpTypePreComponent, CfpTypeMainComponent:
		return nil
	default:
		logger.Set(nil).Warnf(common.InvalidEnumError(v.ToString()))

		return fmt.Errorf(common.InvalidEnumError(v.ToString()))
	}
}

// EnumDqrTypeValidForPutCfp
// Summary: This is function which checks whether the DqrType is valid for PutCfp.
// input: value (interface{}) value
// output: (error) error object
func EnumDqrTypeValidForPutCfp(value interface{}) error {
	v, _ := value.(DqrType)
	switch v {
	case DqrTypePreProcessing, DqrTypeMainProcessing:
		return nil
	default:
		logger.Set(nil).Warnf(common.InvalidEnumError(v.ToString()))

		return fmt.Errorf(common.InvalidEnumError(v.ToString()))
	}
}

// EnumAmountRequiredUnitValidOrNil
// Summary: This is function which checks whether the AmountRequiredUnit is valid or nil.
// input: value (interface{}) value
// output: (error) error object
func EnumAmountRequiredUnitValidOrNil(value interface{}) error {
	if value == nil {
		return nil
	}
	ptr, ok := value.(*string)
	if !ok {
		logger.Set(nil).Warnf("expected a string pointer")

		return fmt.Errorf("expected a string pointer")
	}
	if ptr == nil {
		return nil
	}
	s := *ptr
	_, err := NewAmountRequiredUnit(s)
	if err != nil {
		logger.Set(nil).Warnf(common.InvalidEnumError(s))

		return fmt.Errorf(common.InvalidEnumError(s))
	}
	return nil
}

// EnumGhgDeclaredUnitValid
// Summary: This is function which checks whether the GhgDeclaredUnit is valid.
// input: value (interface{}) value
// output: (error) error object
func EnumGhgDeclaredUnitValid(value interface{}) error {
	s, _ := value.(string)
	_, err := NewGhgDeclaredUnit(s)
	if err != nil {
		logger.Set(nil).Warnf(common.InvalidEnumError(s))

		return fmt.Errorf(common.InvalidEnumError(s))
	}
	return nil
}
