package compressible

import (
	"errors"
	"mime"
	"regexp"

	"github.com/GitbookIO/mimedb"
)

func init() {
	// ensure all extensions and their associated content type of "mimedb" package
	// are stored in "mime" package, ignore the potential returned error of
	// mime.AddExtensionType.
	for ext, mimeEntry := range mimedb.DB {
		mime.AddExtensionType("."+ext, mimeEntry.ContentType)
	}
}

var compressibleTypeRegExp = regexp.MustCompile(`^text\/|\+json$|\+text$|\+xml$`)

// Compress is the interface that wraps basic Compressible method.
type Compress interface {
	// Compressible checks the response Content-Type and Content-Length to
	// determine whether to compress.
	// Recommend to use mime database https://github.com/GitbookIO/mimedb to find
	// which Content-Type is compressible.
	// `length == 0` means response body maybe stream, or will be writed later.
	Compressible(contentType string, contentLength int) bool
}

// Default is a convenient impelementation of interface Compress, using
// https://github.com/GitbookIO/mimedb as mime database.
type Default struct {
	trashold int
}

// SetTrashold set the minimun content length to enable compressible check.
// If you want to skip the trashold check, just set it to 0 (the default value).
// An error will be returned if t is < 0.
func (d *Default) SetTrashold(t int) error {
	if t < 0 {
		return errors.New("compressible-go: trashold should >= 0")
	}

	d.trashold = t
	return nil
}

// Compressible checks whether the given contentType / contentLength pair is
// compressible, using https://github.com/GitbookIO/mimedb as mime database and
// c.Trashold as content length transhold. All types that not in mimedb but
// have the scheme of "text/*", "*/*+json", "*/*+text", "*/*+xml" are
// considered as compressible.
func (d Default) Compressible(contentType string, contentLength int) bool {
	dbMatched := false

	if d.trashold != 0 && contentLength != 0 && contentLength < d.trashold {
		return false
	}

	exts, err := mime.ExtensionsByType(contentType)

	if err != nil {
		return false
	}

	for _, ext := range exts {
		// all exts returned by mime.ExtensionsByType are always
		// start with "."
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
