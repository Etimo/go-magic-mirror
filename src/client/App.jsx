import React,{ Component } from 'react';
import './app.css';
import ReactImage from './react.png';

export default class App extends Component {

state = {simple:null}
componentWillmount() {
  this.setState({simple:"Not gotten"});
  }
  componentDidMount() {
    fetch('/api/pomodoro')
      .then(res => res.json())
      .then(simple => this.setState({ simple: simple.simple }));
  }

  render() {
    const { simple } = this.state;
    return (
      <div>
        {simple ? <h1>{`Hello ${simple}`}</h1> : <h1>Updating the page when fetched.</h1>}
      </div>
    );
  }
}
