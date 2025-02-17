// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/grpc_credential/v2alpha/file_based_metadata.proto

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

// Validate checks the field values on FileBasedMetadataConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FileBasedMetadataConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FileBasedMetadataConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FileBasedMetadataConfigMultiError, or nil if none found.
func (m *FileBasedMetadataConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *FileBasedMetadataConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSecretData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FileBasedMetadataConfigValidationError{
					field:  "SecretData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FileBasedMetadataConfigValidationError{
					field:  "SecretData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSecretData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FileBasedMetadataConfigValidationError{
				field:  "SecretData",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for HeaderKey

	// no validation rules for HeaderPrefix

	if len(errors) > 0 {
		return FileBasedMetadataConfigMultiError(errors)
	}
	return nil
}

// FileBasedMetadataConfigMultiError is an error wrapping multiple validation
// errors returned by FileBasedMetadataConfig.ValidateAll() if the designated
// constraints aren't met.
type FileBasedMetadataConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FileBasedMetadataConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FileBasedMetadataConfigMultiError) AllErrors() []error { return m }

// FileBasedMetadataConfigValidationError is the validation error returned by
// FileBasedMetadataConfig.Validate if the designated constraints aren't met.
type FileBasedMetadataConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FileBasedMetadataConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FileBasedMetadataConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FileBasedMetadataConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FileBasedMetadataConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FileBasedMetadataConfigValidationError) ErrorName() string {
	return "FileBasedMetadataConfigValidationError"
}

// Error satisfies the builtin error interface
func (e FileBasedMetadataConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFileBasedMetadataConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FileBasedMetadataConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FileBasedMetadataConfigValidationError{}
