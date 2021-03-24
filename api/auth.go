package main
import (
    "fmt"
    "time"
    "github.com/Plankiton/SexPistol"
)

type Token struct {
    sex.Token
}

func (token *Token) GetUser() (User, bool) {
    ok := false
    user := User{}
    if (token.Verify()) {
        if db.First(token, "id = ?", token.ID).Error == nil &&
        db.First(&user, "id = ?", token.UserId).Error == nil {
            ok = true
        }
    }

    return user, ok
}

func (model *Token) Create() bool {
    model.ModelType = sex.GetModelType(model)

    user := User{}
    if db.First(&user, "id = ?", model.UserId).Error == nil {
        var order int64
        db.Find(model).Count(&order)

        model.UserId = user.ID
        model.ID = sex.ToHash(fmt.Sprintf(
            "%d;%d;%s;%s;%s", order, user.ID, user.Name, user.Email, time.Now().String(),
        ))

        if sex.ModelCreate(model) == nil {
            ID := model.ID
            ModelType := model.ModelType
            sex.Log("Created", sex.ToLabel(ID, ModelType))
            return true
        }
    }

    return false
}

