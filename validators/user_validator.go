package validators

import (
    "banking-app/models"
    "github.com/go-errors/errors"
    "regexp"
)

const UserEmailPattern string = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

type userValidator struct {
}

func NewUserValidator() *userValidator {
    return &userValidator{}
}

func (this *userValidator) Validate(user *models.User) (err error) {
    // todo: use this method as orchestrator to validate
    regx := regexp.MustCompile(UserEmailPattern)
    validEmail := regx.MatchString(user.Email())
    if !validEmail {
        return errors.New("invalid email")
    }
    return nil

}
