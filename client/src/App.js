import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import 'bulma/css/bulma.css'
import Topics from './components/Topics'


class App extends Component {
  render() {
    return (
      <div className="App">
          <Topics/>
      </div>
    );
  }
}

export default App;
