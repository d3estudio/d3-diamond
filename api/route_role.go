package main

import (
    "fmt"

    "net/url"
    sc "strconv"
    str "strings"
    "github.com/Plankiton/SexPistol"
)

func GetRole(r sex.Request) (sex.Response, int) {
    u := Role {}
    if db.First(&u, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Role not found")
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

func CreateRole(r sex.Request) (sex.Response, int) {
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
        msg := fmt.Sprint("Role create fail, data need to be a object")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    if _, e := data["name"]; !e {
        msg := "Role create fail, Obrigatory field \"name\""
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    if db.First(&Role {}, "name = ?", data["name"]).Error == nil {
        msg := fmt.Sprint("Role create fail, role already registered")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    role := Role {}

    sex.MapTo(data, &role)
    role.Create()

    return sex.Response {
        Type: "Sucess",
        Data: role,
    }, 200
}

func UpdateRole(r sex.Request) (sex.Response, int) {
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
        msg := fmt.Sprint("Role create fail, data need to be a object")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    role := Role{}
    if db.First(&role, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Role update fail, role not found")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    sex.MapTo(data, &role)
    role.Save()

    return sex.Response {
        Type: "Sucess",
        Data: role,
    }, 200
}

func DeleteRole(r sex.Request) (sex.Response, int) {
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

    role := Role{}
    if db.First(&role, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Role delete fail, role not found")
        sex.Err(msg)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    role.Delete()

    return sex.Response {
        Type: "Sucess",
        Message: "Role deleted",
    }, 200
}

func RoleUnsignUser(r sex.Request) (sex.Response, int) {
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

    user := User{}
    if db.First(&user, "id = ?", r.PathVars["uid"]).Error != nil {
        return sex.Response{
            Type: "Error",
            Message: "User not found",
        }, 404
    }

    role := Role{}
    if db.First(&role, "id = ?", r.PathVars["rid"]).Error != nil {
        return sex.Response{
            Type: "Error",
            Message: "Role not found",
        }, 404
    }

    user, role = role.Unsign(user)
    return sex.Response {
        Type: "Sucess",
        Message: fmt.Sprint(user.Name, " Unsigned to ", role.Name),
    }, 200
}

func RoleSignUser(r sex.Request) (sex.Response, int) {
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

    user := User{}
    if db.First(&user, "id = ?", r.PathVars["uid"]).Error != nil {
        return sex.Response{
            Type: "Error",
            Message: "User not found",
        }, 404
    }

    role := Role{}
    if db.First(&role, "id = ?", r.PathVars["rid"]).Error != nil {
        return sex.Response{
            Type: "Error",
            Message: "Role not found",
        }, 404
    }

    user, role = role.Sign(user)
    return sex.Response {
        Type: "Sucess",
        Message: fmt.Sprint(user.Name, " Signed to ", role.Name),
    }, 200
}

func GetUserListByRole(r sex.Request) (sex.Response, int) {
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

    role := Role{}
    if db.First(&role, "id = ?", r.PathVars["id"]).Error != nil {
        return sex.Response{
            Type: "Error",
            Message: "Role not found",
        }, 404
    }

    limit, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("p"))
    r.Conf["query"].(url.Values).Del("l")
    r.Conf["query"].(url.Values).Del("p")

    query := r.Conf["query"].(url.Values).Encode()
    query = str.ReplaceAll(query, "&", " AND ")
    query = str.ReplaceAll(query, "|", " OR ")

    user_list := role.QueryUsers(page, limit, query)

    return sex.Response{
        Type: "Sucess",
        Data: user_list,
    }, 200
}


func GetRoleListByUser(r sex.Request) (sex.Response, int) {
    user := User{}
    if (db.First(&user, "id = ?", r.PathVars["id"]).Error != nil) {
        return sex.Response{
            Type: "Error",
            Message: "User not found",
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

    limit, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("p"))
    r.Conf["query"].(url.Values).Del("l")
    r.Conf["query"].(url.Values).Del("p")

    query := r.Conf["query"].(url.Values).Encode()
    query = str.ReplaceAll(query, "&", " AND ")
    query = str.ReplaceAll(query, "|", " OR ")

    role_list := user.QueryRoles(page, limit, query)

    return sex.Response{
        Type: "Sucess",
        Data: role_list,
    }, 200
}

func GetRoleList(r sex.Request) (sex.Response, int) {
    limit, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("p"))
    r.Conf["query"].(url.Values).Del("l")
    r.Conf["query"].(url.Values).Del("p")

    query := r.Conf["query"].(url.Values).Encode()
    query = str.ReplaceAll(query, "&", " AND ")
    query = str.ReplaceAll(query, "|", " OR ")

    role_list := []Role{}
    offset := (page - 1) * limit
    e := db.Offset(offset).Limit(limit).Order("created_at desc, updated_at, id").Find(&role_list, query)

    if e.Error != nil {
        return sex.Response{
            Type: "Error",
            Message: "Error on creating of profile on database",
        }, 500
    }

    return sex.Response{
        Type: "Sucess",
        Data: role_list,
    }, 200
}
