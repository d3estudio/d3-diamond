package main

import (
    "github.com/Plankiton/SexPistol"
    "os"
)

var con_str string = "host=localhost password=joaojoao dbname=d3diamond port=5432 sslmode=disable TimeZone=America/Araguaina"
func main () {
    var router sex.Pistol

    if os.Getenv("DEBUG_MODE") == "true" {
        sex.Log("Entering on Debug mode, using sqlite database")
        db, err = router.SignDB(":memory:", sex.Sqlite,   // Connection string can be set by env DB_URI too
        &User{}, &Token{}, &Role{}, &UserRole{})       // Models to create on database if not exists

        test_user()
    } else {
        sex.Log("Trying to connect to postgresql")
        db, err = router.SignDB(con_str, sex.Postgres, // Connection string can be set by env DB_URI too
        &User{}, &Token{}, &Role{}, &UserRole{})       // Models to create on database if not exists
    }

    if db.Take(&Role{}).Error != nil {
        user := Role{}
        user.Name = "user"
        user.Desc = "Pode visualizar os próprios resultados e avaliar os colegas"
        user.Create()

        adm := Role{}
        adm.Name = "admin"
        adm.Desc = "Pode fazer oque os usuários podem, além de Criar/Editar/Apagar ou gerar relatórios em CSV sobre os usuários"
        adm.Create()

        founder := Role{}
        founder.Name = "founder"
        founder.Desc = "Tem todas as permissões de admin, e pode mudar o role de qualquer usuário"
        founder.Create()
    }

    if err != nil {
        sex.Die("Database connection fail!")
    }

    sex.Log("Database connection sucessfull!")

    router.Auth = true
    router.

    // Authentication routes
    Add(
        "post", "/login", sex.RouteConf {
            "need-auth": false,
        }, LogIn,
    ).
    Add(
        "post", "/verify", nil, Verify,
    ).
    Add(
        "post", "/logout", nil, LogOut,
    ).

    // User managment routes
    Add(
        "get", "/user/", nil, GetUserList,
    ).
    Add(
        "post", "/user/", sex.RouteConf {
            "need-auth": false,
        }, CreateUser,
    ).
    Add(
        "get", "/user/{id}", nil, GetUser,
    ).
    Add(
        "post", "/user/{id}", nil, UpdateUser,
    ).
    Add(
        "delete", "/user/{id}", nil, DeleteUser,
    ).
    Add(
        "get", "/user/{id}/roles", nil, GetRoleListByUser,
    ).

    // Role managment routes
    Add(
        "get", "/role/", nil, GetRoleList,
    ).
    Add(
        "post", "/role/", sex.RouteConf {
            "need-auth": false,
        }, CreateRole,
    ).
    Add(
        "get", "/role/{id}", nil, GetRole,
    ).
    Add(
        "post", "/role/{id}", nil, UpdateRole,
    ).
    Add(
        "delete", "/role/{id}", nil, DeleteRole,
    ).
    Add(
        "post", "/role/{rid}/sign/{uid}", nil, RoleSignUser,
    ).
    Add(
        "post", "/role/{rid}/unsign/{uid}", nil, RoleUnsignUser,
    ).
    Add(
        "get", "/role/{id}/users", nil, GetUserListByRole,
    )

    router.Run("/", 8000)
}
