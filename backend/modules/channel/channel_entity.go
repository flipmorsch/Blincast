package channel

type ChannelEntity struct {
	ID      string `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Nome    string `json:"nome" gorm:"type:varchar(255);not null"`
	Url     string `json:"url" gorm:"type:varchar(255);not null"`
	Youtube string `json:"youtube" gorm:"type:varchar(255);"`
	Tag     string `json:"tag" gorm:"type:varchar(100);not null"`
	Imagem  string `json:"imagem" gorm:"type:varchar(255)"`
	Site    string `json:"site" gorm:"type:varchar(255);not null"`
}
