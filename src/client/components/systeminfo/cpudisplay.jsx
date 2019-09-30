import React from 'react'
import ProgressBar from './percentbar.jsx'

const CpuDisplay = ({cpus}) => {

  const cpuDisp = Object.keys(cpus).map(type => {
    return (<div className="modelCpuDiv">
      <h1 className="cpuModel">{type}</h1>
      {cpus[type].map(cpu =>
        <ProgressBar key={cpu.cpuIndex} label={cpu.cpuIndex} percent={cpu.utilization}/>
      )}
      </div>)
  })
  console.log(cpuDisp)
  return (
  <div className="cpuDiv">
    {cpuDisp}
  </div>
)}

export default CpuDisplay
