package model

import (
	"fmt"
	"html"
	"strings"

	jsoniter "github.com/json-iterator/go"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"api/database"
)

type IPs []IP

func (ips IPs) Format() string {
	s := []string{}
	for _, ip := range ips {
		s = append(s, ip.Address)
	}
	out, _ := jsoniter.Marshal(&s)
	return string(out)
}

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
	Email    string `gorm:"size:255;not null;unique"`
	IsAdmin  bool   `gorm:"default:false"`
	Role     Role   `gorm:"size:255;not null;default:'standard'"`
	IPs      IPs
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (user *User) Update(username, email, password string, role Role) error {
	if user.Username != username {
		err := database.Database.
			Model(&user).
			Update("username", html.EscapeString(strings.TrimSpace(username))).Error
		if err != nil {
			return err
		}
	}

	if email != "" && user.Email != email {
		err := database.Database.
			Model(&user).
			Update("email", strings.TrimSpace(email)).Error
		if err != nil {
			return err
		}
	}

	if password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		err = database.Database.
			Model(&user).
			Update("password", string(passwordHash)).Error

		if err != nil {
			return err
		}
	}

	err := database.Database.
		Model(&user).
		Update("role", role).Error
	if err != nil {
		return err
	}

	return nil
}

func (user *User) AddIP(address string) error {
	_, err := CreateIP(*user, address)
	return err
}

func (user *User) Delete() error {
	err := database.Database.Unscoped().Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *User) RevokeTokens() error {
	return database.Database.
		Unscoped().
		Where("user_id = ?", user.ID).
		Delete(&APIKey{}).Error
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := database.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserByEmail(email string) (User, error) {
	var user User
	err := database.Database.Where("email=?", email).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var users User
	err := database.Database.Where("id=?", id).First(&users).Error
	if err != nil {
		return User{}, err
	}
	return users, nil
}

func ListUsers() ([]User, error) {
	var users []User
	err := database.Database.Preload("IPs").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func ListAdmins() ([]User, error) {
	var users []User
	err := database.Database.Where("is_admin=1").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func SetupInitialAdmin() {
	username := viper.GetString("app.adminusername")
	password := viper.GetString("app.adminpassword")

	if username == "" {
		username = "admin"
	}

	if password == "" {
		password = "admin"
	}

	user, err := FindUserByUsername(username)
	if err != nil {
		panic(err)
	}

	if user.Username == "" {
		user := User{
			Username: username,
			Password: password,
			Role:     RoleAdmin,
		}

		if _, err := user.Save(); err != nil {
			panic(err)
		}
		return
	}

	if user.Role != RoleAdmin {
		if err := user.Update(username, "", "", RoleAdmin); err != nil {
			panic(err)
		}
		return
	}

	fmt.Println("default admin already exists")
}

func MigrateAdmins() {
	fmt.Println("migrating existing admins")

	admins, err := ListAdmins()
	if err != nil {
		panic(err)
	}

	for _, admin := range admins {
		if admin.Role != RoleAdmin {
			admin.Update(admin.Username, "", "", RoleAdmin)
		}
	}
}
