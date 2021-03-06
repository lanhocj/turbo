package internal

type UserResponse struct {
	Email    string           `json:"email"`
	RoleId   int              `json:"roleId"`
	Locked   bool             `json:"locked"`
	IsAdmin  bool             `json:"isAdmin"`
	Role     string           `json:"role"`
	NodeNum  int              `json:"nodeNum"`
	NodeList NodeListResponse `json:"nodeList"`
	Token    string           `json:"token"`
}

type UserCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserNodeResponse struct {
	Name  string `json:"name"`
	ID    uint   `json:"id"`
	Using bool   `json:"using"`
}
type RequestWithEmail struct {
	Email string `json:"email"`
}

type RequestWithEmailAndPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PutUserToNodeRequest struct {
	Email string `json:"email" form:"email"`
	Role  string `json:"role" form:"role"`
	Node  string `json:"node" form:"node"`
}

type (
	UserNodesResponse []*UserNodeResponse
	UserListResponse  []*UserResponse
)
