package main
import (
    "github.com/Plankiton/SexPistol"
)

func LogIn(r Sex.Request) (Sex.Json, int) {
    var data map[string]interface{}
    if r.JsonBody(&data) != nil {
        msg := Sex.Fmt("Authentication fail, bad request, data need to be a object")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    user := User {}

    if db.First(&user, "email = ?", data["email"]).Error != nil {
        msg := Sex.Fmt("Authentication fail, no users with \"%s\" registered", data["email"])
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 404
    }
    if _, ok := data["pass"];ok {
        if user.CheckPass(data["pass"].(string)) {
            token := Token {}
            token.UserId = user.ID
            db.Add(&token)

            return Sex.Bullet {
                Type: "Sucess",
                Data: map[string]interface{} {
                    "token": token.ID,
                    "user": user,
                },
            }, 200
        }
    }

    msg := Sex.Fmt("Authentication fail, password from <%s:%d> is wrong, permission denied", user.Name, user.ID)
    return Sex.Bullet {
        Message: msg,
        Type:    "Error",
    }, 403
}

func Verify(r Sex.Request) (Sex.Json, int) {
    token := Token {}

    auth := r.Header.Get("Authorization")
    if auth != "" {
        db.First(&token, "id = ?", auth)
        if user, ok := token.GetUser(); ok {
            msg := Sex.Fmt("Token \"%s\" Is valid", auth)
            Sex.Log(msg)
            return Sex.Bullet {
                Type: "Sucess",
                Data: user,
            }, 200

        }
    } else {
        msg := Sex.Fmt("Authentication fail, this route need a token")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    msg := Sex.Fmt("Authentication fail, user not found, permission denied")
    Sex.Err(msg)
    return Sex.Bullet {
        Message: msg,
        Type:    "Error",
    }, 403
}

func LogOut(r Sex.Request) (Sex.Json, int) {
    token := Token {}
    auth := r.Header.Get("Authorization")

    db.First(&token, "id = ?", auth)
    if db.Delete(&token) == nil {
        msg := Sex.Fmt("Token \"%s\" removed", auth)
        Sex.Log(msg)
        return Sex.Bullet {
            Type: "Sucess",
            Message: msg,
        }, 200
    }

    msg := Sex.Fmt("Token \"%s\" can't be removed", auth)
    Sex.Log(msg)
    return Sex.Bullet {
        Type: "Sucess",
        Message: msg,
    }, 500
}
