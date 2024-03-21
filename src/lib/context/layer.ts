export type MapLayerPayload = {
  ui: {
    title: string;
    icon: string;
  };

  config: {
    apiUrl: string;
    layers: string;
    icon: string;
  }
}