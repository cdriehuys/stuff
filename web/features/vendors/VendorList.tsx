"use client";

import { Table } from "@mantine/core";
import Anchor from "../Anchor";
import { useVendorList } from "./queries";

export default function VendorList() {
  const query = useVendorList();

  return (
    <Table highlightOnHover striped>
      <Table.Thead>
        <Table.Tr>
          <Table.Th>Name</Table.Th>
        </Table.Tr>
      </Table.Thead>
      <Table.Tbody>
        {query.data?.data?.items.map((vendor) => (
          <Table.Tr key={vendor.id}>
            <Table.Td>
              <Anchor c="blue" href={`/vendors/${vendor.id}`}>
                {vendor.name}
              </Anchor>
            </Table.Td>
          </Table.Tr>
        ))}
      </Table.Tbody>
    </Table>
  );
}
