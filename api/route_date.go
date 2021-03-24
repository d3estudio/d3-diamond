package main

import (
    "fmt"
    "github.com/Plankiton/SexPistol"
    str "strconv"
)

func CreateDate(r sex.Request) (sex.Response, int) {
    token := Token{}
    token.ID = r.Token
    if _, ok := (token).GetUser();!ok {
        msg := "Authentication fail, your user not exists"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    score := ScoreDate{}
    dt_begin, _ := dateRange("")
    if db.First(&score, "date = ? AND user_id = ?",
    dt_begin, r.PathVars["id"]).Error == nil {
        msg := "This score date already exists"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    UserId, _ := str.Atoi(r.PathVars["id"])
    score.UserId = uint(UserId)

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

func GetDates(r sex.Request) (sex.Response, int) {
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

    score_list := []ScoreDate{}

    query := r.Reader.Header.Get("Query")
    if db.Table("score_dates").
    Where(query).
    Find(&score_list, "user_id = ?",
    r.PathVars["id"]).
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

func GetDate(r sex.Request) (sex.Response, int) {
    score := ScoreDate {}
    if db.First(&score, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("ScoreDate not found")
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

func UpdateDate(r sex.Request) (sex.Response, int) {
    score := ScoreDate {}
    if db.First(&score, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("ScoreDate not found")
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
        Message: "ScoreDate Updated",
    }, 200
}

func DeleteDate(r sex.Request) (sex.Response, int) {
    score := ScoreDate {}
    if db.First(&score, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("ScoreDate not found")
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
        Message: "ScoreDate deleted",
    }, 200
}
