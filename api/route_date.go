package main

import (
    "fmt"
    "github.com/Plankiton/SexPistol"
    str "strconv"
)

func CreateDate(r Sex.Request) (Sex.Json, int) {
    token := Token {}

    auth := r.Header.Get("Authorization")
    if auth != "" {
        db.First(&token, "id = ?", auth)
        if _, ok := (token).GetUser();!ok {
            msg := "Authentication fail, your user not exists"
            Sex.Err(msg)
            return Sex.Bullet {
                Message: msg,
                Type:    "Error",
            }, 405
        }
    }

    score := ScoreDate{}
    dt_begin, _ := dateRange("")
    if db.First(&score, "date = ? AND user_id = ?",
    dt_begin, r.PathVars["id"]).Error == nil {
        msg := "This score date already exists"
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    UserId, _ := str.Atoi(r.PathVars["id"])
    score.UserId = uint(UserId)
    score.Date = dt_begin

    if db.Create(&score) != nil {
        msg := "Unknown error ocurred"
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    return Sex.Bullet {
        Message: "Score created",
        Type:    "Sucess",
    }, 200
}

func GetDates(r Sex.Request) (Sex.Json, int) {
    user := User{}
    if db.First(&user, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Score listing fail, user not found")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    token := Token {}

    auth := r.Header.Get("Authorization")
    if auth != "" {
        db.First(&token, "id = ?", auth)
        if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, user) {
            msg := "Authentication fail, your user not exists or dont have permissions to acess this"
            Sex.Err(msg)
            return Sex.Bullet {
                Message: msg,
                Type:    "Error",
            }, 405
        }
    }

    score_list := []ScoreDate{}

    query := r.Header.Get("Query")
    if db.Table("score_dates").
    Where(query).
    Find(&score_list, "user_id = ?",
    r.PathVars["id"]).
    RowsAffected <= 0 {
        msg := "This user has no avaluations"
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    return Sex.Bullet {
        Type: "Sucess",
        Data: score_list,
    }, 200
}

func GetDate(r Sex.Request) (Sex.Json, int) {
    score := ScoreDate {}
    if db.First(&score, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("ScoreDate not found")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    return Sex.Bullet {
        Type: "Success",
        Data: score,
    }, 200
}

func UpdateDate(r Sex.Request) (Sex.Json, int) {
    score := ScoreDate {}
    if db.First(&score, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("ScoreDate not found")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    if db.Save(&score) != nil {
        msg := "Unknown error ocurred"
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    return Sex.Bullet {
        Type: "Success",
        Message: "ScoreDate Updated",
    }, 200
}

func DeleteDate(r Sex.Request) (Sex.Json, int) {
    score := ScoreDate {}
    if db.First(&score, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("ScoreDate not found")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    if db.Delete(&score) != nil {
        msg := "Unknown error ocurred"
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    return Sex.Bullet {
        Type: "Success",
        Message: "ScoreDate deleted",
    }, 200
}
