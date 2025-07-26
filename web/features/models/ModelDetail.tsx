"use client";

import { Alert, Button, Group, Loader, Stack, Title } from "@mantine/core";
import { IconPencil, IconPlus } from "@tabler/icons-react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import ModelAssetList from "../assets/ModelAssetList";
import ModelDeleteButton from "./ModelDeleteButton";
import { useModelByID } from "./queries";

interface Props {
  modelID: number;
}

export default function ModelDetail({ modelID }: Props) {
  const pathname = usePathname();
  const query = useModelByID(modelID);

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

    const model = query.data.data;

    return (
      <>
        <Stack mb="lg">
          <Title order={2}>{model.name || model.model}</Title>
          {model.name && (
            <Title order={3} size="h4" c="dimmed">
              {model.model}
            </Title>
          )}
        </Stack>
        <Group mb="xl">
          <Button
            component={Link}
            href={`${pathname}/edit`}
            leftSection={<IconPencil />}
            size="compact-md"
          >
            Edit
          </Button>
          <ModelDeleteButton model={model} />
        </Group>
        <Group mb="md">
          <Title flex="1" order={3}>
            Assets
          </Title>
          <Button
            component={Link}
            href={`/vendors/${model.vendorID}/models/${modelID}/new-asset`}
            leftSection={<IconPlus />}
            size="compact-md"
          >
            Create
          </Button>
        </Group>
        <ModelAssetList modelID={modelID} />
      </>
    );
  }

  return null;
}
