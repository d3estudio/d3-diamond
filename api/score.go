package main
import (
    "github.com/Plankiton/SexPistol"
    "time"
)

type Score struct {
    sex.Model
    Name       string    `json:"name,omitempty"`
    Value      float64   `json:"value,omitempty"`
    UserId     uint      `json:"user_id,omitempty"`
    Diff       float64   `json:"diff,omitempty"`
    SenderId   uint      `json:"-"`
}

type ScoreType struct {
    sex.ModelNoID
    ID         string  `json:"name,omitempty" gorm:"primaryKey"`
    Desc       string  `json:"desc,omitempty"`
}

type ScoreDate struct {
    sex.Model
    ID        uint      `json:"-"`
    Date      time.Time `json:"date" gorm:"NOT NULL, index"`
    UserId    uint      `json:"-"`
}

func dateRange(d string) (time.Time, time.Time) {
    sufix := "01T00:00:00.000Z"
    format := "2006-01-02T15:04:05.000Z"
    now := time.Now().Format(format)[:8]+sufix
    if d == "" {
        d = now
    } else if len(d) <= 8 {
        d += sufix
    }

    begin, err := time.Parse(d, format)
    sex.SuperPut(begin,"\n   -> ", err)
    if err != nil {
        now := time.Now()
        y, m, d := now.Date()
        z := now.Location()
        d = 1

        begin = time.Date(y, m, d, 0, 0, 0, 0, z)
    }
    sex.SuperPut(begin)

    end := begin.AddDate(0, 1, 0)
    return begin, end
}
