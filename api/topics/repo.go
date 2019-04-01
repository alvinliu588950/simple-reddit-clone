/*
 *  Wirtten by the author AlvinLiu.
 */
package topics

import (
	"errors"
	"sort"
	"sync"
)

type Topic struct {
	Id int
	Content string
	Votes int	
}

type TopicRepo struct {
	Topics []Topic
	/*
	 * should lock this resource because it's for in-memory storage 
	 * and may be modified by many goroutine concurrently
	 */
	mux sync.Mutex 
}

func NewRepo() *TopicRepo {
	return &TopicRepo{
		Topics: []Topic{},
	}
}

func (tr *TopicRepo) Add(t Topic) (error) {
	tr.mux.Lock()
	if len(t.Content) > 255 {
		return errors.New("Topic content should not execeed 255 characters")
	}
	tr.Topics = append(tr.Topics, t)
	tr.mux.Unlock()
	return nil
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
			tr.mux.Lock()
			tr.Topics[i].Votes += 1
			tr.mux.Unlock()
			return nil
		}
	}
	return errors.New("topic not found")
}

func (tr *TopicRepo) Downvote(id int) (error) {
	for i:=0; i < len(tr.Topics); i++ {
		if tr.Topics[i].Id == id {
			tr.mux.Lock()
			tr.Topics[i].Votes -= 1
			tr.mux.Unlock()
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


