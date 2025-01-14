"use client";

import { browserClient, NewModel } from "@/api/client";
import { createFormActions } from "@mantine/form";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import ModelForm, { ModelFormValues } from "./ModelForm";
import { modelKeys, useModelByID } from "./queries";
import { redirect } from "next/navigation";
import { Alert, Loader } from "@mantine/core";
import { apiErrorAsFormError } from "@/api/errors";
import { ReactNode } from "react";

const formName = "model-edit-form";
const formActions = createFormActions<ModelFormValues>(formName);

interface Props {
  actions?: ReactNode;
  modelID: number;
  vendorID: number;
}

export default function ModelEditForm({ actions, modelID, vendorID }: Props) {
  const query = useModelByID(modelID);

  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: (model: NewModel) =>
      browserClient.updateModelByID(modelID, model),
    onError: (error) => {
      const formErrors = apiErrorAsFormError(error, ["model", "name"]);

      for (const [field, errors] of Object.entries(formErrors)) {
        formActions.setFieldError(field, errors);
      }
    },
    onSuccess: (model) => {
      queryClient.invalidateQueries({ queryKey: modelKeys.lists() });
      queryClient.setQueryData(modelKeys.detail(model.id), { data: model });
    },
  });
  const mutate = (model: Omit<NewModel, "vendorID">) =>
    mutation.mutate({ ...model, vendorID });

  if (mutation.isSuccess) {
    const next = `/vendors/${vendorID}/models/${mutation.data.id}`;
    redirect(next);
  }

  if (query.isError) {
    return <Alert color="red">Failed to load model.</Alert>;
  }

  if (query.isLoading && !query.data) {
    return <Loader size="xl" />;
  }

  if (query.data?.error) {
    return (
      <Alert color="red">{query.data.error.message || "Unknown error"}</Alert>
    );
  }

  if (query.data?.data) {
    return (
      <ModelForm
        actions={actions}
        loading={query.isLoading || mutation.isPending}
        model={query.data.data}
        name={formName}
        onSubmit={mutate}
      />
    );
  }

  return null;
}
