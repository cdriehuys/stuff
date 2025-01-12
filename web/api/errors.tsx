import { components } from "./api";
import { ReactNode } from "react";

export interface FormErrors {
  nonFieldErrors?: ReactNode;
  [x: string]: ReactNode;
}

export const apiErrorAsFormError = (
  error: components["schemas"]["APIError"] | undefined,
  fields: Iterable<string>,
): FormErrors => {
  if (!error?.fields) {
    return {};
  }

  const fieldSet = new Set(fields);

  // First group the errors by field name, discarding unknown fields.
  const grouped = error.fields.reduce(
    (errors: Record<string, string[]>, err) => {
      let field = err.field;
      if (!fieldSet.has(err.field)) {
        field = "nonFieldErrors";
      }

      const prevErrors = errors[field] ?? [];

      return {
        ...errors,
        [field]: [...prevErrors, err.message],
      };
    },
    {},
  );

  // Then keep single errors as a string and transform multiple errors into a
  // list.
  const formErrors: FormErrors = {};
  for (const [field, errors] of Object.entries(grouped)) {
    if (errors.length === 1) {
      formErrors[field] = errors[0];
    } else {
      formErrors[field] = (
        <ul>
          {errors.map((err) => (
            <li key={err}>{err}</li>
          ))}
        </ul>
      );
    }
  }

  return formErrors;
};
