import { Anchor, Container, Text, Title } from "@mantine/core";
import Link from "next/link";

export default function NotFound() {
  return (
    <Container size="xs" py="xl">
      <Title>Not Found</Title>
      <Text>
        This page does not exist. Go{" "}
        <Anchor component={Link} href="/">
          Home
        </Anchor>
        ?
      </Text>
    </Container>
  );
}
