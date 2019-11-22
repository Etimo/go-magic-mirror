import React from 'react'

const ClockTime = ({message}) => {
    const component = message == null ?
    <div className="date" >
        <h1>
            23 : 60
        </h1>
    </div>
    :
    <div className="date" >
        <h1>
            {message.hour} : {message.minute} : {message.second}
        </h1>
    </div>
    return (
        component
    )
}
export default ClockTime
