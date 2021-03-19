import React from "react";

const PhotoSrc = ({ message }) => {
  const component =
    message == null ? (
      <div className="url">
        <h1>No url</h1>
      </div>
    ) : (
      <div className="url">
        <img
          src={message.Url}
          height={message.Height}
          width={message.Width}
        ></img>
        <h1>{/*message.Url*/}</h1>
      </div>
    );
  return component;
};
export default PhotoSrc;
