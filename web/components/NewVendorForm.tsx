"use client";

import apiClient from "@/api/apiClient";
import { redirect } from "next/navigation";
import VendorForm from "./VendorForm";

export default function NewVendorForm() {
  const mutation = apiClient.useMutation("post", "/vendors");

  if (mutation.isSuccess) {
    const next = `/vendors/${mutation.data.id}`;
    redirect(next);
  }

  return (
    <VendorForm
      loading={mutation.isPending}
      onSubmit={(vendor) => mutation.mutate({ body: vendor })}
    />
  );
}
