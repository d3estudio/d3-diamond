package main

func test_user () {
    if db.Take(&User{}).Error != nil {
        joao := User{}
        joao.Name = "Joao da Silva"
        joao.Email = "joao@j.com"
        joao.SetPass("maria")
        joao.Create()

        maria := User{}
        maria.Name = "Maria da Silva"
        maria.Email = "maria@j.com"
        maria.SetPass("joao")
        maria.Create()

        pedro := User{}
        pedro.Name = "Pedro da Silva"
        pedro.Email = "pedro@j.com"
        pedro.SetPass("pedro")
        pedro.Create()

        joao_login := Token{}
        joao_login.UserId = joao.ID
        joao_login.Create()
        joao_login.ID = "joao_login_token"
        joao_login.Save()

        maria_login := Token{}
        maria_login.UserId = maria.ID
        maria_login.Create()
        maria_login.ID = "maria_login_token"
        maria_login.Save()

        pedro_login := Token{}
        pedro_login.UserId = pedro.ID
        pedro_login.Create()
        pedro_login.ID = "pedro_login_token"
        pedro_login.Save()
    }
}
