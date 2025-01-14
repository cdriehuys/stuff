import { getServerClient } from "@/api/client";
import { vendorKeys } from "@/features/vendors/queries";
import {
  dehydrate,
  HydrationBoundary,
  QueryClient,
} from "@tanstack/react-query";
import { ReactNode } from "react";

interface Props {
  children: ReactNode;
  params: Promise<{ vendorID: string }>;
}

export default async function VendorDetailLayout({ children, params }: Props) {
  const vendorID = parseInt((await params).vendorID);

  const queryClient = new QueryClient();
  const api = getServerClient();

  await queryClient.prefetchQuery({
    queryKey: vendorKeys.detail(vendorID),
    queryFn: () => api.getVendorByID(vendorID),
  });

  return (
    <HydrationBoundary state={dehydrate(queryClient)}>
      {children}
    </HydrationBoundary>
  );
}
