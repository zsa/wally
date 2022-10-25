<script lang="ts">
  import { onMount } from "svelte";
  import state, { Step } from "./lib/state";
  import {
    CheckUpdate,
    GetUpdateCheck,
    GetAppVersion,
    InitUSB,
    PromptUpdates,
  } from "../wailsjs/go/state/State";
  import Header from "./components/Header.svelte";
  import Steps from "./components/Steps.svelte";
  import Pills from "./components/Pills.svelte";
  import About from "./components/About.svelte";
  import Console from "./components/Console.svelte";
  import StatusBar from "./components/StatusBar.svelte";

  $: showPills = !!(
    $state.step != Step.PROBING &&
    $state.step != Step.KEYBOARD_SELECT &&
    $state.step != Step.FLASH_COMPLETE &&
    $state.step != Step.WALLY_UPDATE &&
    $state.step != Step.WALLY_UPDATE_COMPLETE
  );

  onMount(async () => {
    InitUSB();
    $state.appVersion = await GetAppVersion();
    await PromptUpdates();
    $state.checkUpdates = await GetUpdateCheck();
    if ($state.checkUpdates) {
      CheckUpdate();
    }
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
