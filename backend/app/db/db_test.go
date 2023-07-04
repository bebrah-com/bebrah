package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEnv(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)

	env := getEnv(dir)
	require.Equal(t, "root", env.DbUser)
	require.Equal(t, "xxxxxx", env.DbPassword)
	require.Equal(t, "host", env.DbHost)
	require.Equal(t, "4000", env.DbPort)
}
