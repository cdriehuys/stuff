"use client";

import { components } from "@/api/api";
import { browserClient } from "@/api/client";
import { Button } from "@mantine/core";
import { IconTrash } from "@tabler/icons-react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { redirect } from "next/navigation";
import { modelKeys } from "./queries";

interface Props {
  model: components["schemas"]["Model"];
}

export default function ModelDeleteButton({ model }: Props) {
  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: () => browserClient.deleteModelByID(model.id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: modelKeys.lists() });
      queryClient.setQueryData(modelKeys.detail(model.id), undefined);
    },
  });

  if (mutation.isSuccess) {
    redirect(`/vendors/${model.vendorID}`);
  }

  return (
    <Button
      color="red"
      leftSection={<IconTrash />}
      loading={mutation.isPending}
      onClick={() => mutation.mutate()}
      size="compact-md"
    >
      Delete
    </Button>
  );
}
