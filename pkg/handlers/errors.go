package handlers

import "errors"

var ErrPermissionDenied = errors.New("don't have access to permit the user with the above permission")
