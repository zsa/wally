<script typoe="ts">
  import state from "../../lib/state";
</script>

{#if $state.selectedDevice}
  <h1>Press your keyboard's reset button</h1>
  {#if $state.selectedDevice.model == "planck"}
    <p>
      You’re going to need a paperclip for this: The reset button is located at
      the top left of the back of your keyboard.
    </p>
  {/if}
  {#if $state.selectedDevice.model == "ergodox"}
    <p>
      You’re going to need a paperclip for this: The reset button is located on
      the right half of your keyboard, next to the three LEDs.
    </p>
  {/if}
  {#if $state.selectedDevice.model == "moonlander"}
    <p>
      You’re going to need a paperclip for this: The reset button is located on
      the left half of your keyboard, next to the three LEDs.
    </p>
  {/if}
{/if}
