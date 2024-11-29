package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// HTTPError
// Summary: This is structure which defines HTTPError.
type HTTPError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

// Below is the Model for Swagger
type HTTP400Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type HTTP401Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type HTTP403Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type HTTP404Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type HTTP500Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type HTTP503Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type HTTP504Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

var (
	// 400 Error Messages
	Err400InvalidRequest  = "Invalid request parameters"
	Err400InvalidJSON     = "Invalid JSON format"
	Err400RequestTooLarge = "Request payload too large"
	Err400Validation      = "Validation failed"
	// 401 Error Messages
	Err401InvalidCredentials = "Invalid credentials"
	Err401Authentication     = "Authentication required"
	Err401InvalidToken       = "Invalid or expired token"
	// 403 Error Messages
	Err403AccessDenied          = "You do not have the necessary privileges"
	Err403InvalidKey            = "Invalid key"
	Err403IPNotAuthorizedForKey = "IP address not authorized for this API key"
	// 404 Error Messages
	Err404ResourceNotFound = "Resource Not Found"
	Err404ItemNotFound     = "Item or record Not Found"
	Err404EndpointNotFound = "Endpoint Not Found"
	// 500 Error Messages
	Err500Unexpected = "Unexpected error occurred"
	// 503 Error Messages
	Err503OuterService = "Unexpected error occurred in outer service"
)

// HTTPErrorSource
// Summary: This is enum which defines HTTPErrorSource.
type HTTPErrorSource string

// ToString
// Summary: This is the function to convert HTTPErrorSource to string.
// output: (string) converted to string
func (t HTTPErrorSource) ToString() string {
	return string(t)
}

var (
	HTTPErrorSourceDataspace    HTTPErrorSource = "dataspace"
	HTTPErrorSourceAuth         HTTPErrorSource = "auth"
	HTTPErrorSourceTraceability HTTPErrorSource = "traceability"
)

// HTTPErrorGenerate
// Summary: This is the function to generate HTTPError.
// input: httpStatusCode(int) http status code
// input: source(HTTPErrorSource) source of error
// input: errorMsg(string) error message
// input: operatorID(string) ID of the operator
// input: dataTarget(string) target of the data
// input: method(string) method of the request
// input: errorDetails(...string) error details
// output: (int) http status code
// output: (HTTPError) HTTPError object
func HTTPErrorGenerate(
	httpStatusCode int,
	source HTTPErrorSource,
	errorMsg string,
	operatorID string,
	dataTarget string,
	method string,
	errorDetails ...string,
) (int, HTTPError) {
	now := time.Now()
	utcNow := now.UTC()
	iso8601Format := "2006-01-02T15:04:05.000Z"
	isoUtcTime := utcNow.Format(iso8601Format)
	var detailMessage string
	if source == HTTPErrorSourceTraceability {
		detailMessage = errorDetails[0]
	} else {
		detailMessage = "id: " + operatorID + ", timeStamp: " + isoUtcTime + ", dataTarget: " + dataTarget + ", method: " + method

		for _, errorDetail := range errorDetails {
			errorMsg += fmt.Sprintf(", %s", errorDetail)
		}
	}

	switch httpStatusCode {
	case 400:
		errorModel := HTTPError{
			Code:    formatErrorCode("BadRequest", source),
			Message: errorMsg,
			Detail:  detailMessage,
		}
		return 400, errorModel
	case 401:
		errorModel := HTTPError{
			Code:    formatErrorCode("Unauthorized", source),
			Message: errorMsg,
			Detail:  detailMessage,
		}
		return 401, errorModel
	case 403:
		errorModel := HTTPError{
			Code:    formatErrorCode("AccessDenied", source),
			Message: errorMsg,
			Detail:  detailMessage,
		}
		return 403, errorModel
	case 404:
		errorModel := HTTPError{
			Code:    formatErrorCode("NotFound", source),
			Message: errorMsg,
			Detail:  detailMessage,
		}
		return 404, errorModel
	case 500:
		errorModel := HTTPError{
			Code:    formatErrorCode("InternalServerError", source),
			Message: errorMsg,
			Detail:  detailMessage,
		}
		return 500, errorModel
	case 503:
		errorModel := HTTPError{
			Code:    formatErrorCode("ServiceUnavailable", source),
			Message: errorMsg,
			Detail:  detailMessage,
		}
		return 503, errorModel
	case 504:
		errorModel := HTTPError{
			Code:    formatErrorCode("Timeout", source),
			Message: errorMsg,
			Detail:  detailMessage,
		}
		return 504, errorModel
	default:
		errorModel := HTTPError{
			Code:    formatErrorCode("InternalServerError", source),
			Message: errorMsg,
			Detail:  detailMessage,
		}
		return 500, errorModel
	}
}

// formatErrorCode
// Summary: This is the function to format error code.
// input: code(string) error code
// input: source(HTTPErrorSource) source of error
// output: (string) formatted error code
func formatErrorCode(code string, source HTTPErrorSource) string {
	return fmt.Sprintf("[%s] %s", source, code)
}

// CustomError
// Summary: This is structure which defines CustomError.
type CustomError struct {
	Code          CustomErrorCode
	Message       string
	MessageDetail *string
	Source        HTTPErrorSource
}

// NewCustomError
// Summary: This is the function to create new CustomError.
// input: code(CustomErrorCode) error code
// input: message(string) error message
// input: messageDetail(*string) error message detail
// input: source(HTTPErrorSource) source of error
// output: (*CustomError) CustomError object
func NewCustomError(code CustomErrorCode, message string, messageDetail *string, source HTTPErrorSource) *CustomError {
	return &CustomError{
		Code:          code,
		Message:       message,
		MessageDetail: messageDetail,
		Source:        source,
	}
}

// Error
// Summary: This is the function to get error message.
// output: (string) error message
func (e CustomError) Error() string {
	return e.Message
}

// IsWarn
// Summary: This is the function to check if the error is a warning.
// output: (bool) true: warning, false: not warning
func (e CustomError) IsWarn() bool {
	return e.Code >= 400 && e.Code < 500
}

// CustomErrorCode
// Summary: This is enum which defines CustomErrorCode.
type CustomErrorCode int

var (
	CustomErrorCode400 CustomErrorCode = http.StatusBadRequest
	CustomErrorCode401 CustomErrorCode = http.StatusUnauthorized
	CustomErrorCode403 CustomErrorCode = http.StatusForbidden
	CustomErrorCode404 CustomErrorCode = http.StatusNotFound
	CustomErrorCode500 CustomErrorCode = http.StatusInternalServerError
	CustomErrorCode503 CustomErrorCode = http.StatusServiceUnavailable
)

// FormatBindErrMsg
// Summary: This is the function to format bind error message.
// input: err(error) error object
// output: (string) formatted error message
func FormatBindErrMsg(err error) string {
	if strings.Contains(err.Error(), "Syntax error") {
		return extractSyntaxErrorMessage(err.Error())
	}
	return extractTypeErrorMessage(err.Error())
}

// extractSyntaxErrorMessage
// Summary: This is the function to extract syntax error message.
// input: errString(string) error string
// output: (string) extracted error message
func extractSyntaxErrorMessage(errString string) string {
	const prefix = "error="
	start := strings.Index(errString, prefix)
	if start == -1 {
		return ""
	}

	message := errString[start+len(prefix):]
	end := strings.Index(message, ", internal=")
	if end != -1 {
		message = message[:end]
	}

	return message
}

// extractTypeErrorMessage
// Summary: This is the function to extract type error message.
// input: errString(string) error string
// output: (string) extracted error message
func extractTypeErrorMessage(errString string) string {
	messageStart := strings.Index(errString, "message=")
	if messageStart == -1 {
		return ""
	}
	message := errString[messageStart+len("message="):]

	fieldStart := strings.Index(message, ", field=")
	if fieldStart == -1 {
		return ""
	}
	field := message[fieldStart+len(", field="):]

	errorDescription := message[:fieldStart]

	nextComma := strings.Index(field, ",")
	if nextComma != -1 {
		field = field[:nextComma]
	}

	formattedMessage := fmt.Sprintf("%s: %s.", field, errorDescription)
	return formattedMessage
}

// UnexpectedQueryParameter
// Summary: This is the function to format unexpected query parameter error message.
// input: param(string) parameter name
// output: (string) formatted error message
func UnexpectedQueryParameter(param string) string {
	return fmt.Sprintf("%v: Unexpected query parameter", param)
}

// InvalidEnumError
// Summary: This is the function to format invalid enum error message.
// input: value(string) value
// output: (string) formatted error message
func InvalidEnumError(value string) string {
	return fmt.Sprintf("cannot be allowed '%s'", value)
}

// InvalidUUIDError
// Summary: This is the function to format invalid UUID error message.
// input: name(string) name
// output: (string) formatted error message
func InvalidUUIDError(name string) string {
	return fmt.Sprintf("%s: invalid UUID.", name)
}

// InconsistentFieldError
// Summary: This is the function to format inconsistent field error message.
// input: name(string) name
// output: (string) formatted error message
func InconsistentFieldError(name string) string {
	return fmt.Sprintf("ensure all objects have the same %v", name)
}

// CfpElementsError
// Summary: This is the function to get cfp elements error message.
// output: (string) error message
func CfpElementsError() string {
	return "cfp models must be 4 elements"
}

// InsufficientCfpType
// Summary: This is the function to format insufficient cfp type error message.
// input: cfpType(string) cfp type
// output: (string) formatted error message
func InsufficientCfpType(cfpType string) string {
	return fmt.Sprintf("cfpType %v is insufficient", cfpType)
}

// OnlyOneCanBeSpecified
// Summary: This is the function to format only one can be specified error message.
// input: name1(string) name1
// input: name2(string) name2
// output: (string) formatted error message
func OnlyOneCanBeSpecified(name1 string, name2 string) string {
	return fmt.Sprintf("only one of %v and %v can be set.", name1, name2)
}

// InvalidCombination
// Summary: This is the function to format invalid combination error message.
// input: name1(string) name1
// input: name2(string) name2
// output: (string) formatted error message
func InvalidCombination(name1 string, name2 string) string {
	return fmt.Sprintf("ensure the combination of %v and %v is correct", name1, name2)
}

// InvalidGhgEmission
// Summary: This is the function to format invalid ghg emission error message.
// input: cfpType1(string) cfp type1
// input: cfpType2(string) cfp type2
// output: (string) formatted error message
func InvalidGhgEmission(cfpType1 string, cfpType2 string) string {
	return fmt.Sprintf("set ghgEmission to a value greater than 0 for cfpType of %v or %v", cfpType1, cfpType2)
}

// CfpCertificateListInconsistent
// Summary: This is the function to get cfp certificate list inconsistent error message.
// output: (string) error message
func CfpCertificateListInconsistentError() string {
	return "cfpCertificateList must be same"
}

// CfpIDsInconsistentError
// Summary: This is the function to get cfp IDs inconsistent error message.
// output: (string) error message
func CfpIDsInconsistentError() string {
	return "cfpIds must be same"
}

// CfpTypeNotFoundError
// Summary: This is the function to format cfp type not found error message.
// input: cfpType(string) cfp type
// output: (string) formatted error message
func CfpTypeNotFoundError(cfpType string) string {
	return fmt.Sprintf("specified CfpType not found: %v", cfpType)
}

// DqrTypeNotFoundError
// Summary: This is the function to format dqr type not found error message.
// input: dqrType(string) dqr type
// output: (string) formatted error message
func DqrTypeNotFoundError(dqrType string) string {
	return fmt.Sprintf("specified DqrType not found: %v", dqrType)
}

// DqrValueInconsistentError
// Summary: This is the function to get dqr value inconsistent error message.
// output: (string) error message
func DqrValueInconsistentError() string {
	return "different dqrValues are set for the same dqrType"
}

// GhgDeclaredUnitsInconsistentError
// Summary: This is the function to get ghg declared units inconsistent error message.
// output: (string) error message
func GhgDeclaredUnitsInconsistentError() string {
	return "ghgDeclaredUnits must be same"
}

// LimitLessThanError
// Summary: This is the function to format limit less than error message.
// input: min(int) min
// input: limit(int) limit
// output: (string) formatted error message
func LimitLessThanError(min int, limit int) string {
	return fmt.Sprintf("limit less than %v error. get value: %v", min, limit)
}

// LimitUpperError
// Summary: This is the function to format limit upper error message.
// input: limit(int) limit
// output: (string) formatted error message
func LimitUpperError(limit int) string {
	return fmt.Sprintf("limit upper limit error. get value: %v.", limit)
}

// TraceIDNotFoundError
// Summary: This is the function to format trace ID not found error message.
// input: traceID(string) ID of the trace
// output: (string) formatted error message
func TraceIDNotFoundError(traceID string) string {
	return fmt.Sprintf("traceId %v not found", traceID)
}

// TraceIDAlreadyHasCfpsError
// Summary: This is the function to format trace ID already has cfps error message.
// input: traceID(string) ID of the trace
// output: (string) formatted error message
func TraceIDAlreadyHasCfpsError(traceID string) string {
	return fmt.Sprintf("traceId %v already has cfps", traceID)
}

// TraceIDsInconsistentError
// Summary: This is the function to get trace IDs inconsistent error message.
// output: (string) error message
func TraceIDsInconsistentError() string {
	return "traceIds must be same"
}

// TraceIDsUpperLimitError
// Summary: This is the function to format trace IDs upper limit error message.
// input: length(int) length
// output: (string) formatted error message
func TraceIDsUpperLimitError(length int) string {
	return fmt.Sprintf("The upper limit for traceIds is 50. get length: %v", length)
}

// UnexpectedEnumError
// Summary: This is the function to format unexpected enum error message.
// input: name(string) name
// input: value(string) value
// output: (string) formatted error message
func UnexpectedEnumError(name string, value string) string {
	return fmt.Sprintf("unexpected %v. get value: %v", name, value)
}

// NotFoundInResponseError
// Summary: This is the function to format not found response error message.
// input: partsItem(string) partsItem
// input: supportPartsItem(*string) supportPartsItem
// output: (string) formatted error message
func NotFoundInResponseError(partsItem string, supportPartsItem *string) string {
	return fmt.Sprintf("set of partsItem %v, supportPartsItem %v not found in response", partsItem, supportPartsItem)
}

// InconsistentError
// Summary: This is the function to get values inconsistent error message.
// output: (string) error message
func InconsistentError(value1 string, value2 string) string {
	return fmt.Sprintf("%v and %v must be equal", value1, value2)
}

// NotHaveValuesError
// Summary: This is the function to get not have values error message.
// output: (string) error message
func NotHaveValuesError(value1 string, value2 string, value3 string) string {
	return fmt.Sprintf("%v, %v, and %v must all have values or all be null", value1, value2, value3)
}

// UnexpectedResponse
// Summary: This is the function to format unexpected response error message.
// input: param(string) parameter name
// output: (string) formatted error message
func UnexpectedResponse(system string) string {
	return fmt.Sprintf("Unexpected %v api response", system)
}

// NotFoundError
// Summary: This is the function to format not found error message.
// input: value(string) target value
// output: (string) formatted error message
func NotFoundError(value string) string {
	return fmt.Sprintf("%v not found", value)
}

// DeleteTableError
// Summary: This is the function to format delete table error message.
// input: name(string) table name
// input: err(error) error object
// output: (string) formatted error message
func DeleteTableError(name string, err error) string {
	return fmt.Sprintf("failed to physically delete record from table %v : %v", name, err)
}

// TraceabilityAPIError
// Summary: This is structure which defines TraceabilityAPIError.
type TraceabilityAPIError struct {
	Errors  []TraceabilityAPIErrorDetail `json:"errors"`
	Message *string                      `json:"message"`
}

// TraceabilityAPIErrorDelete
// Summary: This is structure which defines TraceabilityAPIErrorDelete.
type TraceabilityAPIErrorDelete struct {
	Errors  []TraceabilityAPIErrorDetailDelete `json:"errors"`
	Message *string                            `json:"message"`
}

// TraceabilityAPIErrorDetail
// Summary: This is structure which defines TraceabilityAPIErrorDetail.
type TraceabilityAPIErrorDetail struct {
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
}

// TraceabilityAPIErrorDetailDelete
// Summary: This is structure which defines TraceabilityAPIErrorDetailDelete.
type TraceabilityAPIErrorDetailDelete struct {
	ErrorCode        string    `json:"errorCode"`
	ErrorDescription string    `json:"errorDescription"`
	RelevantData     *[]string `json:"relevantData"`
}

// ToTracebilityAPIError
// Summary: This is the function to convert json string to TraceabilityAPIError.
// input: jsonStr(string) json string
// output: (*TraceabilityAPIError) TraceabilityAPIError object
func ToTracebilityAPIError(jsonStr string) *TraceabilityAPIError {
	var traceabilityAPIError TraceabilityAPIError
	if err := json.Unmarshal([]byte(jsonStr), &traceabilityAPIError); err != nil {
		return nil
	}

	return &traceabilityAPIError
}

// TraceabilityAPIErrorDelete
// Summary: This is the function to convert json string to TraceabilityAPIErrorDelete.
// input: jsonStr(string) json string
// output: (*TraceabilityAPIErrorDelete) TraceabilityAPIErrorDelete object
func ToTracebilityAPIErrorDelete(jsonStr string) *TraceabilityAPIErrorDelete {
	var traceabilityAPIErrorDelete TraceabilityAPIErrorDelete
	if err := json.Unmarshal([]byte(jsonStr), &traceabilityAPIErrorDelete); err != nil {
		return nil
	}

	return &traceabilityAPIErrorDelete
}

// ErrorCodesToString
// Summary: This is the function to convert error codes to string.
// output: (string) converted to string
func (e *TraceabilityAPIError) ErrorCodesToString() string {
	var errorCodes []string
	for _, errorDetail := range e.Errors {
		errorCodes = append(errorCodes, errorDetail.ErrorCode)
	}

	return strings.Join(errorCodes, ", ")
}

// ErrorCodesToString
// Summary: This is the function to convert error codes to string.
// output: (string) converted to string
func (e *TraceabilityAPIErrorDelete) ErrorCodesToString() string {
	var errorCodes []string
	for _, errorDetail := range e.Errors {
		errorCodes = append(errorCodes, errorDetail.ErrorCode)
	}

	return strings.Join(errorCodes, ", ")
}

// ErrorDescriptionsToString
// Summary: This is the function to convert error descriptions to string.
// output: (string) converted to string
func (e *TraceabilityAPIError) ErrorDescriptionsToString() string {
	var errorDescriptions []string
	for _, errorDetail := range e.Errors {
		errorDescriptions = append(errorDescriptions, errorDetail.ErrorDescription)
	}

	return strings.Join(errorDescriptions, ", ")
}

// ErrorDescriptionsToString
// Summary: This is the function to convert error descriptions to string.
// output: (string) converted to string
func (e *TraceabilityAPIErrorDelete) ErrorDescriptionsToString() string {
	var errorDescriptions []string
	for _, errorDetail := range e.Errors {
		errorDescriptions = append(errorDescriptions, errorDetail.ErrorDescription)
	}

	return strings.Join(errorDescriptions, ", ")
}

// ToCustomError
// Summary: This is the function to convert TraceabilityAPIError to CustomError.
// input: httpStatus(int) http status code
// output: (*CustomError) CustomError object
func (e *TraceabilityAPIError) ToCustomError(httpStatus int) *CustomError {
	message := e.Message
	if message != nil {
		return NewCustomError(CustomErrorCode(httpStatus), "", message, HTTPErrorSourceTraceability)
	} else {
		errorCodes := e.ErrorCodesToString()
		errorDescriptions := e.ErrorDescriptionsToString()
		return NewCustomError(CustomErrorCode(httpStatus), errorDescriptions, &errorCodes, HTTPErrorSourceTraceability)
	}

}

// ToCustomError
// Summary: This is the function to convert TraceabilityAPIErrorDelete to CustomError.
// input: httpStatus(int) http status code
// output: (*CustomError) CustomError object
func (e *TraceabilityAPIErrorDelete) ToCustomError(httpStatus int) *CustomError {
	message := e.Message
	if message != nil {
		return NewCustomError(CustomErrorCode(httpStatus), "", message, HTTPErrorSourceTraceability)
	} else {
		errorCodes := e.ErrorCodesToString()
		errorDescriptions := e.ErrorDescriptionsToString()
		return NewCustomError(CustomErrorCode(httpStatus), errorDescriptions, &errorCodes, HTTPErrorSourceTraceability)
	}
}

// ToAuthAPIError
// Summary: This is the function to convert json string to HTTPError.
// input: jsonStr(string) json string
// output: (*HTTPError) HTTPError object
func ToAuthAPIError(jsonStr string) *HTTPError {
	var httpError HTTPError
	if err := json.Unmarshal([]byte(jsonStr), &httpError); err != nil {
		return nil
	}

	return &httpError
}

// ToCustomError
// Summary: This is the function to convert HTTPError to CustomError.
// input: httpStatus(int) http status code
// output: (*CustomError) CustomError object
func (e *HTTPError) ToCustomError(httpStatus int) *CustomError {
	return NewCustomError(CustomErrorCode(httpStatus), e.Message, &e.Detail, HTTPErrorSourceAuth)
}
