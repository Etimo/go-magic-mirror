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
  const [layout, setLayout] = useState({});
  const [error, setError] = useState(false);
  function connectWs() {
    console.log("Setting up websocket")
    const socket = new WebSocket("ws://localhost:8080/ws");
    socket.onopen = () => {
      // this.sendMessages(socket);
      setError(false);
    }

    socket.onclose = function (e) {
      console.log("Will reconnect in one second", e.reason);
      setTimeout(function () {
        connectWs();
      }, 1000);
    };

    socket.onerror = function (err) {
      setError(true);
      console.error('Error: ', err.message, 'Closing socket');
      socket.close();
    };

    socket.onmessage = (event) => {
      console.log("message here:", event.data);
      try {
        const data = JSON.parse(event.data)

        if ("Id" in data) {
          setWidgets((widgets) => {
            return { ...widgets, [data.Id]: data }
          });
        } else if ("pluginId" in data) {
          setLayout((layout) => {
            return { ...layout, [data.pluginId]: data }
          });         // Layout message
        }
      } catch (e) {
        console.error("Exception in socket ", e);
      }

    };
  }

  useEffect(() => { connectWs(); }, [])

  const backgroundColor = error ? "red" : "black";

  return (
    <div>
      <div className="grid"
        style={{
          backgroundColor,
          display: "grid",
          gridColumnGap: "5px",
          gridRowGap: "5px",
          width: "100vw",
          height: "100vh",
          gridTemplateColumns: `repeat(${layout.cols}, 1fr)`,
          gridTemplateRows: `repeat(${layout.rows}, 1fr)`,
          gridAutoRows: "100px",
          gridAutoColumns: "100px"
        }}
      >
        {Object.keys(widgets).map((id) => {
          return (
            <Widget key={id} data={widgets[id]} layout={layout[id]}></Widget>
          )
        })}
      </div>
    </div>
  );
};
