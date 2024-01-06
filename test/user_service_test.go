package test

//Please run tests one by one
import (
	"os"
	"testing"
	"user-management/config"
	"user-management/database"
	"user-management/models"
	"user-management/service"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var userService service.IUserService
var db *gorm.DB

func TestMain(m *testing.M) {
	db = database.ConnectDB(config.TestDB)
	userService = service.NewUserService(db)
	exitCode := m.Run()
	os.Exit(exitCode)

}

//TestData initialization

func TestDataInitialize(t *testing.T) {
	db.Create(&models.User{Name: "Test User 1", Age: 20})
	db.Create(&models.User{Name: "Test User 2", Age: 30})
	db.Create(&models.User{Name: "Test User 3", Age: 40})
}

// Delete all users from database

func TestDeleteAllUsers(t *testing.T) {
	var users []models.User
	// Delete all users from database if exists
	db.Find(&users)
	if len(users) > 0 {
		db.Delete(&users)
	}
	// Drop users table before running each test
	db.Migrator().DropTable(&models.User{})
}

func setup(t *testing.T) {
	TestDataInitialize(t)
}

func clear(t *testing.T) {
	TestDeleteAllUsers(t)
}

func TestGetAllUsers(t *testing.T) {
	setup(t)

	expectedUsers := []models.User{
		{ID: 1,
			Name: "Test User 1", Age: 20},
		{
			ID:   2,
			Name: "Test User 2", Age: 30},
		{ID: 3,
			Name: "Test User 3", Age: 40},
	}
	t.Run("TestGetAllUsers", func(t *testing.T) {
		actualUsers, err := userService.GetAllUsers()
		if err != nil {
			t.Error("Error while fetching all users: ", err)
		}
		assert.Equal(t, expectedUsers, actualUsers, "All users fetched successfully")
		assert.Equal(t, len(expectedUsers), len(actualUsers), "All users fetched successfully")
	})

	log.Info("TestGetAllUsers success")

	clear(t)
}

func TestGetUserByID(t *testing.T) {
	setup(t)

	expectedUser := models.User{
		ID:   1,
		Name: "Test User 1", Age: 20}
	t.Run("TestGetUserByID", func(t *testing.T) {
		actualUser, err := userService.GetUserByID(1)
		if err != nil {
			t.Error("Error while fetching user by id: ", err)
		}
		assert.Equal(t, expectedUser, actualUser, "User fetched successfully")
	})

	log.Info("TestGetUserByID success")

	clear(t)
}

func TestCreateUser(t *testing.T) {

	setup(t)

	user := models.User{
		ID:   4,
		Name: "Test User 4", Age: 50}
	t.Run("TestCreateUser", func(t *testing.T) {
		err := userService.CreateUser(user)
		if err != nil {
			t.Error("Error while creating user: ", err)
		}
		users, err := userService.GetAllUsers()
		if err != nil {
			t.Error("Error while fetching all users: ", err)
		}
		assert.Equal(t, 4, len(users), "User created successfully")
		assert.Equal(t, user, users[3], "User created successfully")
	})

	log.Info("TestCreateUser success")

	clear(t)

}

func TestUpdateUserByID(t *testing.T) {
	setup(t)

	user := models.User{
		ID:   1,
		Name: "Updated User", Age: 80}
	t.Run("TestUpdateUserByID", func(t *testing.T) {
		err := userService.UpdateUserByID(1, user)
		if err != nil {
			t.Error("Error while updating user: ", err)
		}
		users, err := userService.GetAllUsers()
		if err != nil {
			t.Error("Error while fetching all users: ", err)
		}
		assert.Equal(t, user, users[0], "User updated successfully")
	})
	clear(t)

	log.Info("TestUpdateUserByID success")

}

func TestDeleteByID(t *testing.T) {
	setup(t)

	t.Run("TestDeleteByID", func(t *testing.T) {
		err := userService.DeleteByID(1)
		if err != nil {
			t.Error("Error while deleting user: ", err)
		}
		users, err := userService.GetAllUsers()
		if err != nil {
			t.Error("Error while fetching all users: ", err)
		}
		assert.Equal(t, 2, len(users), "User deleted successfully")
	})

	log.Info("TestDeleteByID success")

	clear(t)
}
