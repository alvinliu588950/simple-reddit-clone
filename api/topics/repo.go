package topics

type Topic struct {
	Id int
	Content string
	Votes int	
}

type TopicRepo struct {
	Topics []Topic
}

func NewRepo() *TopicRepo {
	return &TopicRepo{
		Topics: []Topic{},
	}
}

func (tr *TopicRepo) GetAll() []Topic {
	return tr.Topics
}
