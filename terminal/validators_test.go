package terminal

import (
	"testing"

	"github.com/nicolasdscp/giwow/internal/exception"
	"github.com/stretchr/testify/require"
)

func TestNotEmpty(t *testing.T) {
	a := ""
	b := "some text"

	require.EqualError(t, NotEmpty("a")(a), exception.StringEmpty("a").Error())
	require.NoError(t, NotEmpty("b")(b))
}
