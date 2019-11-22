import React from 'react'

const ClockDate = ({message}) => {
    const component = message == null ?
    <div className="date" >
        <h1>
            Dec 32
        </h1>
    </div>
    :
    <div className="date" >
        <h1>
            {message.day} {message.month} - {message.year}
        </h1>
    </div>
    return (
        component
    )
}
export default ClockDate
