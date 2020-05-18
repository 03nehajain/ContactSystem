package contracts

type GetSingleContactResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}
