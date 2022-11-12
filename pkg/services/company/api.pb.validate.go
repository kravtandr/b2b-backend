// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/services/company/api.proto

package company_service

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

// Validate checks the field values on GetCompanyRequestByInn with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCompanyRequestByInn) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCompanyRequestByInn with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCompanyRequestByInnMultiError, or nil if none found.
func (m *GetCompanyRequestByInn) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCompanyRequestByInn) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Inn

	if len(errors) > 0 {
		return GetCompanyRequestByInnMultiError(errors)
	}
	return nil
}

// GetCompanyRequestByInnMultiError is an error wrapping multiple validation
// errors returned by GetCompanyRequestByInn.ValidateAll() if the designated
// constraints aren't met.
type GetCompanyRequestByInnMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCompanyRequestByInnMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCompanyRequestByInnMultiError) AllErrors() []error { return m }

// GetCompanyRequestByInnValidationError is the validation error returned by
// GetCompanyRequestByInn.Validate if the designated constraints aren't met.
type GetCompanyRequestByInnValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCompanyRequestByInnValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCompanyRequestByInnValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCompanyRequestByInnValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCompanyRequestByInnValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCompanyRequestByInnValidationError) ErrorName() string {
	return "GetCompanyRequestByInnValidationError"
}

// Error satisfies the builtin error interface
func (e GetCompanyRequestByInnValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCompanyRequestByInn.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCompanyRequestByInnValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCompanyRequestByInnValidationError{}

// Validate checks the field values on GetCompanyRequestById with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCompanyRequestById) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCompanyRequestById with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCompanyRequestByIdMultiError, or nil if none found.
func (m *GetCompanyRequestById) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCompanyRequestById) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetCompanyRequestByIdMultiError(errors)
	}
	return nil
}

// GetCompanyRequestByIdMultiError is an error wrapping multiple validation
// errors returned by GetCompanyRequestById.ValidateAll() if the designated
// constraints aren't met.
type GetCompanyRequestByIdMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCompanyRequestByIdMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCompanyRequestByIdMultiError) AllErrors() []error { return m }

// GetCompanyRequestByIdValidationError is the validation error returned by
// GetCompanyRequestById.Validate if the designated constraints aren't met.
type GetCompanyRequestByIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCompanyRequestByIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCompanyRequestByIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCompanyRequestByIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCompanyRequestByIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCompanyRequestByIdValidationError) ErrorName() string {
	return "GetCompanyRequestByIdValidationError"
}

// Error satisfies the builtin error interface
func (e GetCompanyRequestByIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCompanyRequestById.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCompanyRequestByIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCompanyRequestByIdValidationError{}

// Validate checks the field values on GetCompanyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCompanyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCompanyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCompanyResponseMultiError, or nil if none found.
func (m *GetCompanyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCompanyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for LegalName

	// no validation rules for Itn

	// no validation rules for Psrn

	// no validation rules for Address

	// no validation rules for LegalAddress

	// no validation rules for Email

	// no validation rules for Phone

	// no validation rules for Link

	// no validation rules for Activity

	// no validation rules for OwnerId

	// no validation rules for Rating

	// no validation rules for Verified

	if len(errors) > 0 {
		return GetCompanyResponseMultiError(errors)
	}
	return nil
}

// GetCompanyResponseMultiError is an error wrapping multiple validation errors
// returned by GetCompanyResponse.ValidateAll() if the designated constraints
// aren't met.
type GetCompanyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCompanyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCompanyResponseMultiError) AllErrors() []error { return m }

// GetCompanyResponseValidationError is the validation error returned by
// GetCompanyResponse.Validate if the designated constraints aren't met.
type GetCompanyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCompanyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCompanyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCompanyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCompanyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCompanyResponseValidationError) ErrorName() string {
	return "GetCompanyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetCompanyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCompanyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCompanyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCompanyResponseValidationError{}

// Validate checks the field values on GetCompanyAndPostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCompanyAndPostResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCompanyAndPostResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCompanyAndPostResponseMultiError, or nil if none found.
func (m *GetCompanyAndPostResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCompanyAndPostResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for LegalName

	// no validation rules for Itn

	// no validation rules for Psrn

	// no validation rules for Address

	// no validation rules for LegalAddress

	// no validation rules for Email

	// no validation rules for Phone

	// no validation rules for Link

	// no validation rules for Activity

	// no validation rules for OwnerId

	// no validation rules for Rating

	// no validation rules for Verified

	// no validation rules for Post

	if len(errors) > 0 {
		return GetCompanyAndPostResponseMultiError(errors)
	}
	return nil
}

// GetCompanyAndPostResponseMultiError is an error wrapping multiple validation
// errors returned by GetCompanyAndPostResponse.ValidateAll() if the
// designated constraints aren't met.
type GetCompanyAndPostResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCompanyAndPostResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCompanyAndPostResponseMultiError) AllErrors() []error { return m }

// GetCompanyAndPostResponseValidationError is the validation error returned by
// GetCompanyAndPostResponse.Validate if the designated constraints aren't met.
type GetCompanyAndPostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCompanyAndPostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCompanyAndPostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCompanyAndPostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCompanyAndPostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCompanyAndPostResponseValidationError) ErrorName() string {
	return "GetCompanyAndPostResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetCompanyAndPostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCompanyAndPostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCompanyAndPostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCompanyAndPostResponseValidationError{}

// Validate checks the field values on UpdateCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateCompanyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateCompanyRequestMultiError, or nil if none found.
func (m *UpdateCompanyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCompanyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Address

	// no validation rules for LegalAddress

	// no validation rules for Itn

	// no validation rules for Phone

	// no validation rules for Link

	// no validation rules for Activity

	// no validation rules for OwnerId

	// no validation rules for Post

	if len(errors) > 0 {
		return UpdateCompanyRequestMultiError(errors)
	}
	return nil
}

// UpdateCompanyRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateCompanyRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateCompanyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCompanyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCompanyRequestMultiError) AllErrors() []error { return m }

// UpdateCompanyRequestValidationError is the validation error returned by
// UpdateCompanyRequest.Validate if the designated constraints aren't met.
type UpdateCompanyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCompanyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCompanyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCompanyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCompanyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCompanyRequestValidationError) ErrorName() string {
	return "UpdateCompanyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCompanyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCompanyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCompanyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCompanyRequestValidationError{}
