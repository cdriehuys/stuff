"use client";

import apiClient from "@/api/apiClient";
import { Button } from "@mantine/core";
import { IconTrash } from "@tabler/icons-react";
import { redirect } from "next/navigation";

export default function VendorDeleteButton({ vendorID }: { vendorID: number }) {
  const mutation = apiClient.useMutation("delete", "/vendors/{vendorID}");
  const mutate = () => mutation.mutate({ params: { path: { vendorID } } });

  if (mutation.isSuccess) {
    redirect("/vendors");
  }

  return (
    <Button
      color="red"
      leftSection={<IconTrash />}
      loading={mutation.isPending}
      onClick={mutate}
    >
      Delete
    </Button>
  );
}
