import { Component, h } from "preact";
import FileGlyph from "../images/file.svg";

export default class FirmwareSelect extends Component {
  handleButtonClick = e => {
    e.preventDefault();
    window.external.invoke("openFirmwareFile");
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
            Select a <strong>bin file</strong> compatible with your Planck EZ.
          </p>
        )}
        {model === 1 && (
          <p>
            Select a <strong>hex file</strong> compatible with your ErgoDox EZ.
          </p>
        )}
        <button className="button" onClick={this.handleButtonClick}>
          Select File
        </button>
      </div>
    );
  }
}
