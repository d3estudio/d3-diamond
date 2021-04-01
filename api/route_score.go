package main

import (
    "fmt"
    "github.com/Plankiton/SexPistol"
    str "strconv"
)

func CreateScore(r Sex.Request) (Sex.Json, int) {
    var data map[string]interface{}
    if r.JsonBody(&data) != nil {
        msg := fmt.Sprint("Role create fail, data need to be a object")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    token := Token {}

    auth := r.Header.Get("Authorization")
    curr := User{}; if auth != "" {
        db.First(&token, "id = ?", auth)
        {
            ok := false
            if curr, ok = (token).GetUser();!ok {
                msg := "Authentication fail, your user not exists"
                Sex.Err(msg)
                return Sex.Bullet {
                    Message: msg,
                    Type:    "Error",
                }, 405
            }
        }
    }

    score := Score{}

    dt_begin, dt_end := dateRange("")
    if db.First(&score, "name = ? AND sender_id = ? AND user_id = ? AND created_at BETWEEN ? AND ?",
    data["name"], curr.ID, r.PathVars["id"], dt_begin, dt_end).Error == nil {
        msg := "This score already exists"
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    if db.First(&ScoreType{}, "id = ?", data["name"]).Error != nil {
        msg := "Score type invalid"
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    Sex.Copy(data, &score)
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

func GetScoreList(r Sex.Request) (Sex.Json, int) {
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

    score_list := []Score{}

    query := r.Header.Get("Query")
    dt_begin, dt_end := dateRange(r.PathVars["date"])
    dt_begin = dt_begin.AddDate(0, -1, 0)

    if db.Table("scores").
    Where(query).
    Find(&score_list, "user_id = ? AND created_at BETWEEN ? AND ?",
    r.PathVars["id"], dt_begin, dt_end).
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

func GetScore(r Sex.Request) (Sex.Json, int) {
    score := Score {}
    if db.First(&score, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Score not found")
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

func UpdateScore(r Sex.Request) (Sex.Json, int) {
    score := Score {}
    if db.First(&score, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Score not found")
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
        Message: "Score Updated",
    }, 200
}

func DeleteScore(r Sex.Request) (Sex.Json, int) {
    score := Score {}
    if db.First(&score, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Score not found")
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
        Message: "Score deleted",
    }, 200
}
