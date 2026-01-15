package repository

import "errors"

var ErrFKConstraint = errors.New("fk constraint failed")
var ErrObjectNotFound = errors.New("object not found")
var ErrUserNotFound = errors.New("user not found")
