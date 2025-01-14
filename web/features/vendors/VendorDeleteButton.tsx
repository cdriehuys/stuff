"use client";

import { components } from "@/api/api";
import { browserClient } from "@/api/client";
import { asAPIError } from "@/api/errors";
import { Alert, Button, Modal } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { IconTrash } from "@tabler/icons-react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { redirect } from "next/navigation";
import { useState } from "react";
import { vendorKeys } from "./queries";

export default function VendorDeleteButton({ vendorID }: { vendorID: number }) {
  const [error, setError] = useState<components["schemas"]["APIError"]>({});
  const [opened, { open, close }] = useDisclosure(false);

  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: () => browserClient.deleteVendorByID(vendorID),
    onError: (error) => {
      setError(asAPIError(error));
      open();
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: vendorKeys.lists() });
      queryClient.setQueryData(vendorKeys.detail(vendorID), undefined);
    },
  });

  if (mutation.isSuccess) {
    redirect("/vendors");
  }

  return (
    <>
      <Modal
        opened={opened}
        onClose={close}
        centered
        withCloseButton={false}
        padding="0"
      >
        <Alert
          color="red"
          withCloseButton
          onClose={close}
          title="Can't delete vendor"
        >
          {error.message || "An unknown error occurred."}
        </Alert>
      </Modal>

      <Button
        color="red"
        leftSection={<IconTrash />}
        loading={mutation.isPending}
        onClick={() => mutation.mutate()}
      >
        Delete
      </Button>
    </>
  );
}
