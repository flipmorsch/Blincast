package channel

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ChannelEntity struct {
	gorm.Model
	Nome    *string  `json:"nome" gorm:"type:varchar(255);not null" validate:"required"`
	Url     *string  `json:"url" gorm:"type:varchar(255);not null" validate:"required,url"`
	Youtube *string  `json:"youtube" gorm:"type:varchar(255);" validate:"omitempty,url"`
	Tag     *string  `json:"tag" gorm:"type:varchar(100);not null" validate:"required"`
	Imagem  *string `json:"imagem" gorm:"type:varchar(255)"`
	Site    *string  `json:"site" gorm:"type:varchar(255);not null" validate:"required,url"`
}

var (
	validateOnce sync.Once
	validateInst *validator.Validate
)

func getValidator() *validator.Validate {
	validateOnce.Do(func() {
		validateInst = validator.New()
	})
	return validateInst
}

func ValidateChannelEntity(ch *ChannelEntity) error {
	if ch == nil {
		return nil
	}
	return getValidator().Struct(ch)
}
