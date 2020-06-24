import React from "react";

import PlanckGlyph from "../images/planck.svg";
import ErgodoxGlyph from "../images/ergodox.svg";

export default class DeviceSelect extends React.Component {
  handleDeviceSelect = (e, device) => {
    e.preventDefault();
    window.backend.State.SelectDevice(device.model, device.bus, device.port);
  };

  renderDevices() {
    const { devices } = this.props;
    return devices.map((device, idx) => {
      let model;
      let glyph;
      switch (device.model) {
        case 0:
          model = "Planck EZ";
          glyph = PlanckGlyph;
          break;
        case 1:
          model = "Ergodox EZ";
          glyph = ErgodoxGlyph;
          break;
        default:
          break;
      }
      return (
        <div
          className="media clickable"
          key={idx}
          onClick={e => this.handleDeviceSelect(e, device)}
        >
          <img alt={model} className="glyph" src={glyph} />
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
