package main

import (
    "fmt"
    "github.com/Plankiton/SexPistol"
)

func GetScore(r sex.Request) (sex.Response, int) {
    score := Score {}
    if db.First(&score, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Score not found")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    return sex.Response {
        Type: "Success",
        Data: score,
    }, 200
}

func CreateScore(r sex.Request) (sex.Response, int) {
    if _, e := r.PathVars["id"]; !e {
        msg := fmt.Sprint("Score create fail, where is \"id\" path variable? neede to be \"", "/", r.Conf["path-template"],"\"")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }
    if !sex.ValidateData(r.Data, sex.GenericJsonObj){
        msg := fmt.Sprint("Score create fail, data need to be a object")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    receiver := User{}
    if db.First(&receiver, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Score create fail, user not found")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    curr, ok := (token).GetUser()
    if !ok {
        msg := "Authentication fail, your user not exists"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    if curr.ID == receiver.ID {
        msg := fmt.Sprint("Score create fail, You can't make a self-avaluation")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    if _, e := data["name"]; !e {
        msg := "Score create fail, Obrigatory field \"name\""
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    if db.First(&ScoreType{}, "id = ?", data["name"]).Error != nil {
        msg := fmt.Sprint("Score create fail, Invalid Score Type \"", data["name"],"\"")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    score := Score {}

    sex.MapTo(data, &score)
    aval := curr.CreateAvaluation(receiver)
    score.AvalId = aval.ID

    if db.First(&Score{}, "aval_id = ? AND type_id = ?", aval.ID, data["name"]).Error == nil {
        msg := fmt.Sprint("Score create fail, Score already exists")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    if score.Create() {
        return sex.Response {
            Type: "Sucess",
            Data: score,
        }, 200
    }

    msg := "Unknown error ocurred"
    sex.Err(msg)
    return sex.Response {
        Message: msg,
        Type:    "Error",
    }, 500
}

func UpdateScore(r sex.Request) (sex.Response, int) {
    score := Score {}
    if db.First(&score, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Score not found")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    aval := Avaluation{}
    if db.First(&aval, "id = ?", score.AvalId).Error != nil {
        msg := fmt.Sprint("Avaluation not found")
        sex.Err(msg)
        score.Delete()
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    user := User{}
    if db.First(&user, "id = ?", aval.UserId).Error != nil {
        msg := fmt.Sprint("User not found")
        sex.Err(msg)
        score.Delete()
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    curr, ok := (token).GetUser()
    if !ok {
        msg := "Authentication fail, your user not exists"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    if !CheckPermissions(curr, user) {
        msg := "Authentication fail, you dont have permissions to acess this"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    if curr.ID == user.ID {
        msg := fmt.Sprint("Score create fail, You can't make a self-avaluation")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    if _, e := data["value"]; !e {
        msg := "Score create fail, Obrigatory field \"value\""
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    sex.MapTo(data, &score)
    if score.Save() {
        return sex.Response {
            Type: "Success",
            Message: "Score saved",
        }, 200
    }

    msg := "Unknown error ocurred"
    sex.Err(msg)
    return sex.Response {
        Message: msg,
        Type:    "Error",
    }, 500
}

func DeleteScore(r sex.Request) (sex.Response, int) {
    score := Score {}
    if db.First(&score, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Score not found")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    aval := Avaluation{}
    if db.First(&aval, "id = ?", score.AvalId).Error != nil {
        msg := fmt.Sprint("Avaluation not found")
        sex.Err(msg)
        score.Delete()
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    user := User{}
    if db.First(&user, "id = ?", aval.UserId).Error != nil {
        msg := fmt.Sprint("User not found")
        sex.Err(msg)
        score.Delete()
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, user) {
        msg := "Authentication fail, your user not exists or dont have permissions to acess this"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    if score.Delete() {
        return sex.Response {
            Type: "Success",
            Message: "Score deleted",
        }, 200
    }

    msg := "Unknown error ocurred"
    sex.Err(msg)
    return sex.Response {
        Message: msg,
        Type:    "Error",
    }, 500
}

func GetScoreList(r sex.Request) (sex.Response, int) {
    if _, e := r.PathVars["id"]; !e {
        msg := fmt.Sprint("Score create fail, where is \"id\" path variable? neede to be \"", "/", r.Conf["path-template"],"\"")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }
    if !sex.ValidateData(r.Data, sex.GenericJsonObj){
        msg := fmt.Sprint("Score create fail, data need to be a object")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    receiver := User{}
    if db.First(&receiver, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Score create fail, user not found")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    curr, ok := (token).GetUser()
    if !ok {
        msg := "Authentication fail, your user not exists"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    if curr.ID == receiver.ID {
        msg := fmt.Sprint("Score create fail, You can't make a self-avaluation")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    aval := curr.CreateAvaluation(receiver)

    score_list := []Score{}
    if db.Find(&score_list, "aval_id = ?", aval.ID).Error == nil {
        msg := fmt.Sprint("Score create fail, Score already exists")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    return sex.Response {
        Type: "Sucess",
        Data: score_list,
    }, 200
}
