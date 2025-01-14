"use client";

import { Breadcrumbs, Skeleton, Text } from "@mantine/core";
import { ReactNode } from "react";
import Anchor from "../Anchor";
import { useVendorByID } from "./queries";

interface Props {
  children?: ReactNode;
  vendorID: number;
}

export default function VendorBreadcrumbs({ children, vendorID }: Props) {
  const query = useVendorByID(vendorID);

  return (
    <Breadcrumbs mb="lg">
      <Anchor href="/vendors">Vendors</Anchor>
      {query.data?.data?.name ? (
        children ? (
          <Anchor href={`/vendors/${vendorID}`}>{query.data.data.name}</Anchor>
        ) : (
          <Text>{query.data.data.name}</Text>
        )
      ) : query.isLoading ? (
        <Skeleton width={60} height={16} />
      ) : (
        <Text>Unknown</Text>
      )}
      {children}
    </Breadcrumbs>
  );
}
