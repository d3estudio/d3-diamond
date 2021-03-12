package main
import (
    "github.com/Plankiton/SexPistol"
    str "strings"
)

func CheckPermissions(curr User, object interface {}) bool {
    if object == nil {
        return len(curr.QueryRoles(1,1, "name = 'Founder' OR name = 'Admin'")) == 1
    }

    switch sex.GetModelType(object) {
    case "Role":
        role := object.(Role)
        if str.ToLower(role.Name) == "founder" {
            return true
        }
    case "User":
        user := object.(User)

        if len(curr.QueryRoles(1,1, "name = 'Founder' OR name = 'Admin'")) > 0 ||
           curr.ID == user.ID {
            return true
        }
    }

    return false
}
