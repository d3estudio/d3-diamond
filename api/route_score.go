package main

import (
    "fmt"
    "github.com/Plankiton/SexPistol"
    str "strconv"
)

func CreateScore(r sex.Request) (sex.Response, int) {
    if !sex.ValidateData(r.Data, sex.GenericJsonObj) {
        msg := fmt.Sprint("Role create fail, data need to be a object")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    token := Token{}
    token.ID = r.Token
    curr := User{}; {
        ok := false
        if curr, ok = (token).GetUser();!ok {
            msg := "Authentication fail, your user not exists"
            sex.Err(msg)
            return sex.Response {
                Message: msg,
                Type:    "Error",
            }, 405
        } }

    data := r.Data.(map[string]interface{})
    score := Score{}

    dt_begin, dt_end := dateRange("")
    if db.First(&score, "name = ? AND sender_id = ? AND user_id = ? AND created_at BETWEEN ? AND ?",
    data["name"], curr.ID, r.PathVars["id"], dt_begin, dt_end).Error == nil {
        msg := "This score already exists"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    if db.First(&ScoreType{}, "id = ?", data["name"]).Error != nil {
        msg := "Score type invalid"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    sex.MapTo(data, &score)
    UserId, _ := str.Atoi(r.PathVars["id"])
    score.UserId = uint(UserId)
    score.SenderId = curr.ID

    last := 0.0
    if db.Table("scores").

    Select("value").
    Where("user_id = ? AND sender_id = ? AND created_at < ?",
    score.UserId, score.SenderId, dt_begin).
    Last(&last).
    Error == nil {
        score.Diff = score.Value - last
    }

    if !score.Create() {
        msg := "Unknown error ocurred"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    return sex.Response {
            Message: "Score created",
            Type:    "Sucess",
        }, 200
}

func GetScoreList(r sex.Request) (sex.Response, int) {
    user := User{}
    if db.First(&user, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Score listing fail, user not found")
        sex.Err(msg)
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

    score_list := []Score{}

    query := r.Reader.Header.Get("Query")
    dt_begin, dt_end := dateRange(r.PathVars["date"])
    dt_begin = dt_begin.AddDate(0, -1, 0)

    if db.Table("scores").
    Where(query).
    Find(&score_list, "user_id = ? AND created_at BETWEEN ? AND ?",
    r.PathVars["id"], dt_begin, dt_end).
    RowsAffected <= 0 {
        msg := "This user has no avaluations"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    return sex.Response {
        Type: "Sucess",
        Data: score_list,
    }, 200
}

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

    if !score.Save() {
        msg := "Unknown error ocurred"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    return sex.Response {
        Type: "Success",
        Message: "Score Updated",
    }, 200
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

    if !score.Delete() {
        msg := "Unknown error ocurred"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    return sex.Response {
        Type: "Success",
        Message: "Score deleted",
    }, 200
}
