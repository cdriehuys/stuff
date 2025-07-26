import { browserClient } from "@/api/client";
import { useQuery } from "@tanstack/react-query";

export const modelKeys = {
  all: ["models"] as const,
  lists: () => [...modelKeys.all, "lists"],
  list: (filters: unknown) => [...modelKeys.lists(), filters] as const,
  details: () => [...modelKeys.all, "details"] as const,
  detail: (id: number) => [...modelKeys.details(), id] as const,
};

export const useModelByID = (id: number) =>
  useQuery({
    queryKey: modelKeys.detail(id),
    queryFn: () => browserClient.getModelByID(id),
  });

export const useVendorModels = (vendorID: number) =>
  useQuery({
    queryKey: modelKeys.list({ vendorID }),
    queryFn: () => browserClient.getModelsByVendor(vendorID),
  });

export const useModelDisplayName = (id: number): string | undefined => {
  const query = useModelByID(id);

  if (query.data?.data) {
    return query.data.data.name || query.data.data.model;
  }

  if (query.isError || query.data?.error) {
    return "Unknown";
  }

  return undefined;
};
