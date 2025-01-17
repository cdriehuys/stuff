import { getServerClient } from "@/api/client";
import { modelKeys } from "@/features/models/queries";
import {
  dehydrate,
  HydrationBoundary,
  QueryClient,
} from "@tanstack/react-query";
import { ReactNode } from "react";

interface Props {
  children: ReactNode;
  params: Promise<{ modelID: string }>;
}

export default async function ModelDetailLayout({ children, params }: Props) {
  const modelID = parseInt((await params).modelID);

  const queryClient = new QueryClient();
  const api = getServerClient();

  await queryClient.prefetchQuery({
    queryKey: modelKeys.detail(modelID),
    queryFn: () => api.getModelByID(modelID),
  });

  return (
    <HydrationBoundary state={dehydrate(queryClient)}>
      {children}
    </HydrationBoundary>
  );
}
