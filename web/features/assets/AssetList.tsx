"use client";

import { Alert, Loader, Text } from "@mantine/core";
import AssetTable from "./AssetTable";
import { useAssets } from "./queries";

export default function AssetList() {
  const query = useAssets();

  if (query.isError) {
    return (
      <Alert color="red" title="Error">
        Request failed
      </Alert>
    );
  }

  if (query.isLoading && !query.data) {
    return <Loader size="xl" />;
  }

  if (query.data?.error) {
    return (
      <Alert color="red" title="Error">
        {query.data.error.message || "Request failed for an unknown reason."}
      </Alert>
    );
  }

  if (!query.data?.data?.items.length) {
    return (
      <Alert title="No assets">
        <Text mb="md">No assets are currently being tracked.</Text>
      </Alert>
    );
  }

  return <AssetTable assets={query.data.data.items} />;
}
