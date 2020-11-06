import React from "react";

import PhotoSrc from "./photoSrc.jsx";

const Photo = ({ message, id }) => {
  const component =
    message == null ? (
      <div className="photo" name={id}>
        <h1>no photo</h1>
      </div>
    ) : (
      <div className="widget fadeIn" name={id}>
        <PhotoSrc message={message} />
      </div>
    );
  return component;
};
export default Photo;
