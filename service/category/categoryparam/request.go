package categoryparam

type Request struct {
	category struct {
		title string
		color string
	}

	authenticatedUserId uint
}

func NewRequest(title string, color string, authenticatedUserId uint) *Request {
	return &Request{
		category: struct {
			title string
			color string
		}{title: title, color: color},
		authenticatedUserId: authenticatedUserId,
	}
}

func (r *Request) GetTitle() string {
	return r.category.title
}
func (r *Request) GetColor() string {
	return r.category.color
}
func (r *Request) GetAuthenticatedUserId() uint {
	return r.authenticatedUserId
}

type ListRequest struct {
	userId uint
}

func NewListRequest(userId uint) *ListRequest {
	return &ListRequest{userId: userId}
}
func (lr *ListRequest) GetUserId() uint {
	return lr.userId
}
