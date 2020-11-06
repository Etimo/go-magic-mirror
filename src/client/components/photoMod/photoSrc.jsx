import React from "react";

const PhotoSrc = ({ url }) => {
  const component =
    url == null ? (
      <div className="url">
        <h1>No url</h1>
      </div>
    ) : (
      <div className="url">
        <img src={url}></img>
        <h1>{url}</h1>
      </div>
    );
  return component;
};
export default PhotoSrc;
