<script lang="ts">
  import downloadIcon from "$lib/assets/download.svg";
  import closeIcon from "$lib/assets/close.svg";
  import nextIcon from "$lib/assets/next.svg";
  import prevIcon from "$lib/assets/previous.svg";

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

<div class="carousel-root">
  <img class="carousel-content" src={items[index]} alt="selected moment" />
  <div class="carousel-controls">
    <button
      class="carousel-prev {hasPrev ? '' : 'hidden'}"
      type="button"
      on:click={prev}
    >
      <img alt="next" src={prevIcon} />
    </button>
    <button
      class="carousel-next {hasNext ? '' : 'hidden'}"
      type="button"
      on:click={next}
    >
      <img alt="next" src={nextIcon} />
    </button>
    <button type="button" class="carousel-close" on:click={close}>
      <img alt="close carousel" src={closeIcon} />
    </button>
    <a class="carousel-download" href={items[index]} download>
      <img alt="save this moment" src={downloadIcon} />
    </a>
  </div>
</div>

<style>
  .hidden {
    visibility: hidden;
  }

  .carousel-root {
    position: fixed;
    top: 0px;
    left: 0px;
    width: 100vw;
    height: 100vh;
    background-color: color-mix(in srgb, var(--sapphire) 75%, transparent);
  }

  .carousel-controls {
    display: grid;
    grid-template-columns: 50% 50%;
    grid-template-rows: 50px 1fr 50px;
    width: 100%;
    height: 100%;
    position: absolute;
  }

  .carousel-controls button {
    color: var(--gold);
    cursor: pointer;
    border: none;
    background: transparent;
  }

  .carousel-controls button img {
    height: 24px;
  }

  .carousel-prev {
    grid-column: 1;
    grid-row: 2;
  }

  .carousel-next {
    grid-column: 2;
    grid-row: 2;
  }

  .carousel-close {
    grid-column: 1 / span 2;
    grid-row: 1;
  }

  .carousel-close img {
    float: right;
    margin: 8px;
  }

  .carousel-download {
    grid-column: 1 / span 2;
    grid-row: 3;
  }

  .carousel-download img {
    height: 24px;
  }

  .carousel-content {
    max-width: 100%;
    max-height: 100%;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }
</style>
