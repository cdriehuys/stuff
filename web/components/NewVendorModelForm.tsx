"use client";

import apiClient from "@/api/apiClient";
import { redirect } from "next/navigation";
import ModelForm from "./ModelForm";
import { components } from "@/api/api";

interface Props {
  vendorID: number;
}

export default function NewVendorModelForm({ vendorID }: Props) {
  const mutation = apiClient.useMutation("post", "/vendors/{vendorID}/models");
  const mutate = (body: components["schemas"]["NewModel"]) =>
    mutation.mutate({ params: { path: { vendorID } }, body });

  if (mutation.isSuccess) {
    const next = `/models/${mutation.data.id}`;
    redirect(next);
  }

  return <ModelForm loading={mutation.isPending} onSubmit={mutate} />;
}
