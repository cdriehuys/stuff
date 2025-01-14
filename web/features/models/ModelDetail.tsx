"use client";

import { Alert, Loader, Stack, Title } from "@mantine/core";
import ModelDeleteButton from "./ModelDeleteButton";
import { useModelByID } from "./queries";

interface Props {
  modelID: number;
}

export default function ModelDetail({ modelID }: Props) {
  const query = useModelByID(modelID);

  if (query.isError) {
    return (
      <Alert color="red" title="Error">
        {"Request failed"}
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
        <ModelDeleteButton model={model} />
      </>
    );
  }

  return null;
}
