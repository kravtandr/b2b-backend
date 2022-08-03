package userUseCase

import (
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/domain"
	"log"

	"github.com/asaskevich/govalidator"
	"github.com/valyala/fasthttp"
)

func NewUserUseCase(userStorage domain.UserStorage) domain.UserUseCase {
	return userUseCase{userStorage: userStorage}
}

type userUseCase struct {
	userStorage domain.UserStorage
}

// func (c userUseCase) SearchCompanies(key domain.CompanySearch) (value []byte, err error) {
// 	companies, err := c.companyStorage.SearchCompanies(key.Name)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	bytes, err := chttp.ApiResp(companies)
// 	if err != nil {
// 		log.Printf("error while marshalling JSON: %s", err)
// 	}
// 	return bytes, err
// }

func (u userUseCase) Login(user *domain.User) (int, error) {
	foundUser, err := u.GetByEmail(user.Email)
	if err != nil {
		log.Printf("error while login-GetByEmail")
		log.Print(err)
		return fasthttp.StatusNotFound, err
	}

	if foundUser.Password != user.Password {
		return fasthttp.StatusBadRequest, err
	}

	return fasthttp.StatusOK, err
}

func (u userUseCase) Registration(user *domain.User) (int, error) {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		log.Printf("error while validating user")
		return fasthttp.StatusBadRequest, err
	}
	log.Printf(user.Email)

	_, err = u.GetByEmail(user.Email)
	if err == nil {
		log.Printf("user with this email already exists")
		return fasthttp.StatusBadRequest, err
	}

	err = u.Add(*user)
	if err != nil {
		log.Printf("error while adding user")
		return fasthttp.StatusBadRequest, err
	}

	return fasthttp.StatusOK, err
}

func (u userUseCase) RegistrationCompany(company *domain.Company) (int, error) {
	_, err := govalidator.ValidateStruct(company)
	if err != nil {
		log.Printf("error while validating user")
		return fasthttp.StatusBadRequest, err
	}
	err = u.AddCompany(*company)
	if err != nil {
		log.Printf("error while adding user")
		return fasthttp.StatusBadRequest, err
	}

	return fasthttp.StatusOK, err
}

func (c userUseCase) GetByEmail(key string) (value domain.User, err error) {
	return c.userStorage.GetByEmail(key)
}

func (c userUseCase) GetPublicUserByEmail(key string) (value []byte, err error) {
	user, err := c.userStorage.GetPublicUserByEmail(key)
	if err != nil {
		log.Printf("error while GetPublicUserByEmail: %s", err)
	}
	if (user == domain.PublicUser{}) {
		return []byte{}, err
	} else {
		bytes, err := chttp.ApiResp(user)
		if err != nil {
			log.Printf("error while marshalling JSON: %s", err)
		}
		return bytes, err
	}

}

func (c userUseCase) Add(user domain.User) error {
	return c.userStorage.Add(user)
}

func (c userUseCase) AddCompany(company domain.Company) error {
	return c.userStorage.AddCompany(company)
}

func (c userUseCase) Validate(user *domain.User) bool {
	// if !govalidator.IsEmail(company.Email) ||
	// 	!govalidator.StringLength(company.Password, cnst.MinPasswordLength, cnst.MaxPasswordLength) ||
	// 	!govalidator.MaxStringLength(company.Email, cnst.MaxEmailLength) {
	// 	return false
	// }
	return true
}

// func (c companyUseCase) GetCompanyById(key string) (value []byte, err error) {
// 	company, err := c.companyStorage.GetCompanyById(key)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	bytes, err := chttp.ApiResp(company)
// 	if err != nil {
// 		log.Printf("error while marshalling JSON: %s", err)
// 	}
// 	return bytes, err
// }

// func (c companyUseCase) GetCompaniesByCategoryId(key string) (value []byte, err error) {
// 	companies, err := c.companyStorage.GetCompaniesByCategoryId(key)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	bytes, err := chttp.ApiResp(companies)
// 	if err != nil {
// 		log.Printf("error while marshalling JSON: %s", err)
// 	}
// 	return bytes, err
// }

// func (u CompanyUseCase) Update(id int, updatedUser domain.User) error {

// 	user, err := u.GetByEmail(updatedUser.Email)
// 	if err == nil && user.Id != id {
// 		return errors.New("user with this email already exists") // change later
// 	}

// 	return u.companyStorage.Update(id, updatedUser)
// }

// func (u CompanyUseCase) Delete(id int) error {
// 	return u.companyStorage.Delete(id)
// }
