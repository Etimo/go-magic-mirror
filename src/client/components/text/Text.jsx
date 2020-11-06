import React from 'react'
export default ({ message, id }) => (
    message ?
        <div style={{
            gridColumn: `span ${message.width}`,
            gridRow: `span ${message.height}`,
            fontSize: "400%"
        }} className="text" name={id}>
            {message.icon ? <span>{message.icon}</span> : ""}
            {message.value}
        </div> : "")