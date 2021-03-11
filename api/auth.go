package main
import (
    "fmt"
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

func (model *Token) Create() {
    model.ModelType = sex.GetModelType(model)

    user := User{}
    user.ID = model.UserId

    if db.First(&user).Error == nil {
        var order int64
        db.Find(model).Count(&order)

        model.UserId = user.ID
        model.ID = sex.ToHash(fmt.Sprintf(
            "%d;%d;%s;%s;%s", order, user.ID, user.Name, user.Email, user.Phone,
        ))

        if sex.ModelCreate(model) == nil {
            ID := model.ID
            ModelType := model.ModelType
            sex.Log("Created", sex.ToLabel(ID, ModelType))
        }
    }
}

