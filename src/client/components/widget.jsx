import React from "react";
import Clock from "./clock/clock";
import GoogleCalendar from "./googlecalendar/calendarbase";
// import List from "./List/List";
import Photo from "./photoMod/photo";
import SystemInfo from "./systeminfo/systeminfo";
import Text from "./text/Text";

const components = {
  Text: Text,
//   List: List,
//   SystemInfo: SystemInfo,
  Clock: Clock,
//   GoogleCalendar: GoogleCalendar,
  Photo: Photo,
};

export default ({ data }) => {
  const component = components[data.type];
  const el = component ? (
    React.createElement(component, {
      data,
    })
  ) : (
    <p>Invalid component {data.type}</p>
  );

  return (
    <div
      style={{
        gridColumn: `${data.x} / ${data.x + data.width}`,
        gridRow: `${data.y} / ${data.y + data.height}`,
      }}
      className={`widget type-${data.type.toLowerCase()}`}
    >
        {el}
    </div>
  );
};
