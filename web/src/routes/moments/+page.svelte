<script lang="ts">
  import { PUBLIC_BASE_API_URL } from "$env/static/public";
  import upload from "$lib/assets/upload.svg";
  import download from "$lib/assets/download.svg";

  // Data Loading

  let uploadURL: string;
  let items: string[];
  let vmPromise = fetch(PUBLIC_BASE_API_URL).then(async (resp) => {
    if (!resp.ok) {
      throw "failed to retrieve feed";
    }
    let body = await resp.json();
    uploadURL = body["upload_url"];
    items = [...body["download_urls"]];
  });

  // File Upload

  let fileInput: HTMLInputElement;
  let files: FileList;

  $: if (files && files[0]) {
    fetch(uploadURL, {
      method: "PUT",
      body: files[0],
    }).then((resp) => {
      if (!resp.ok) {
        throw "failed to post";
      }
      fileInput.value = "";
    });
  }

  // UI Management

  let index = 0;
  $: hasNext = index + 1 < items?.length;
  $: hasPrev = index > 0;
  $: hasAnything = items?.length > 0;

  const handleNext = () => {
    index = hasNext ? index + 1 : index;
  };
  const handlePrev = () => {
    index = hasPrev ? index - 1 : index;
  };
</script>

{#await vmPromise then}
  <div class="feed">
    <input
      bind:this={fileInput}
      id="feed-input"
      type="file"
      accept="image/*"
      class="display-none"
      bind:files
    />
    {#if hasAnything}
      <div class="feed-header">
        {index} of {items.length}
      </div>
      <img class="feed-content" src={items[index]} alt="selected moment" />
      <button
        class="feed-nav feed-nav-prev {hasPrev ? '' : 'hidden'}"
        type="button"
        on:click={handlePrev}
      >
        &lt;
      </button>
      <button
        class="feed-nav feed-nav-next {hasNext ? '' : 'hidden'}"
        type="button"
        on:click={handleNext}
      >
        &gt;
      </button>
      <button type="button" class="feed-upload" on:click={fileInput.click}>
        <img alt="share your moment" src={upload} />
      </button>
      <a class="feed-download" href={items[index]} download>
        <img alt="save this moment" src={download} />
      </a>
    {:else}
      <button
        type="button"
        class="feed-upload feed-upload-first"
        on:click={fileInput.click}
      >
        <img alt="share your moment" src={upload} />
      </button>
    {/if}
  </div>
{/await}

<style>
  .display-none {
    display: none !important;
  }

  .hidden {
    visibility: hidden;
  }

  .feed {
    display: grid;
    grid-template-columns: 50% 50%;
    grid-template-rows: 25px 1fr 50px;
  }

  .feed .feed-header {
    grid-column: 1 / span 2;
    grid-row: 1;
  }

  .feed .feed-nav {
    border: none;
    background: transparent;
    z-index: 10;
  }

  .feed .feed-nav-prev {
    grid-column: 1;
    grid-row: 2;
  }

  .feed .feed-nav-next {
    grid-column: 2;
    grid-row: 2;
  }

  .feed .feed-upload {
    grid-column: 1;
    grid-row: 3;
    cursor: pointer;
    border: none;
    background: transparent;
  }

  .feed .feed-upload img {
    max-width: 50%;
    max-height: 50%;
  }

  .feed .feed-download img {
    max-width: 50%;
    max-height: 50%;
  }

  .feed .feed-upload-first {
    grid-column: 1 / span 2;
    grid-row: 2;
  }

  .feed .feed-download {
    grid-column: 2;
    grid-row: 3;
  }
  .feed .feed-content {
    max-width: 100%;
    max-height: 100%;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }
</style>
