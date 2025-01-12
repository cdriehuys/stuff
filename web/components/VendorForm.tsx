"use client";

import { components } from "@/api/api";
import { Alert, Button } from "@mantine/core";
import { hasLength, useForm } from "@mantine/form";
import TextInput from "./TextInput";

interface Props {
  loading?: boolean;
  name?: string;
  onSubmit: (vendor: components["schemas"]["NewVendor"]) => void;
}

export default function VendorForm({ loading, name, onSubmit }: Props) {
  const form = useForm({
    name,
    mode: "uncontrolled",
    initialValues: {
      name: "",
    },
    validate: {
      name: hasLength(
        { min: 1, max: 150 },
        "Name must be 1-150 characters long.",
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
        key={form.key("name")}
        label="Name"
        mb="md"
        withAsterisk
        {...form.getInputProps("name")}
      />
      <Button loading={loading} type="submit">
        Submit
      </Button>
    </form>
  );
}
