import React from "react";
import FileGlyph from "../images/file.svg";

export default class FirmwareSelect extends React.Component {
  handleButtonClick = e => {
    e.preventDefault();
    window.backend.State.SelectFirmware();
  };

  render() {
    const { model } = this.props.device;
    return (
      <div>
        <div className="media-container">
          <div className="media">
            <img alt="Search" className="glyph" src={FileGlyph} />
          </div>
        </div>
        <h3>Select firmware</h3>
        {model === 0 && (
          <p>
            Select or drop a <strong>bin file</strong> compatible with your
            Planck EZ.
          </p>
        )}
        {model === 1 && (
          <p>
            Select or drop a <strong>hex file</strong> compatible with your
            ErgoDox EZ.
          </p>
        )}
        {model === 2 && (
          <p>
            Select or drop a <strong>bin file</strong> compatible with your Moonlander MK1.
          </p>
        )}
        {model === 3 && (
          <p>
            Select or drop a <strong>bin file</strong> compatible with your board
          </p>
        )}
        <button className="button" role="button" onClick={this.handleButtonClick}>
          Select File
        </button>
      </div>
    );
  }
}
