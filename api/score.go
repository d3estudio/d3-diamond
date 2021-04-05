package main
import (
    SexDB "github.com/Plankiton/SexPistol/Database"
    "time"
)

type Score struct {
    SexDB.Model
    Name       string    `json:"name,omitempty"`
    Value      float64   `json:"value,omitempty"`
    UserId     uint      `json:"user_id,omitempty"`
    Diff       float64   `json:"diff,omitempty"`
    SenderId   uint      `json:"-"`
}

type ScoreType struct {
    SexDB.Model
    ID         string  `json:"name,omitempty" gorm:"primaryKey"`
    Desc       string  `json:"desc,omitempty"`
}

type ScoreDate struct {
    SexDB.Model
    ID        uint      `json:"-"`
    Date      time.Time `json:"date" gorm:"NOT NULL, index"`
    UserId    uint      `json:"-"`
}

func (model *ScoreDate) TableName() string {
    return "score_dates"
}

func (model *ScoreType) TableName() string {
    return "score_types"
}

func (model *Score) TableName() string {
    return "scores"
}

func (model *Score) New() error {
    dt_begin, _ := dateRange("")
    if db.First(&ScoreDate{}, "user_id = ? and date = ?", model.UserId, dt_begin).Error != nil {
        date := ScoreDate { Date: dt_begin, UserId: model.UserId}
        db.Create(&date)
    }

    return nil
}

func dateRange(date string) (time.Time, time.Time) {
    format := "2006-01-02"
    begin, err := time.Parse(format, date)
    if err != nil {
        begin = time.Now()
    }

    end := begin.AddDate(0, 1, 0)
    return begin, end
}

