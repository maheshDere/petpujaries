package repository

import (
	"context"
	"petpujaris/config"
	"petpujaris/logger"
	"petpujaris/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"
)

var ctx context.Context
var db *sqlx.DB
var dbRegistry DatabaseRegistry
var pgClient PgClient

func init() {
	var err error
	logger.Setup()
	config.SetupConfig()
	config.LoadConfig()
	ctx = context.Background()

	dbconfig := config.GetDBConfig()
	db, err = GetDBConnection(dbconfig)
	if err != nil {
		panic(err)
	}
	pgClient, err := NewPgClient(dbconfig)
	if err != nil {
		panic(err)
	}

	dbRegistry = NewDBRegistry(pgClient)
}

func TestCreateUser(t *testing.T) {
	t.Run("when create user with valid parameters", func(t *testing.T) {
		time := time.Now()
		user := models.User{
			Name:             "test user 12",
			Email:            "adminwe12@barbqeacoaq.com",
			MobileNumber:     "919096317275",
			IsActive:         true,
			Password:         "password1",
			RoleID:           1,
			ResourceableID:   1,
			ResourceableType: "Company",
			CreatedAt:        time,
			UpdatedAt:        time,
		}
		t.Run("It shoud return error equal to nil", func(t *testing.T) {
			err := dbRegistry.CreateUser(ctx, user)
			assert.NoError(t, err)
		})
	})
}
