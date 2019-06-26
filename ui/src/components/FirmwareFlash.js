import React from "react";
import Loader from "./Loader";
import FlashGlyph from "../images/flash.svg";

export default class FirmwareFlash extends React.Component {
  componentDidMount() {
    this.startPolling();
  }

  componentWillUnmount() {
    this.stopPolling();
  }

  startPolling = () => {
    this.pollInterval = setInterval(() => {
      this.props.pollFlashProgress();
    }, 100);
  };

  stopPolling = e => {
    clearInterval(this.pollInterval);
  };

  render() {
    const { step, sent, total } = this.props.flashProgress;
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
  }
}
