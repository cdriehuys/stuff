"use client";

import apiClient from "@/api/apiClient";
import { Alert, Loader, Title } from "@mantine/core";
import VendorDeleteButton from "./VendorDeleteButton";

interface Props {
  vendorID: number;
}

export default function VendorDetail({ vendorID }: Props) {
  const query = apiClient.useQuery("get", "/vendors/{vendorID}", {
    params: {
      path: {
        vendorID,
      },
    },
  });

  if (query.isLoading && !query.data) {
    return <Loader color="blue" size="xl" />;
  }

  if (query.isError) {
    return (
      <Alert color="red" title="Error">
        {query.error.message ?? "Request failed"}
      </Alert>
    );
  }

  if (query.data) {
    return (
      <>
        <Title order={2} mb="lg">
          {query.data.name}
        </Title>
        <VendorDeleteButton vendorID={vendorID} />
      </>
    );
  }

  return null;
}
