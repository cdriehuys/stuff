"use client";

import { Alert, Button, Loader, Text } from "@mantine/core";
import { IconPlus } from "@tabler/icons-react";
import Link from "next/link";
import { useParams } from "next/navigation";
import AssetTable from "./AssetTable";
import { useModelAssets } from "./queries";

interface Props {
  modelID: number;
}

export default function ModelAssetList({ modelID }: Props) {
  const params = useParams<{ vendorID: string }>();
  const query = useModelAssets(modelID);

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
        <Text mb="md">No assets have been registered for this model.</Text>
        <Button
          component={Link}
          href={`/vendors/${params.vendorID}/models/${modelID}/new-asset`}
          leftSection={<IconPlus />}
          variant="outline"
        >
          Create
        </Button>
      </Alert>
    );
  }

  return <AssetTable assets={query.data.data.items} />;
}
