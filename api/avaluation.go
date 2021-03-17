package main
import (
    "github.com/Plankiton/SexPistol"
    "time"
)

type Avaluation struct {
    sex.Model
    UserId     uint  `json:"-"`
    SenderId   uint  `json:"-"`
}

func (self User) CreateAvaluation(user User) Avaluation {
    link := Avaluation{}
    link.Sign(user, self)
    return link
}

func (self *Avaluation) Sign(receiver User, sender User) error {
    now := time.Now()
    y, m, _ := now.Date()
    dt := time.Date(y, m, 1, 0, 0, 0, 0, now.Location())
    date := dt.String()
    dateend := dt.AddDate(0, 1, 0).String()

    e := db.First(self, "user_id = ? AND sender_id = ? AND created_at BETWEEN ? AND ?", receiver.ID, sender.ID,
    date[:10], dateend[:10])
    if e.Error != nil {
        self.UserId = receiver.ID
        self.SenderId = sender.ID
        return sex.ModelCreate(self)
    }

    return e.Error
}

func (self *Avaluation) QueryScores(page int, limit int, query ...interface{}) []Score {
    e := db.First(self)
    if e.Error == nil {
        score_list := []uint{}
        scores := []Score{}
        e := db.Raw("SELECT s.id FROM scores s JOIN avaluations a on s.aval_id = a.id AND a.id = ?", self.ID)
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
