<script lang="ts">
  import download from "$lib/assets/download.svg";
  import upload from "$lib/assets/upload.svg";

  export let index: number;
  export let items: string[];
  export let add: () => void;

  $: hasNext = index + 1 < items?.length;
  $: hasPrev = index > 0;

  const next = () => {
    index = index + 1;
  };
  const prev = () => {
    index = index - 1;
  };
</script>

<div class="feed-header">
  {index} of {items.length}
</div>
<img class="feed-content" src={items[index]} alt="selected moment" />
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
<button type="button" class="feed-upload" on:click={add}>
  <img alt="share your moment" src={upload} />
</button>
<a class="feed-download" href={items[index]} download>
  <img alt="save this moment" src={download} />
</a>

<style>
  .hidden {
    visibility: hidden;
  }

  .feed-header {
    grid-column: 1 / span 2;
    grid-row: 1;
  }

  .feed-nav {
    border: none;
    background: transparent;
    z-index: 10;
  }

  .feed-nav-prev {
    grid-column: 1;
    grid-row: 2;
  }

  .feed-nav-next {
    grid-column: 2;
    grid-row: 2;
  }

  .feed-upload {
    grid-column: 1;
    grid-row: 3;
    cursor: pointer;
    border: none;
    background: transparent;
  }

  .feed-upload img {
    max-width: 50%;
    max-height: 50%;
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
