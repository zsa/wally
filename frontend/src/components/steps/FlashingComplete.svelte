<script type="ts">
  import { Reset, Quit } from "../../../wailsjs/go/state/State";
</script>

<h1>All done!</h1>
<p>
  Your keyboard was successfully flashed and rebooted. <br />
  Enjoy the new firmware!
</p>
<button class="btn" role="button" on:click={Reset}> Flash again </button>
<button class="btn" role="button" on:click={Quit}> Close </button>
