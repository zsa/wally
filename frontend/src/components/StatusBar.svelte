<script lang="ts">
  import state, { Step } from "../lib/state";
  import ZSALogo from "../img/zsa-logo.png";
  import { Reset } from "../../wailsjs/go/state/State";
  function toggleConsole() {
    $state.showConsole = !$state.showConsole;
    $state.showAbout = false;
  }

  function toggleAbout() {
    $state.showAbout = !$state.showAbout;
    $state.showConsole = false;
  }

  function handleReset() {
    Reset();
  }

  $: showReset = !!(
    $state.step != Step.PROBING &&
    $state.step !== Step.FIRMWARE_FLASHING &&
    $state.step !== Step.WALLY_UPDATE
  );
</script>

<footer class="statusbar">
  <div class="title">KEYBOARD:</div>
  <div class="status">
    <div class="model">
      {#if $state.step == Step.PROBING}
        LOOKING...
      {:else if $state.step == Step.KEYBOARD_SELECT}
        -SELECT-
      {:else if $state.selectedDevice}
        {$state.selectedDevice.friendlyName}
      {:else}
        -
      {/if}
    </div>
    {#if showReset}
      <button class="unstyled reset" on:click={handleReset}>Restart</button>
    {/if}
  </div>
  <div class="log-toggle" on:click={toggleConsole}>
    >_
    {#if $state.hasError}
      <span class="bubble">!</span>
    {/if}
  </div>
  <div>
    <img alt="ZSA" class="logo zsa-logo" src={ZSALogo} />
  </div>
  <button class="unstyled version" on:click={toggleAbout}
    >V{$state.appVersion}</button
  >
</footer>

<style>
  .statusbar {
    height: 40px;
    width: 100vw;
    font-size: 0.9em;
    background-color: #393838;
    display: flex;
    align-items: center;
    padding: 0px 16px;
  }
  .statusbar .title {
    color: #b2b2b2;
  }

  .statusbar .status {
    align-items: center;
    color: #ffffff;
    display: flex;
    flex: 1;
    font-weight: bold;
    text-transform: uppercase;
  }

  .statusbar .status .model {
    margin: 0 10px;
  }

  .statusbar .status .reset {
    color: #fff;
    text-decoration: underline;
  }

  .statusbar .logo {
    position: relative;
    top: 1px;
  }

  .statusbar .zsa-logo {
    top: 3px;
    height: 25px;
  }

  .statusbar .log-toggle {
    background-color: rgba(0, 0, 0, 0.1);
    border-radius: 2px;
    color: white;
    cursor: pointer;
    font-family: monospace;
    font-size: 1.2em;
    font-weight: bold;
    margin-right: 20px;
    padding: 1px 5px 7px 5px;
    position: relative;
  }

  .statusbar .log-toggle .bubble {
    position: absolute;
    top: 0px;
    right: -8px;
    height: 16px;
    width: 16px;
    background-color: #ef5253;
    color: #fff;
    font-size: 0.65em;
    line-height: 1.8em;
    text-align: center;
    border-radius: 50%;
  }

  .statusbar .log-toggle:hover {
    background-color: rgba(0, 0, 0, 0.2);
  }

  .statusbar .version {
    color: #ffee97;
    text-decoration: underline;
  }
</style>
