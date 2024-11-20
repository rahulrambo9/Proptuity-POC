package stderrors

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Title   string      `json:"title"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}

// ErrorCode type is used to define error codes
type ErrorCode int

const (
	BadRequest ErrorCode = iota + 2000
	InvalidUser
	NotFound
	Unauthorized
	Forbidden
	InternalServerError
	InvalidContent
	ContentGone
	ContentNotAvailable
	Conflict
	ConfigurationError
	FailedDependency
	AwardError
	UnprocessableEntity
	InvalidToken
	DecryptionError
	EncryptionError
	MissingReqHeaders
)

var StdErrors = map[ErrorCode]string{
	BadRequest:          "bad request",
	InvalidUser:         "invalid user",
	NotFound:            "not found",
	Unauthorized:        "unauthorized",
	Forbidden:           "forbidden",
	InternalServerError: "internal server error",
	InvalidContent:      "invalid content id",
	ContentGone:         "content already completed",
	ContentNotAvailable: "content not available",
	Conflict:            "conflict",
	ConfigurationError:  "configuration error",
	FailedDependency:    "failed dependency",
	AwardError:          "award error",
	UnprocessableEntity: "unprocessable entity",
	InvalidToken:        "invalid token",
	DecryptionError:     "decryption error",
	EncryptionError:     "encryption error",
	MissingReqHeaders:   "missing request headers",
}

type ErrorResponse struct {
	HttpStatusCode int
	Code           int         `json:"code"`
	Title          string      `json:"title"`
	Message        string      `json:"message"`
	Detail         interface{} `json:"detail"`
}

func (e *ErrorResponse) Error() string {
	return e.Title + ": " + e.Message
}

func New(c context.Context, httpStatusCode int, errorCode int, errorTitle string, errorMessage string) error {
	// notify new relic
	// txn := newrelic.FromContext(c)
	// if txn != nil {
	// 	txn.NoticeError(newrelic.Error{
	// 		Message: errorMessage,
	// 		Class:   errorTitle,
	// 		Attributes: map[string]interface{}{
	// 			"error_detail": errorMessage,
	// 		},
	// 	})
	// }

	//log error
	// if logger != nil {
	// 	logger.Error(errorTitle + ": " + errorMessage)
	// }

	// return error response
	return &ErrorResponse{
		HttpStatusCode: httpStatusCode,
		Code:           errorCode,
		Title:          errorTitle,
		Message:        errorMessage,
	}
}

// Anything that is wrong with client request: missing parameters, values out of range, invalid data types.
func ErrBadRequest(c context.Context, message string) error {
	return New(c, http.StatusBadRequest, int(BadRequest), StdErrors[BadRequest], message)
}

// The user in the JWT was no identified as a valid guest or registered user
func ErrInvalidUser(c context.Context, message string) error {
	return New(c, http.StatusUnauthorized, int(InvalidUser), StdErrors[InvalidUser], message)
}

// Whatever the client was trying to access and it is not there: a game, an inventory item, a prize.
func ErrNotFound(c context.Context, message string) error {
	return New(c, http.StatusNotFound, int(NotFound), StdErrors[NotFound], message)
}

// When the user cannot perform a certain action or does not have a valid JWT
func ErrUnauthorized(c context.Context, message string) error {
	return New(c, http.StatusUnauthorized, int(Unauthorized), StdErrors[Unauthorized], message)
}

// When the user is not allowed to perform a certain action
func ErrForbidden(c context.Context, message string) error {
	return New(c, http.StatusForbidden, int(Forbidden), StdErrors[Forbidden], message)
}

// Anything that goes wrong server side
func ErrInternalServerError(c context.Context, message string) error {
	return New(c, http.StatusInternalServerError, int(InternalServerError), StdErrors[InternalServerError], message)
}

// Trying to access unpublished, not yet published content or simply content that is not valid for any business reason
func ErrInvalidContent(c context.Context, message string) error {
	return New(c, http.StatusBadRequest, int(InvalidContent), StdErrors[InvalidContent], message)
}

// Trying to access content that has already been completed
func ErrContentGone(c context.Context, message string) error {
	return New(c, http.StatusGone, int(ContentGone), StdErrors[ContentGone], message)
}

// When the user is not eligible to access a given piece of content, due to VIP level, influencer level, etc
func ErrContentNotAvailable(c context.Context, message string) error {
	return New(c, http.StatusForbidden, int(ContentNotAvailable), StdErrors[ContentNotAvailable], message)
}

// User was trying to modify the state of a resource in a way that is not valid
func ErrConflict(c context.Context, message string) error {
	return New(c, http.StatusConflict, int(Conflict), StdErrors[Conflict], message)
}

// The application parameters are not valid/not present
func ErrConfigurationError(c context.Context, message string) error {
	return New(c, http.StatusInternalServerError, int(ConfigurationError), StdErrors[ConfigurationError], message)
}

// Something went wrong when trying to interact with any of the dependecie: offers, gps, userapi, rf, etc.
func ErrFailedDependency(c context.Context, message string) error {
	return New(c, http.StatusFailedDependency, int(FailedDependency), StdErrors[FailedDependency], message)
}

// When any reward was not possible: tokens, entries, inventory due to any reason
func ErrAwardError(c context.Context, message string) error {
	return New(c, http.StatusInternalServerError, int(AwardError), StdErrors[AwardError], message)
}

// When there is a problem with the resource state and the action cannot be performed
func ErrUnprocessableEntity(c context.Context, message string) error {
	return New(c, http.StatusUnprocessableEntity, int(UnprocessableEntity), StdErrors[UnprocessableEntity], message)
}

// The provided token or game code was not valid
func ErrInvalidToken(c context.Context, message string) error {
	return New(c, http.StatusUnprocessableEntity, int(InvalidToken), StdErrors[InvalidToken], message)
}

// ErrDecryption is returned when there is an error during decryption.
func ErrDecryption(c context.Context, message string) error {
	return New(c, http.StatusInternalServerError, int(DecryptionError), StdErrors[DecryptionError], message)
}

// ErrEncryption is returned when there is an error during encryption.
func ErrEncryption(c context.Context, message string) error {
	return New(c, http.StatusInternalServerError, int(EncryptionError), StdErrors[EncryptionError], message)
}

// ErrMissingReqHeaders is returned when required headers are missing in the request.
func ErrMissingReqHeaders(c context.Context, message string) error {
	return New(c, http.StatusBadRequest, int(MissingReqHeaders), StdErrors[MissingReqHeaders], message)
}

// Handler returns a standard error to the client
// adhering the conventions of the Appservices APIs
//
// Usage:
//
//	app := fiber.New(fiber.Config{
//	    ReadBufferSize: 16384,
//	    ErrorHandler:   stderror.Handler(),
//	})
func Handler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {

		// Check if the error is stderror.ErrorResponse type
		if e, ok := err.(*ErrorResponse); ok {
			detail := e.Detail
			if detail == nil {
				detail = e.Message
			}
			// Return HTTP status code and send JSON error message
			return c.Status(e.HttpStatusCode).JSON(ApiResponse{
				Code:    int(e.Code),
				Title:   e.Title,
				Message: e.Message,
				Detail:  detail,
			})
		}

		// error was unexpected, return a generic error response
		return c.Status(fiber.StatusInternalServerError).JSON(ApiResponse{
			Code:    5000,
			Title:   "Internal Server Error",
			Message: err.Error(),
			Detail:  err.Error(),
		})
	}
}
