"use client";

import { Skeleton, Text } from "@mantine/core";
import { useParams } from "next/navigation";
import { ReactNode } from "react";
import VendorBreadcrumbs from "../vendors/VendorBreadcrumbs";
import Anchor from "../Anchor";
import { useModelByID } from "./queries";

interface Props {
  children?: ReactNode;
  modelID: number;
}

export default function ModelBreadcrumbs({ children, modelID }: Props) {
  const { vendorID } = useParams<{ vendorID: string }>();
  const query = useModelByID(modelID);

  return (
    <VendorBreadcrumbs vendorID={parseInt(vendorID)}>
      {query.data?.data ? (
        children ? (
          <Anchor
            href={`/vendors/${query.data.data.vendorID}/models/${modelID}`}
          >
            {query.data.data.name || query.data.data.model}
          </Anchor>
        ) : (
          <Text>{query.data.data.name || query.data.data.model}</Text>
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
