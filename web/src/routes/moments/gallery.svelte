<script lang="ts">
  import upload from "$lib/assets/upload.svg";

  export let items: string[];
  export let add: () => void;
  export let select: (index: number) => void;
</script>

<ul>
  <li>
    <button type="button" class="moments-upload" on:click={add}>
      <img alt="share your moment" src={upload} />
    </button>
  </li>
  {#each items as item, i}
    <li>
      <img src={item} alt="" on:click={() => select(i)} />
    </li>
  {/each}
</ul>

<style>
  ul {
    list-style: none;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    gap: 8px;
  }

  ul li {
    height: 200px;
    position: relative;
    cursor: pointer;
  }

  ul li img {
    object-fit: cover;
    width: 100%;
    height: 100%;
    vertical-align: middle;
  }

  .moments-upload {
    background: transparent;
    border: none;
    cursor: pointer;
    width: 100%;
    height: 100%;
  }
</style>
