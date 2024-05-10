<script lang="ts">
  import cover from "$lib/assets/cover.jpg";
</script>

<div class="cover-container">
  <img class="cover-image" alt="" src={cover} />
  <div class="cover-text">
    <div class="cover-names">
      <div>ALEX</div>
      <div>&</div>
      <div>LYNN</div>
    </div>
    <div class="cover-date">October 11th, 2024 | Buffalo, NY</div>
  </div>
</div>

<style>
  .cover-container {
    width: 100%;
  }

  .cover-image {
    width: 100%;
  }

  .cover-text {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }

  .cover-names {
    font-size: xx-large;
  }
</style>
