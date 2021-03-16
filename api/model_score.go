package main
import (
    "github.com/Plankiton/SexPistol"
)

func (model *Score) Create() bool {
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

func (model *Score) Delete() bool {
    ID := model.ID
    ModelType := model.ModelType

    if sex.ModelDelete(model) == nil {
        sex.Log("Deleted", sex.ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *Score) Save() bool {
    ID := model.ID
    ModelType := model.ModelType

    if sex.ModelSave(model) == nil {
        sex.Log("Updated", sex.ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *Score) Update(columns sex.Dict) bool {
    ID := model.ID
    ModelType := model.ModelType

    if sex.ModelUpdate(model, columns) == nil {
        sex.Log("Updated", sex.ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *ScoreType) Create() bool {
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

func (model *ScoreType) Delete() bool {
    ID := model.ID
    ModelType := model.ModelType

    if sex.ModelDelete(model) == nil {
        sex.Log("Deleted", sex.ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *ScoreType) Save() bool {
    ID := model.ID
    ModelType := model.ModelType

    if sex.ModelSave(model) == nil {
        sex.Log("Updated", sex.ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *ScoreType) Update(columns sex.Dict) bool {
    ID := model.ID
    ModelType := model.ModelType

    if sex.ModelUpdate(model, columns) == nil {
        sex.Log("Updated", sex.ToLabel(ID, ModelType))
        return true
    }

    return false
}
