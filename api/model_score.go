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

    if sex.ModelCreate(model) == nil {
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

func (model *UserScore) Create() bool {
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

func (model *UserScore) Delete() bool {
    ID := model.ID
    ModelType := model.ModelType

    if sex.ModelCreate(model) == nil {
        sex.Log("Deleted", sex.ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *UserScore) Save() bool {
    ID := model.ID
    ModelType := model.ModelType

    if sex.ModelSave(model) == nil {
        sex.Log("Updated", sex.ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *UserScore) Update(columns sex.Dict) bool {
    ID := model.ID
    ModelType := model.ModelType

    if sex.ModelUpdate(model, columns) == nil {
        sex.Log("Updated", sex.ToLabel(ID, ModelType))
        return true
    }

    return false
}
