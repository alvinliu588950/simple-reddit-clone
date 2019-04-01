/*
 *  Wirtten by the author AlvinLiu.
 */
import React, { Component } from 'react';
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
