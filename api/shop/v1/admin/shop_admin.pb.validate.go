// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/shop/v1/admin/shop_admin.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on CreateShopAdminRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateShopAdminRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CreateShopAdminRequestValidationError is the validation error returned by
// CreateShopAdminRequest.Validate if the designated constraints aren't met.
type CreateShopAdminRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateShopAdminRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateShopAdminRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateShopAdminRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateShopAdminRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateShopAdminRequestValidationError) ErrorName() string {
	return "CreateShopAdminRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateShopAdminRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateShopAdminRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateShopAdminRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateShopAdminRequestValidationError{}

// Validate checks the field values on CreateShopAdminReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateShopAdminReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CreateShopAdminReplyValidationError is the validation error returned by
// CreateShopAdminReply.Validate if the designated constraints aren't met.
type CreateShopAdminReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateShopAdminReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateShopAdminReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateShopAdminReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateShopAdminReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateShopAdminReplyValidationError) ErrorName() string {
	return "CreateShopAdminReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateShopAdminReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateShopAdminReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateShopAdminReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateShopAdminReplyValidationError{}

// Validate checks the field values on UpdateShopAdminRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateShopAdminRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateShopAdminRequestValidationError is the validation error returned by
// UpdateShopAdminRequest.Validate if the designated constraints aren't met.
type UpdateShopAdminRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateShopAdminRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateShopAdminRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateShopAdminRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateShopAdminRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateShopAdminRequestValidationError) ErrorName() string {
	return "UpdateShopAdminRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateShopAdminRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateShopAdminRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateShopAdminRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateShopAdminRequestValidationError{}

// Validate checks the field values on UpdateShopAdminReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateShopAdminReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateShopAdminReplyValidationError is the validation error returned by
// UpdateShopAdminReply.Validate if the designated constraints aren't met.
type UpdateShopAdminReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateShopAdminReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateShopAdminReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateShopAdminReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateShopAdminReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateShopAdminReplyValidationError) ErrorName() string {
	return "UpdateShopAdminReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateShopAdminReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateShopAdminReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateShopAdminReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateShopAdminReplyValidationError{}

// Validate checks the field values on DeleteShopAdminRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteShopAdminRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteShopAdminRequestValidationError is the validation error returned by
// DeleteShopAdminRequest.Validate if the designated constraints aren't met.
type DeleteShopAdminRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteShopAdminRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteShopAdminRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteShopAdminRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteShopAdminRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteShopAdminRequestValidationError) ErrorName() string {
	return "DeleteShopAdminRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteShopAdminRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteShopAdminRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteShopAdminRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteShopAdminRequestValidationError{}

// Validate checks the field values on DeleteShopAdminReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteShopAdminReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteShopAdminReplyValidationError is the validation error returned by
// DeleteShopAdminReply.Validate if the designated constraints aren't met.
type DeleteShopAdminReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteShopAdminReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteShopAdminReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteShopAdminReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteShopAdminReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteShopAdminReplyValidationError) ErrorName() string {
	return "DeleteShopAdminReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteShopAdminReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteShopAdminReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteShopAdminReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteShopAdminReplyValidationError{}

// Validate checks the field values on GetShopAdminRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetShopAdminRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetShopAdminRequestValidationError is the validation error returned by
// GetShopAdminRequest.Validate if the designated constraints aren't met.
type GetShopAdminRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetShopAdminRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetShopAdminRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetShopAdminRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetShopAdminRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetShopAdminRequestValidationError) ErrorName() string {
	return "GetShopAdminRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetShopAdminRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetShopAdminRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetShopAdminRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetShopAdminRequestValidationError{}

// Validate checks the field values on GetShopAdminReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetShopAdminReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetShopAdminReplyValidationError is the validation error returned by
// GetShopAdminReply.Validate if the designated constraints aren't met.
type GetShopAdminReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetShopAdminReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetShopAdminReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetShopAdminReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetShopAdminReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetShopAdminReplyValidationError) ErrorName() string {
	return "GetShopAdminReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetShopAdminReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetShopAdminReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetShopAdminReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetShopAdminReplyValidationError{}

// Validate checks the field values on ListShopAdminRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListShopAdminRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListShopAdminRequestValidationError is the validation error returned by
// ListShopAdminRequest.Validate if the designated constraints aren't met.
type ListShopAdminRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListShopAdminRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListShopAdminRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListShopAdminRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListShopAdminRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListShopAdminRequestValidationError) ErrorName() string {
	return "ListShopAdminRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListShopAdminRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListShopAdminRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListShopAdminRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListShopAdminRequestValidationError{}

// Validate checks the field values on ListShopAdminReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListShopAdminReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Age

	return nil
}

// ListShopAdminReplyValidationError is the validation error returned by
// ListShopAdminReply.Validate if the designated constraints aren't met.
type ListShopAdminReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListShopAdminReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListShopAdminReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListShopAdminReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListShopAdminReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListShopAdminReplyValidationError) ErrorName() string {
	return "ListShopAdminReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListShopAdminReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListShopAdminReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListShopAdminReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListShopAdminReplyValidationError{}
