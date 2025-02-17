// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/on_demand/v2/on_demand.proto

package on_demandv2

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

// Validate checks the field values on OnDemand with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OnDemand) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OnDemand with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OnDemandMultiError, or nil
// if none found.
func (m *OnDemand) ValidateAll() error {
	return m.validate(true)
}

func (m *OnDemand) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OnDemandMultiError(errors)
	}
	return nil
}

// OnDemandMultiError is an error wrapping multiple validation errors returned
// by OnDemand.ValidateAll() if the designated constraints aren't met.
type OnDemandMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OnDemandMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OnDemandMultiError) AllErrors() []error { return m }

// OnDemandValidationError is the validation error returned by
// OnDemand.Validate if the designated constraints aren't met.
type OnDemandValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OnDemandValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OnDemandValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OnDemandValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OnDemandValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OnDemandValidationError) ErrorName() string { return "OnDemandValidationError" }

// Error satisfies the builtin error interface
func (e OnDemandValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOnDemand.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OnDemandValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OnDemandValidationError{}
