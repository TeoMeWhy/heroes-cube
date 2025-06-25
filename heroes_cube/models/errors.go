package models

import "errors"

var ErrorItemTypeAlreadyExists = errors.New("tipo de item já existente no inventário")
var ErrorItemIdNotFoundOnInventory = errors.New("item não encontrado no inventário")
