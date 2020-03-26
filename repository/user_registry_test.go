package repository

import (
	"context"
	"petpujaris/config"
	"petpujaris/logger"
	"petpujaris/models"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

var ctx context.Context
var db *sqlx.DB
var ur UserRegistry
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
	pgClient, err = NewPgClient(dbconfig)
	if err != nil {
		panic(err)
	}

	ur = NewUserRegistry(pgClient)
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
			userID, err := ur.Save(ctx, user)
			assert.NotEmpty(t, userID)
			assert.NoError(t, err)
			deleteUser(int(userID))
		})

	})
}

func TestGetResourceableID(t *testing.T) {
	t.Run("when invalid id pass to method", func(t *testing.T) {
		userID := uint64(0)
		t.Run("It shoud return error", func(t *testing.T) {
			_, err := ur.GetResourceableID(ctx, userID)
			assert.Error(t, err)
		})
	})
	t.Run("when valid id pass to method", func(t *testing.T) {
		id, err := createTestAdmin()
		assert.NoError(t, err)
		userID := uint64(id)
		expectedResourceableID := uint64(1)
		t.Run("It shoud return resourceable id", func(t *testing.T) {
			resourceableID, err := ur.GetResourceableID(ctx, userID)
			assert.NoError(t, err)
			assert.Equal(t, expectedResourceableID, resourceableID)
		})
		deleteUser(id)
	})
}

func createTestAdmin() (int, error) {
	var id int
	adminUser := models.User{
		Name:             "Test admin",
		Email:            "test@test.com",
		MobileNumber:     "4157819647",
		IsActive:         true,
		Password:         "password",
		RoleID:           3,
		ResourceableID:   1,
		ResourceableType: "Company",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	creatTestAdminUserQuery := Command{
		Table:       "users",
		Description: "INSERT USERS DETAILS",
		Query:       "insert into %s (name, email, mobile_number, is_active, password_digest,  role_id,  resourceable_id,  resourceable_type, created_at, updated_at) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id",
	}

	row := pgClient.QueryRow(context.Background(), creatTestAdminUserQuery, adminUser.Name, adminUser.Email, adminUser.MobileNumber, adminUser.IsActive, adminUser.Password, adminUser.RoleID, adminUser.ResourceableID, adminUser.ResourceableType, adminUser.CreatedAt, adminUser.UpdatedAt)
	err := row.Scan(&id)
	return id, err
}

func deleteUser(ID int) {
	deleteUserQuery := Command{
		Table:       "users",
		Description: "DELETE USERS DETAILS",
		Query:       "delete from %s where id = $1",
	}
	pgClient.Exec(context.Background(), deleteUserQuery, ID)
}
