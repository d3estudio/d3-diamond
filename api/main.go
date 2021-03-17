package main

import (
    "github.com/Plankiton/SexPistol"
    "os"
)

var con_str string = "host=localhost password=joaojoao dbname=d3diamond port=5432 sslmode=disable TimeZone=America/Araguaina"
func main () {
    var router sex.Pistol
    models := []interface{}{
        &User{},
        &Token{},
        &Role{},
        &UserRole{},
        &Score{},
        &ScoreType{},
        &Avaluation{},
    }

    if os.Getenv("DEBUG_MODE") == "true" {
        sex.Log("Entering on Debug mode, using sqlite database")
        db, err = router.SignDB(":memory:", sex.Sqlite,   // Connection string can be set by env DB_URI too
        models...)       // Models to create on database if not exists

        test_user()
    } else {
        sex.Log("Trying to connect to postgresql")
        db, err = router.SignDB(con_str, sex.Postgres, // Connection string can be set by env DB_URI too
        models...)       // Models to create on database if not exists
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

    if db.Take(&ScoreType{}).Error != nil {
        desbravamento := ScoreType{}
        desbravamento.ID = "desbravamento"
        desbravamento.Desc = "Abrir mata virgem com facão. Coragem para explorar fronteiras além das já mapeadas."
        desbravamento.Create()

        comprometimento := ScoreType{}
        comprometimento.ID = "comprometimento"
        comprometimento.Desc = "Honrar o compromisso. Absorver diferentes inteligências para crescer."
        comprometimento.Create()

        criatividade := ScoreType{}
        criatividade.ID = "criatividade"
        criatividade.Desc = "Sentir a sua veia criativa pulsar. Potência para criar com multidisciplinaridade e assim expandir soluções."
        criatividade.Create()

        adaptabilidade := ScoreType{}
        adaptabilidade.ID = "adaptabilidade"
        adaptabilidade.Desc = "N˜ão seja uma samambaia louca, seja um bambu. Conseguir operar em diferentes contextos e fluir."
        adaptabilidade.Create()

        contundencia := ScoreType{}
        contundencia.ID = "contundência"
        contundencia.Desc = "Quebrar o existe. Fôlego para questionar o existente, quebrar se necessário e recriar com confianca."
        contundencia.Create()

        excelencia := ScoreType{}
        excelencia.ID = "excelência"
        excelencia.Desc = "Busca constante para aperfeiçoar a paixão. Realização das tarefas com o nível de excelência dentro da sua fase."
        excelencia.Create()

        comunicacao := ScoreType{}
        comunicacao.ID = "comunicação"
        comunicacao.Desc = "Não contar com telepatia. Prática da comunicação clara dentro do time."
        comunicacao.Create()

        autonomia := ScoreType{}
        autonomia.ID = "autonomia"
        autonomia.Desc = "Ser independente. Assumir responsabilidades, manter-se informado e saber para onde o time está indo."
        autonomia.Create()

        realizacao := ScoreType{}
        realizacao.ID = "realização"
        realizacao.Desc = "Decolar e colocar na rua. Evoluir o projeto através da materialização das ideias criativas."
        realizacao.Create()

        maturidade_emocional := ScoreType{}
        maturidade_emocional.ID = "maturidade emocional"
        maturidade_emocional.Desc = "Tato com as pessoas e chamego com a comunidade. Aprendizados através do poder da escuta com os outros."
        maturidade_emocional.Create()
    }

    if db.Take(&Token{}).Error != nil {
        token := Token{}
        token.UserId = 1
        token.Create()
        token.ID = "3ae3c630a26b2695974a9bae2b2fd0492e9fc81f"
        token.Save()
    }

    if db.Take(&UserRole{}).Error != nil {
        founder := Role{}
        db.First(&founder, "name = ?", "founder")

        root := UserRole{}
        root.UserId = 1
        root.RoleId = founder.ID
        root.Create()
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
    ).

    // ScoreType managment routes
    Add(
        "get", "/score-type/", nil, GetScoreTypeList,
    ).
    Add(
        "post", "/score-type/", sex.RouteConf {
            "need-auth": false,
        }, CreateScoreType,
    ).
    Add(
        "get", "/score-type/{id}", nil, GetScoreType,
    ).
    Add(
        "post", "/score-type/{id}", nil, UpdateScoreType,
    ).
    Add(
        "delete", "/score-type/{id}", nil, DeleteScoreType,
    ).

    // Score managment routes
    Add(
        "get", "/score/{id}", nil, GetScore,
    ).
    Add(
        "post", "/score/{id}", nil, DeleteScore,
    ).
    Add(
        "post", "/score/{id}", nil, UpdateScore,
    ).
    Add(
        "post", "/user/{id}/score/", sex.RouteConf {
            "need-auth": false,
        }, CreateScore,
    )

    router.Run("/", 8000)
}
