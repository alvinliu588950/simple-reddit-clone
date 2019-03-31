import React, { Component } from 'react';
import axios from 'axios';
import '../css/Topics.css'

class Topics extends Component {
    constructor(props) {
        super(props)
        this.state = {
            topics: [],
            content: "",
        }

        this.handleInputChange = this.handleInputChange.bind(this);
        this.addTopics = this.addTopics.bind(this);
        this.upvote = this.upvote.bind(this);        
        this.downvote = this.downvote.bind(this);
    }
    handleInputChange(event) {
        const target = event.target;
        const value = target.value;
        const name = target.name;

        this.setState({
            [name]: value
        });
    }
    componentDidMount() {
        this.fetchTopics()
    }
    render() {
        return (
            <div>
                <div className="columns">
                    <div className="column is-6 is-offset-3">
                        <h1 className="title has-text-centered">Topics</h1>
                    </div>
                </div>
                <div className="columns">
                    <div className="column is-6 is-offset-3">
                        <form>
                            <div className="field">
                                <label className="label">New Topic</label>
                                <div className="control">
                                    <input 
                                        className="input" 
                                        type="text" 
                                        placeholder="Text input"
                                        name="content"
                                        checked={this.state.content}
                                        onChange={this.handleInputChange}
                                    />
                                </div>
                            </div>
                            <div className="field">
                                <div className="control has-text-right">
                                    <button 
                                        className="button is-link" 
                                        type="button"
                                        onClick={() => this.addTopics()}
                                        >Add
                                    </button>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
                <div className="columns">
                    <div className="column is-6 is-offset-3">
                        {
                            this.state.topics.map((item, index) => (
                                <div className="card" key={index}>
                                    <header className="card-header">
                                        <p className="card-header-title">Topic number: {item.Id} </p>
                                    </header>
                                    <div className="card-content">
                                        <div className="vote">
                                            <button className="button is-light"  onClick={() => this.upvote(item.Id)}>
                                                <i className="fas fa-arrow-up"></i>
                                            </button>
                                            <p>{item.Votes}</p>
                                            <button className="button is-light" onClick={() => this.downvote(item.Id)}>
                                                <i className="fas fa-arrow-down" ></i>
                                            </button>                                    
                                        </div>
                                        <div className="content">
                                            <p>{item.Content}</p>
                                        </div>
                                    </div>
                                </div>
                            ))
                        }
                    </div>
                </div>
            </div>          
        );
    }
    fetchTopics() {
        var self = this;
        axios.get('/api/v1/topics')
            .then((resp) => {
                self.setState({ 
                    topics: resp.data.topics
                });
            })
            .catch((err) => {
                alert('Something went wrong.Please check your network or contact us.')
            })
    }
    addTopics() {
        var self = this;
        axios({
            method: 'post',
            url: '/api/v1/topics/add',
            data: {
                content: self.state.content
            }
        })
        .then((resp) => {
            self.fetchTopics()
        })
        .catch((err) => {
            alert('Something went wrong.Please check your network or contact us.')
        })      
    }
    upvote(id) {
        var self = this;
        axios({
            method: 'patch',
            url: '/api/v1/topics/upvote',
            data: {
                id: id
            }
        })
        .then((resp) => {
            // after user vote, the page should display the updated list of topics
            self.fetchTopics()
        })
        .catch((err) => {
            alert('Something went wrong.Please check your network or contact us.')
        })     
    }
    downvote(id) {
        var self = this;
        axios({
            method: 'patch',
            url: '/api/v1/topics/downvote',
            data: {
                id: id
            }
        })
        .then((resp) => {
            // after user vote, the page should display shows the updated list of topics
            self.fetchTopics()
        })
        .catch((err) => {
            alert('Something went wrong.Please check your network or contact us.')
        })     
    }
}

export default Topics;
