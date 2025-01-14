import { getServerClient } from "@/api/client";
import ModelBreadcrumbs from "@/features/models/ModelBreadcrumbs";
import ModelDetail from "@/features/models/ModelDetail";
import { modelKeys } from "@/features/models/queries";
import {
  dehydrate,
  HydrationBoundary,
  QueryClient,
} from "@tanstack/react-query";

interface Props {
  params: Promise<{ modelID: string }>;
}

export default async function ModelDetailPage({ params }: Props) {
  const modelID = parseInt((await params).modelID);

  const queryClient = new QueryClient();
  const api = getServerClient();

  await queryClient.prefetchQuery({
    queryKey: modelKeys.detail(modelID),
    queryFn: () => api.getModelByID(modelID),
  });

  return (
    <HydrationBoundary state={dehydrate(queryClient)}>
      <ModelBreadcrumbs modelID={modelID} />
      <ModelDetail modelID={modelID} />
    </HydrationBoundary>
  );
}
