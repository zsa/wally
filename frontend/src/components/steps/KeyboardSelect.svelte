<script type="ts">
  import state from "../../lib/state";
  import { SelectDevice } from "../../../wailsjs/go/state/State";
  import Planck from "../../img/planck-logo.png";
  import Ergodox from "../../img/ergodox-logo.png";
  import Moonlander from "../../img/moonlander-logo.png";

  function handleDeviceClick(fingerprint: number) {
    SelectDevice(fingerprint);
  }
</script>

<div class="device-select">
  <div>
    <div class="devices">
      {#each [...$state.devices] as [fingerprint, device]}
        <button
          class="device unstyled"
          on:click={() => {
            handleDeviceClick(device.fingerprint);
          }}
        >
          {#if device.model == "planck"}
            <img src={Planck} alt={device.friendlyName} />
          {/if}
          {#if device.model == "ergodox"}
            <img src={Ergodox} alt={device.friendlyName} />
          {/if}
          {#if device.model == "moonlander"}
            <img src={Moonlander} alt={device.friendlyName} />
          {/if}
        </button>
      {/each}
    </div>

    <h1>Select keyboard</h1>
    <p>
      There are several keyboards connected that are compatible, please select
      one.
    </p>
  </div>
</div>

<style>
  .device-select {
    display: flex;
    align-items: center;
    height: 100%;
  }
  .devices {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 30px 0px 24px 0px;
    justify-content: space-around;
  }
  .device img {
    max-width: 128px;
  }
  .device img:hover {
    opacity: 0.8;
  }
</style>
