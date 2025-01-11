"use client";

import { useModelByID } from "@/api/queries";
import { Alert, Loader, Stack, Title } from "@mantine/core";

interface Props {
  modelID: number;
}

export default function ModelDetail({ modelID }: Props) {
  const query = useModelByID(modelID);

  if (query.isError) {
    return (
      <Alert color="red" title="Error">
        {query.error.message ?? "Request failed"}
      </Alert>
    );
  }

  if (!query.data && query.isLoading) {
    return <Loader color="blue" size="xl" />;
  }

  if (query.data) {
    return (
      <>
        <Stack mb="lg">
          <Title order={2}>{query.data.name || query.data.model}</Title>
          {query.data.name && (
            <Title order={3} size="h4" c="dimmed">
              {query.data.model}
            </Title>
          )}
        </Stack>
      </>
    );
  }

  return null;
}
