package taskparam

type Request struct {
	task struct {
		title      string
		dueDate    string
		categoryId uint
	}

	authenticatedUserId uint
}

func NewRequest(title string, dueDate string, categoryId uint, authenticatedUserId uint) *Request {
	return &Request{
		task: struct {
			title      string
			dueDate    string
			categoryId uint
		}{title: title, dueDate: dueDate, categoryId: categoryId},
		authenticatedUserId: authenticatedUserId,
	}
}
func (req *Request) GetTitle() string {
	return req.task.title
}
func (req *Request) GetDueDate() string {
	return req.task.dueDate
}
func (req *Request) GetCategoryId() uint {
	return req.task.categoryId
}
func (req *Request) GetAuthenticatedUserId() uint {
	return req.authenticatedUserId
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
