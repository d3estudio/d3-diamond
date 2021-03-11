package main

import (
    "github.com/Plankiton/SexPistol"
)

var con_str string = "host=localhost password=joaojoao dbname=d3diamond port=5432 sslmode=disable TimeZone=America/Araguaina"
func main () {
    router := new(sex.Pistol)
    db, err = router.SignDB(con_str, sex.Postgres,   // Connection string can be set by env DB_URI too
        &User{}, &Token{}, &Role{}, &UserRole{})     // Models to create on database if not exists
    if err != nil {
        sex.Die("Database connection fail!")
    }

    sex.Log("Database connection sucessfull!")

    router.
    Add("get", "/{name}", nil, func(r sex.Request) ([]byte, int) {
        return []byte(
            "Hello, "+ r.PathVars["name"],
        ), 200
    })

    router.Run("/", 8000)
}
