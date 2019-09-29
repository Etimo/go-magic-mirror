
import React from "react"

class SimpleSocket extends React.Component{
  //Default methods, will always attempt reconnect.
  onClose = (e) => {
    console.log("Socket closed.. reconnecting: ",e.reason)
    setTimeout(() => {this.connect();},500)
  }
  onError = (err) => {
    console.log("Socket experienced error: ",err);
    setTimeout(() => {this.connect();},500)
    this.state.socket.close()
  }

  connect= () => {
        const url = this.buildUrl(this.props.url);
        console.log("Connecting socket ",url," : ",this.props.onmessage);
        let socket = new WebSocket(url);
        socket.onmessage = this.props.onmessage;
        socket.onerror = this.onError;
        socket.onclose = this.onclose;
        console.log(this.onClose);
        this.setState({
            socket:socket
        });
  }
    buildUrl(url){
        if(url.startsWith("ws://")||url.startsWith("wss://")){
            return url;
        }
        const location = window.location;
        const start = location.scheme === "https:" ? "wss://" : "ws://";
        const urlToo =  start + location.host + (url.startsWith("/") ? "" :"/")+url;
        return urlToo;
    }
	constructor(props) {
        super(props)
        this.state={};
    }
    componentDidMount(){
      this.connect()
    }
    render(){
        return <span className="SocketSpan"></span>
    }

}
export default SimpleSocket;
