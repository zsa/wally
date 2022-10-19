<section class="pills">
  <div class="pill active" />
  <div class="pill" />
  <div class="pill" />
  <div class="pill" />
</section>

<style>
  .pills {
    align-items: center;
    display: flex;
    justify-content: center;
    padding: 20px 0px;
    height: 40px;
  }

  .pills .pill {
    background: rgba(0, 0, 0, 0.25);
    border-radius: 10px;
    height: 10px;
    margin: 0 8px;
    transform: scale(0.75);
    transition: transform 0.4s cubic-bezier(0.25, 0.8, 0.25, 1),
      background 0.4s cubic-bezier(0.25, 0.8, 0.25, 1);
    width: 10px;
  }

  .pills .pill.active {
    background: white;
    transform: scale(1.2);
    transition-delay: 0.4s;
  }
</style>
