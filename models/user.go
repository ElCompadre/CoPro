package models

import (
	"regexp"
	"strings"

	u "github.com/elcompadre/elcompadre/copro/utils"
)

type User struct {
	Base
	Email    string `json:"email" gorm:"type: varchar(200)"`
	Password string `json:"password" gorm:"type: varchar(500)"`
	Token    string `json:"token" gorm:"type: varchar(500)"`
}

func (user *User) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if matched, _ := regexp.MatchString(`[A-Za-z]*\d`, user.Password); user.Password == "" || len(user.Password) < 6 || !matched {
		return u.Message(false, "Password must have at least 6 characters, must contain one upper case, one lower case and one digit."), false
	}

	return u.Message(false, "User validated"), true
}
