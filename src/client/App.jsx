import React, { useState, useEffect } from "react";
import "./app.scss";
import Widget from './components/widget'


export default () => {
  const [creationMessages, setCreationMessages] = useState([
    {
      name: "googlecalendar",
      id: "meetingCalendar",
      calendars: ["Etvrimo Event-bokning"],
    },
  ]);
  const [widgets, setWidgets] = useState({});
  const [layout, setLayout] = useState({ cols: 1, rows: 1 });

  useEffect(() => {
    console.log("Setting up websocket")
    const socket = new WebSocket("ws://localhost:8080/ws");
    socket.onopen = () => {
      // this.sendMessages(socket);
    }
    socket.onmessage = (event) => {
      // console.log("message here:", event.data);
      try {
        const data = JSON.parse(event.data);
  
        if ("Id" in data) {
          setWidgets((widgets) => {
            return { ...widgets, [data.Id]: data }
          });
        } else if ("cols" in data) {
          // Layout message
          console.log("Layout message");
          setLayout(data);
        }
      } catch (e) {
        console.error("Unable to parse json");
      }
  
    };
  }, [])

  return (
    <div>
      <div className="grid"
        style={{
          display: "grid",
          gridColumnGap: "5px",
          gridRowGap: "5px",
          width: "100vw",
          height: "100vh",
          gridTemplateColumns: `repeat(${layout.cols}, 1fr)`,
          gridTemplateRows: `repeat(${layout.rows}, 1fr)`,
        }}
      >
        {Object.keys(widgets).map((id) => {
          return (
            <Widget key={id} data={widgets[id]}></Widget>
          )
        })}
      </div>
    </div>
  );
};
