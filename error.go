package main

import "github.com/pkg/errors"

var (
	ErrStartDateBeforeEndDate = errors.New("start date must be earlier than end date")
)
