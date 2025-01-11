"use client";

import apiClient from "@/api/apiClient";
import { redirect } from "next/navigation";
import ModelForm from "./ModelForm";
import { components } from "@/api/api";

interface Props {
  vendorID: number;
}

export default function NewVendorModelForm({ vendorID }: Props) {
  const mutation = apiClient.useMutation("post", "/models");
  const mutate = (body: components["schemas"]["NewModel"]) =>
    mutation.mutate({ body: { ...body, vendorID } });

  if (mutation.isSuccess) {
    const next = `/vendors/${vendorID}/models/${mutation.data.id}`;
    redirect(next);
  }

  return <ModelForm loading={mutation.isPending} onSubmit={mutate} />;
}
