package main

import (
    "fmt"

    "net/url"
    "net/http"
    sc "strconv"
    str "strings"

    "github.com/Plankiton/SexPistol"
)

func GetUser(r Sex.Request) (Sex.Json, int) {
    u := User {}
    id, _ := sc.Atoi(r.PathVars["id"])
    if db.First(&u, "id = ?", id).Error != nil {
        msg := fmt.Sprint("User not found")
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
        if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, u) {
            msg := "Authentication fail, your user not exists or dont have permissions to acess this"
            Sex.Err(msg)
            return Sex.Bullet {
                Message: msg,
                Type:    "Error",
            }, 405
        }
    }

    return Sex.Bullet {
        Type: "Success",
        Data: u,
    }, 200
}

func CreateUser(r Sex.Request) (Sex.Json, int) {
    token := Token {}

    auth := r.Header.Get("Authorization")
    if auth != "" {
        db.First(&token, "id = ?", auth)
        if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
            msg := "Authentication fail, logged user not found or dont have permissions to acess this"
            Sex.Err(msg)
            return Sex.Bullet {
                Message: msg,
                Type:    "Error",
            }, 405
        }
    }

    var data map[string]interface{}
    if r.JsonBody(&data) != nil {
        msg := fmt.Sprint("User create fail, data need to be a object")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }


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
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    if db.First(&User {}, "email = ?", data["email"]).Error == nil {
        msg := fmt.Sprint("User create fail, user already registered")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    user := User {}

    Sex.Copy(data, &user)
    user.SetPass(data["pass"].(string))
    db.Add(&user)

    role := Role{}
    db.First(&role, "name = ?", "user")
    role.Sign(user)

    return Sex.Bullet {
        Type: "Sucess",
        Data: user,
    }, 200
}

func UpdateUser(r Sex.Request) (Sex.Json, int) {
    var data map[string]interface{}
    if r.JsonBody(&data) != nil {
        msg := fmt.Sprint("User create fail, data need to be a object")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }


    user := User{}
    if db.First(&user, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("User update fail, user not found")
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
            msg := "Update fail, permission denied"
            Sex.Err(msg)
            return Sex.Bullet {
                Message: msg,
                Type:    "Error",
            }, 405
        }
    }


    Sex.Copy(data, &user)
    if _, e := data["pass"];e {
        user.SetPass(data["pass"].(string))
    }

    if db.Sav(&user) == nil {
        return Sex.Bullet {
            Type: "Sucess",
            Data: user,
        }, 200
    }

    return Sex.Bullet {
        Type: "Error",
        Message: "Tryed to update with field already existent",
    }, 500
}

func DeleteUser(r Sex.Request) (Sex.Json, int) {
    user := User{}
    if db.First(&user, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("User delete fail, user not found")
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
            msg := "Authentication fail, permission denied"
            Sex.Err(msg)
            return Sex.Bullet {
                Message: msg,
                Type:    "Error",
            }, 405
        }
    }


    db.Del(&user)

    return Sex.Bullet {
        Type: "Sucess",
        Message: "User deleted",
    }, 200
}

func GetUserList(r Sex.Request) (Sex.Json, int) {
    token := Token {}

    auth := r.Header.Get("Authorization")
    if auth != "" {
        db.First(&token, "id = ?", auth)
        if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
            msg := "Authentication fail, permission denied"
            Sex.Err(msg)
            return Sex.Bullet {
                Message: msg,
                Type:    "Error",
            }, 405
        }
    }

    limit, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("p"))
    r.Conf["query"].(url.Values).Del("l")
    r.Conf["query"].(url.Values).Del("p")

    user_list := []User{}
    offset := (page - 1) * limit
    e := db.Offset(offset).Limit(limit).Order("created_at desc, updated_at, id").Find(&user_list)

    if e.Error != nil {
        msg := "Error on getting of users"
        Sex.Log(msg, e.Error)
        return Sex.Bullet {
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

    return Sex.Bullet {
        Type: "Sucess",
        Data: query_response,
    }, 200
}

func GetRoleListByUser(r Sex.Request) (Sex.Json, int) {
    user := User{}
    if (db.First(&user, "id = ?", r.PathVars["id"]).Error != nil) {
        return Sex.Bullet {
            Type: "Error",
            Message: "User not found",
        }, 404
    }

    token := Token {}

    auth := r.Header.Get("Authorization")
    if auth != "" {
        db.First(&token, "id = ?", auth)
        if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, user) {
            msg := "Authentication fail, permission denied"
            Sex.Err(msg)
            return Sex.Bullet {
                Message: msg,
                Type:    "Error",
            }, 405
        }
    }

    limit, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

    role_list := []map[string]interface{}{}

    offset := (page - 1) * limit
    e := db.Table("roles r").Select("r.*").
    Order("r.created_at desc, r.updated_at, r.id").
    Joins("join user_roles l").
    Joins("join users u on l.user_id = u.id AND l.role_id = r.id AND u.id = ?", user.ID).
    Offset(offset).Limit(limit).
    Find(&role_list).Error
    if e != nil {
        msg := "Unknown error ocurred"
        Sex.Err(msg, e)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    return Sex.Bullet {
        Type: "Sucess",
        Data: role_list,
    }, 200
}

