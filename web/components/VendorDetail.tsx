"use client";

import { useVendorByID } from "@/api/queries";
import { Alert, Button, Group, Loader, Space, Title } from "@mantine/core";
import { IconPlus } from "@tabler/icons-react";
import Link from "next/link";
import VendorDeleteButton from "./VendorDeleteButton";
import VendorModelList from "./VendorModelList";

interface Props {
  vendorID: number;
}

export default function VendorDetail({ vendorID }: Props) {
  const query = useVendorByID(vendorID);

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
        <Title order={2} mb="md">
          {query.data.name}
        </Title>
        <VendorDeleteButton vendorID={vendorID} />
        <Space h="xl" />
        <Group mb="md">
          <Title flex="1" order={3}>
            Models
          </Title>
          <Button
            component={Link}
            href={`/vendors/${vendorID}/new-model`}
            leftSection={<IconPlus />}
            size="compact-md"
          >
            Create
          </Button>
        </Group>
        <VendorModelList vendorID={vendorID} />
      </>
    );
  }

  return null;
}
