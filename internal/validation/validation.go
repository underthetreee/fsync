package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	fs "github.com/underthetreee/fsync/pkg/proto"
)

func ValidateFile(file *fs.File) error {
	return validation.ValidateStruct(file,
		validation.Field(&file.Filename, validation.Required),
		validation.Field(&file.Content, validation.Required),
	)
}
