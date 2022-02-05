package model

type TaskCreat struct {
	Assignee  string `json:"assignee"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Deadline  string `json:"deadline"`
	Status    string `json:"status"`
}

type Task struct {
	Id        string `json:"id"`
	Assignee  string `json:"assignee"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Deadline  string `json:"deadline"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

type ListResp struct{
	Tasks []Task `json:"tasks"` 
	Count int64 `json:"count"`
}
  