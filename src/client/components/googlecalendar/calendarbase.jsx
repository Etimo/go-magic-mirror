
import React from 'react'
import Calendar from './calendarview.jsx'

const GoogleCalendar = ({message,id}) => {
  return <div key={id} className="calendarContainer">
    {
      (message ==null ? [] : message.calendars).map(c =>
      <Calendar name={c.calendarName} events={c.currentEvents}/>)
    }
  </div>
}
export default GoogleCalendar;
