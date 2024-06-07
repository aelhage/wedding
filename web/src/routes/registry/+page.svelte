<script>
  import Buttonlink from "$lib/components/buttonlink.svelte";
  import richardson from "$lib/assets/richardson.jpg";
</script>

<div>
  We are so grateful to have you join us on our special day. Your presence is
  the best gift we could ask for. If you would like to contribute to our
  honeymoon fund, you can do so here

  <div>
    <Buttonlink
      content="Honeymoon Fund"
      href="https://www.zola.com/registry/lynnandalex2024"
    />
  </div>

  We have also registered at the following stores for those who prefer
  traditional gifts:

  <div>
    <Buttonlink
      content="Crate & Barrel"
      href="https://www.crateandbarrel.com/gift-registry/alex-elhage-and-lynn-goosman/r7106929"
    />
    <Buttonlink
      content="Target"
      href="https://www.target.com/gift-registry/gift-giver?registryId=92d309d0-169f-11ef-bfbe-09bc9319adab&type=WEDDING"
    />
  </div>
  Thank you for your love and support!
</div>

<style>
  img {
    max-width: 100%;
  }
  li {
    list-style-type: none;
  }
</style>
