import { components } from "@/api/api";
import { Alert, Group, Modal, Text } from "@mantine/core";
import { IconExclamationCircle } from "@tabler/icons-react";
import { ReactNode } from "react";

interface Props {
  error?: components["schemas"]["APIError"];
  opened: boolean;
  onClose: () => void;
  title: ReactNode;
}

export default function ErrorOverlay({ error, opened, onClose, title }: Props) {
  const wrappedTitle = typeof title === "string" ? <Text>{title}</Text> : title;

  return (
    <Modal
      opened={opened}
      onClose={onClose}
      title={
        <Group c="red">
          <IconExclamationCircle />
          {wrappedTitle}
        </Group>
      }
    >
      <Alert color="red" variant="white">
        {error?.message || "An unknown error occurred."}
      </Alert>
    </Modal>
  );
}
