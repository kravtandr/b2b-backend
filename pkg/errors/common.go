package errors

import "errors"

var (
	UserDoesNotExist         = errors.New("user does not exist")
	ChatDoesNotExist         = errors.New("chat does not exist")
	ProductDoesNotExist      = errors.New("product does not exist")
	CategoryDoesNotExist     = errors.New("category does not exist")
	OrderDoesNotExist        = errors.New("order does not exist")
	SessionDoesNotExist      = errors.New("session does not exist")
	WrongUserPassword        = errors.New("wrong user password")
	CompanyDoesNotExist      = errors.New("company does not exist")
	CompanyUsersLinkNotExist = errors.New("CompanyUsersLinkdoes not exist")
	NoBytesToWrite           = errors.New("No bytes to write file")
	CantWriteTmpFile         = errors.New("Error while write tmp file")
	CantCreateTmpFile        = errors.New("Error while create tmp file")
	CantDecodeImgFromBase64  = errors.New("Error while DecodeImgFromBase64")
	ErrorMinioFPutObject     = errors.New("Error in FPutObject")
)

var DeniedAccess = errors.New("user doesn't have permission for this action")

var (
	GatewayErrorMsgSessionNotFound = "сессия отсутствует"
)
