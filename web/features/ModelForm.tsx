"use client";

import { components } from "@/api/api";
import { Alert, Button } from "@mantine/core";
import { hasLength, useForm } from "@mantine/form";
import TextInput from "./TextInput";

interface Props {
  loading?: boolean;
  name?: string;
  onSubmit: (
    model: Omit<components["schemas"]["NewModel"], "vendorID">,
  ) => void;
}

export interface ModelFormValues {
  model: string;
  name: string;
}

export default function ModelForm({ loading, name, onSubmit }: Props) {
  const form = useForm<ModelFormValues>({
    name,
    mode: "uncontrolled",
    initialValues: {
      model: "",
      name: "",
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
      <Button loading={loading} type="submit">
        Submit
      </Button>
    </form>
  );
}
