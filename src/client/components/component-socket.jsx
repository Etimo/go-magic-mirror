import React from "react"
import SimpleSocket from "./simple-socket.jsx"

const MopidySocket = ({url,onmessage,writeMessages}) =>  (
        <SimpleSocket onmessage={onmessage} url={url} writeMessages={writeMessages} />
);
export default MopidySocket;
