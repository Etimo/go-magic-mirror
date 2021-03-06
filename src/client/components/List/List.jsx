import React from 'react'
export default({ message, id }) => (
    message ?
        <div style={{
            gridColumn: `span ${message.width}`,
            gridRow: `span ${message.height}`,
        }} className="text" name={id}>
            {message.values.map(v => (<div><span>{v.icon}</span> {v.value}</div>)) }
        </div> : "")