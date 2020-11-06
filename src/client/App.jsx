
import React from 'react'
import { Component } from 'react';
import './app.scss';
import ComponentSocket from './components/component-socket.jsx'
import SystemInfo from './components/systeminfo/systeminfo.jsx'
import Clock from './components/clock/clock';
import Text from './components/text/Text';
import GoogleCalendar from './components/googlecalendar/calendarbase.jsx';
import List from './components/List/List';
import Photo from './components/photoMod/photo'

const containerStyles = {
  display: "grid",
  gridGap: "50px",
  gridTemplateColumns: "auto auto auto auto"
};

const components = {
  "Text": Text,
  "List": List,
  "SystemInfo": SystemInfo,
  "Clock": Clock,
  "GoogleCalendar": GoogleCalendar,
  "Photo": Photo,
}

export default class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      creationMessages: [
        /*{
         name:"systeminfo",
         id:"systeminfo",
         delay:500
       },*/
        {
          name: "googlecalendar",
          id: "meetingCalendar",
          calendars: ["Etvrimo Event-bokning"]
        }
      ]
    };
  }

  onmessage = event => {
    console.log("message here:", event.data);
    const data = JSON.parse(event.data);
    var stateUpdate = {};
    stateUpdate[data.Id] = data;
    this.setState(stateUpdate);
  };

  render() {
    return (
      <div style={containerStyles}>
        <ComponentSocket
          url="ws://localhost:8080/ws"
          onmessage={this.onmessage}
          writeMessages={this.state.creationMessages} />
        {Object.keys(this.state).map(id => {
          const component = components[this.state[id].type];
          return component ? React.createElement(component, { message: this.state[id], id: this.state[id].id, key: this.state[id].id}) : ""
        })}
      </div>
    );
  }
}
