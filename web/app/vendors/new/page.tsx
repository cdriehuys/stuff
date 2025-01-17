import NewVendorForm from "@/features/vendors/NewVendorForm";
import { Container, Title } from "@mantine/core";

export default function NewVendorPage() {
  return (
    <Container size="sm">
      <Title mb="lg">New Vendor</Title>
      <NewVendorForm />
    </Container>
  );
}
