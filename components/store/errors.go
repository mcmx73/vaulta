package store

import "errors"

var (
	// ErrSectionNotFound is returned when trying to access a section that has
	// not been created yet.
	ErrSectionNotFound = errors.New("section not found")

	// ErrSectionExists is returned when creating a Section that already exists.
	ErrSectionExists = errors.New("section already exists")

	// ErrSectionNameRequired is returned when creating a Section with a blank name.
	ErrSectionNameRequired = errors.New("section name required")

	// ErrKeyRequired is returned when inserting a zero-length key.
	ErrKeyRequired = errors.New("key required")

	// ErrKeyTooLarge is returned when inserting a key that is larger than MaxKeySize.
	ErrKeyTooLarge = errors.New("key too large")

	// ErrValueTooLarge is returned when inserting a value that is larger than MaxValueSize.
	ErrValueTooLarge = errors.New("value too large")

	// ErrIncompatibleValue is returned when trying create or delete a Section
	// on an existing non-Section key or when trying to create or delete a
	// non-Section key on an existing Section key.
	ErrIncompatibleValue = errors.New("incompatible value")

	// ErrSectionNotFound is returned when trying to access a section that has
	// not been created yet.
	ErrKeyNotFound = errors.New("key not found")

	ErrPtrNeeded = errors.New("pointer needed")

	ErrEmtpySection = errors.New("empty section")

	ErrNotIndexed = errors.New("section index error")

	ErrInvalidIndex = errors.New("index error created")
)

