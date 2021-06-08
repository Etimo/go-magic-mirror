import React from 'react'

const ClockDate = ({data}) => {
    const component = data == null ?
    <div className="date" >
        <h1>
            Dec 32
        </h1>
    </div>
    :
    <div className="date" >
        <h1>
            {data.day} {data.month} - {data.year}
        </h1>
    </div>
    return (
        component
    )
}
export default ClockDate
