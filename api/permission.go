package main

func CheckPermissions(curr User, object interface {}) bool {
    if object == nil {
        roles := curr.QueryRoles(1,1, "name = 'Founder'")
        return len(roles) == 1
    }

    user := object.(User)

    roles := curr.QueryRoles(1,1, "name = 'Founder' OR name = 'Admin'")

    if len(roles) > 0 ||
    curr.ID == user.ID {
        return true
    }

    return false
}
