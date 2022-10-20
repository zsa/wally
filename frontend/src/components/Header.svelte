<script lang="ts">
  import state, { Step } from "../lib/state";

  import Probing from "../img/search.svg";
  import FirmwareSelect from "../img/file.svg";
  import PlanckReset from "../img/planck-reset.png";
  import ErgodoxReset from "../img/ergodox-reset.png";
  import MoonlanderReset from "../img/moonlander-reset.png";
  import FirmwareFlashing from "../img/flash.svg";
  import FlashingComplete from "../img/complete.svg";
  import Pulse from "../img/pulse.svg";
  import Error from "../img/error.png";
  $: hasHeader = $state.step !== Step.KEYBOARD_SELECT;
</script>

<!-- The Keyboard selection / Wally update views don't require a header icon -->
{#if hasHeader}
  <section class="header">
    {#if $state.step == Step.PROBING}
      <img alt="Searching for keyboards icon" src={Probing} />
    {:else if $state.step == Step.FIRMWARE_SELECT}
      <img alt="Select a firmware icon" src={FirmwareSelect} />
    {:else if $state.step == Step.KEYBOARD_RESET}
      {#if $state.selectedDevice}
        {#if $state.selectedDevice.model == "planck"}
          <img alt="Reset Planck icon" src={PlanckReset} />
        {/if}
        {#if $state.selectedDevice.model == "ergodox"}
          <img alt="Reset Ergodox icon" class="offset" src={ErgodoxReset} />
        {/if}
        {#if $state.selectedDevice.model == "moonlander"}
          <img alt="Reset Moonlander icon" src={MoonlanderReset} />
        {/if}
      {/if}
    {:else if $state.step == Step.FIRMWARE_FLASHING}
      <img alt="Flashing in progress icon" src={FirmwareFlashing} />
    {:else if $state.step == Step.FLASH_COMPLETE}
      <img alt="Flashing complete icon" src={FlashingComplete} />
    {:else if $state.step == Step.FATAL_ERROR}
      <img alt="Error icon" src={Error} />
    {:else if $state.step == Step.WALLY_UPDATE}
      <img alt="Loading icon" src={Pulse} />
    {:else if $state.step == Step.WALLY_UPDATE_COMPLETE}
      <img alt="Update complete icon" src={FlashingComplete} />
    {/if}
  </section>
{/if}

<style>
  .header {
    height: 45%;
    display: flex;
    align-items: flex-end;
    justify-content: center;
  }
  .header img {
    width: 120px;
  }
  .header img.offset {
    width: 180px;
    margin-left: 60px;
  }
</style>
