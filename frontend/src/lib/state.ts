import { get, writable, type Writable } from "svelte/store";

import { EventsOn } from "../../wailsjs/runtime";

export enum Step { // Each step matches a view in Wally
  PROBING,
  KEYBOARD_SELECT,
  FIRMWARE_SELECT,
  KEYBOARD_RESET,
  FIRMWARE_FLASHING,
  FLASH_COMPLETE,
  FATAL_ERROR,
  WALLY_UPDATE,
  WALLY_UPDATE_COMPLETE,
}

export enum FirmwareFormat {
  HEX,
  BIN,
}

type TProgress = {
  current: number;
  total: number;
};

type Device = {
  bootloader: boolean; // is the device currently in bootloader mode.
  fingerprint: number; // unique id equals to the libusb handle pointer casted to an int
  firmwareFormat: number;
  friendlyName: string; // Friendly name ex: Moonlander MK1
  model: string; // a model identifier used to identify the keyboard ex: moonlander, ergodox, planck, stm32, halfkay, gd32
};

export type TLog = {
  timestamp: number;
  level: "info" | "warning" | "fatal";
  message: string;
};

type TState = {
  appVersion: string;
  checkUpdates: boolean;
  devices: Map<number, Device>;
  flashProgress: number;
  hasError: boolean;
  logs: TLog[];
  selectedDevice: Device | null;
  showAbout: boolean;
  showConsole: boolean;
  step: Step;
  updateProgress: number;
};

const state = writable<TState>({
  appVersion: "",
  checkUpdates: false,
  devices: new Map(),
  flashProgress: 0,
  hasError: false,
  logs: [],
  selectedDevice: null,
  showAbout: false,
  showConsole: false,
  step: Step.PROBING,
  updateProgress: 0,
});

//Attaches the events from the Go process to the UI state
function attachToEvents(state: Writable<TState>) {
  EventsOn("log", ({ log }: { log: TLog }) => {
    state.update((_state: TState) => {
      _state.logs.push(log);
      if (log.level == "fatal") {
        _state.hasError = true;
      }
      return _state;
    });
  });

  EventsOn("reset", () => {
    state.update((_state: TState) => {
      _state.hasError = false;
      _state.selectedDevice = null;
      _state.flashProgress = 0;
      _state.updateProgress = 0;
      return _state;
    });
  });

  EventsOn("stepChanged", ({ step }: { step: Step }) => {
    state.update((_state: TState) => {
      _state.step = step;
      return _state;
    });
  });

  EventsOn("deviceConnected", ({ device }: { device: Device }) => {
    state.update((_state) => {
      _state.devices.set(device.fingerprint, device);
      return _state;
    });
  });

  EventsOn("deviceSelected", ({ device }: { device: Device }) => {
    console.log(device);
    state.update((_state) => {
      _state.selectedDevice = device;
      return _state;
    });
  });

  EventsOn("updateProgress", (progress: TProgress) => {
    state.update((_state) => {
      if (progress.total > 0) {
        _state.updateProgress = Math.floor(
          (progress.current / progress.total) * 100
        );
      }
      return _state;
    });
  });

  EventsOn("flashProgress", (progress: TProgress) => {
    state.update((_state) => {
      if (progress.total > 0) {
        _state.flashProgress = Math.floor(
          (progress.current / progress.total) * 100
        );
      }
      return _state;
    });
  });

  EventsOn("deviceDisconnected", ({ fingerprint }: { fingerprint: number }) => {
    state.update((_state) => {
      _state.devices.delete(fingerprint);
      // If there's no devices connected after removing it from the map
      // Redirect to probing view if the current step is Keyboard select or Firmware select

      if (
        _state.devices.size == 0 &&
        (_state.step == Step.KEYBOARD_SELECT ||
          _state.step == Step.FIRMWARE_SELECT)
      )
        _state.step = Step.PROBING;
      return _state;
    });
  });

  EventsOn("promptUpdateCheck", () => {
    const update = window.confirm(
      "Would you like Wally to check for updates on startup?"
    );
  });
}

attachToEvents(state);

export default state;
