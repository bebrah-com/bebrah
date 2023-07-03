package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestInitDB(t *testing.T) {
// 	// InitDb("/root/code/okjiang/bebrah/backend")

// 	gomysql.RegisterTLSConfig("tidb", &tls.Config{
// 		MinVersion: tls.VersionTLS12,
// 		ServerName: "gateway01.eu-central-1.prod.aws.tidbcloud.com",
// 	})

// 	dsn := "2FRug3riV65YHEr.root:ptxaTmXLzPgG5wjF@tcp(gateway01.eu-central-1.prod.aws.tidbcloud.com:4000)/test?tls=tidb"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	require.NoError(t, err)

// 	db.AutoMigrate(&Test{})

// 	db.Create(&Test{ID: 1})
// 	// Db().Get("")
// }

func TestGetEnv(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)

	env := getEnv(dir)
	require.Equal(t, "root", env.DbUser)
	require.Equal(t, "xxxxxx", env.DbPassword)
	require.Equal(t, "host", env.DbHost)
	require.Equal(t, "4000", env.DbPort)
}
