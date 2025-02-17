// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/event_reporting/v2alpha/event_reporting_service.proto

package v2alpha

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

// Validate checks the field values on StreamEventsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StreamEventsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamEventsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StreamEventsRequestMultiError, or nil if none found.
func (m *StreamEventsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamEventsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetIdentifier()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StreamEventsRequestValidationError{
					field:  "Identifier",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StreamEventsRequestValidationError{
					field:  "Identifier",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIdentifier()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamEventsRequestValidationError{
				field:  "Identifier",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetEvents()) < 1 {
		err := StreamEventsRequestValidationError{
			field:  "Events",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetEvents() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StreamEventsRequestValidationError{
						field:  fmt.Sprintf("Events[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StreamEventsRequestValidationError{
						field:  fmt.Sprintf("Events[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamEventsRequestValidationError{
					field:  fmt.Sprintf("Events[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return StreamEventsRequestMultiError(errors)
	}
	return nil
}

// StreamEventsRequestMultiError is an error wrapping multiple validation
// errors returned by StreamEventsRequest.ValidateAll() if the designated
// constraints aren't met.
type StreamEventsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamEventsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamEventsRequestMultiError) AllErrors() []error { return m }

// StreamEventsRequestValidationError is the validation error returned by
// StreamEventsRequest.Validate if the designated constraints aren't met.
type StreamEventsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamEventsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamEventsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamEventsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamEventsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamEventsRequestValidationError) ErrorName() string {
	return "StreamEventsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StreamEventsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamEventsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamEventsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamEventsRequestValidationError{}

// Validate checks the field values on StreamEventsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StreamEventsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamEventsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StreamEventsResponseMultiError, or nil if none found.
func (m *StreamEventsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamEventsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StreamEventsResponseMultiError(errors)
	}
	return nil
}

// StreamEventsResponseMultiError is an error wrapping multiple validation
// errors returned by StreamEventsResponse.ValidateAll() if the designated
// constraints aren't met.
type StreamEventsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamEventsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamEventsResponseMultiError) AllErrors() []error { return m }

// StreamEventsResponseValidationError is the validation error returned by
// StreamEventsResponse.Validate if the designated constraints aren't met.
type StreamEventsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamEventsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamEventsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamEventsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamEventsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamEventsResponseValidationError) ErrorName() string {
	return "StreamEventsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StreamEventsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamEventsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamEventsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamEventsResponseValidationError{}

// Validate checks the field values on StreamEventsRequest_Identifier with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StreamEventsRequest_Identifier) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamEventsRequest_Identifier with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// StreamEventsRequest_IdentifierMultiError, or nil if none found.
func (m *StreamEventsRequest_Identifier) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamEventsRequest_Identifier) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetNode() == nil {
		err := StreamEventsRequest_IdentifierValidationError{
			field:  "Node",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StreamEventsRequest_IdentifierValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StreamEventsRequest_IdentifierValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamEventsRequest_IdentifierValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StreamEventsRequest_IdentifierMultiError(errors)
	}
	return nil
}

// StreamEventsRequest_IdentifierMultiError is an error wrapping multiple
// validation errors returned by StreamEventsRequest_Identifier.ValidateAll()
// if the designated constraints aren't met.
type StreamEventsRequest_IdentifierMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamEventsRequest_IdentifierMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamEventsRequest_IdentifierMultiError) AllErrors() []error { return m }

// StreamEventsRequest_IdentifierValidationError is the validation error
// returned by StreamEventsRequest_Identifier.Validate if the designated
// constraints aren't met.
type StreamEventsRequest_IdentifierValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamEventsRequest_IdentifierValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamEventsRequest_IdentifierValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamEventsRequest_IdentifierValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamEventsRequest_IdentifierValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamEventsRequest_IdentifierValidationError) ErrorName() string {
	return "StreamEventsRequest_IdentifierValidationError"
}

// Error satisfies the builtin error interface
func (e StreamEventsRequest_IdentifierValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamEventsRequest_Identifier.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamEventsRequest_IdentifierValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamEventsRequest_IdentifierValidationError{}
