<script lang="ts">
  import state from "../lib/state";
  import { SetUpdateCheck } from "../../wailsjs/go/state/State";
  import A from "./BrowserLink.svelte";

  let checked = false;
  function toggleAbout() {
    $state.showAbout = !$state.showAbout;
    $state.showConsole = false;
  }

  function toggleCheck() {
    checked = !checked;
    SetUpdateCheck(checked);
  }
</script>

<div class="about">
  <h1>Wally</h1>
  <p class="version">Version {$state.appVersion}</p>
  <p>
    The official flashing tool for <A href="https://zsa.io">ZSA</A> keyboards.
  </p>
  <p>
    This software is licensed under the <A
      href="https://github.com/zsa/wally/blob/master/license.md">MIT Licence</A
    >.<br />Source code is available on <A href="https://github.com/zsa/wally"
      >Github</A
    >.
  </p>
  <p>
    Visit <A href="https://zsa.io">zsa.io</A> for more info, and email
    <A href="mailto:contact@zsa.io">contact@zsa.io</A> with any questions.
  </p>
  <p class="check">
    <input
      type="checkbox"
      name="check"
      {checked}
      on:click|preventDefault={toggleCheck}
    /> Automatically check for updates.
  </p>
  <div class="close">
    <button
      class="btn small"
      name="update_check"
      value={1}
      on:click={toggleAbout}>Close</button
    >
  </div>
</div>

<style>
  .about {
    background-color: white;
    border-radius: 5px;
    position: fixed;
    left: 30px;
    right: 30px;
    top: 70px;
    padding: 20px;
  }
  p {
    font-size: 16px;
    margin: 1em 0 0 0;
    padding: 0;
  }
  .version {
    opacity: 0.6;
    font-size: 0.9em;
  }
  .check {
    align-items: center;
    display: flex;
  }
  .check input {
    margin-right: 8px;
  }
  .close {
    display: flex;
    flex-direction: row-reverse;
  }
</style>
