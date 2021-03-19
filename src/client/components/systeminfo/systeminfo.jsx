import React from 'react'
import MemoryDisplay from './memorydisplay.jsx'
import CpuDisplay from './cpudisplay.jsx'

//Props example
/*message here:
 * {"id":"systeminfo",
 * "os":"linux",
 * "hostName":"erik-etimo-laptop",
 * "memoryTotal":"16 GB",
 * "memoryUsedPercent":"30.02 %","usedMemory":"","cpus":[{"ModelName":"Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz","Mhz":3500,"Utilization":8.333333332659633},{"ModelName":"Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz","Mhz":3500,"Utilization":26.000000000203727},{"ModelName":"Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz","Mhz":3500,"Utilization":12.24489795875942},{"ModelName":"Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz","Mhz":3500,"Utilization":22.91666666622455}],"uptime":339526}
*/
const SystemInfo = ({message,id}) => {
  const component = message == null ?

    <div className="systeminfo" name={id}>
      <h1> Awaiting first update</h1>
    </div> :

     <div className="widget systeminfo fadeIn" name={id}>
    <div className="osinfo">
      <p>{message.hostName+" running "+message.os }</p>
      <p>Uptime: {message.uptime} seconds</p>
    </div>
       <MemoryDisplay totalMemory={message.memoryTotal} memoryUsedPercent={message.memoryUsedPercent}/>
       <CpuDisplay cpus={message.cpus}/>
    </div>
  return (
      component
     )
}
export default SystemInfo
