<script lang="ts">
  import state from "../../lib/state";
  import ProgressBar from "../ProgressBar.svelte";
</script>

<h3>Flashing firmware</h3>
<p>Please don’t unplug your board.</p>
<ProgressBar progress={$state.flashProgress} />
