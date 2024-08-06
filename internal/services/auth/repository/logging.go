package repository

import (
	company_models "b2b/m/internal/services/company/models"
	"context"

	"b2b/m/internal/services/auth/models"

	"go.uber.org/zap"
)

const (
	module = "auth_repo"
)

type loggingMiddleware struct {
	logger *zap.SugaredLogger

	next AuthRepository
}

func NewLoggingMiddleware(logger *zap.SugaredLogger, next AuthRepository) AuthRepository {
	return &loggingMiddleware{
		logger: logger,
		next:   next,
	}
}

func (l *loggingMiddleware) CountUsersPayments(ctx context.Context, userID int64) (count int, err error) {
	l.logger.Infow(module,
		"Action", "CountUsersPayments",
		"Request", userID,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "CountUsersPayments",
				"Request", userID,
				"Error", err,
			)
		}
	}()
	return l.next.CountUsersPayments(ctx, userID)
}

func (l *loggingMiddleware) AddPayment(ctx context.Context, payment *models.Payment) (p *models.Payment, err error) {
	l.logger.Infow(module,
		"Action", "AddPayment",
		"Request", payment,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "AddPayment",
				"Request", payment,
				"Error", err,
			)
		}
	}()
	return l.next.AddPayment(ctx, payment)
}

func (l *loggingMiddleware) UpdatePayment(ctx context.Context, payment *models.Payment) (p *models.Payment, err error) {
	l.logger.Infow(module,
		"Action", "UpdatePayment",
		"Request", payment,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "UpdatePayment",
				"Request", payment,
				"Error", err,
			)
		}
	}()
	return l.next.UpdatePayment(ctx, payment)
}

func (l *loggingMiddleware) GetPayment(ctx context.Context, paymentID string) (p *models.Payment, err error) {
	l.logger.Infow(module,
		"Action", "GetPayment",
		"Request", paymentID,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetPayment",
				"Request", paymentID,
				"Error", err,
			)
		}
	}()
	return l.next.GetPayment(ctx, paymentID)
}

func (l *loggingMiddleware) GetUsersPayments(ctx context.Context, userID int64) (p *models.Payments, err error) {
	l.logger.Infow(module,
		"Action", "GetUsersPayments",
		"Request", userID,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetUsersPayments",
				"Request", userID,
				"Error", err,
			)
		}
	}()
	return l.next.GetUsersPayments(ctx, userID)
}

func (l *loggingMiddleware) UpdateUserBalance(ctx context.Context, userID int64, newBalance int64) (u *models.User, err error) {
	l.logger.Infow(module,
		"Action", "UpdateUserBalance",
		"Request", userID,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "UpdateUserBalance",
				"Request", userID,
				"Error", err,
			)
		}
	}()
	return l.next.UpdateUserBalance(ctx, userID, newBalance)
}

func (l *loggingMiddleware) GetUserByEmail(ctx context.Context, email string) (u *models.User, err error) {
	l.logger.Infow(module,
		"Action", "GetUserByEmail",
		"Request", email,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetUserByEmail",
				"Request", email,
				"Error", err,
			)
		}
	}()

	return l.next.GetUserByEmail(ctx, email)
}
func (l *loggingMiddleware) GetUserByID(ctx context.Context, ID int64) (u *models.User, err error) {
	l.logger.Infow(module,
		"Action", "GetUserByID",
		"Request", ID,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetUserByID",
				"Request", ID,
				"Error", err,
			)
		}
	}()

	return l.next.GetUserByID(ctx, ID)
}

func (l *loggingMiddleware) UpdateUser(ctx context.Context, user *models.User) (u *models.User, err error) {
	l.logger.Infow(module,
		"Action", "UpdateUser",
		"Request", user,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "UpdateUser",
				"Request", user,
				"Error", err,
			)
		}
	}()

	return l.next.UpdateUser(ctx, user)
}

func (l *loggingMiddleware) GetUsersCompany(ctx context.Context, userId int64) (u *company_models.Company, err error) {
	l.logger.Infow(module,
		"Action", "GetUsersCompany",
		"Request", userId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetUsersCompany",
				"Request", userId,
				"Error", err,
			)
		}
	}()

	return l.next.GetUsersCompany(ctx, userId)
}

func (l *loggingMiddleware) GetCompanyUserLink(ctx context.Context, userId int64, companyId int64) (u *company_models.CompaniesUsersLink, err error) {
	l.logger.Infow(module,
		"Action", "GetCompanyUserLink",
		"Request", userId, companyId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetCompanyUserLink",
				"Request", userId, companyId,
				"Error", err,
			)
		}
	}()

	return l.next.GetCompanyUserLink(ctx, userId, companyId)
}
func (l *loggingMiddleware) CreateUser(ctx context.Context, user *models.User) (u *models.User, err error) {
	l.logger.Infow(module,
		"Action", "CreateUser",
		"Request", user,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "CreateUser",
				"Request", user,
				"Error", err,
			)
		}
	}()

	return l.next.CreateUser(ctx, user)
}

func (l *loggingMiddleware) FastRegistration(ctx context.Context, newCompany *company_models.Company, user *models.User, post string) (err error) {
	l.logger.Infow(module,
		"Action", "FastRegistration",
		"Request", newCompany, user, post,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "FastRegistration",
				"Request", newCompany, user, post,
				"Error", err,
			)
		}
	}()

	return l.next.FastRegistration(ctx, newCompany, user, post)
}

//	func (l *loggingMiddleware) UpdateUser(ctx context.Context, user *models.User) (u *models.User, err error) {
//		l.logger.Infow(module,
//			"Action", "UpdateUser",
//			"Request", user,
//		)
//		defer func() {
//			if err != nil {
//				l.logger.Infow(module,
//					"Action", "UpdateUser",
//					"Request", user,
//					"Error", err,
//				)
//			}
//		}()
//
//		return l.next.UpdateUser(ctx, user)
//	}
func (l *loggingMiddleware) CreateUserSession(ctx context.Context, userID int64, hash string) (err error) {
	l.logger.Infow(module,
		"Action", "CreateUserSession",
		"Request", userID,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "CreateUserSession",
				"Request", userID,
				"Error", err,
			)
		}
	}()

	return l.next.CreateUserSession(ctx, userID, hash)
}
func (l *loggingMiddleware) ValidateUserSession(ctx context.Context, hash string) (ID int64, err error) {
	l.logger.Infow(module,
		"Action", "ValidateUserSession",
		"Request", hash,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "ValidateUserSession",
				"Request", hash,
				"Error", err,
			)
		}
	}()

	return l.next.ValidateUserSession(ctx, hash)
}
func (l *loggingMiddleware) RemoveUserSession(ctx context.Context, hash string) (err error) {
	l.logger.Infow(module,
		"Action", " RemoveUserSession",
		"Request", hash,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", " RemoveUserSession",
				"Request", hash,
				"Error", err,
			)
		}
	}()

	return l.next.RemoveUserSession(ctx, hash)
}

func (l *loggingMiddleware) GetUserInfo(ctx context.Context, id int64) (u *models.User, err error) {
	l.logger.Infow(module,
		"Action", "GetUserInfo",
		"Request", id,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetUserInfo",
				"Request", id,
				"Error", err,
			)
		}
	}()

	return l.next.GetUserInfo(ctx, id)
}
