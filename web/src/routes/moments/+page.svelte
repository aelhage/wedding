<script lang="ts">
  import Active from "./active.svelte";

  const startDate = new Date(2024, 10, 10);
  const endDate = new Date(2024, 10, 12, 12);
  const currentDate = new Date();
</script>

{#if currentDate > startDate}
  <Active allowUpload={currentDate < endDate} />
{:else}
  <p>
    Visit this page during the celebration to share your photos with all of our
    guests!
  </p>
{/if}
