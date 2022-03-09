package got

import (
	"errors"
	"testing"

	"github.com/golang-must/must"
)

func TestOk(t *testing.T) {
	t.Run("value present", func(t *testing.T) {
		must := must.New(t)

		result := Ok(1)
		must.Equal(result.Data(), 1)
		must.Nil(result.Err())
	})
}

func FuzzOk(f *testing.F) {
	f.Add("abcde")
	f.Fuzz(func(t *testing.T, value string) {
		must := must.New(t)
		result := Ok(value)

		must.Equal(result.Data(), value)
	})
}

func TestErr(t *testing.T) {
	t.Run("error present", func(t *testing.T) {
		must := must.New(t)

		result := Err(errors.New("don't cry :)"))
		must.Nil(result.data)
		must.NotNil(result.err)
	})
}

func FuzzErr(f *testing.F) {
	f.Add("abcde")
	f.Fuzz(func(t *testing.T, err string) {
		must := must.New(t)
		result := Err(errors.New(err))

		must.Nil(result.data)
		must.NotNil(result.err)
		must.Equal(result.err.Error(), err)
	})
}
