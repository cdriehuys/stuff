"use client";

import { Alert, Button, Group, Loader, Space, Title } from "@mantine/core";
import { IconPlus } from "@tabler/icons-react";
import Link from "next/link";
import VendorDeleteButton from "./VendorDeleteButton";
import VendorModelList from "../models/VendorModelList";
import { useVendorByID } from "./queries";

interface Props {
  vendorID: number;
}

export default function VendorDetail({ vendorID }: Props) {
  const query = useVendorByID(vendorID);

  if (query.isError) {
    return (
      <Alert color="red" title="Error">
        Request failed
      </Alert>
    );
  }

  if (query.isLoading && !query.data) {
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

    const vendor = query.data.data;

    return (
      <>
        <Title order={2} mb="md">
          {vendor.name}
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
