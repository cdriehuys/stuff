import { browserClient } from "@/api/client";
import { useQuery } from "@tanstack/react-query";

export const assetKeys = {
  all: ["assets"] as const,
  lists: () => [...assetKeys.all, "lists"],
  list: (filters: unknown) => [...assetKeys.lists(), filters] as const,
  details: () => [...assetKeys.all, "details"] as const,
  detail: (id: number) => [...assetKeys.details(), id] as const,
};

export const useAssets = () =>
  useQuery({
    queryKey: assetKeys.list({}),
    queryFn: () => browserClient.getAssets(),
  });

export const useAssetByID = (id: number) =>
  useQuery({
    queryKey: assetKeys.detail(id),
    queryFn: () => browserClient.getAssetByID(id),
  });

export const useModelAssets = (modelID: number) =>
  useQuery({
    queryKey: assetKeys.list({ modelID }),
    queryFn: () => browserClient.getAssetsByModel(modelID),
  });
