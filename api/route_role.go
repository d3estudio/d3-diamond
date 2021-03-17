package main

import (
    "fmt"

    "net/url"
    "net/http"
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
    if role.ID != 0 {
        return sex.Response {
            Type: "Sucess",
            Message: fmt.Sprint(user.Name, " Signed to ", role.Name),
        }, 200
    }

    msg := "This user already signed to this role"
    sex.Err(msg)
    return sex.Response {
        Type: "Sucess",
        Message: "This user isn't signed to this role",
    }, 500
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
    if role.ID != 0 {
        return sex.Response {
            Type: "Sucess",
            Message: fmt.Sprint(user.Name, " Unsigned to ", role.Name),
        }, 200
    }

    sex.Err("This user isn't signed to this role")
    return sex.Response {
        Type: "Sucess",
        Message: "This user isn't signed to this role",
    }, 500
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

    query := r.Conf["headers"].(http.Header).Get("Query")
    query = str.ReplaceAll(query, "&&", " AND ")
    query = str.ReplaceAll(query, "||", " OR ")

    if query != "" {
        query = " AND "+query
    }

    role_list := []map[string]interface{}{}

    offset := (page - 1) * limit
    e := db.Table("users u").Select("u.name, u.genre, u.updated_at, u.created_at, u.email").
    Order("r.created_at desc, r.updated_at, r.id").
    Joins("join user_roles l").
    Joins("join roles r on l.user_id = u.id AND l.role_id = r.id AND r.id = ?"+ query, role.ID).
    Offset(offset).Limit(limit).
    Find(&role_list).Error
    if e != nil {
        msg := "Query error, query \""+r.Conf["headers"].(http.Header).Get("Query")+"\" is not valid"
        sex.Err(msg, e)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    return sex.Response{
        Type: "Sucess",
        Data: role_list,
    }, 200

}

func GetRoleList(r sex.Request) (sex.Response, int) {
    limit, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

    query := r.Conf["headers"].(http.Header).Get("Query")
    query = str.ReplaceAll(query, "&&", " AND ")
    query = str.ReplaceAll(query, "||", " OR ")

    if query != "" {
        query = " AND "+query
    }

    role_list := []map[string]interface{}{}

    offset := (page - 1) * limit
    e := db.Table("roles r").Select("r.*").
    Order("r.created_at desc, r.updated_at, r.id").
    Joins("join user_roles l").
    Joins("join users u on l.user_id = u.id AND l.role_id = r.id"+ query).
    Offset(offset).Limit(limit).
    Find(&role_list).Error
    if e != nil {
        msg := "Query error, query \""+r.Conf["headers"].(http.Header).Get("Query")+"\" is not valid"
        sex.Err(msg, e)
        return sex.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    return sex.Response{
        Type: "Sucess",
        Data: role_list,
    }, 200
}
