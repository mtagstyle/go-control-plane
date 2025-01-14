// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/percent.proto

package _type

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

// Validate checks the field values on Percent with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Percent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Percent with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PercentMultiError, or nil if none found.
func (m *Percent) ValidateAll() error {
	return m.validate(true)
}

func (m *Percent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetValue(); val < 0 || val > 100 {
		err := PercentValidationError{
			field:  "Value",
			reason: "value must be inside range [0, 100]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PercentMultiError(errors)
	}
	return nil
}

// PercentMultiError is an error wrapping multiple validation errors returned
// by Percent.ValidateAll() if the designated constraints aren't met.
type PercentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PercentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PercentMultiError) AllErrors() []error { return m }

// PercentValidationError is the validation error returned by Percent.Validate
// if the designated constraints aren't met.
type PercentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PercentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PercentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PercentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PercentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PercentValidationError) ErrorName() string { return "PercentValidationError" }

// Error satisfies the builtin error interface
func (e PercentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPercent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PercentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PercentValidationError{}

// Validate checks the field values on FractionalPercent with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FractionalPercent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FractionalPercent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FractionalPercentMultiError, or nil if none found.
func (m *FractionalPercent) ValidateAll() error {
	return m.validate(true)
}

func (m *FractionalPercent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Numerator

	if _, ok := FractionalPercent_DenominatorType_name[int32(m.GetDenominator())]; !ok {
		err := FractionalPercentValidationError{
			field:  "Denominator",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return FractionalPercentMultiError(errors)
	}
	return nil
}

// FractionalPercentMultiError is an error wrapping multiple validation errors
// returned by FractionalPercent.ValidateAll() if the designated constraints
// aren't met.
type FractionalPercentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FractionalPercentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FractionalPercentMultiError) AllErrors() []error { return m }

// FractionalPercentValidationError is the validation error returned by
// FractionalPercent.Validate if the designated constraints aren't met.
type FractionalPercentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FractionalPercentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FractionalPercentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FractionalPercentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FractionalPercentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FractionalPercentValidationError) ErrorName() string {
	return "FractionalPercentValidationError"
}

// Error satisfies the builtin error interface
func (e FractionalPercentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFractionalPercent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FractionalPercentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FractionalPercentValidationError{}
