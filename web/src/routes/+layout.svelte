<script>
  import { page } from "$app/stores";
</script>

<div class="app-root">
  <div class="app-header">
    <span>A & L</span>
    <nav>
      <a href="/" class:active={$page.url.pathname === "/"}>Home</a>
      <a href="/day" class:active={$page.url.pathname === "/day"}>Day Of</a>
      <a href="/travel" class:active={$page.url.pathname === "/travel"}
        >Travel</a
      >
      <a href="/rsvp" class:active={$page.url.pathname === "/rsvp"}>RSVP</a>
      <a href="/registry" class:active={$page.url.pathname === "/registry"}
        >Registry</a
      >
      <a href="/moments" class:active={$page.url.pathname === "/moments"}
        >Moments</a
      >
    </nav>
  </div>
  <div class="app-content">
    <slot />
  </div>
</div>

<style>
  :root {
    --gold: #ab8d6c;
    --emerald: #24341f;
    --amethyst: #39023c;
    --ruby: #610629;
    --sapphire: #0d0e38;
  }

  :global(body) {
    margin: auto;
    padding: 0;
    width: 100vh;
    color: var(--gold);
    background-color: var(--emerald);
  }

  .app-root {
    width: 90vw;
    max-width: 800px;
    text-align: center;
    position: absolute;
    height: 100%;
    display: grid;
    grid-template-rows: auto 1fr;
    row-gap: 16px;
    margin: 16px;
  }

  .app-header {
    grid-row: 1;
  }

  .app-header > nav {
    display: inline;
    float: right;
  }

  .app-header > nav > * {
    margin: 0 4px;
  }

  .app-content {
    grid-row: 2;
  }

  a {
    text-decoration: none;
    color: var(--gold);
  }

  .active {
    text-decoration: underline;
    text-underline-offset: 8px;
  }
</style>
