package main

import (
	"fmt"

  "net/url"
  sc "strconv"
  str "strings"

  "github.com/Plankiton/SexPistol"
)

func GetUser(r sex.Request) (sex.Response, int) {
    u := User {}
    if db.First(&u, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("User not found")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, u) {
        msg := "Authentication fail, permission denied"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    return sex.Response {
        Type: "Success",
        Data: u,
    }, 200
}

func CreateUser(r sex.Request) (sex.Response, int) {
    if !sex.ValidateData(r.Data, sex.GenericJsonObj) {
        msg := fmt.Sprint("User create fail, data need to be a object")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    needed := []string{
            "email", "name", "pass",
        }
    if (len(data) < len(needed)){
        msg := "User create fail, Obrigatory field"
        if (len(data)==4) {
            msg += "s"
        }
        msg += " missing: "
        for _, k := range needed {
            if _, exist := data[k]; !exist {
                msg += fmt.Sprintf(`"%s", `, k)
            }
        }
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    if db.First(&User {}, "email = ?", data["email"]).Error == nil {
        msg := fmt.Sprint("User create fail, user already registered")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    user := User {}

    sex.MapTo(data, &user)
    user.SetPass(data["pass"].(string))
    user.Create()

    return sex.Response {
        Type: "Sucess",
        Data: user,
    }, 200
}

func UpdateUser(r sex.Request) (sex.Response, int) {
    if !sex.ValidateData(r.Data, sex.GenericJsonObj) {
        msg := fmt.Sprint("User create fail, data need to be a object")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    user := User{}
    if db.First(&user, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("User update fail, user not found")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, user) {
        msg := "Update fail, permission denied"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }


    sex.MapTo(data, &user)
    if _, e := data["pass"];e {
        user.SetPass(data["pass"].(string))
    }

    if user.Save() {
        return sex.Response {
            Type: "Sucess",
            Data: user,
        }, 200
    }

    return sex.Response {
        Type: "Error",
        Message: "Tryed to update with field already existent",
    }, 500
}

func DeleteUser(r sex.Request) (sex.Response, int) {
    user := User{}
    if db.First(&user, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("User delete fail, user not found")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, user) {
        msg := "Authentication fail, permission denied"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }


    user.Delete()

    return sex.Response {
        Type: "Sucess",
        Message: "User deleted",
    }, 200
}

func GetUserList(r sex.Request) (sex.Response, int) {
    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
        msg := "Authentication fail, permission denied"
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    limit, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("p"))
    r.Conf["query"].(url.Values).Del("l")
    r.Conf["query"].(url.Values).Del("p")

    query := r.Conf["query"].(url.Values).Encode()
    query = str.ReplaceAll(query, "&", " AND ")
    query = str.ReplaceAll(query, "|", " OR ")

    user_list := []User{}
    offset := (page - 1) * limit
    e := db.Offset(offset).Limit(limit).Order("created_at desc, updated_at, id").Find(&user_list, query)

    if e.Error != nil {
        msg := "Error on getting of users"
        sex.Log(msg, e.Error)
        return sex.Response{
            Type: "Error",
            Message: msg,
        }, 500
    }

    query_response := []map[string]interface{}{}
    for _, u := range user_list {
        item := map[string]interface{}{}

        item["name"] = u.Name
        item["id"] = u.ID

        query_response = append(query_response, item)
    }

    return sex.Response{
        Type: "Sucess",
        Data: query_response,
    }, 200
}
