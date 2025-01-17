import { getServerClient } from "@/api/client";
import { modelKeys } from "@/features/models/queries";
import VendorBreadcrumbs from "@/features/vendors/VendorBreadcrumbs";
import VendorDetail from "@/features/vendors/VendorDetail";
import {
  dehydrate,
  HydrationBoundary,
  QueryClient,
} from "@tanstack/react-query";

interface Props {
  params: Promise<{ vendorID: string }>;
}

export default async function VendorDetailPage({ params }: Props) {
  const vendorID = parseInt((await params).vendorID);

  const queryClient = new QueryClient();
  const api = getServerClient();

  await queryClient.prefetchQuery({
    queryKey: modelKeys.list({ vendorID }),
    queryFn: () => api.getModelsByVendor(vendorID),
  });

  return (
    <HydrationBoundary state={dehydrate(queryClient)}>
      <VendorBreadcrumbs vendorID={vendorID} />
      <VendorDetail vendorID={vendorID} />
    </HydrationBoundary>
  );
}
