import { Asset } from "@/api/client";
import { Table, Anchor } from "@mantine/core";

interface Props {
  assets: Asset[];
}

export default function AssetTable({ assets }: Props) {
  return (
    <Table highlightOnHover striped>
      <Table.Thead>
        <Table.Tr>
          <Table.Th>Serial</Table.Th>
          <Table.Th>Added</Table.Th>
        </Table.Tr>
      </Table.Thead>
      <Table.Tbody>
        {assets.map((asset) => (
          <Table.Tr key={asset.id}>
            <Table.Td>
              <Anchor href={`/assets/${asset.id}`}>{asset.serial}</Anchor>
            </Table.Td>
            <Table.Td>{asset.createdAt}</Table.Td>
          </Table.Tr>
        ))}
      </Table.Tbody>
    </Table>
  );
}
