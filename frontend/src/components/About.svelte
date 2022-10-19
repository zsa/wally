<script lang="ts">
  import state from "../lib/state";
  function toggleAbout() {
    $state.showAbout = !$state.showAbout;
    $state.showConsole = false;
  }
</script>

<div class="about">
  <h1>Wally</h1>
  <p class="version">Version {$state.appVersion}</p>
  <p>
    The official flashing tool for <a href="https://zsa.io" target="_blank"
      >ZSA</a
    > keyboards.
  </p>
  <p>
    This software is licensed under the <a href="">MIT Licence</a>.<br />Source
    code is available on <a href="">Github</a>.
  </p>
  <p>
    Visit <a href="https://zsa.io">zsa.io</a> for more info, and email
    <a href="mailto:contact@zsa.io">contact@zsa.io</a> with any questions.
  </p>
  <p class="check">
    <input type="checkbox" name="check" /> Automatically check for updates.
  </p>
  <div class="close">
    <button class="btn small" on:click={toggleAbout}>Close</button>
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
