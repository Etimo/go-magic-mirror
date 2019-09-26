import React from "react"
import SimpleSocket from "./simple-socket.jsx"
const UpdateSocket = ({url,onmessage}) =>  (
        <SimpleSocket onmessage={onmessage} url={url} />
);
export default UpdateSocket;
