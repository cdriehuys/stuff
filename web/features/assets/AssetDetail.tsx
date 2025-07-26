"use client";

import { Alert, Button, Group, Loader, Title } from "@mantine/core";
import { useAssetByID } from "./queries";
import AssetDeleteButton from "./AssetDeleteButton";
import Link from "next/link";
import { IconPencil } from "@tabler/icons-react";

interface Props {
  assetID: number;
}

export default function AssetDetail({ assetID }: Props) {
  const query = useAssetByID(assetID);

  if (query.isError) {
    return (
      <Alert color="red" title="Error">
        Request failed
      </Alert>
    );
  }

  if (!query.data && query.isLoading) {
    return <Loader color="blue" size="xl" />;
  }

  if (query.data) {
    if (query.data.error) {
      return (
        <Alert color="red" title="Error">
          {query.data.error.message || "Request failed for an unknown reason."}
        </Alert>
      );
    }

    const asset = query.data.data;

    return (
      <>
        <Title mb="lg" order={2}>
          {asset.serial}
        </Title>
        <Group>
          <Button
            component={Link}
            href={`/assets/${asset.id}/edit`}
            leftSection={<IconPencil />}
            size="compact-md"
          >
            Edit
          </Button>
          <AssetDeleteButton asset={asset} />
        </Group>
      </>
    );
  }
}
