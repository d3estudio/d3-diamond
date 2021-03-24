package main
import (
    "github.com/Plankiton/SexPistol"
    "time"
    "fmt"
)

func LogIn(r sex.Request) (sex.Response, int) {
    var data map[string]interface{}
    if sex.ValidateData(r.Data, sex.GenericJsonObj) {
        data = r.Data.(map[string]interface{})
    } else {
        msg := fmt.Sprint("Authentication fail, bad request, data need to be a object")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    user := User {}

    if db.First(&user, "email = ?", data["email"]).Error != nil {
        msg := fmt.Sprint("Authentication fail, no users with \"", data["email"], "\" registered")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }
    if _, ok := data["pass"];ok {
        if user.CheckPass(data["pass"].(string)) {
            token := Token {}
            token.UserId = user.ID
            token.Create()

            return sex.Response {
                Type: "Sucess",
                Data: map[string]interface{} {
                    "token": token.ID,
                    "user": user,
                },
            }, 200
        }
    }

    msg := fmt.Sprint("Authentication fail, password from <", user.Name, ":", user.ID,"> is wrong, permission denied")
    res := sex.Response {
        Message: msg,
        Type:    "Error",
    }
    res.SetCookie("token", r.Token, time.Hour*24*360*10, r)
    return res, 403
}

func Verify(r sex.Request) (sex.Response, int) {
    token := Token {}
    if r.Token != "" {
        db.First(&token, "id = ?", r.Token)
        if user, ok := token.GetUser(); ok {
            msg := fmt.Sprint("Token \"", r.Token, "\" Is valid")
            sex.Log(msg)
            res := sex.Response {
                Type: "Sucess",
                Data: user,
            }

            res.SetCookie("token", r.Token, time.Hour*24*360*10, r)
            return res, 200
        }
    } else {
        msg := fmt.Sprint("Authentication fail, this route need a token")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    msg := fmt.Sprint("Authentication fail, user not found, permission denied")
    sex.Err(msg)
    return sex.Response {
        Message: msg,
        Type:    "Error",
    }, 403
}

func LogOut(r sex.Request) (sex.Response, int) {
    token := Token {}
    token.ID = r.Token

    db.First(&token)
    if token.Delete() {
        msg := fmt.Sprint("Token \"", r.Token, "\" removed")
        sex.Log(msg)
        return sex.Response {
            Type: "Sucess",
            Message: msg,
        }, 200
    }

    msg := fmt.Sprint("Token \"", r.Token, "\" can't be removed")
    sex.Log(msg)
    return sex.Response {
        Type: "Sucess",
        Message: msg,
    }, 500
}
