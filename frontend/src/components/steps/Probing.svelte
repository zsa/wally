<script lang="ts">
  import { onMount } from "svelte";
  import state, { Step } from "../../lib/state";
  import { SetStep, SelectDevice } from "../../../wailsjs/go/state/State";

  const TICKS_THRESHOLD = 3;

  let ticks: number = 0;

  let ticker = null;

  onMount(() => {
    ticker = setInterval(() => {
      ticks += 1;
      if ($state.devices.size == 1) {
        const [fingerprint] = $state.devices.keys();
        SelectDevice(fingerprint);
        SetStep(Step.FIRMWARE_SELECT);
      }
      if ($state.devices.size > 1) {
        SetStep(Step.KEYBOARD_SELECT);
      }
    }, 1000);

    return () => {
      clearInterval(ticker);
    };
  });
</script>

<h1>Looking for compatible keyboards</h1>

{#if ticks < TICKS_THRESHOLD}
  <p>Hold on while your keyboard is being detected</p>
{:else}
  <p>Make sure your keyboard is firmly connected</p>
{/if}
