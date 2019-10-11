
import React from 'react'
import posed from 'react-pose'

const ProgressBar = ({label,percent}) => (
  <div className="progressDiv">
    <p className="progressSpan">{label}</p>
    <div className="progress" pose={percent} style={{width:percent+"%"}}></div>
  </div>
)
export default ProgressBar


