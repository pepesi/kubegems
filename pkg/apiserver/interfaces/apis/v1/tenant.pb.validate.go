// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kubegems/datas/v1/tenant.proto

package v1

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

// Validate checks the field values on Tenant with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Tenant) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tenant with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TenantMultiError, or nil if none found.
func (m *Tenant) ValidateAll() error {
	return m.validate(true)
}

func (m *Tenant) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(m.GetName()) > 32 {
		err := TenantValidationError{
			field:  "Name",
			reason: "value length must be at most 32 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_Tenant_Name_Pattern.MatchString(m.GetName()) {
		err := TenantValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Remark

	// no validation rules for Enabled

	if len(errors) > 0 {
		return TenantMultiError(errors)
	}
	return nil
}

// TenantMultiError is an error wrapping multiple validation errors returned by
// Tenant.ValidateAll() if the designated constraints aren't met.
type TenantMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TenantMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TenantMultiError) AllErrors() []error { return m }

// TenantValidationError is the validation error returned by Tenant.Validate if
// the designated constraints aren't met.
type TenantValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TenantValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TenantValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TenantValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TenantValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TenantValidationError) ErrorName() string { return "TenantValidationError" }

// Error satisfies the builtin error interface
func (e TenantValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTenant.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TenantValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TenantValidationError{}

var _Tenant_Name_Pattern = regexp.MustCompile("^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$")

// Validate checks the field values on CreateTenantRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTenantRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTenantRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTenantRequestMultiError, or nil if none found.
func (m *CreateTenantRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTenantRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTenant()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateTenantRequestValidationError{
					field:  "Tenant",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateTenantRequestValidationError{
					field:  "Tenant",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTenant()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTenantRequestValidationError{
				field:  "Tenant",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateTenantRequestMultiError(errors)
	}
	return nil
}

// CreateTenantRequestMultiError is an error wrapping multiple validation
// errors returned by CreateTenantRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateTenantRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTenantRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTenantRequestMultiError) AllErrors() []error { return m }

// CreateTenantRequestValidationError is the validation error returned by
// CreateTenantRequest.Validate if the designated constraints aren't met.
type CreateTenantRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTenantRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTenantRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTenantRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTenantRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTenantRequestValidationError) ErrorName() string {
	return "CreateTenantRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTenantRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTenantRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTenantRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTenantRequestValidationError{}

// Validate checks the field values on CreateTenantResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTenantResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTenantResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTenantResponseMultiError, or nil if none found.
func (m *CreateTenantResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTenantResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTenant()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateTenantResponseValidationError{
					field:  "Tenant",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateTenantResponseValidationError{
					field:  "Tenant",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTenant()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTenantResponseValidationError{
				field:  "Tenant",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Succeed

	// no validation rules for Message

	if len(errors) > 0 {
		return CreateTenantResponseMultiError(errors)
	}
	return nil
}

// CreateTenantResponseMultiError is an error wrapping multiple validation
// errors returned by CreateTenantResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateTenantResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTenantResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTenantResponseMultiError) AllErrors() []error { return m }

// CreateTenantResponseValidationError is the validation error returned by
// CreateTenantResponse.Validate if the designated constraints aren't met.
type CreateTenantResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTenantResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTenantResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTenantResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTenantResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTenantResponseValidationError) ErrorName() string {
	return "CreateTenantResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTenantResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTenantResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTenantResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTenantResponseValidationError{}

// Validate checks the field values on DeleteTenantRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteTenantRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTenantRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTenantRequestMultiError, or nil if none found.
func (m *DeleteTenantRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTenantRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	if len(errors) > 0 {
		return DeleteTenantRequestMultiError(errors)
	}
	return nil
}

// DeleteTenantRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteTenantRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteTenantRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTenantRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTenantRequestMultiError) AllErrors() []error { return m }

// DeleteTenantRequestValidationError is the validation error returned by
// DeleteTenantRequest.Validate if the designated constraints aren't met.
type DeleteTenantRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTenantRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTenantRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTenantRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTenantRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTenantRequestValidationError) ErrorName() string {
	return "DeleteTenantRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTenantRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTenantRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTenantRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTenantRequestValidationError{}

// Validate checks the field values on DeleteTenantResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteTenantResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTenantResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTenantResponseMultiError, or nil if none found.
func (m *DeleteTenantResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTenantResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Succeed

	// no validation rules for Message

	if len(errors) > 0 {
		return DeleteTenantResponseMultiError(errors)
	}
	return nil
}

// DeleteTenantResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteTenantResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteTenantResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTenantResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTenantResponseMultiError) AllErrors() []error { return m }

// DeleteTenantResponseValidationError is the validation error returned by
// DeleteTenantResponse.Validate if the designated constraints aren't met.
type DeleteTenantResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTenantResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTenantResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTenantResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTenantResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTenantResponseValidationError) ErrorName() string {
	return "DeleteTenantResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTenantResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTenantResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTenantResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTenantResponseValidationError{}
