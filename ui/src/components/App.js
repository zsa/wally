import { Component, h } from "preact";

import DeviceProbe from "./DeviceProbe";
import DeviceSelect from "./DeviceSelect";
import FirmwareSelect from "./FirmwareSelect";
import DeviceReset from "./DeviceReset";
import FirmwareFlash from "./FirmwareFlash";
import FlashComplete from "./FlashComplete";
import Console from "./Console";
import Footer from "./Footer";

export default class App extends Component {
  state = {
    logToggle: false
  };

  toggleLog = () => {
    this.setState({ logToggle: !this.state.logToggle });
  };

  render() {
    const {
      state: {
        data: { device, devices, logs, step, flashProgress },
        completeFlash,
        probeDevices,
        selectDevice,
        selectFirmware,
        flashFirmware,
        pollFlashProgress,
        resetState
      }
    } = this.props;

    const hasError = logs.some(log => log.level == "error");

    return (
      <div className="frame">
        <div className="body">
          <ul className="screens">
            <li className={step === 0 ? "screen active" : "screen"}>
              {step === 0 && (
                <DeviceProbe devices={devices} probeDevices={probeDevices} />
              )}
            </li>
            <li className={step === 1 ? "screen active" : "screen"}>
              {step === 1 && (
                <DeviceSelect devices={devices} selectDevice={selectDevice} />
              )}
            </li>
            <li className={step === 2 ? "screen active" : "screen"}>
              {step === 2 && (
                <FirmwareSelect
                  device={device}
                  selectFirmware={selectFirmware}
                />
              )}
            </li>
            <li className={step === 3 ? "screen active" : "screen"}>
              {step === 3 && (
                <DeviceReset
                  device={device}
                  pollFlashProgress={pollFlashProgress}
                />
              )}
            </li>
            <li className={step === 4 ? "screen active" : "screen"}>
              {step === 4 && (
                <FirmwareFlash
                  device={device}
                  completeFlash={completeFlash}
                  flashFirmware={flashFirmware}
                  pollFlashProgress={pollFlashProgress}
                  flashProgress={flashProgress}
                />
              )}
            </li>
            <li className={step === 5 ? "screen active" : "screen"}>
              {step === 5 && <FlashComplete resetState={resetState} />}
            </li>
          </ul>
        </div>
        {step > 1 && (
          <div className="dots">
            <span className={step === 2 ? "dot active" : "dot"} />
            <span className={step === 3 ? "dot active" : "dot"} />
            <span className={step === 4 ? "dot active" : "dot"} />
            <span className={step === 5 ? "dot active" : "dot"} />
          </div>
        )}
        {this.state.logToggle === true && <Console logs={logs} />}
        <Footer
          hasError={hasError}
          step={step}
          model={device.model}
          toggleLog={this.toggleLog}
        />
      </div>
    );
  }
}
