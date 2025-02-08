package dto

type (
	MessageDto struct {
		Text string `json:"text" binding:"required"`
	}

	ChatMessageDto struct {
		Role string `json:"role" binding:"required"`
		Text string `json:"text" binding:"required"`
	}

	ChatDto struct {
		Messages []ChatMessageDto `json:"messages" binding:"required"`
	}


	DbMessageDto struct {
		Email     string `db:"email"`
		Text      string `db:"text"`
		Role      string `db:"role"`
		CreatedAt string `db:"created_at"`
  }
  
	ChatQueryDto struct {
		Limit  int `form:"limit" binding:"required,min=1,max=100"`
		Offset int `form:"offset" binding:"required,min=0"`
	}
)
