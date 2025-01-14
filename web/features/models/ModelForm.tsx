"use client";

import { components } from "@/api/api";
import { Alert, Button, Group } from "@mantine/core";
import { hasLength, useForm } from "@mantine/form";
import TextInput from "../TextInput";
import { NewModel } from "@/api/client";
import { ReactNode } from "react";

interface Props {
  actions?: ReactNode;
  loading?: boolean;
  model?: Partial<NewModel>;
  name?: string;
  onSubmit: (
    model: Omit<components["schemas"]["NewModel"], "vendorID">,
  ) => void;
}

export interface ModelFormValues {
  model: string;
  name: string;
}

export default function ModelForm({
  actions,
  loading,
  model,
  name,
  onSubmit,
}: Props) {
  const form = useForm<ModelFormValues>({
    name,
    mode: "uncontrolled",
    initialValues: {
      model: model?.model ?? "",
      name: model?.name ?? "",
    },
    validate: {
      model: hasLength(
        { min: 1, max: 150 },
        "Models must contain 1-150 characters.",
      ),
      name: hasLength({ max: 150 }, "Names must be 150 characters or less."),
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
        key={form.key("model")}
        label="Model"
        description="The model number or other unique identifier the manufacturer uses for this product."
        mb="md"
        withAsterisk
        {...form.getInputProps("model")}
      />
      <TextInput
        key={form.key("name")}
        label="Name"
        description="A human-readable name for the product."
        mb="md"
        {...form.getInputProps("name")}
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
