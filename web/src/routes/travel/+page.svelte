<script>
  import Buttonlink from "$lib/components/buttonlink.svelte";
  import richardson from "$lib/assets/richardson.jpg";
</script>

<div>
  <div>
    <h2>The Richardson Hotel</h2>

    <img alt="" src={richardson} />

    <h3>444 Forest Ave, Buffalo, NY 14213</h3>
    <h3>(716) 493-2610</h3>
  </div>

  <div>
    A room block has been schedule at The Richardson Hotel in the heart of
    Buffalo. Call to book under the room block for Goosman-Elhage Wedding by
    September 12th, 2024 or use the link below. Rooms can be canceled up to 24
    hours prior to arrival. Rates are as follows:
    <ul>
      <li>Daily rate: $219.00 + tax</li>
      <li>Check in: 4:00 pm</li>
      <li>Check out: 11:00 am</li>
    </ul>
  </div>

  <div>
    <Buttonlink
      content="Website"
      href="https://therichardsonhotelbuffalo.com/"
    />
    <Buttonlink
      content="Book Now"
      href="https://be.synxis.com/?Hotel=40297&Chain=30877&group=2410GOOELH"
    />
  </div>
</div>

<style>
  img {
    max-width: 100%;
  }
  li {
    list-style-type: none;
  }
</style>
