import React from 'react'
export default({ message, id }) => (
    message ?
        <div style={{
            gridColumn: `span ${message.width}`,
            gridRow: `span ${message.height}`,
        }} className="text" name={id}>
            <ul>
                { message.values.map(v => (<li>{v}</li>)) }
            </ul>
        </div> : "")