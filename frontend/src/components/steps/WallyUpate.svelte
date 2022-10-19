<script lang="ts">
  import state from "../../lib/state";
  import ProgressBar from "../ProgressBar.svelte";
</script>

<h1>Downloading update</h1>
<ProgressBar progress={$state.updateProgress} />
