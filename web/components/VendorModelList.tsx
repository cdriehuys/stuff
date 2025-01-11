import apiClient from "@/api/apiClient";
import { Alert, Button, Loader, Table, Text } from "@mantine/core";
import Anchor from "./Anchor";
import Link from "next/link";
import { IconPlus } from "@tabler/icons-react";

interface Props {
  vendorID: number;
}

export default function VendorModelList({ vendorID }: Props) {
  const query = apiClient.useQuery("get", "/vendors/{vendorID}/models", {
    params: { path: { vendorID } },
  });

  if (query.isError) {
    return (
      <Alert color="red" title="Error">
        {query.error.message ?? "Error retrieving models."}
      </Alert>
    );
  }

  if (query.isLoading && !query.data) {
    return <Loader size="xl" />;
  }

  if (!query.data?.items.length) {
    return (
      <Alert title="No models">
        <Text mb="md">No models have been registered for this vendor.</Text>
        <Button
          component={Link}
          href={`/vendors/${vendorID}/new-model`}
          leftSection={<IconPlus />}
          variant="outline"
        >
          Create
        </Button>
      </Alert>
    );
  }

  return (
    <Table highlightOnHover striped>
      <Table.Thead>
        <Table.Tr>
          <Table.Th>Model</Table.Th>
          <Table.Th>Name</Table.Th>
        </Table.Tr>
      </Table.Thead>
      <Table.Tbody>
        {query.data.items.map((model) => (
          <Table.Tr key={model.id}>
            <Table.Td>
              <Anchor href={`/vendors/${model.vendorID}/models/${model.id}`}>
                {model.model}
              </Anchor>
            </Table.Td>
            <Table.Td>{model.name}</Table.Td>
          </Table.Tr>
        ))}
      </Table.Tbody>
    </Table>
  );
}