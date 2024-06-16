// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: hello/hello.proto

package hello

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

// Validate checks the field values on HelloRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HelloRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HelloRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HelloRequestMultiError, or
// nil if none found.
func (m *HelloRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *HelloRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return HelloRequestMultiError(errors)
	}

	return nil
}

// HelloRequestMultiError is an error wrapping multiple validation errors
// returned by HelloRequest.ValidateAll() if the designated constraints aren't met.
type HelloRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HelloRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HelloRequestMultiError) AllErrors() []error { return m }

// HelloRequestValidationError is the validation error returned by
// HelloRequest.Validate if the designated constraints aren't met.
type HelloRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelloRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelloRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelloRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelloRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelloRequestValidationError) ErrorName() string { return "HelloRequestValidationError" }

// Error satisfies the builtin error interface
func (e HelloRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelloRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelloRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelloRequestValidationError{}

// Validate checks the field values on HelloReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HelloReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HelloReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HelloReplyMultiError, or
// nil if none found.
func (m *HelloReply) ValidateAll() error {
	return m.validate(true)
}

func (m *HelloReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return HelloReplyMultiError(errors)
	}

	return nil
}

// HelloReplyMultiError is an error wrapping multiple validation errors
// returned by HelloReply.ValidateAll() if the designated constraints aren't met.
type HelloReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HelloReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HelloReplyMultiError) AllErrors() []error { return m }

// HelloReplyValidationError is the validation error returned by
// HelloReply.Validate if the designated constraints aren't met.
type HelloReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelloReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelloReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelloReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelloReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelloReplyValidationError) ErrorName() string { return "HelloReplyValidationError" }

// Error satisfies the builtin error interface
func (e HelloReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelloReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelloReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelloReplyValidationError{}

// Validate checks the field values on PingRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PingRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PingRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PingRequestMultiError, or
// nil if none found.
func (m *PingRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PingRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PingRequestMultiError(errors)
	}

	return nil
}

// PingRequestMultiError is an error wrapping multiple validation errors
// returned by PingRequest.ValidateAll() if the designated constraints aren't met.
type PingRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PingRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PingRequestMultiError) AllErrors() []error { return m }

// PingRequestValidationError is the validation error returned by
// PingRequest.Validate if the designated constraints aren't met.
type PingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PingRequestValidationError) ErrorName() string { return "PingRequestValidationError" }

// Error satisfies the builtin error interface
func (e PingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PingRequestValidationError{}

// Validate checks the field values on PingReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PingReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PingReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PingReplyMultiError, or nil
// if none found.
func (m *PingReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PingReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PingReplyMultiError(errors)
	}

	return nil
}

// PingReplyMultiError is an error wrapping multiple validation errors returned
// by PingReply.ValidateAll() if the designated constraints aren't met.
type PingReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PingReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PingReplyMultiError) AllErrors() []error { return m }

// PingReplyValidationError is the validation error returned by
// PingReply.Validate if the designated constraints aren't met.
type PingReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PingReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PingReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PingReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PingReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PingReplyValidationError) ErrorName() string { return "PingReplyValidationError" }

// Error satisfies the builtin error interface
func (e PingReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPingReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PingReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PingReplyValidationError{}
