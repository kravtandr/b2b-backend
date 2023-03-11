package grpc_errors

import (
	"net/http"

	"b2b/m/pkg/error_adapter"
	"b2b/m/pkg/errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	PreparedAuthServiceErrorMap = map[error]error{
		errors.UserDoesNotExist:    status.Error(codes.NotFound, "user not found"),
		errors.SessionDoesNotExist: status.Error(codes.NotFound, "session not found"),
		errors.WrongUserPassword:   status.Error(codes.PermissionDenied, "wrong password"),
	}
	PreparedFastOrderServiceErrorMap = map[error]error{
		errors.CompanyDoesNotExist: status.Error(codes.NotFound, "not found"),
	}
	PreparedCompanyServiceErrorMap = map[error]error{
		errors.CompanyDoesNotExist: status.Error(codes.NotFound, "not found"),
	}
	PreparedProductsServiceErrorMap = map[error]error{
		errors.ProductDoesNotExist: status.Error(codes.NotFound, "not found"),
	}
)

var (
	PreparedAuthErrors = map[codes.Code]error_adapter.HttpError{
		codes.NotFound: {
			MSG:  "Пользователь не авторизован",
			Code: http.StatusUnauthorized,
		},
		codes.InvalidArgument: {
			MSG:  "Пользователь не авторизован",
			Code: http.StatusUnauthorized,
		},
	}
	CommonAuthError = error_adapter.HttpError{
		MSG:  "Произошла ошибка обращения во внутренний сервис",
		Code: http.StatusBadRequest,
	}

	UserGatewayError = map[codes.Code]error_adapter.HttpError{
		codes.NotFound: {
			MSG:  "Запись отсутствует",
			Code: http.StatusNotFound,
		},
		codes.PermissionDenied: {
			MSG:  "Доступ к данным запрешен",
			Code: http.StatusBadRequest,
		},
	}
	CommonError = error_adapter.HttpError{
		MSG:  "Произошла ошибка обращения во внутренний сервис",
		Code: http.StatusBadRequest,
	}
	Fail = error_adapter.HttpError{
		MSG:  "Total fail",
		Code: http.StatusBadRequest,
	}

	PreparedCountryErrors = map[error]error_adapter.HttpError{
		errors.CountryDoesNotExist: {
			MSG:  "Запись отсутствует",
			Code: http.StatusNotFound,
		},
	}
)
