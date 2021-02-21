package validator

import (
	_ "fmt"

	"github.com/ema-corp/go-validation-sample/pkg/models"
	"github.com/go-playground/validator"
)

type Custom struct {
    validator *validator.Validate
}

func CustomNew() *Custom {
    v := validator.New()
    v.RegisterValidation("is_icon_id", isIconID)

    c := &Custom{v}
    return c
}

func isIconID(fl validator.FieldLevel) bool {
    num := fl.Field().Int()
    if num < 1 || 6 < num {
        return false
    }

    return true
}

func (c *Custom) CheckUser(user models.User) bool {
    err := c.validator.Struct(user)
    if err != nil {
        // fmt.Println(err)
        return false
    }

    return true
}

func (c *Custom) CheckUserJson(userJson models.UserJson) bool {
    err := c.validator.Struct(userJson)
    if err != nil {
        // fmt.Println(err)
        return false
    }

    return true
}
