import React from 'react'
export default ({ message, id }) => (
    message ?
        <div style={{
            gridColumn: `span ${message.width}`,
            gridRow: `span ${message.height}`,
        }} className="text" name={id}>
            {message.icon ? <span>{message.icon}</span> : ""}
            {message.value}
        </div> : "")