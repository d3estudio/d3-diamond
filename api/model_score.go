package main
import (
    "github.com/Plankiton/SexPistol"
)

func (model *Score) Create() {
    if (model.ModelType == "") {
        model.ModelType = sex.GetModelType(model)
    }

    if sex.ModelCreate(model) == nil {
        ID := model.ID
        ModelType := model.ModelType
        sex.Log("Created", sex.ToLabel(ID, ModelType))
    }
}

func (model *Score) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    if sex.ModelCreate(model) == nil {
        sex.Log("Deleted", sex.ToLabel(ID, ModelType))
    }
}

func (model *Score) Save() {
    ID := model.ID
    ModelType := model.ModelType

    if sex.ModelSave(model) == nil {
        sex.Log("Updated", sex.ToLabel(ID, ModelType))
    }
}

func (model *Score) Update(columns sex.Dict) {
    ID := model.ID
    ModelType := model.ModelType

    if sex.ModelUpdate(model, columns) == nil {
        sex.Log("Updated", sex.ToLabel(ID, ModelType))
    }
}
