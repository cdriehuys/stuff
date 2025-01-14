import { browserClient } from "@/api/client";
import { useQuery } from "@tanstack/react-query";

export const vendorKeys = {
  all: ["vendors"] as const,
  lists: () => [...vendorKeys.all, "lists"] as const,
  list: (params?: unknown) => [...vendorKeys.lists(), params] as const,
  details: () => [...vendorKeys.all, "details"] as const,
  detail: (id: number) => [...vendorKeys.details(), id] as const,
};

export const useVendorByID = (id: number) =>
  useQuery({
    queryKey: vendorKeys.detail(id),
    queryFn: () => browserClient.getVendorByID(id),
  });

export const useVendorList = () =>
  useQuery({
    queryKey: vendorKeys.list(),
    queryFn: () => browserClient.getVendors(),
  });
