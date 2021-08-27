// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-quiz-api/ocp_quiz_api.proto

package ocp_quiz_api

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

// Validate checks the field values on Quiz with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Quiz) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for ClassroomId

	// no validation rules for UserId

	// no validation rules for Link

	return nil
}

// QuizValidationError is the validation error returned by Quiz.Validate if the
// designated constraints aren't met.
type QuizValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuizValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuizValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuizValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuizValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuizValidationError) ErrorName() string { return "QuizValidationError" }

// Error satisfies the builtin error interface
func (e QuizValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuiz.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuizValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuizValidationError{}

// Validate checks the field values on CreateQuizV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateQuizV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetClassroomId() <= 0 {
		return CreateQuizV1RequestValidationError{
			field:  "ClassroomId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetUserId() <= 0 {
		return CreateQuizV1RequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for Link

	return nil
}

// CreateQuizV1RequestValidationError is the validation error returned by
// CreateQuizV1Request.Validate if the designated constraints aren't met.
type CreateQuizV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateQuizV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateQuizV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateQuizV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateQuizV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateQuizV1RequestValidationError) ErrorName() string {
	return "CreateQuizV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateQuizV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateQuizV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateQuizV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateQuizV1RequestValidationError{}

// Validate checks the field values on CreateQuizV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateQuizV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for QuizId

	return nil
}

// CreateQuizV1ResponseValidationError is the validation error returned by
// CreateQuizV1Response.Validate if the designated constraints aren't met.
type CreateQuizV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateQuizV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateQuizV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateQuizV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateQuizV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateQuizV1ResponseValidationError) ErrorName() string {
	return "CreateQuizV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateQuizV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateQuizV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateQuizV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateQuizV1ResponseValidationError{}

// Validate checks the field values on DescribeQuizV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeQuizV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetQuizId() <= 0 {
		return DescribeQuizV1RequestValidationError{
			field:  "QuizId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeQuizV1RequestValidationError is the validation error returned by
// DescribeQuizV1Request.Validate if the designated constraints aren't met.
type DescribeQuizV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeQuizV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeQuizV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeQuizV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeQuizV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeQuizV1RequestValidationError) ErrorName() string {
	return "DescribeQuizV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeQuizV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeQuizV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeQuizV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeQuizV1RequestValidationError{}

// Validate checks the field values on DescribeQuizV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeQuizV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetQuiz()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeQuizV1ResponseValidationError{
				field:  "Quiz",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeQuizV1ResponseValidationError is the validation error returned by
// DescribeQuizV1Response.Validate if the designated constraints aren't met.
type DescribeQuizV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeQuizV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeQuizV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeQuizV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeQuizV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeQuizV1ResponseValidationError) ErrorName() string {
	return "DescribeQuizV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeQuizV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeQuizV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeQuizV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeQuizV1ResponseValidationError{}

// Validate checks the field values on ListQuizV1Request with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListQuizV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLimit() < 1 {
		return ListQuizV1RequestValidationError{
			field:  "Limit",
			reason: "value must be greater than or equal to 1",
		}
	}

	if m.GetOffset() < 0 {
		return ListQuizV1RequestValidationError{
			field:  "Offset",
			reason: "value must be greater than or equal to 0",
		}
	}

	return nil
}

// ListQuizV1RequestValidationError is the validation error returned by
// ListQuizV1Request.Validate if the designated constraints aren't met.
type ListQuizV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListQuizV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListQuizV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListQuizV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListQuizV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListQuizV1RequestValidationError) ErrorName() string {
	return "ListQuizV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListQuizV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListQuizV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListQuizV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListQuizV1RequestValidationError{}

// Validate checks the field values on ListQuizV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListQuizV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetQuizes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListQuizV1ResponseValidationError{
					field:  fmt.Sprintf("Quizes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for CurrentPage

	return nil
}

// ListQuizV1ResponseValidationError is the validation error returned by
// ListQuizV1Response.Validate if the designated constraints aren't met.
type ListQuizV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListQuizV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListQuizV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListQuizV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListQuizV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListQuizV1ResponseValidationError) ErrorName() string {
	return "ListQuizV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListQuizV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListQuizV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListQuizV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListQuizV1ResponseValidationError{}

// Validate checks the field values on RemoveQuizV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveQuizV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetQuizId() <= 0 {
		return RemoveQuizV1RequestValidationError{
			field:  "QuizId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveQuizV1RequestValidationError is the validation error returned by
// RemoveQuizV1Request.Validate if the designated constraints aren't met.
type RemoveQuizV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveQuizV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveQuizV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveQuizV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveQuizV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveQuizV1RequestValidationError) ErrorName() string {
	return "RemoveQuizV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveQuizV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveQuizV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveQuizV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveQuizV1RequestValidationError{}

// Validate checks the field values on RemoveQuizV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveQuizV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Found

	return nil
}

// RemoveQuizV1ResponseValidationError is the validation error returned by
// RemoveQuizV1Response.Validate if the designated constraints aren't met.
type RemoveQuizV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveQuizV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveQuizV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveQuizV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveQuizV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveQuizV1ResponseValidationError) ErrorName() string {
	return "RemoveQuizV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveQuizV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveQuizV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveQuizV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveQuizV1ResponseValidationError{}
