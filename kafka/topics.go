package kafka

type Topics struct {
	TODOS string

	CurrentTopic string
}

func NewTopics() *Topics {
	return &Topics{
		TODOS: "todos",
	}
}
