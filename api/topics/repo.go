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

func (tr *TopicRepo) Add(t Topic) {
	tr.Topics = append(tr.Topics, t)
}

func (tr *TopicRepo) Find(id int) *Topic {
	for i:=0; i < len(tr.Topics); i++ {
		if tr.Topics[i].Id == id {
			return &tr.Topics[i]
		}
	}
	return nil
}

func (tr *TopicRepo) Upvote(id int) {
	for i:=0; i < len(tr.Topics); i++ {
		if tr.Topics[i].Id == id {
			tr.Topics[i].Votes += 1
		}
	}
}
