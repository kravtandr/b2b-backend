// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/services/fastOrder/api.proto

package fastOrder_service

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on FastOrderRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FastOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FastOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FastOrderRequestMultiError, or nil if none found.
func (m *FastOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FastOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Role

	if l := utf8.RuneCountInString(m.GetProductCategory()); l < 3 || l > 128 {
		err := FastOrderRequestValidationError{
			field:  "ProductCategory",
			reason: "value length must be between 3 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetProductName()); l < 3 || l > 256 {
		err := FastOrderRequestValidationError{
			field:  "ProductName",
			reason: "value length must be between 3 and 256 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetOrderText()); l < 3 || l > 1000 {
		err := FastOrderRequestValidationError{
			field:  "OrderText",
			reason: "value length must be between 3 and 1000 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetOrderComments()); l < 3 || l > 1000 {
		err := FastOrderRequestValidationError{
			field:  "OrderComments",
			reason: "value length must be between 3 and 1000 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetFio()); l < 7 || l > 768 {
		err := FastOrderRequestValidationError{
			field:  "Fio",
			reason: "value length must be between 7 and 768 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetEmail()); l < 3 || l > 256 {
		err := FastOrderRequestValidationError{
			field:  "Email",
			reason: "value length must be between 3 and 256 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = FastOrderRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetPhone()); l < 11 || l > 13 {
		err := FastOrderRequestValidationError{
			field:  "Phone",
			reason: "value length must be between 11 and 13 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetCompanyName()); l < 3 || l > 128 {
		err := FastOrderRequestValidationError{
			field:  "CompanyName",
			reason: "value length must be between 3 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetItn()); l < 10 || l > 12 {
		err := FastOrderRequestValidationError{
			field:  "Itn",
			reason: "value length must be between 10 and 12 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return FastOrderRequestMultiError(errors)
	}
	return nil
}

func (m *FastOrderRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *FastOrderRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// FastOrderRequestMultiError is an error wrapping multiple validation errors
// returned by FastOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type FastOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FastOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FastOrderRequestMultiError) AllErrors() []error { return m }

// FastOrderRequestValidationError is the validation error returned by
// FastOrderRequest.Validate if the designated constraints aren't met.
type FastOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FastOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FastOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FastOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FastOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FastOrderRequestValidationError) ErrorName() string { return "FastOrderRequestValidationError" }

// Error satisfies the builtin error interface
func (e FastOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFastOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FastOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FastOrderRequestValidationError{}

// Validate checks the field values on LandingOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LandingOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LandingOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LandingOrderRequestMultiError, or nil if none found.
func (m *LandingOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LandingOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetProductCategory()); l < 3 || l > 128 {
		err := LandingOrderRequestValidationError{
			field:  "ProductCategory",
			reason: "value length must be between 3 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetDeliveryAddress()); l < 3 || l > 256 {
		err := LandingOrderRequestValidationError{
			field:  "DeliveryAddress",
			reason: "value length must be between 3 and 256 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetDeliveryDate()); l < 3 || l > 256 {
		err := LandingOrderRequestValidationError{
			field:  "DeliveryDate",
			reason: "value length must be between 3 and 256 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetOrderText()); l < 3 || l > 1000 {
		err := LandingOrderRequestValidationError{
			field:  "OrderText",
			reason: "value length must be between 3 and 1000 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetEmail()); l < 3 || l > 256 {
		err := LandingOrderRequestValidationError{
			field:  "Email",
			reason: "value length must be between 3 and 256 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = LandingOrderRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetItn()); l < 10 || l > 12 {
		err := LandingOrderRequestValidationError{
			field:  "Itn",
			reason: "value length must be between 10 and 12 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return LandingOrderRequestMultiError(errors)
	}
	return nil
}

func (m *LandingOrderRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *LandingOrderRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// LandingOrderRequestMultiError is an error wrapping multiple validation
// errors returned by LandingOrderRequest.ValidateAll() if the designated
// constraints aren't met.
type LandingOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LandingOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LandingOrderRequestMultiError) AllErrors() []error { return m }

// LandingOrderRequestValidationError is the validation error returned by
// LandingOrderRequest.Validate if the designated constraints aren't met.
type LandingOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LandingOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LandingOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LandingOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LandingOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LandingOrderRequestValidationError) ErrorName() string {
	return "LandingOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e LandingOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLandingOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LandingOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LandingOrderRequestValidationError{}
