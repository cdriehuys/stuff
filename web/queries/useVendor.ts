import apiClient from "@/api/apiClient";

const useVendorByID = (vendorID: number) =>
  apiClient.useQuery("get", "/vendors/{vendorID}", {
    params: { path: { vendorID } },
  });

export default useVendorByID;
