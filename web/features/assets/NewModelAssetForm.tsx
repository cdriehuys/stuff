"use client";

import { createFormActions } from "@mantine/form";
import AssetForm, { AssetFormValues } from "./AssetForm";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { browserClient, NewAsset } from "@/api/client";
import { apiErrorAsFormError } from "@/api/errors";
import { assetKeys } from "./queries";
import { redirect } from "next/navigation";
import Anchor from "../Anchor";

interface Props {
  modelID: number;
  vendorID: number;
}

const formName = "new-model-asset-form";
const formActions = createFormActions<AssetFormValues>(formName);

export default function NewModelAssetForm({ modelID, vendorID }: Props) {
  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: (asset: NewAsset) => browserClient.createAsset(asset),
    onError: (error) => {
      const formErrors = apiErrorAsFormError(error, ["serial", "comments"]);

      for (const [field, errors] of Object.entries(formErrors)) {
        formActions.setFieldError(field, errors);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: assetKeys.all });
    },
  });
  const mutate = (body: Omit<NewAsset, "modelID">) =>
    mutation.mutate({ ...body, modelID });

  if (mutation.isSuccess) {
    const next = `/assets/${mutation.data.id}`;
    redirect(next);
  }

  return (
    <AssetForm
      loading={mutation.isPending}
      name={formName}
      onSubmit={mutate}
      actions={
        <Anchor href={`/vendors/${vendorID}/models/${modelID}`}>Cancel</Anchor>
      }
    />
  );
}
