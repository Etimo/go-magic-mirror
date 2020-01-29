import React from 'react'
import Event from './event.jsx'

const Calendar = ({events,name}) =>
  <div className="calendar">
    <div className="calendarHeader">
      <h1>{name}</h1>
    </div>
    <div className="eventDiv">
      {events.sort((a,b) => new Date(a.startTime).getTime()> new Date(b.startTime).getTime())
          .map(e => <Event key={e.summary+e.startTime} event={e}/>)}
      {events.map(e => <div>{JSON.stringify(e)} key={e.summary+e.startTime}</div>)}
    </div>
  </div>

    export default Calendar;
