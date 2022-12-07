package entity

type RegisterView struct {
	ID       string `json: "id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginView struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	Account      string `json:"account"`
}

type CreateUserView struct {
	ID       string `json: "id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Point	 int `json:"point"`
}

type GetUserCountView struct {
	TotalCount	int	`json:"totalCount"`
}

type GetTransactionsCountView struct {
	TotalTransactions			int	`json:"totalTransactions"`
	TotalFailedTransactions		int	`json:"totalFailedTransactions"`
	TotalOnProgressTransactions	int	`json:"totalOnProgressTransactions"`
	TotalSuccessTransactions	int	`json:"totalSuccessTransactions"`
}