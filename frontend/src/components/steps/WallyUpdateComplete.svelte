<script lang="ts">
  import { InstallUpdate } from "../../../wailsjs/go/state/State";
  $: loading = false;

  const handleClick = () => {
    loading: true;
    InstallUpdate();
  };
</script>

<h1>Download complete</h1>
<p>Press the install button below to start the install</p>
{#if loading}
  <button class="btn" disabled>Opening update</button>
{:else}
  <button class="btn" on:click={handleClick}>Install</button>
{/if}
