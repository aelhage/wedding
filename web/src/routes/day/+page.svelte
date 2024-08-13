<script>
  import time from "$lib/assets/time.svg";
  import location from "$lib/assets/location.svg";
  import museum from "$lib/assets/museum.jpg";
  import trinity from "$lib/assets/trinity.png";
</script>

<div class="root">
  <h1>Friday, October 11th, 2024</h1>

  <h2>Ceremony</h2>
  <img class="header" alt="trinity" src={trinity} />
  <h3>Trinity Episcopal Church</h3>
  <div>
    <img class="icon" alt="ceremony time" src={time} />
    <address>3:30 pm</address>
  </div>
  <div>
    <img class="icon" alt="ceremony location" src={location} />
    <address>371 Deleware Avenue, Buffalo, NY 14202</address>
  </div>
  <p>
    Parking lot access on Franklin Street at rear of church. Disregard any signs
    to pay.
  </p>

  <h2>Reception</h2>
  <img class="header" alt="museum" src={museum} />
  <h3>The Buffalo History Museum</h3>
  <div>
    <img class="icon" alt="reception time" src={time} />
    <address>6:00 - 11:00 pm</address>
  </div>
  <div>
    <img class="icon" alt="reception location" src={location} />
    <address>1 Museum Court Buffalo, NY 14202</address>
  </div>
</div>

<style>
  address {
    white-space: pre;
    display: inline-block;
  }
  .root {
    width: 100%;
  }
  .icon {
    width: 24px;
    height: 24px;
  }
  .header {
    max-width: 100%;
  }
</style>
