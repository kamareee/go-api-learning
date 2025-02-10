package todo

type Item struct {
	Id   string `json:"id"`
	Item string `json:"item"`
}

type Service struct {
	todos map[string]string
}

func NewService() *Service {
	return &Service{
		todos: make(map[string]string),
	}
}

func (svc *Service) Add(todo Item) {
	svc.todos[todo.Id] = todo.Item
}

func (svc *Service) Get(id string) string {
	return svc.todos[id]
}

func (svc *Service) GetAll() map[string]string {
	return svc.todos
}

func (svc *Service) Delete(id string) {
	delete(svc.todos, id)
}
