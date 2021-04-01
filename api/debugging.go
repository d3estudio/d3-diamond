package main

func test_user () {
    if db.Take(&User{}).Error != nil {
        joao := User{}
        joao.Name = "Joao da Silva"
        joao.Email = "joao@j.com"
        joao.SetPass("maria")
        db.Create(&joao)

        maria := User{}
        maria.Name = "Maria da Silva"
        maria.Email = "maria@j.com"
        maria.SetPass("joao")
        db.Create(&maria)

        pedro := User{}
        pedro.Name = "Pedro da Silva"
        pedro.Email = "pedro@j.com"
        pedro.SetPass("pedro")
        db.Create(&pedro)

        joao_login := Token{}
        joao_login.UserId = joao.ID
        db.Create(&joao_login)
        joao_login.ID = "joao_login_token"
        db.Save(&joao_login)

        maria_login := Token{}
        maria_login.UserId = maria.ID
        db.Create(&maria_login)
        maria_login.ID = "maria_login_token"
        db.Save(&maria_login)

        pedro_login := Token{}
        pedro_login.UserId = pedro.ID
        db.Create(&pedro_login)
        pedro_login.ID = "pedro_login_token"
        db.Save(&pedro_login)
    }
}
