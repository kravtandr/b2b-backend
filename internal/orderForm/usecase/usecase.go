package OrderFormUseCase

import (
	cnst "b2b/m/pkg/constants"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/domain"
	"log"

	"github.com/asaskevich/govalidator"
	"github.com/valyala/fasthttp"
)

func NewOrderFormUseCase(OrderFormStorage domain.OrderFormStorage) domain.OrderFormUseCase {
	return OrderFormUseCase{OrderFormStorage: OrderFormStorage}
}

type OrderFormUseCase struct {
	OrderFormStorage domain.OrderFormStorage
}

func (c OrderFormUseCase) Add(orderForm domain.OrderForm) (int, error) {

	err := c.OrderFormStorage.Add(orderForm)
	if err != nil {
		log.Printf("error while adding user")
		return fasthttp.StatusBadRequest, err
	}

	return fasthttp.StatusOK, err
}

func (c OrderFormUseCase) Validate(orderForm *domain.OrderForm) (bool, []byte) {
	var validate_errors domain.ValidateErrors
	_, err := govalidator.ValidateStruct(orderForm)
	if err != nil {
		log.Printf("error while validating orderForm")
		validate_errors.Errors = append(validate_errors.Errors, "Ошибка неверная структура")
	}
	if !govalidator.IsEmail(orderForm.Email) {
		validate_errors.Errors = append(validate_errors.Errors, "Некорректный Email")
	}
	if !govalidator.StringLength(orderForm.Product_category, cnst.MinCategoryLength, cnst.MaxCategoryLength) {
		validate_errors.Errors = append(validate_errors.Errors, "Некорректная категория")
	}
	if !govalidator.StringLength(orderForm.Product_name, cnst.MinProductNameLength, cnst.MaxProductNameLength) {
		validate_errors.Errors = append(validate_errors.Errors, "Некорректное имя товара")
	}
	if !govalidator.StringLength(orderForm.Order_text, cnst.MinOrderTextLength, cnst.MaxOrderTextLength) {
		validate_errors.Errors = append(validate_errors.Errors, "Некорректный текст заказа")
	}
	if !govalidator.StringLength(orderForm.Order_comments, cnst.MinOrderCommentsLength, cnst.MaxOrderCommentsLength) {
		validate_errors.Errors = append(validate_errors.Errors, "Некорректный комменатрий заказа")
	}
	if !govalidator.StringLength(orderForm.Phone, cnst.MinPhoneLength, cnst.MaxPhoneLength) {
		validate_errors.Errors = append(validate_errors.Errors, "Некорректный телефон")
	}
	if !govalidator.StringLength(orderForm.Company_name, cnst.MinCompanyNameLength, cnst.MaxCompanyNameLength) {
		validate_errors.Errors = append(validate_errors.Errors, "Некорректное имя компании")
	}
	if !govalidator.StringLength(orderForm.Itn, cnst.MinItnLength, cnst.MaxItnLength) {
		validate_errors.Errors = append(validate_errors.Errors, "Некорректный ИНН")
	}
	if !govalidator.StringLength(orderForm.Fio, cnst.MinFioLength, cnst.MaxFioLength) {
		validate_errors.Errors = append(validate_errors.Errors, "Некорректное ФИО")
	}
	if !govalidator.MaxStringLength(orderForm.Email, cnst.MaxEmailLength) {
		validate_errors.Errors = append(validate_errors.Errors, "Email слишком длинный")
	}
	bytes, Marchalerr := chttp.ApiResp(validate_errors, err)
	if Marchalerr != nil {
		log.Printf("error while marshalling JSON: %s", err)
		return false, bytes
	}
	if len(validate_errors.Errors) > 0 {
		return false, bytes
	}
	return true, bytes
}
