package main

import (
    SexDB "github.com/Plankiton/SexPistol/Database"
    "github.com/Plankiton/SexPistol"
    "os"
)

var con_str string = "host=localhost password=joaojoao dbname=d3diamond port=5432 sslmode=disable TimeZone=America/Araguaina"
func main () {
    var router Sex.Pistol
    models := []SexDB.ModelSkel {
        &User{},
        &Token{},
        &Role{},
        &UserRole{},
        &Score{},
        &ScoreType{},
        &ScoreDate{},
    }

    if os.Getenv("DEBUG_MODE") == "true" {
        Sex.Log("Entering on Debug mode, using sqlite database")
        db, err = SexDB.Open(":memory:", SexDB.Sqlite)
        db.AddModels(models...)

        test_user()
    } else {
        Sex.Log("Trying to connect to postgresql")
        db, err = SexDB.Open(con_str, SexDB.Postgres)
        db.AddModels(models...)
    }

    if db.Take(&Role{}).Error != nil {
        user := new(Role)
        user.Name = "user"
        user.Desc = "Pode visualizar os próprios resultados e avaliar os colegas"
        db.Create(user)

        adm := new(Role)
        adm.Name = "admin"
        adm.Desc = "Pode fazer oque os usuários podem, além de Criar/Editar/Apagar ou gerar relatórios em CSV sobre os usuários"
        db.Create(adm)

        founder := new(Role)
        founder.Name = "founder"
        founder.Desc = "Tem todas as permissões de admin, e pode mudar o role de qualquer usuário"
        db.Create(founder)
    }

    if db.Take(&ScoreType{}).Error != nil {
        desbravamento := new(ScoreType)
        desbravamento.ID = "desbravamento"
        desbravamento.Desc = "Abrir mata virgem com facão. Coragem para explorar fronteiras além das já mapeadas."
        db.Create(desbravamento)

        comprometimento := new(ScoreType)
        comprometimento.ID = "comprometimento"
        comprometimento.Desc = "Honrar o compromisso. Absorver diferentes inteligências para crescer."
        db.Create(comprometimento)

        criatividade := new(ScoreType)
        criatividade.ID = "criatividade"
        criatividade.Desc = "Sentir a sua veia criativa pulsar. Potência para criar com multidisciplinaridade e assim expandir soluções."
        db.Create(criatividade)

        adaptabilidade := new(ScoreType)
        adaptabilidade.ID = "adaptabilidade"
        adaptabilidade.Desc = "N˜ão seja uma samambaia louca, seja um bambu. Conseguir operar em diferentes contextos e fluir."
        db.Create(adaptabilidade)

        contundencia := new(ScoreType)
        contundencia.ID = "contundência"
        contundencia.Desc = "Quebrar o existe. Fôlego para questionar o existente, quebrar se necessário e recriar com confianca."
        db.Create(contundencia)

        excelencia := new(ScoreType)
        excelencia.ID = "excelência"
        excelencia.Desc = "Busca constante para aperfeiçoar a paixão. Realização das tarefas com o nível de excelência dentro da sua fase."
        db.Create(excelencia)

        comunicacao := new(ScoreType)
        comunicacao.ID = "comunicação"
        comunicacao.Desc = "Não contar com telepatia. Prática da comunicação clara dentro do time."
        db.Create(comunicacao)

        autonomia := new(ScoreType)
        autonomia.ID = "autonomia"
        autonomia.Desc = "Ser independente. Assumir responsabilidades, manter-se informado e saber para onde o time está indo."
        db.Create(autonomia)

        realizacao := new(ScoreType)
        realizacao.ID = "realização"
        realizacao.Desc = "Decolar e colocar na rua. Evoluir o projeto através da materialização das ideias criativas."
        db.Create(realizacao)

        maturidade_emocional := new(ScoreType)
        maturidade_emocional.ID = "maturidade emocional"
        maturidade_emocional.Desc = "Tato com as pessoas e chamego com a comunidade. Aprendizados através do poder da escuta com os outros."
        db.Create(maturidade_emocional)
    }

    if db.Take(&Token{}).Error != nil {
        token := new(Token)
        token.UserId = 1
        db.Create(token)
        token.ID = "3ae3c630a26b2695974a9bae2b2fd0492e9fc81f"
        db.Create(token)
    }

    if db.Take(&UserRole{}).Error != nil {
        founder := new(Role)
        db.First(&founder, "name = ?", "founder")

        root := new(UserRole)
        root.UserId = 1
        root.RoleId = founder.ID
        db.Create(root)
    }

    if err != nil {
        Sex.Die("Database connection fail!")
    }

    Sex.Log("Database connection sucessfull!")

    router.Auth = true
    router.

    // Authentication routes
    Add("/login", LogIn, "post").
    Add("/verify", Verify, "post").
    Add("/logout", LogOut, "post").

    // User managment routes
    Add("/user/", GetUserList, "get").
    Add("/user/", CreateUser, "post").
    Add("/user/{id}", GetUser, "get").
    Add("/user/{id}", UpdateUser, "post").
    Add("/user/{id}", DeleteUser, "delete").
    Add("/user/{id}/roles", GetRoleListByUser, "get").

    // Role managment routes
    Add("/role/", GetRoleList, "get").
    Add("/role/", CreateRole, "post").
    Add("/role/{id}", GetRole, "get").
    Add("/role/{id}", UpdateRole, "post").
    Add("/role/{id}", DeleteRole, "delete").
    Add("/role/{rid}/sign/{uid}", RoleSignUser, "post").
    Add("/role/{rid}/unsign/{uid}", RoleUnsignUser, "post").
    Add("/role/{id}/users", GetUserListByRole, "get").

    // ScoreType managment routes
    Add("/score-type/", GetScoreTypeList, "get").
    Add("/score-type/", CreateScoreType, "post").
    Add("/score-type/{id}", GetScoreType, "get").
    Add("/score-type/{id}", UpdateScoreType, "post").
    Add("/score-type/{id}", DeleteScoreType, "delete").

    // Score managment routes
    Add("/user/{id}/score/", CreateScore, "post").
    Add("/user/{id}/scores/", GetScoreList, "get").

    Add("/score/{id}", GetScore, "get").
    Add("/score/{id}", DeleteScore, "post").
    Add("/score/{id}", UpdateScore, "post").

    // Date registered routes
    Add("/user/{id}/dates", GetDates, "get").
    Add("/user/{id}/date", CreateDate, "post").
    Add("/date/{id}", GetDate, "get").
    Add("/date/{id}", DeleteDate, "delete").
    Add("/date/{id}", UpdateDate, "post")

    router.Run("/", 8000)
}
