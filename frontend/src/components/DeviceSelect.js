import React from "react";

import PlanckGlyph from "../images/planck-logo.png";
import ErgodoxGlyph from "../images/ergodox-logo.png";
import MoonlanderGlyph from "../images/moonlander-logo.png";

export default class DeviceSelect extends React.Component {
  handleDeviceSelect = (e, device) => {
    e.preventDefault();
    window.backend.State.SelectDevice(device.model, device.bus, device.port);
  };

  renderDevices() {
    const { devices } = this.props;
    return devices.map((device, idx) => {
      let model;
      let glyph = null;
      switch (device.model) {
        case 0:
          model = "Select your Planck EZ";
          glyph = PlanckGlyph;
          break;
        case 1:
          model = "Select your Ergodox EZ";
          glyph = ErgodoxGlyph;
          break;
        case 2:
          model = "Select your Moonlander MK1";
          glyph = MoonlanderGlyph;
          break;
        case 3:
          model = "Board in reset mode";
          break;
        default:
          break;
      }
      return (
        <div
          aria-label={model}
          className="media clickable"
          key={idx}
          role="button"
          onClick={e => this.handleDeviceSelect(e, device)}
        >
          {glyph && <img alt={model} className="logo glyph" src={glyph} />}
          {!glyph && <p>{model} </p>}
        </div>
      );
    });
  }

  render() {
    const devices = this.renderDevices();
    return (
      <div>
        <div className="media-container list">{devices}</div>
        <h3>Select keyboard</h3>
        <p>
          There are several keyboards connected that are compatible, please
          select one.
        </p>
      </div>
    );
  }
}
