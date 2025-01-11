"use client";

import { useModelByID } from "@/api/queries";
import { Skeleton, Text } from "@mantine/core";
import { useParams } from "next/navigation";
import { ReactNode } from "react";
import Anchor from "./Anchor";
import VendorBreadcrumbs from "./VendorBreadcrumbs";

interface Props {
  children?: ReactNode;
  modelID: number;
}

export default function ModelBreadcrumbs({ children, modelID }: Props) {
  const { vendorID } = useParams<{ vendorID: string }>();
  const query = useModelByID(modelID);

  return (
    <VendorBreadcrumbs vendorID={parseInt(vendorID)}>
      {query.data ? (
        children ? (
          <Anchor href={`/vendors/${query.data.vendorID}/models/${modelID}`}>
            {query.data.name || query.data.model}
          </Anchor>
        ) : (
          <Text>{query.data.name || query.data.model}</Text>
        )
      ) : query.isLoading ? (
        <Skeleton width={120} height={16} />
      ) : (
        <Text>Unknown</Text>
      )}
      {children}
    </VendorBreadcrumbs>
  );
}
