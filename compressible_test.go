package compressible

import (
	"testing"

	"github.com/GitbookIO/mimedb"
	"github.com/stretchr/testify/assert"
)

var dc = Default{}

func TestDefaultCompressible(t *testing.T) {
	t.Run("Should return error when set transhold < 0", func(t *testing.T) {
		assert := assert.New(t)

		err := dc.SetTrashold(-1)
		assert.Equal(err.Error(), "compressible-go: trashold should >= 0")
	})

	t.Run("Should set transhold >= 0 successfully", func(t *testing.T) {
		assert := assert.New(t)

		err := dc.SetTrashold(1024)
		assert.Nil(err)
	})

	t.Run("All mimedb compressible types reflect in compressible-go", func(t *testing.T) {
		assert := assert.New(t)

		for _, mimeEntry := range mimedb.DB {
			assert.Equal(mimeEntry.Compressible,
				dc.Compressible(mimeEntry.ContentType, dc.trashold))
		}
	})

	t.Run("Invalid types should return false", func(t *testing.T) {
		assert := assert.New(t)

		assert.False(dc.Compressible("foo/bar", dc.trashold))
	})

	t.Run("Types have the specified schemes should be compressible", func(t *testing.T) {
		assert := assert.New(t)

		types := [...]string{
			"text/foobar",
			"foo/bar+json",
			"foo/bar+text",
			"foo/bar+xml",
		}

		for _, t := range types {
			assert.True(dc.Compressible(t, dc.trashold))
		}
	})

	t.Run("should not be compressible if contentLength is smaller than transhold", func(t *testing.T) {
		assert := assert.New(t)

		err := dc.SetTrashold(1024)
		assert.Nil(err)

		assert.False(dc.Compressible("text/html", dc.trashold-1))
		assert.True(dc.Compressible("text/html", dc.trashold))
	})
}
