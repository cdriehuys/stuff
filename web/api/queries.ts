import apiClient from "@/api/apiClient";

export const useModelByID = (modelID: number) =>
  apiClient.useQuery("get", "/models/{modelID}", {
    params: { path: { modelID } },
  });

export const useVendorByID = (vendorID: number) =>
  apiClient.useQuery("get", "/vendors/{vendorID}", {
    params: { path: { vendorID } },
  });
