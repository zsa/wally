import React from "react";
import CompleteGlyph from "../images/complete.svg";

export default class FlashComplete extends React.Component {
  handleResetClick = e => {
    e.preventDefault();
    window.backend.State.ResetState();
  };

  handleCloseClick = e => {
    e.preventDefault();
    window.backend.State.Shutdown();
  };

  render() {
    return (
      <div>
        <div className="media-container">
          <div className="media">
            <img alt="Search" className="glyph" src={CompleteGlyph} />
          </div>
        </div>
        <h3>All done!</h3>
        <p>
          Your keyboard was successfully flashed and rebooted. <br />
          Enjoy the new firmware!
        </p>
        <button className="button" onClick={this.handleResetClick}>
          Flash again
        </button>
        <button className="button" onClick={this.handleCloseClick}>
          Close
        </button>
      </div>
    );
  }
}
