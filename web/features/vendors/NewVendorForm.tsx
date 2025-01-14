"use client";

import { browserClient, NewVendor } from "@/api/client";
import { apiErrorAsFormError } from "@/api/errors";
import { createFormActions } from "@mantine/form";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { redirect } from "next/navigation";
import VendorForm from "../VendorForm";
import { vendorKeys } from "./queries";

const formName = "new-vendor-form";
const vendorFormActions = createFormActions(formName);

export default function NewVendorForm() {
  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: (vendor: NewVendor) => browserClient.createVendor(vendor),
    onError: (error) => {
      const formErrors = apiErrorAsFormError(error, ["name"]);

      for (const [field, errors] of Object.entries(formErrors)) {
        vendorFormActions.setFieldError(field, errors);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: vendorKeys.lists() });
    },
  });

  if (mutation.isSuccess) {
    const next = `/vendors/${mutation.data.id}`;
    redirect(next);
  }

  return (
    <VendorForm
      loading={mutation.isPending}
      name={formName}
      onSubmit={mutation.mutate}
    />
  );
}
