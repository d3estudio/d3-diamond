package main

import (
    "github.com/Plankiton/SexPistol"
)

func main () {
    router := sex.Pistol{}

    router.
    Add("get", "/{name}", nil, func(r sex.Request) ([]byte, int) {
        return []byte(
            "Hello, "+ r.PathVars["name"],
        ), 200
    })

    router.Run("/", 8000)
}
