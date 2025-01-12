"use client";

import apiClient from "@/api/apiClient";
import { redirect } from "next/navigation";
import VendorForm from "./VendorForm";
import { apiErrorAsFormError } from "@/api/errors";
import { createFormActions } from "@mantine/form";

const formName = "new-vendor-form";
const vendorFormActions = createFormActions(formName);

export default function NewVendorForm() {
  const mutation = apiClient.useMutation("post", "/vendors", {
    onError: (error) => {
      const formErrors = apiErrorAsFormError(error, ["model", "name"]);

      for (const [field, errors] of Object.entries(formErrors)) {
        vendorFormActions.setFieldError(field, errors);
      }
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
      onSubmit={(vendor) => mutation.mutate({ body: vendor })}
    />
  );
}
