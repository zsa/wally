<script type="ts">
  import state from "../../lib/state";
  import { logsToString } from "../../lib/utils";
  import A from "../BrowserLink.svelte";
  const COPY_INITIAL_TEXT = "copy the logs";
  let copyText = COPY_INITIAL_TEXT;

  const handleCopy = () => {
    if (navigator.clipboard) {
      navigator.clipboard.writeText(logsToString($state.logs));
      copyText = "copied!";
    } else {
      copyText = "copy error :(";
    }
    window.setTimeout(() => (copyText = COPY_INITIAL_TEXT), 3000);
  };
</script>

<h1>Something went wrong:</h1>
<p>
  You can retry the flashing proces from the start. If Wally doesn't recognize
  your keyboard, make sure you press the reset button with a paperclip first.
</p>
<p>
  If the error persists please <a
    href="#logcopy"
    on:click|preventDefault={handleCopy}>{copyText}</a
  >
  and send them to <A href="mailto:contact@zsa.io">contact@zsa.io</A>
</p>

<style>
  p {
    text-align: left;
  }
  a {
    color: #403c3a;
  }
</style>
