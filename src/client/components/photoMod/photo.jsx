import React from "react";

import PhotoSrc from "./photoSrc.jsx";

const Photo = ({ message, id }) => {
  const component =
    message == null ? (
      <div className="photo" name={id}>
        <h1>no photo</h1>
      </div>
    ) : (
      <div className="photo" name={id}>
        <PhotoSrc url={message.Url} />
      </div>
    );
  return component;
};
export default Photo;
