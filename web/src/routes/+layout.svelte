<script>
  import { page } from "$app/stores";
</script>

<div class="app-root">
  <div class="app-header">
    <span>A & L</span>
    <nav>
      <a href="/" class:active={$page.url.pathname === "/"}>Home</a>
      <a href="/schedule" class:active={$page.url.pathname === "/schedule"}
        >Schedule</a
      >
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

  <slot />
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
    min-height: 100vh;
    color: var(--gold);
    background-color: var(--emerald);
  }

  .app-root {
    width: 100vw;
    max-width: 800px;
    text-align: center;
    position: absolute;
  }

  .app-header {
    margin: 16px;
  }

  .app-header > nav {
    display: inline;
    float: right;
  }

  .app-header > nav > * {
    margin: 0 4px;
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
