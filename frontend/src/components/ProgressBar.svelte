<script lang="ts">
  export let progress = 0;
</script>

<div class="rail">
  <div class="progress" style={`width: ${progress}%`} />
</div>

<style>
  .rail {
    width: 80%;
    margin: 20px auto;
    border-radius: 88px;
    height: 20px;
    padding: 2px;
    background-color: rgba(0, 0, 0, 0.1);
  }

  .rail .progress {
    border-radius: 88px;
    height: 100%;
    background-color: #ffffff;
    transition: width 200ms ease-in-out;
  }
</style>
