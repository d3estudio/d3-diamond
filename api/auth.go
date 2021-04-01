package main
import (
    SexDB "github.com/Plankiton/SexPistol/Database"

    "fmt"
    "time"
)

type Token struct {
    SexDB.Token
}

func (token *Token) GetUser() (User, bool) {
    ok := false
    user := User{}
    if db.First(token, "id = ?", token.ID).Error == nil &&
    db.First(&user, "id = ?", token.UserId).Error == nil {
        ok = true
    }

    return user, ok
}

func (model *Token) New() error {
    user := User{}
    err := db.First(&user, "id = ?", model.UserId).Error
    if err == nil {
        var order int64
        db.Find(model).Count(&order)

        model.UserId = user.ID
        model.ID = SexDB.ToHash(fmt.Sprintf(
            "%d;%d;%s;%s;%s", order, user.ID, user.Name, user.Email, time.Now().String(),
        ))

        return nil
    }

    return err
}

