<script lang="ts">
  import state, { Step } from "../lib/state";

  import Probing from "../img/search.svg";
  import FirmwareSelect from "../img/file.svg";
  import KeyboardReset from "../img/search.svg";
  import FirmwareFlashing from "../img/flash.svg";
  import FlashingComplete from "../img/complete.svg";
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
      <img alt="Reset keyboard icon" src={KeyboardReset} />
    {:else if $state.step == Step.FIRMWARE_FLASHING}
      <img alt="Flashing in progress icon" src={FirmwareFlashing} />
    {:else if $state.step == Step.FLASH_COMPLETE}
      <img alt="Flashing complete icon" src={FlashingComplete} />
    {:else if $state.step == Step.FATAL_ERROR}
      <img alt="Error icon" src={Error} />
    {:else if $state.step == Step.WALLY_UPDATE}
      <svg
        width="120"
        height="120"
        viewBox="0 0 120 120"
        xmlns="http://www.w3.org/2000/svg"
        ><style>
          .spinner_7NYg {
            animation: spinner_0KQs 1.2s cubic-bezier(0.52, 0.6, 0.25, 0.99)
              infinite;
          }
          @keyframes spinner_0KQs {
            0% {
              r: 20px;
              opacity: 1;
            }
            100% {
              r: 40px;
              opacity: 0;
            }
          }
        </style><circle
          class="spinner_7NYg"
          fill="white"
          cx="60"
          cy="60"
          r="10"
        /></svg
      >
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
</style>
