package datasource

import "errors"

// ErrBlankRegionReceived ...
var ErrBlankRegionReceived error = errors.New("blank origin received to process in trie")
