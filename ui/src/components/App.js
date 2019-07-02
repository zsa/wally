import React, { useCallback, useState } from "react";
import { useDropzone } from "react-dropzone";

import DeviceProbe from "./DeviceProbe";
import DeviceSelect from "./DeviceSelect";
import FirmwareSelect from "./FirmwareSelect";
import DeviceReset from "./DeviceReset";
import FirmwareFlash from "./FirmwareFlash";
import FlashComplete from "./FlashComplete";
import Console from "./Console";
import Footer from "./Footer";

export default function App(props) {
  const {
    state: {
      data: { device, appVersion, devices, logs, step, flashProgress },
      log,
      completeFlash,
      probeDevices,
      selectDevice,
      selectFirmware,
      flashFirmware,
      pollFlashProgress,
      resetState,
      selectFirmwareWithData
    }
  } = props;

  const onDrop = useCallback(files => {
    const allowedExtension = device.model === 0 ? "bin" : "hex";
    const file = files[0];
    const fileExtension = file.name.split(".").pop();
    const isValidExtension = fileExtension === allowedExtension;
    if (isValidExtension === true) {
      const reader = new FileReader();
      reader.readAsArrayBuffer(file);
      reader.onload = function() {
        const view = new Int8Array(reader.result);
        const bin = view.map(n => n.toString(10)).join(" ");
        selectFirmwareWithData(bin);
      };
      reader.onerror = function(evt) {
        log("error", "Error while reading the firmware file.");
      };
    } else {
      log(
        "error",
        `The file "${
          file.name
        }" is not a valid firmware file, a .${allowedExtension} is expected.`
      );
    }
  });

  const hasError = logs.some(log => log.level == "error");

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop
  });

  const [toggleLog, setToggleLog] = useState(false);

  const allowedExtension =
    step === 2 ? (device.model === 0 ? ".bin" : ".hex") : "";

  return (
    <div
      className="frame"
      {...getRootProps({
        onClick: e => e.stopPropagation()
      })}
    >
      <input allow={allowedExtension} {...getInputProps()} />
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
              <FirmwareSelect device={device} selectFirmware={selectFirmware} />
            )}
          </li>
          <li className={step === 3 ? "screen active" : "screen"}>
            {step === 3 && (
              <DeviceReset
                device={device}
                flashFirmware={flashFirmware}
                pollFlashProgress={pollFlashProgress}
              />
            )}
          </li>
          <li className={step === 4 ? "screen active" : "screen"}>
            {step === 4 && (
              <FirmwareFlash
                device={device}
                completeFlash={completeFlash}
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
      {toggleLog === true && <Console logs={logs} />}
      <Footer
        appVersion={appVersion}
        hasError={hasError}
        step={step}
        model={device.model}
        toggleLog={() => {
          setToggleLog(!toggleLog);
        }}
      />
      {isDragActive && step === 2 ? (
        <div className="dnd-overlay">
          {device.model === 0 && (
            <p>
              Drop a <strong>bin file</strong> compatible with your Planck EZ.
            </p>
          )}
          {device.model === 1 && (
            <p>
              Drop a <strong>hex file</strong> compatible with your ErgoDox EZ.
            </p>
          )}
        </div>
      ) : null}
    </div>
  );
}
