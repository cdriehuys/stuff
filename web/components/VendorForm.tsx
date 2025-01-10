"use client";

import { components } from "@/api/api";
import { Button, TextInput } from "@mantine/core";
import { hasLength, useForm } from "@mantine/form";

interface Props {
  loading?: boolean;
  onSubmit: (vendor: components["schemas"]["NewVendor"]) => void;
}

export default function VendorForm({ loading, onSubmit }: Props) {
  const form = useForm({
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
