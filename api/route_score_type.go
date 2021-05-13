package main

import (
    "fmt"
    "net/url"
    sc "strconv"
    str "strings"
    "github.com/Plankiton/SexPistol"
)

func GetScoreType(r Sex.Request) (Sex.Json, int) {
    u := ScoreType {}
    if db.First(&u, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("ScoreType not found")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    return Sex.Bullet {
        Type: "Success",
        Data: u,
    }, 200
}

func CreateScoreType(r Sex.Request) (Sex.Json, int) {
    token := Token {}

    auth := r.Header.Get("Authorization")
    if auth != "" {
        db.First(&token, "id = ?", auth)
        if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
            msg := "Authentication fail, your user not exists or dont have permissions to acess this"
            Sex.Err(msg)
            return Sex.Bullet {
                Message: msg,
                Type:    "Error",
            }, 405
        }
    }

    var data map[string]interface{}
    if r.JsonBody(&data) != nil {
        msg := fmt.Sprint("ScoreType create fail, data need to be a object")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }


    if _, e := data["name"]; !e {
        msg := "ScoreType create fail, Obrigatory field \"name\""
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data["name"] = str.ToLower(data["name"].(string))
    if db.First(&ScoreType {}, "id = ?", data["name"]).Error == nil {
        msg := fmt.Sprint("ScoreType create fail, score_type already registered")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    score_type := ScoreType {}

    Sex.Copy(data, &score_type)
    db.Add(&score_type)

    return Sex.Bullet {
        Type: "Sucess",
        Data: score_type,
    }, 200
}

func UpdateScoreType(r Sex.Request) (Sex.Json, int) {
    token := Token {}

    auth := r.Header.Get("Authorization")
    if auth != "" {
        db.First(&token, "id = ?", auth)
        if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
            msg := "Authentication fail, your user not exists or dont have permissions to acess this"
            Sex.Err(msg)
            return Sex.Bullet {
                Message: msg,
                Type:    "Error",
            }, 405
        }
    }

    var data map[string]interface{}
    if r.JsonBody(&data) != nil {
        msg := fmt.Sprint("ScoreType create fail, data need to be a object")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }


    score_type := ScoreType{}
    if db.First(&score_type, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("ScoreType update fail, score_type not found")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    Sex.Copy(data, &score_type)
    db.Sav(&score_type)

    return Sex.Bullet {
        Type: "Sucess",
        Data: score_type,
    }, 200
}

func DeleteScoreType(r Sex.Request) (Sex.Json, int) {
    token := Token {}

    auth := r.Header.Get("Authorization")
    if auth != "" {
        db.First(&token, "id = ?", auth)
        if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
            msg := "Authentication fail, your user not exists or dont have permissions to acess this"
            Sex.Err(msg)
            return Sex.Bullet {
                Message: msg,
                Type:    "Error",
            }, 405
        }
    }

    score_type := ScoreType{}
    if db.First(&score_type, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("ScoreType delete fail, score_type not found")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    if db.Del(&score_type) == nil {
        return Sex.Bullet {
            Type: "Sucess",
            Message: "ScoreType deleted",
        }, 200
    }

    msg := fmt.Sprint("ScoreType delete fail")
    Sex.Err(msg)
    return Sex.Bullet {
        Message: msg,
        Type:    "Error",
    }, 500
}

func GetScoreTypeList(r Sex.Request) (Sex.Json, int) {
    limit, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("p"))
    r.Conf["query"].(url.Values).Del("l")
    r.Conf["query"].(url.Values).Del("p")

    score_list := []ScoreType{}
    offset := (page - 1) * limit
    e := db.Offset(offset).Limit(limit).Order("created_at desc, updated_at, id").Find(&score_list)

    if e.Error != nil {
        return Sex.Bullet {
            Type: "Error",
            Message: "Error on creating of profile on database",
        }, 500
    }

    return Sex.Bullet {
        Type: "Sucess",
        Data: score_list,
    }, 200
}
