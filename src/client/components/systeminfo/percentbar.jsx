
import React from 'react'

const ProgressBar = ({label,percent}) => (
  <div className="progressDiv">
    <p className="progressSpan">{label}</p>
    <div className="progress" style={{width:percent+"%"}}></div>
  </div>
)
export default ProgressBar


