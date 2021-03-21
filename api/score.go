package main
import "github.com/Plankiton/SexPistol"

type Score struct {
    sex.Model
    AvalId     uint      `json:"-"`
    TypeId     string    `json:"name,omitempty"`
    Value      float64   `json:"value,omitempty"`
    Ended      bool      `json:"-"`
    Total      int       `json:"total,omitempty"`
}

type ScoreType struct {
    sex.ModelNoID
    ID         string  `json:"name,omitempty" gorm:"primaryKey"`
    Desc       string  `json:"desc,omitempty"`
}
