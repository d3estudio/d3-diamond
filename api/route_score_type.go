package main

import (
    "fmt"
    "net/url"
    sc "strconv"
    str "strings"
    "github.com/Plankiton/SexPistol"
)

func GetScoreType(r sex.Request) (sex.Response, int) {
    u := ScoreType {}
    if db.First(&u, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("ScoreType not found")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    return sex.Response {
        Type: "Success",
        Data: u,
    }, 200
}

func CreateScoreType(r sex.Request) (sex.Response, int) {
    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
        msg := "Authentication fail, your user not exists or dont have permissions to acess this"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    if !sex.ValidateData(r.Data, sex.GenericJsonObj) {
        msg := fmt.Sprint("ScoreType create fail, data need to be a object")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    if _, e := data["name"]; !e {
        msg := "ScoreType create fail, Obrigatory field \"name\""
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data["name"] = str.ToLower(data["name"].(string))
    if db.First(&ScoreType {}, "id = ?", data["name"]).Error == nil {
        msg := fmt.Sprint("ScoreType create fail, score_type already registered")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    score_type := ScoreType {}

    sex.MapTo(data, &score_type)
    score_type.Create()

    return sex.Response {
        Type: "Sucess",
        Data: score_type,
    }, 200
}

func UpdateScoreType(r sex.Request) (sex.Response, int) {
    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
        msg := "Authentication fail, your user not exists or dont have permissions to acess this"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    if !sex.ValidateData(r.Data, sex.GenericJsonObj) {
        msg := fmt.Sprint("ScoreType create fail, data need to be a object")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    score_type := ScoreType{}
    if db.First(&score_type, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("ScoreType update fail, score_type not found")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    sex.MapTo(data, &score_type)
    score_type.Save()

    return sex.Response {
        Type: "Sucess",
        Data: score_type,
    }, 200
}

func DeleteScoreType(r sex.Request) (sex.Response, int) {
    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
        msg := "Authentication fail, your user not exists or dont have permissions to acess this"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    score_type := ScoreType{}
    if db.First(&score_type, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("ScoreType delete fail, score_type not found")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    if score_type.Delete() {
        return sex.Response {
            Type: "Sucess",
            Message: "ScoreType deleted",
        }, 200
    }

    msg := fmt.Sprint("ScoreType delete fail")
    sex.Err(msg)
    return sex.Response {
        Message: msg,
        Type:    "Error",
    }, 500
}

func GetScoreTypeList(r sex.Request) (sex.Response, int) {
    limit, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("p"))
    r.Conf["query"].(url.Values).Del("l")
    r.Conf["query"].(url.Values).Del("p")

    query := r.Conf["query"].(url.Values).Encode()
    query = str.ReplaceAll(query, "&", " AND ")
    query = str.ReplaceAll(query, "|", " OR ")

    score_list := []ScoreType{}
    offset := (page - 1) * limit
    e := db.Offset(offset).Limit(limit).Order("created_at desc, updated_at, id").Find(&score_list, query)

    if e.Error != nil {
        return sex.Response{
            Type: "Error",
            Message: "Error on creating of profile on database",
        }, 500
    }

    return sex.Response{
        Type: "Sucess",
        Data: score_list,
    }, 200
}
