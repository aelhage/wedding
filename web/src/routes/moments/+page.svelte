<script lang="ts">
  import { PUBLIC_BASE_API_URL } from "$env/static/public";
  let vmPromise = fetch(PUBLIC_BASE_API_URL);
</script>

{#await vmPromise then vm}
  <pre>
    {JSON.stringify(vm, null, 2)}
</pre>
{/await}
