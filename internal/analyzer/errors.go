package analyzer

import "errors"

var (
	ErrFileNotFound = errors.New("fichier introuvable")
	ErrFileNotReadable = errors.New("fichier non accessible")
	ErrParsingFailed = errors.New("erreur de parsing")
)

type FileNotFoundError struct {
	Path string
}

func (e *FileNotFoundError) Error() string {
	return "fichier introuvable: " + e.Path
}

func (e *FileNotFoundError) Unwrap() error {
	return ErrFileNotFound
}

type ParsingError struct {
	Details string
}

func (e *ParsingError) Error() string {
	return "erreur de parsing: " + e.Details
}

func (e *ParsingError) Unwrap() error {
	return ErrParsingFailed
}
