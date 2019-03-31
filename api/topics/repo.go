package topics

import (
	"errors"
	"sort"
)

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

func (tr *TopicRepo) Find(id int) (*Topic, error) {
	for i:=0; i < len(tr.Topics); i++ {
		if tr.Topics[i].Id == id {
			return &tr.Topics[i], nil
		}
	}
	return nil, errors.New("topic not found")
}

func (tr *TopicRepo) Upvote(id int) (error){
	for i:=0; i < len(tr.Topics); i++ {
		if tr.Topics[i].Id == id {
			tr.Topics[i].Votes += 1
			return nil
		}
	}
	return errors.New("topic not found")
}

func (tr *TopicRepo) Downvote(id int) (error) {
	for i:=0; i < len(tr.Topics); i++ {
		if tr.Topics[i].Id == id {
			tr.Topics[i].Votes -= 1
			return nil
		}
	}
	return errors.New("topic not found")
}

func (tr *TopicRepo) GetSorted() ([]Topic, error){
	topics := tr.Topics
	sort.Slice(topics, func(i, j int) bool {
		return topics[i].Votes > topics[j].Votes
	})
	return topics, nil
}


