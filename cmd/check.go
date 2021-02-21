package main

import (
    "encoding/json"
    "fmt"

    "github.com/ema-corp/go-validation-sample/pkg/models"
    "github.com/ema-corp/go-validation-sample/pkg/validator"
)

const userStruct = `User {
     Name   string  validate:"required"
     IconID int  validate:"is_icon_id,required"
}`

const userJsonStruct = `UserJson {
     Name string json:"name" validate:"required"
}`

func main() {
    v := validator.CustomNew()

    fmt.Printf("\n### start User validation ###\n%v\n", userStruct)

    correctUser := models.User {
        Name: "ok",
        IconID: 2,
    }
    if ok := v.CheckUser(correctUser); !ok {
        fmt.Printf("check failed: %v\n", correctUser)
    } else {
        fmt.Printf("check success: %v\n", correctUser)
    }

    incorrectUser := models.User {
        Name: "bad",
        IconID: 100,
    }
    if ok := v.CheckUser(incorrectUser); !ok {
        fmt.Printf("check failed: %v\n", incorrectUser)
    } else {
        fmt.Printf("check success: %v\n", incorrectUser)
    }

    fmt.Printf("\n### start UserJson validation ###\n%v\n", userJsonStruct)
    correctUserJson := models.UserJson{}
    correctOrigin := `{"name": "ok"}`
    if err := json.Unmarshal([]byte(correctOrigin), &correctUserJson); err != nil {
        fmt.Println(err)
    }
    if ok := v.CheckUserJson(correctUserJson); !ok {
        fmt.Printf("check failed: %v\n", correctOrigin)
    } else {
        fmt.Printf("check success: %v\n", correctOrigin)
    }

    incorrectUserJson := models.UserJson{}
    incorrectOrigin := `{"userName": "bad"}`
    if err := json.Unmarshal([]byte(incorrectOrigin), &incorrectUserJson); err != nil {
        fmt.Println(err)
    }
    if ok := v.CheckUserJson(incorrectUserJson); !ok {
        fmt.Printf("check failed: %v\n", incorrectOrigin)
    } else {
        fmt.Printf("check success: %v\n", incorrectOrigin)
    }
}
