package main

import (
    "fmt"

    "net/url"
    "net/http"
    sc "strconv"
    str "strings"
    "github.com/Plankiton/SexPistol"
)

func GetRole(r Sex.Request) (Sex.Json, int) {
    u := Role {}
    if db.First(&u, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Role not found")
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

func CreateRole(r Sex.Request) (Sex.Json, int) {
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
        msg := fmt.Sprint("Role create fail, data need to be a object")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    if _, e := data["name"]; !e {
        msg := "Role create fail, Obrigatory field \"name\""
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    if db.First(&Role {}, "name = ?", data["name"]).Error == nil {
        msg := fmt.Sprint("Role create fail, role already registered")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    role := Role {}

    Sex.Copy(data, &role)
    db.Add(&role)

    return Sex.Bullet {
        Type: "Sucess",
        Data: role,
    }, 200
}

func UpdateRole(r Sex.Request) (Sex.Json, int) {
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
        msg := fmt.Sprint("Role create fail, data need to be a object")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    role := Role{}
    if db.First(&role, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Role update fail, role not found")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    Sex.Copy(data, &role)
    db.Sav(&role)

    return Sex.Bullet {
        Type: "Sucess",
        Data: role,
    }, 200
}

func DeleteRole(r Sex.Request) (Sex.Json, int) {
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

    role := Role{}
    if db.First(&role, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Role delete fail, role not found")
        Sex.Err(msg)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    db.Delete(&role)

    return Sex.Bullet {
        Type: "Sucess",
        Message: "Role deleted",
    }, 200
}

func RoleSignUser(r Sex.Request) (Sex.Json, int) {
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

    user := User{}
    if db.First(&user, "id = ?", r.PathVars["uid"]).Error != nil {
        return Sex.Bullet {
            Type: "Error",
            Message: "User not found",
        }, 404
    }

    role := Role{}
    if db.First(&role, "id = ?", r.PathVars["rid"]).Error != nil {
        return Sex.Bullet {
            Type: "Error",
            Message: "Role not found",
        }, 404
    }

    user, role = role.Sign(user)
    if role.ID != 0 {
        return Sex.Bullet {
            Type: "Sucess",
            Message: fmt.Sprint(user.Name, " Signed to ", role.Name),
        }, 200
    }

    msg := "This user already signed to this role"
    Sex.Err(msg)
    return Sex.Bullet {
        Type: "Sucess",
        Message: "This user isn't signed to this role",
    }, 500
}

func RoleUnsignUser(r Sex.Request) (Sex.Json, int) {
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

    user := User{}
    if db.First(&user, "id = ?", r.PathVars["uid"]).Error != nil {
        return Sex.Bullet {
            Type: "Error",
            Message: "User not found",
        }, 404
    }

    role := Role{}
    if db.First(&role, "id = ?", r.PathVars["rid"]).Error != nil {
        return Sex.Bullet {
            Type: "Error",
            Message: "Role not found",
        }, 404
    }

    user, role = role.Unsign(user)
    if role.ID != 0 {
        return Sex.Bullet {
            Type: "Sucess",
            Message: fmt.Sprint(user.Name, " Unsigned to ", role.Name),
        }, 200
    }

    Sex.Err("This user isn't signed to this role")
    return Sex.Bullet {
        Type: "Sucess",
        Message: "This user isn't signed to this role",
    }, 500
}

func GetUserListByRole(r Sex.Request) (Sex.Json, int) {
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

    role := Role{}
    if db.First(&role, "id = ?", r.PathVars["id"]).Error != nil {
        return Sex.Bullet {
            Type: "Error",
            Message: "Role not found",
        }, 404
    }

    limit, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

    role_list := []map[string]interface{}{}

    offset := (page - 1) * limit
    e := db.Table("users u").Select("u.name, u.genre, u.updated_at, u.created_at, u.email").
    Order("r.created_at desc, r.updated_at, r.id").
    Joins("join user_roles l").
    Joins("join roles r on l.user_id = u.id AND l.role_id = r.id AND r.id = ?", role.ID).
    Offset(offset).Limit(limit).
    Find(&role_list).Error
    if e != nil {
        msg := "Unknown error"
        Sex.Err(msg, e)
        return Sex.Bullet {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    return Sex.Bullet {
        Type: "Sucess",
        Data: role_list,
    }, 200

}

func GetRoleList(r Sex.Request) (Sex.Json, int) {
    limit, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

    role_list := []map[string]interface{}{}

    offset := (page - 1) * limit
    e := db.Table("roles r").Select("r.*").
    Order("r.created_at desc, r.updated_at, r.id").
    Joins("join user_roles l").
    Joins("join users u on l.user_id = u.id AND l.role_id = r.id").
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
