package main
import (
    "github.com/Plankiton/SexPistol"
)

func (model *Avaluation) Create() bool {
    if (model.ModelType == "") {
        model.ModelType = sex.GetModelType(model)
    }

    if sex.ModelCreate(model) == nil {
        ID := model.ID
        ModelType := model.ModelType
        sex.Log("Created", sex.ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *Avaluation) Delete() bool {
    ID := model.ID
    ModelType := model.ModelType

    if sex.ModelCreate(model) == nil {
        sex.Log("Deleted", sex.ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *Avaluation) Save() bool {
    ID := model.ID
    ModelType := model.ModelType

    if sex.ModelSave(model) == nil {
        sex.Log("Updated", sex.ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *Avaluation) Update(columns sex.Dict) bool {
    ID := model.ID
    ModelType := model.ModelType

    if sex.ModelUpdate(model, columns) == nil {
        sex.Log("Updated", sex.ToLabel(ID, ModelType))
        return true
    }

    return false
}
