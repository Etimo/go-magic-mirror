import React from 'react'
import ProgressBar from './percentbar.jsx'

const MemoryDisplay = ({totalMemory,memoryUsedPercent}) => (
       <div className="memoryDiv">
         <h1>Memory: {totalMemory}</h1>
         <ProgressBar label="Memory used " percent={memoryUsedPercent}/>
       </div>
)

export default MemoryDisplay
