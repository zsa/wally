import React from "react";
import FlashGlyph from "../images/flash.svg";

export default props => {
  const { sent, total } = props.flashProgress;
  const percentage = total === 0 ? 0 : Math.floor((sent / total) * 100);
  return (
    <div>
      <div className="media-container">
        <div className="media">
          <img alt="Flash" src={FlashGlyph} />
        </div>
      </div>
      <h3>Flashing firmware</h3>
      <p>Please don’t unplug your board.</p>
      <div className="rail">
        <div className="progress" style={{ width: `${percentage}%` }} />
      </div>
    </div>
  );
};
