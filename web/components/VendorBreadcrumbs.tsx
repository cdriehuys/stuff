"use client";

import useVendorByID from "@/queries/useVendor";
import { Breadcrumbs, Skeleton, Text } from "@mantine/core";
import { ReactNode } from "react";
import Anchor from "./Anchor";

export default function VendorBreadcrumbs({ vendorID }: { vendorID: number }) {
  const query = useVendorByID(vendorID);

  const crumbs: ReactNode[] = [
    <Anchor key="/vendors" href="/vendors">
      Vendors
    </Anchor>,
  ];

  if (query.data) {
    crumbs.push(
      <Text key={`/vendors/${query.data.id}`}>{query.data.name}</Text>,
    );
  }

  return (
    <Breadcrumbs mb="lg">
      <Anchor href="/vendors">Vendors</Anchor>
      {query.data?.name ? (
        <Text>{query.data.name}</Text>
      ) : query.isLoading ? (
        <Skeleton width={60} height={16} />
      ) : (
        <Text>Unknown</Text>
      )}
    </Breadcrumbs>
  );
}
