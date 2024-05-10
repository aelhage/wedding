<script lang="ts">
  import download from "$lib/assets/download.svg";
  import upload from "$lib/assets/upload.svg";

  export let index: number;
  export let items: string[];
  export let close: () => void;

  $: hasNext = index + 1 < items?.length;
  $: hasPrev = index > 0;

  const next = () => {
    index = index + 1;
  };
  const prev = () => {
    index = index - 1;
  };
</script>

<div class="feed-root">
  <div class="feed-header">
    {index} of {items.length}
  </div>
  <button
    class="feed-nav feed-nav-prev {hasPrev ? '' : 'hidden'}"
    type="button"
    on:click={prev}
  >
    &lt;
  </button>
  <button
    class="feed-nav feed-nav-next {hasNext ? '' : 'hidden'}"
    type="button"
    on:click={next}
  >
    &gt;
  </button>
  <button type="button" class="feed-close" on:click={close}> X </button>
  <a class="feed-download" href={items[index]} download>
    <img alt="save this moment" src={download} />
  </a>
  <img class="feed-content" src={items[index]} alt="selected moment" />
</div>

<style>
  .hidden {
    visibility: hidden;
  }

  .feed-root {
    display: grid;
    grid-template-columns: 50% 50%;
    grid-template-rows: 25px 1fr 50px;
  }

  .feed-header {
    grid-column: 1 / span 2;
    grid-row: 1;
  }

  .feed-nav {
    border: none;
    background: transparent;
  }

  .feed-nav-prev {
    grid-column: 1;
    grid-row: 2;
  }

  .feed-nav-next {
    grid-column: 2;
    grid-row: 2;
  }

  .feed-close {
    grid-column: 1;
    grid-row: 3;
    cursor: pointer;
    border: none;
    background: transparent;
  }

  .feed-download img {
    max-width: 50%;
    max-height: 50%;
  }

  .feed-download {
    grid-column: 2;
    grid-row: 3;
  }

  .feed-content {
    max-width: 100%;
    max-height: 100%;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }
</style>
