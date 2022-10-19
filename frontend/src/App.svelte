<script lang="ts">
  import { onMount } from "svelte";
  import state, { Step } from "./lib/state";
  import { GetAppVersion, StateStart } from "../wailsjs/go/state/State";
  import Header from "./components/Header.svelte";
  import Steps from "./components/Steps.svelte";
  import Pills from "./components/Pills.svelte";
  import About from "./components/About.svelte";
  import Console from "./components/Console.svelte";
  import StatusBar from "./components/StatusBar.svelte";

  $: showPills = !!(
    $state.step != Step.PROBING &&
    $state.step != Step.FLASH_COMPLETE &&
    $state.step != Step.WALLY_UPDATE &&
    $state.step != Step.WALLY_UPDATE_COMPLETE
  );

  onMount(async () => {
    StateStart();
    $state.appVersion = await GetAppVersion();
  });
</script>

{#if $state.showConsole}
  <Console />
{:else}
  <Header />
  <Steps />
  {#if showPills}<Pills />{/if}
{/if}

{#if $state.showAbout}
  <About />
{/if}

<StatusBar />
