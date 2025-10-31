package channel

type UpdateChannelDTO struct {
	Nome    *string `json:"nome"`
	Url     *string `json:"url"`
	Youtube *string `json:"youtube"`
	Tag     *string `json:"tag"`
	Imagem  *string `json:"imagem"`
	Site    *string `json:"site"`
}
