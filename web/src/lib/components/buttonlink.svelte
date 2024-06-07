<script lang="ts">
  export let href: string;
  export let content: string;
</script>

<a {href}>{content}</a>

<style>
  a {
    display: inline-block;
    text-decoration: none;
    text-align: center;
    background-color: var(--gold);
    color: var(--sapphire);
    padding: 12px;
  }
</style>
