"use client";

import { components } from "@/api/api";
import apiClient from "@/api/apiClient";
import { apiErrorAsFormError } from "@/api/errors";
import { createFormActions } from "@mantine/form";
import { redirect } from "next/navigation";
import ModelForm, { ModelFormValues } from "./ModelForm";

interface Props {
  vendorID: number;
}

const formName = "new-vendor-model-form";
const modelFormActions = createFormActions<ModelFormValues>(formName);

export default function NewVendorModelForm({ vendorID }: Props) {
  const mutation = apiClient.useMutation("post", "/models", {
    onError: (error) => {
      const formErrors = apiErrorAsFormError(error, ["model", "name"]);

      for (const [field, errors] of Object.entries(formErrors)) {
        modelFormActions.setFieldError(field, errors);
      }
    },
  });
  const mutate = (body: Omit<components["schemas"]["NewModel"], "vendorID">) =>
    mutation.mutate({ body: { ...body, vendorID } });

  if (mutation.isSuccess) {
    const next = `/vendors/${vendorID}/models/${mutation.data.id}`;
    redirect(next);
  }

  return (
    <ModelForm loading={mutation.isPending} name={formName} onSubmit={mutate} />
  );
}
