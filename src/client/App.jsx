
import React from 'react'
import { Component } from 'react';
import './app.css';
import ComponentSocket from './components/component-socket.jsx'
export default class App extends Component {

state = {
  onmessage : (event) => {
    console.log("message here:"+event.data);
    const data = JSON.parse(event.data);
    var stateUpdate = {}
    stateUpdate[data.id]=data;
    this.setState(stateUpdate);
  }
};

componentWillmount() {
  this.setState({simple:"Not gotten"});
  }
  componentDidMount() {
  }

  render() {
    return (
      <div>
        <ComponentSocket url="ws://localhost:8080/ws" onmessage={this.state.onmessage}/>
      </div>
    );
  }
}
