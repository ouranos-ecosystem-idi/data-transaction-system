package traceability

import (
	"github.com/google/uuid"
)

// ToDo: Modelと処理をリプレース後にファイルごと削除する
// GetPartsStructure
type GetPartsStructure struct {
	TraceID    uuid.UUID `json:"traceId" gorm:"type:uuid;not null" example:"d9a38406-cae2-4679-b052-15a75f5531e6"`
	DataTarget string    `json:"dataTarget" gorm:"type:string;not null" example:"partsStructure"`
}

// PutPartsStructure
type PutPartsStructure struct {
	TraceID        uuid.UUID             `json:"traceID" gorm:"type:uuid;not null" example:"8c53aeff-0452-23d9-1007-141efbd977c2"`
	DataTarget     string                `json:"dataTarget" gorm:"type:string;not null" example:"partsStructure"`
	PartsStructure []PartsStructureModel `json:"partsStructure" gorm:"-"`
}
