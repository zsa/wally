<script type="ts">
  import { BrowserOpenURL } from "../../wailsjs/runtime/";
  export let href = "";

  function handleClick() {
    BrowserOpenURL(href);
  }
</script>

<button class="button-link" on:click={handleClick}> <slot /></button>

<style>
  .button-link {
    appearance: none;
    background: none;
    border: none;
    cursor: pointer;
    font-size: 1em;
    margin: 0;
    padding: 0;
  }
</style>
