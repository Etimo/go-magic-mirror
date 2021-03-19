import React from 'react'
const isImage = (icon) => !!icon && (icon.endsWith('.jpg') || icon.endsWith('.svg') || icon.endsWith('.png')); 
export default ({ message, id }) => (
    message ?
        <div className="widget fadeIn text" style={{
            gridColumn: `span ${message.width}`,
            gridRow: `span ${message.height}`,
            fontSize: "400%"
        }} name={id}>
            {isImage(message.icon) ? <img src={message.icon}/> : message.icon ? <span className={message.icon}></span> : ""}
            {message.value}
        </div> : "")