package models

import "errors"

var ErrorItemTypeAlreadyExists = errors.New("tipo de item já existente no inventário")
var ErrorItemIdNotFoundOnInventory = errors.New("item não encontrado no inventário")

var ErrorRaceNotFound = errors.New("raça não encontrada")
var ErrorClassNotFound = errors.New("classe não encontrada")

var ErrorCreatureAlreadyExists = errors.New("criatura já existente")

var ErrorNotEnoughSkillPoints = errors.New("pontos de habilidade insuficientes")
