import React from "react";
import Clock from "./clock/clock";
import GoogleCalendar from "./googlecalendar/calendarbase";
// import List from "./List/List";
import Photo from "./photoMod/photo";
import SystemInfo from "./systeminfo/systeminfo";
import Text from "./text/Text";
import SlackChannel from "./slackbasis/slackchannel";

const components = {
  Text: Text,
  //   List: List,
  SystemInfo: SystemInfo,
  Clock: Clock,
  //   GoogleCalendar: GoogleCalendar,
  Slack: SlackChannel,
  Photo: Photo,
};

export default ({ data, layout }) => {
  const component = components["type" in data ? data.type : layout.pluginType];
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
        gridColumn: `${layout?.x} / ${layout?.x + layout?.width}`,
        gridRow: `${layout?.y} / ${layout?.y + layout?.height}`,
      }}
      className={`widget type-${data?.type?.toLowerCase()}`}
    >
      {el}
    </div>
  );
};
