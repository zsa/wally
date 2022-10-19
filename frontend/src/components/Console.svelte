<script lang="ts">
  import state from "../lib/state";
</script>

<div class="console">
  <div class="lines">
    {#each $state.logs as log}
      <div class="line">
        <div class="timestamp">
          {new Date(log.timestamp * 1000).toTimeString().substr(0, 8)}
        </div>
        <div class={`level ${log.level}`}>
          <div class="message">{log.level}</div>
        </div>
        <div class="message">{log.message}</div>
      </div>
    {/each}
  </div>
</div>

<style>
  .console {
    background: #18262f;
    color: #a6afb8;
    overflow-y: auto;
    font-family: monospace;
    font-size: 0.8em;
    flex: 1;
    padding: 10px 20px;
  }

  .console .lines {
    display: flex;
    flex-direction: column;
  }

  .console .lines .line {
    display: flex;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  }

  .console .lines .line .timestamp {
    color: #a363d5;
    width: 70px;
  }

  .console .lines .line .level {
    text-align: center;
    width: 70px;
  }

  .console .lines .line .level.info {
    color: #33b5e1;
  }

  .console .lines .line .level.fatal {
    color: #ef5253;
  }

  .console .lines .line .level.warning {
    color: #ff6200;
  }

  .console .lines .line .message {
    flex: 1;
  }
</style>
