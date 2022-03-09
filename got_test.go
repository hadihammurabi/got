package got

import (
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

func TestErr(t *testing.T) {
	t.Run("error present", func(t *testing.T) {
		must := must.New(t)

		result := Err("don't cry :)")
		must.Nil(result.data)
		must.NotNil(result.err)
	})
}
