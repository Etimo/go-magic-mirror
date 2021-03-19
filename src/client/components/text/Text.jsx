import React from 'react'
const isImage = (icon) => !!icon && (icon.endsWith('.jpg') || icon.endsWith('.svg') || icon.endsWith('.png')); 

export default ({ data }) => {
    return (
        <div className="widget fadeIn text" style={{
            gridColumn: `${data.x}`,
            gridRow: `${data.y}`,
            fontSize: "400%"
        }} name={data.id}>
            {data.icon ? 
                <span className={data.icon}></span> : ""
            }
            {data.text}
        </div>
    )
}