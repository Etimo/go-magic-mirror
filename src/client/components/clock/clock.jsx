import React from 'react'
import ClockDate from './clockDate.jsx'
import ClockTime from './clockTime.jsx'

const Clock = ({message,id}) => {
    const component = message == null ?
    <div className="clock" name={id}>
        <h1>
        <ClockDate message={null}/>
        <ClockTime message={null}/>
        </h1>
    </div>
    :
    <div className="widget fadeIn clock" name={id}>
        <ClockDate message={message.date}/>
        <ClockTime message={message.time}/>
    </div>
    return (
        component
    )
}
export default Clock
