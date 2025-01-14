import { TextInput as BaseTextInput, TextInputProps } from "@mantine/core";

/**
 * An extension of {@link BaseTextInput} that uses a `<div>` to wrap the input's
 * errors and avoid hydration errors for cases where the errors are a list.
 */
export default function TextInput(props: TextInputProps) {
  return <BaseTextInput errorProps={{ component: "div" }} {...props} />;
}
