import React from 'react'

const Calendar = ({events,name}) =>
  <div className="calendar">
  <div className="calendarHeader">
    <h1>{name}</h1>
  </div>
  {events.map(e => <div>{JSON.stringify(e)}</div>)}
</div>

export default Calendar;
