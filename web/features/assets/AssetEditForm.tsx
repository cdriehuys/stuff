"use client";

import { browserClient, NewAsset } from "@/api/client";
import { apiErrorAsFormError } from "@/api/errors";
import { Alert, Loader } from "@mantine/core";
import { createFormActions } from "@mantine/form";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { redirect } from "next/navigation";
import { ReactNode } from "react";
import AssetForm, { AssetFormValues } from "./AssetForm";
import { assetKeys, useAssetByID } from "./queries";

const formName = "asset-edit-form";
const formActions = createFormActions<AssetFormValues>(formName);

interface Props {
  actions?: ReactNode;
  assetID: number;
}

export default function AssetEditForm({ actions, assetID }: Props) {
  const query = useAssetByID(assetID);

  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: (asset: NewAsset) =>
      browserClient.updateAssetByID(assetID, asset),
    onError: (error) => {
      const formErrors = apiErrorAsFormError(error, ["serial", "comments"]);

      for (const [field, errors] of Object.entries(formErrors)) {
        formActions.setFieldError(field, errors);
      }
    },
    onSuccess: (asset) => {
      queryClient.invalidateQueries({ queryKey: assetKeys.lists() });
      queryClient.setQueryData(assetKeys.detail(asset.id), { data: asset });
    },
  });
  const mutate = (asset: Omit<NewAsset, "modelID">) => {
    const modelID = query.data?.data?.modelID;
    if (!modelID) {
      throw new Error("Form submitted without model ID available.");
    }

    return mutation.mutate({ ...asset, modelID });
  };

  if (mutation.isSuccess) {
    const next = `/assets/${assetID}`;
    redirect(next);
  }

  if (query.isError) {
    return <Alert color="red">Failed to load asset.</Alert>;
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
      <AssetForm
        actions={actions}
        loading={query.isLoading || mutation.isPending}
        asset={query.data.data}
        name={formName}
        onSubmit={mutate}
      />
    );
  }

  return null;
}
