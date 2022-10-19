<script lang="ts">
  import state, { Step } from "../lib/state";
  import FatalError from "./steps/FatalError.svelte";
  import FirmwareFlashing from "./steps/FirmwareFlashing.svelte";
  import FirmwareSelect from "./steps/FirmwareSelect.svelte";
  import FlashingComplete from "./steps/FlashingComplete.svelte";
  import KeyboardReset from "./steps/KeyboardReset.svelte";
  import KeyboardSelect from "./steps/KeyboardSelect.svelte";
  import Probing from "./steps/Probing.svelte";
  import WallyUpate from "./steps/WallyUpate.svelte";
  import WallyUpateComplete from "./steps/WallyUpdateComplete.svelte";
</script>

<div class="steps">
  {#if $state.step == Step.PROBING}
    <Probing />
  {:else if $state.step == Step.KEYBOARD_SELECT}
    <KeyboardSelect />
  {:else if $state.step == Step.FIRMWARE_SELECT}
    <FirmwareSelect />
  {:else if $state.step == Step.KEYBOARD_RESET}
    <KeyboardReset />
  {:else if $state.step == Step.FIRMWARE_FLASHING}
    <FirmwareFlashing />
  {:else if $state.step == Step.FLASH_COMPLETE}
    <FlashingComplete />
  {:else if $state.step == Step.FATAL_ERROR}
    <FatalError />
  {:else if $state.step == Step.WALLY_UPDATE}
    <WallyUpate />
  {:else if $state.step == Step.WALLY_UPDATE_COMPLETE}
    <WallyUpateComplete />
  {/if}
</div>

<style>
  .steps {
    flex: 1;
    padding: 20px 10px;
    text-align: center;
  }
</style>
