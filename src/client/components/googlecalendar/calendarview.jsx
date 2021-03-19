import React from 'react'
import Event from './event.jsx'

const Calendar = ({ events, name }) =>
  events == null ? (
    <div className="calendar">
      <div className="calendarHeader">
        <h1>{name}</h1>
      </div>
    </div>
  ) :
    <div className="widget fadeIn calendar">
      <div className="calendarHeader">
        <h1>{name}</h1>
      </div>
      <div className="eventDiv">
        {events.sort((a, b) => new Date(a.startTime).getTime() > new Date(b.startTime).getTime())
          .map(e => <Event key={e.summary + e.startTime} event={e} />)}
        {events.sort((a, b) => new Date(a.startTime).getTime() > new Date(b.startTime).getTime())
          .map(e => <div key={e.summary + e.startTime}>{JSON.stringify(e)}</div>)}
      </div>
    </div>

export default Calendar;
