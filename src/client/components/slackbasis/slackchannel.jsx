
import React from 'react'
import Message from './message.jsx'

const SlackChannel = ({ data }) => {
  return <div key={data.Id} className="slackChannelContainer">
    <div className="slackMessageList">
      {data.slackMessages.map(m =>
        <Message message={m} />)}
    </div>
  </div>
}
export default SlackChannel;
