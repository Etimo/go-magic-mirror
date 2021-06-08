import React, { Fragment } from "react";
import ClockDate from "./clockDate.jsx";
import ClockTime from "./clockTime.jsx";

export default ({ data }) => {
  return (
    <Fragment>
      <ClockDate data={data.date} />
      <br />
      <ClockTime data={data.time} />
    </Fragment>
  );
};
