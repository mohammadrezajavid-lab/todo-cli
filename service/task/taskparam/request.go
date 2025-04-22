package taskparam

type RequestTask struct {
	task struct {
		title      string
		dueDate    string
		categoryId uint
	}

	authenticatedUserId uint
}

func NewRequest(title string, dueDate string, categoryId uint, authenticatedUserId uint) *RequestTask {
	return &RequestTask{
		task: struct {
			title      string
			dueDate    string
			categoryId uint
		}{title: title, dueDate: dueDate, categoryId: categoryId},
		authenticatedUserId: authenticatedUserId,
	}
}
func (req *RequestTask) GetTitle() string {
	return req.task.title
}
func (req *RequestTask) GetDueDate() string {
	return req.task.dueDate
}
func (req *RequestTask) GetCategoryId() uint {
	return req.task.categoryId
}
func (req *RequestTask) GetAuthenticatedUserId() uint {
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

type ListByDateRequest struct {
	userId  uint
	dueDate string
}

func NewListByDateRequest(userId uint, dueDate string) *ListByDateRequest {
	return &ListByDateRequest{
		userId:  userId,
		dueDate: dueDate,
	}
}
func (lr *ListByDateRequest) GetUserId() uint {
	return lr.userId
}
func (lr *ListByDateRequest) GetDueDate() string {
	return lr.dueDate
}
