// Package compressible provides Compressible Content-Type / mime checking for Go.
package compressible

import (
	"mime"
	"regexp"

	"github.com/GitbookIO/mimedb"
)

var compressibleTypeRegExp = regexp.MustCompile(`(?i)^text\/|\+json$|\+text$|\+xml$`)

// Is checks the response Content-Type to determine whether to compress.
// Using mime database https://github.com/GitbookIO/mimedb to find
// which Content-Type is compressible. All types that not in mimedb but
// have the scheme of "text/*", "*/*+json", "*/*+text", "*/*+xml" are
// considered as compressible.
func Is(contentType string) bool {
	dbMatched := false

	exts, err := mime.ExtensionsByType(contentType)

	if err != nil {
		return false
	}

	for _, ext := range exts {
		// all exts returned by mime.ExtensionsByType are always start with ".".
		if entry, ok := mimedb.DB[ext[1:]]; ok {
			dbMatched = true

			if entry.Compressible {
				return true
			}
		}
	}

	if !dbMatched && compressibleTypeRegExp.MatchString(contentType) {
		return true
	}

	return false
}

// WithThreshold is an impelementation with transhold. The transhold defines the // minimun content length to enable compressible check.
type WithThreshold int

// Compressible checks the response Content-Type to determine whether to
// compress. Using mime database https://github.com/GitbookIO/mimedb to find
// which Content-Type is compressible and WithThreshold as content length
// transhold All types that not in mimedb but have the scheme of "text/*",
// "*/*+json", "*/*+text", "*/*+xml" are considered as compressible.
func (wt WithThreshold) Compressible(contentType string, contentLength int) bool {
	if contentLength != 0 && wt > 0 && contentLength < int(wt) {
		return false
	}

	return Is(contentType)
}

// Load loads all extensions and their content types of
// https://github.com/GitbookIO/mimedb to offical "mime" package.
// Recommond to apply this function in your main package's init function.
func Load() error {
	for ext, mimeEntry := range mimedb.DB {
		if err := mime.AddExtensionType("."+ext, mimeEntry.ContentType); err != nil {
			return err
		}
	}

	return nil
}
