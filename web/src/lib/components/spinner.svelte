<script lang="ts">
  export let duration: string = "0.75s";
</script>

<div class="circle" style="--duration: {duration}" />

<style>
  .circle {
    height: 60px;
    width: 60px;
    border-color: var(--gold) transparent var(--gold) var(--gold);
    border-width: 4px;
    border-style: solid;
    border-image: initial;
    border-radius: 50%;
    animation: var(--duration) linear 0s infinite normal none running rotate;
  }
  @keyframes rotate {
    0% {
      transform: rotate(0);
    }
    100% {
      transform: rotate(360deg);
    }
  }
</style>
