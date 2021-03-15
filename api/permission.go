package main

func CheckPermissions(curr User, object interface {}) bool {
    if object == nil {
        return len(curr.QueryRoles(1,1, "name = 'Founder'")) == 1
    }

    user := object.(User)

    if len(curr.QueryRoles(1,1, "name = 'Founder' OR name = 'Admin'")) > 0 ||
    curr.ID == user.ID {
        return true
    }

    return false
}
