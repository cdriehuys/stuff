"use client";

import { NewAsset } from "@/api/client";
import { Alert, Button, Group, Textarea } from "@mantine/core";
import { hasLength, useForm } from "@mantine/form";
import { ReactNode } from "react";
import TextInput from "../TextInput";

interface Props {
  actions?: ReactNode;
  loading?: boolean;
  asset?: Partial<NewAsset>;
  name?: string;
  onSubmit: (asset: Omit<NewAsset, "modelID">) => void;
}

export interface AssetFormValues {
  serial: string;
  comments: string;
}

export default function AssetForm({
  actions,
  loading,
  asset,
  name,
  onSubmit,
}: Props) {
  const form = useForm<AssetFormValues>({
    name,
    mode: "uncontrolled",
    initialValues: {
      serial: asset?.serial ?? "",
      comments: asset?.comments ?? "",
    },
    validate: {
      serial: hasLength(
        { min: 1, max: 150 },
        "Serial numbers must contain 1-150 characters.",
      ),
      comments: hasLength(
        { max: 150 },
        "Comments must be 1,000 characters or less.",
      ),
    },
  });

  return (
    <form onSubmit={form.onSubmit(onSubmit)}>
      {form.errors.nonFieldErrors && (
        <Alert mb="md" color="red">
          {form.errors.nonFieldErrors}
        </Alert>
      )}
      <TextInput
        key={form.key("serial")}
        label="Serial"
        description="The serial number or other unique identifier the manufacturer uses for this product."
        mb="md"
        withAsterisk
        {...form.getInputProps("serial")}
      />
      <Textarea
        key={form.key("comments")}
        label="Comments"
        mb="md"
        {...form.getInputProps("comments")}
      />
      <Group justify="space-between">
        <Button loading={loading} type="submit">
          Submit
        </Button>
        {actions}
      </Group>
    </form>
  );
}
