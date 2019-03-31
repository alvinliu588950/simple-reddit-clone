import React, { Component } from 'react';
import axios from 'axios';

class Topics extends Component {
    constructor(props) {
        super(props)
        this.state = {
            topics: [],
            content: "",
        }

        this.handleInputChange = this.handleInputChange.bind(this);
        this.addTopics = this.addTopics.bind(this);
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
        var cardStyle = {
            marginBottom: '10px'
        }
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
                                        onClick={this.addTopics}
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
                                <div className="card" key={index} style={cardStyle}>
                                    <div className="card-content">
                                        <p>{item.Content}</p>
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
                console.log(resp)
                self.setState({ 
                    topics: resp.data.topics
                });
            })
            .catch((err) => {
                console.log(err);
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
            console.log(err);
        })      
    }
}

export default Topics;
