
import React from 'react'

const Message = ({ message }) =>
    <div className="slackMessage">
        <p>{message.text}</p>
    </div>

export default Message;
