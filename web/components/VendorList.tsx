"use client";

import { Anchor, Table } from "@mantine/core";
import { useQuery } from "@tanstack/react-query";
import Link from "next/link";

async function getVendors() {
  const response = await fetch("/api/vendors");
  const data = await response.json();

  return data.items;
}

export default function VendorList() {
  const query = useQuery({ queryKey: ["vendors"], queryFn: getVendors });

  return (
    <Table highlightOnHover striped>
      <Table.Thead>
        <Table.Tr>
          <Table.Th>Name</Table.Th>
        </Table.Tr>
      </Table.Thead>
      <Table.Tbody>
        {/* eslint-disable-next-line @typescript-eslint/no-explicit-any */}
        {query.data?.map((vendor: any) => (
          <Table.Tr key={vendor.id}>
            <Table.Td>
              <Anchor component={Link} c="blue" href={`/vendors/${vendor.id}`}>
                {vendor.name}
              </Anchor>
            </Table.Td>
          </Table.Tr>
        ))}
      </Table.Tbody>
    </Table>
  );
}
