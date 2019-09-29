import React from "react"
import SimpleSocket from "./simple-socket.jsx"

const MopidySocket = ({url,onmessage}) =>  (
        <SimpleSocket onmessage={onmessage} url={url} />
);
export default MopidySocket;
