<script lang="ts">
  import LayerContainer from "$lib/components/map/layer/LayerContainer.svelte";
import { onMount } from "svelte";
  
  let map: any; // Mapbox mapboxgl.Maps
  let lng: number, lat: number, zoom: number;

  onMount(async () => {
    const { Map } = (await import("mapbox-gl")).default;

    map = new Map({
      container: 'map', // container ID
      style: 'mapbox://styles/mapbox/standard',
      center: [126.93, 37.53], // starting position [lng, lat]
      zoom: 15,
      accessToken: import.meta.env.VITE_MAPBOX_TOKEN,
    });

    lat = map.getCenter().lat;
    lng = map.getCenter().lng;
    zoom = map.getZoom();
  });
</script>

<svelte:head>
  <link rel="stylesheet" href="/css/mapbox-gl.css" />
</svelte:head>
<!-- Map container -->
<div class="map" id="map" />

<!-- Displayable Layer Control -->
<LayerContainer />
<slot></slot>

<style>
  .map {
    position: absolute;
    width: 100%;
    height: 100%;
  }
</style>

    
<!-- Your layout content here -->