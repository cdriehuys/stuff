"use client";

import { components } from "@/api/api";
import { APIError, browserClient } from "@/api/client";
import { Button } from "@mantine/core";
import { IconTrash } from "@tabler/icons-react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { redirect } from "next/navigation";
import { useDisclosure } from "@mantine/hooks";
import { useState } from "react";
import { asAPIError } from "@/api/errors";
import ErrorOverlay from "../errors/ErrorOverlay";
import { assetKeys } from "./queries";

interface Props {
  asset: components["schemas"]["Asset"];
}

export default function AssetDeleteButton({ asset }: Props) {
  const [error, setError] = useState<APIError>({});
  const [opened, { open, close }] = useDisclosure(false);

  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: () => browserClient.deleteAssetByID(asset.id),
    onError: (error) => {
      setError(asAPIError(error));
      open();
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: assetKeys.lists() });
      queryClient.setQueryData(assetKeys.detail(asset.id), undefined);
    },
  });

  if (mutation.isSuccess) {
    redirect(`/vendors/${asset.vendorID}/models/${asset.modelID}`);
  }

  return (
    <>
      <ErrorOverlay
        error={error}
        onClose={close}
        opened={opened}
        title="Can't Delete Asset"
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
