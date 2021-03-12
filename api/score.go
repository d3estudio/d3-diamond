package main

import (
    "github.com/Plankiton/SexPistol"
)

type Score struct {
    sex.Model
    Name       string `json:"name"`
}

type UserScore struct {
    sex.Model
    UserId     uint
    ScoreId    uint
}

func (self UserScore) Sign(user User, score Score) (User, Score) {
    self.UserId = user.ID
    self.ScoreId = score.ID

    self.Create()
    sex.Log("Linked", sex.ToLabel(user.ID, user.ModelType), user.Name, "to", sex.ToLabel(score.ID, score.ModelType), score.Name)

    return user, score
}

func (self UserScore) Unsign(user User, score Score) (User, Score) {
    self.UserId = user.ID
    self.ScoreId = score.ID

    self.Delete()
    sex.Log("Unlinked", sex.ToLabel(user.ID, user.ModelType), user.Name, "from", sex.ToLabel(score.ID, score.ModelType), score.Name)

    return user, score
}

func (self Score) Sign(user User) (User, Score) {
    link := UserScore{}
    user, self = link.Sign(user, self)

    return user, self
}

func (self Score) Unsign(user User) (User, Score) {
    link := UserScore{}
    e := db.Where("user_id = ? AND score_id = ?", user.ID, self.ID).First(&link)
    if e.Error == nil {
        user, self = link.Unsign(user, self)
    }

    return user, self
}

func (self *Score) GetUsers(page int, limit int) []User {
    e := db.First(self)
    if e.Error == nil {
        user_list := []uint{}
        users := []User{}
        e := db.Raw("SELECT u.id FROM users u INNER JOIN user_scores ur INNER JOIN scores r ON ur.score_id = r.id AND ur.user_id = u.id AND r.id = ?", self.ID)
        if limit > 0 && page > 0 {
            e = e.Offset((page-1)*limit).Limit(limit)
        }
        e = e.Find(&user_list)

        if e.Error == nil {
            e := db.Find(&users, "id in ?", user_list)
            if e.Error == nil {
                return users
            }
        }
        if e.Error == nil {
            e := db.Find(&users, "id in ?", user_list)
            if e.Error == nil {
                return users
            }
        }
    }

    return []User{}
}

func (self *User) GetScores(page int, limit int) []Score {
    e := db.First(self)
    if e.Error == nil {
        score_list := []uint{}
        scores := []Score{}
        e := db.Raw("SELECT r.id FROM scores r INNER JOIN user_scores ur INNER JOIN users u ON ur.score_id = r.id AND ur.user_id = u.id AND u.id = ?", self.ID)
        if limit > 0 && page > 0 {
            e = e.Offset((page-1)*limit).Limit(limit)
        }
        e = e.Find(&score_list)

        if e.Error == nil {
            e := db.Find(&scores, "id in ?", score_list)
            if e.Error == nil {
                return scores
            }
        }
    }

    return []Score{}
}


func (self *Score) QueryUsers(page int, limit int, query ...interface{}) []User {
    e := db.First(self)
    if e.Error == nil {
        user_list := []uint{}
        users := []User{}
        e := db.Raw("SELECT u.id FROM users u INNER JOIN user_scores ur INNER JOIN scores r ON ur.score_id = r.id AND ur.user_id = u.id AND r.id = ?", self.ID)
        if limit > 0 && page > 0 {
            e = e.Offset((page-1)*limit).Limit(limit)
        }

        e = e.Find(&user_list, query...)

        if e.Error == nil {
            e := db.Find(&users, "id in ?", user_list)
            if e.Error == nil {
                return users
            }
        }
    }

    return []User{}
}


func (self *User) QueryScores(page int, limit int, query...interface{}) []Score {
    e := db.First(self)
    if e.Error == nil {
        score_list := []uint{}
        scores := []Score{}
        e := db.Raw("SELECT r.id FROM scores r INNER JOIN user_scores ur INNER JOIN users u ON ur.score_id = r.id AND ur.user_id = u.id AND u.id = ?", self.ID)
        if limit > 0 && page > 0 {
            e = e.Offset((page-1)*limit).Limit(limit)
        }
        e = e.Find(&score_list, query...)

        if e.Error == nil {
            e := db.Find(&scores, "id in ?", score_list)
            if e.Error == nil {
                return scores
            }
        }
    }

    return []Score{}
}
