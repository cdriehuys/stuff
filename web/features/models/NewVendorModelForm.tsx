"use client";

import { browserClient, NewModel } from "@/api/client";
import { apiErrorAsFormError } from "@/api/errors";
import { createFormActions } from "@mantine/form";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { redirect } from "next/navigation";
import ModelForm, { ModelFormValues } from "../ModelForm";
import { modelKeys } from "./queries";

interface Props {
  vendorID: number;
}

const formName = "new-vendor-model-form";
const modelFormActions = createFormActions<ModelFormValues>(formName);

export default function NewVendorModelForm({ vendorID }: Props) {
  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: (model: NewModel) => browserClient.createModel(model),
    onError: (error) => {
      const formErrors = apiErrorAsFormError(error, ["model", "name"]);

      for (const [field, errors] of Object.entries(formErrors)) {
        modelFormActions.setFieldError(field, errors);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: modelKeys.lists() });
    },
  });
  const mutate = (body: Omit<NewModel, "vendorID">) =>
    mutation.mutate({ ...body, vendorID });

  if (mutation.isSuccess) {
    const next = `/vendors/${vendorID}/models/${mutation.data.id}`;
    redirect(next);
  }

  return (
    <ModelForm loading={mutation.isPending} name={formName} onSubmit={mutate} />
  );
}
