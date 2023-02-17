import React from 'react'

const MessageList = ({ messages }) =>
  messages == null ? (
    <div className="slackList">
      <div className="slackHeader">
      </div>
    </div>
  ) :
    <div className="widget fadeIn calendar">
      <div className="slackHeader">
        {messages.sort(
          (a, b) => a.timestamp > b.timestamp)
          .map(m => <p>m.UserName</p>)}
      </div>
    </div>

export default MessageList;
