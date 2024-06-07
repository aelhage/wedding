<script lang="ts">
  import { PUBLIC_BASE_API_URL } from "$env/static/public";
  import upload from "$lib/assets/upload.svg";
  import Carousel from "./carousel.svelte";
  import Gallery from "./gallery.svelte";

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

  let index = -1;
  $: hasAnything = items?.length > 0;

  const add = () => {
    fileInput.click();
  };
</script>

{#await vmPromise then}
  <div class="moments-root">
    <input
      bind:this={fileInput}
      id="moments-input"
      type="file"
      accept="image/*"
      class="display-none"
      capture
      bind:files
    />
    {#if hasAnything}
      {#if index >= 0}
        <Carousel bind:index {items} close={() => (index = -1)} />
      {:else}
        <Gallery {items} {add} select={(i) => (index = i)} />
      {/if}
    {:else}
      <button type="button" class="moments-upload-first" on:click={add}>
        <img alt="share your moment" src={upload} />
      </button>
    {/if}
  </div>
{/await}

<style>
  .display-none {
    display: none !important;
  }

  .moments-upload-first {
    grid-column: 1 / span 2;
    grid-row: 2;
    cursor: pointer;
    border: none;
    background: transparent;
  }

  .moments-upload-first img {
    max-width: 50%;
    max-height: 50%;
  }
</style>
