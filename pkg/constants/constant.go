package constants

import "errors"

// Define new error here
var (
	ErrPasswordIsNotValid  = errors.New("password must have from 8 to 30 characters")
	ErrEmailIsNotValid     = errors.New("email is not valid")
	ErrEmailHasExisted     = errors.New("email has existed")
	ErrLoginFailed         = errors.New("email and password are not valid")
	ErrFirstNameIsEmpty    = errors.New("first name can not be blank")
	ErrFirstNameTooLong    = errors.New("first name too long, max character is 30")
	ErrLastNameIsEmpty     = errors.New("last name can not be blank")
	ErrLastNameTooLong     = errors.New("last name too long, max character is 30")
	ErrCannotRegister      = errors.New("cannot register")
	ErrConvertJsonToStruct = errors.New("can't convert json string to struct")
	ErrWrongPass           = errors.New("password wrong")
	ErrUserNotExist        = errors.New("user not found")
)

// MessageError Define message error here
var (
	// ============== [START] DEFINE COMMON MESSAGE ERROR -------------------
	MessageError           = "error"
	MissingToken           = "missing token"
	MissingTokenHeader     = "missing authorization header"
	InvalidToken           = "invalid token"
	InvalidTokenClaim      = "invalid token claims"
	InvalidUserIDClaim     = "invalid userId claim"
	InvalidTokenHeader     = "invalid authorization header"
	ErrorGenerateRandomStr = "generate salt fail"
	InvalidPassword        = "password wrong"
	ParamPassedWrong       = "Params passed wrong"
	BindingError           = "Binding param error"
	NotFound               = "Not found"
	// ============== [END] DEFINE COMMON MESSAGE ERROR -------------------

	// GetDetailDone ============== [START] DEFINE COMMON MESSAGE Success -------------------
	GetDetailDone = "Get detail %s done"

	// ============== [END] DEFINE COMMON MESSAGE Success -------------------

	// CategoryFail ============== [START] DEFINE CATEGORY MESSAGE ERROR ------------------
	/**
	FAIL
	*/
	CategoryFail         = "Category fail"
	WrongGetListCategory = "Get list category fail!!"

	/**
	SUCCESS
	*/
	GetListCategoryDone = "Get list category done"
	// ============== [END] DEFINE CATEGORY MESSAGE ERROR ------------------
)
