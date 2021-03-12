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
        db, err = router.SignDB(con_str, sex.Sqlite,   // Connection string can be set by env DB_URI too
        &User{}, &Token{}, &Role{}, &UserRole{})       // Models to create on database if not exists
    } else {
        sex.Log("Trying to connect to postgresql")
        db, err = router.SignDB(con_str, sex.Postgres, // Connection string can be set by env DB_URI too
        &User{}, &Token{}, &Role{}, &UserRole{})       // Models to create on database if not exists
    }

    if err != nil {
        sex.Die("Database connection fail!")
    }

    sex.Log("Database connection sucessfull!")

    router.Auth = true
    router.
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
    )

    test_user()
    router.Run("/", 8000)
}
