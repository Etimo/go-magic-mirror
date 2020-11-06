import React from "react";
import { Component } from "react";
import "./app.scss";
import ComponentSocket from "./components/component-socket.jsx";
import SystemInfo from "./components/systeminfo/systeminfo.jsx";
import Clock from "./components/clock/clock";
import Text from "./components/text/Text";
import GoogleCalendar from "./components/googlecalendar/calendarbase.jsx";
import Photo from "./components/photoMod/photo.jsx";

const containerStyles = {
  display: "grid",
  gridGap: "50px",
  gridTemplateColumns: "auto auto auto auto"
};

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
          writeMessages={this.state.creationMessages}
        />
        <Clock id="clock" message={this.state.clock} />
        <Photo id="photo" message={this.state.photo} />
        <Text message={this.state.weather} />
        <GoogleCalendar
          id="meetinCalendar"
          message={this.state.meetingCalendar}
        />
        <SystemInfo id="systeminfo" message={this.state.systeminfo} />
      </div>
    );
  }
}
