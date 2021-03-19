import React from 'react'

const ClockTime = ({data}) => {
    const component = data == null ?
    <div className="date" >
        <h1>
            23 : 60
        </h1>
    </div>
    :
    <div className="date" >
        <h1>
            {data.hour} : {data.minute} : {data.second}
        </h1>
    </div>
    return (
        component
    )
}
export default ClockTime
