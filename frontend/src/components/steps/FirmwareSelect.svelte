<script lang="ts">
  import {
    SelectDevice,
    SelectFirmware,
  } from "../../../wailsjs/go/state/State";
  import state, { FirmwareFormat } from "../../lib/state";

  function handleClick() {
    SelectFirmware();
  }

  let firmwareExtension = "";
  if ($state.selectedDevice) {
    console.log($state.selectedDevice);
    switch ($state.selectedDevice.firmwareFormat) {
      case FirmwareFormat.HEX:
        firmwareExtension = ".hex";
        break;
      case FirmwareFormat.BIN:
        firmwareExtension = ".bin";
        break;

      default:
        break;
    }
  }
</script>

<h1>Select firmware</h1>
{#if $state.selectedDevice}
  <p>
    Select or drop a <b> {firmwareExtension} file</b>
    compatible with your {$state.selectedDevice.friendlyName}.
  </p>
{/if}
<button class="btn" on:click={handleClick}>Select file</button>
