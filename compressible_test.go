package compressible

import (
	"testing"

	"github.com/GitbookIO/mimedb"
	"github.com/stretchr/testify/assert"
)

func TestDefaultCompressible(t *testing.T) {
	Load()

	t.Run("All mimedb compressible types reflect in compressible-go", func(t *testing.T) {
		assert := assert.New(t)

		for _, mimeEntry := range mimedb.DB {
			assert.Equal(mimeEntry.Compressible, Is(mimeEntry.ContentType))
		}
	})

	t.Run("Invalid types should return false", func(t *testing.T) {
		assert := assert.New(t)

		assert.False(Is("foo/bar"))
	})

	t.Run("Types have the specified schemes should be compressible", func(t *testing.T) {
		assert := assert.New(t)

		types := [...]string{
			"Text/foobar",
			"foo/bar+jSOn",
			"foo/bar+text",
			"foo/bar+XML",
		}

		for _, t := range types {
			assert.True(Is(t))
		}
	})

	t.Run("should not be compressible if contentLength is smaller than transhold", func(t *testing.T) {
		assert := assert.New(t)
		var wt WithTrashold = 1024

		assert.False(wt.Compressible("text/html", int(wt)-1))
		assert.True(wt.Compressible("text/html", int(wt)))
	})
}
