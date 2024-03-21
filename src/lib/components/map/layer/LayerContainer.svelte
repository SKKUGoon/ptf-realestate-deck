<script lang="ts">
  let expanded = false;
  let layersAvail: string[] = [];

  const handleExpansion = () => { expanded = !expanded };
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="container" on:click={ handleExpansion } class:expanded= { expanded }>
  {#if !expanded}
    <img src="/icons/researchLayer.png" alt="Layer Icon" class="icon" />
  {/if}

  <div class="section-title">
    Map Type
  </div>

  <div class="divider"></div>
</div>

<style lang="scss">
  .container {
    position: fixed;
    bottom: 40px;
    left: 50px;
  
    border-radius: 8px;
    width: 75px;
    height: 75px;
    overflow: hidden;
    
    background-color: #fff;
    box-shadow: 0 2px 4px rgba(0,0,0,0.2);
    z-index: 10;
    
    transition: all 0.3s ease;
    
    &.expanded {
      width: 270px; // Expanded width
      // left: 600px;
      min-height: 75px;
      height: auto; 
      background-color: #fff;
      padding: 15px;
      box-sizing: border-box;
      // Adjust styles for expanded state
    }
  }
  
  .section-title {
    font-size: 16px;
    font-weight: bold;
    margin: 10px 0;
    width: 100%;
    box-sizing: border-box;
    color: #333;
  }
  
  .divider {
    height: 1px;
    background-color: #ccc;
    width: calc(100% - 30px);
    margin: 15px 15px;
  }
  
  .icon {
    width: 100%;
    height: 100%;
    object-fit: contain;
    display: block;
    padding: 1.5px;
    box-sizing: border-box;
    border-radius: 8px;
  }
  
  // CSS Grid for LayerSelector components
//   .layers-container {
//     display: grid;
//     grid-template-columns: repeat(3, 1fr); // Creates 3 columns
//     grid-gap: 10px; // Adjust gap size as needed
//     width: calc(100% - 30px); // Adjust based on container padding
//     margin: auto;
//   }
</style>