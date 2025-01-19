"use client";

import { components } from "@/api/api";
import { browserClient } from "@/api/client";
import { Button } from "@mantine/core";
import { IconTrash } from "@tabler/icons-react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { redirect } from "next/navigation";
import { modelKeys } from "./queries";
import { useDisclosure } from "@mantine/hooks";
import { useState } from "react";
import { asAPIError } from "@/api/errors";
import ErrorOverlay from "../errors/ErrorOverlay";

interface Props {
  model: components["schemas"]["Model"];
}

export default function ModelDeleteButton({ model }: Props) {
  const [error, setError] = useState<components["schemas"]["APIError"]>({});
  const [opened, { open, close }] = useDisclosure(false);

  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: () => browserClient.deleteModelByID(model.id),
    onError: (error) => {
      setError(asAPIError(error));
      open();
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: modelKeys.lists() });
      queryClient.setQueryData(modelKeys.detail(model.id), undefined);
    },
  });

  if (mutation.isSuccess) {
    redirect(`/vendors/${model.vendorID}`);
  }

  return (
    <>
      <ErrorOverlay
        error={error}
        onClose={close}
        opened={opened}
        title="Can't Delete Model"
      />
      <Button
        color="red"
        leftSection={<IconTrash />}
        loading={mutation.isPending}
        onClick={() => mutation.mutate()}
        size="compact-md"
      >
        Delete
      </Button>
    </>
  );
}
