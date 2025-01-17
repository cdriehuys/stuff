import NewVendorModelForm from "@/features/models/NewVendorModelForm";
import VendorBreadcrumbs from "@/features/vendors/VendorBreadcrumbs";
import { Container, Text, Title } from "@mantine/core";

interface Props {
  params: Promise<{ vendorID: string }>;
}

export default async function NewVendorPage({ params }: Props) {
  const vendorID = parseInt((await params).vendorID);

  return (
    <>
      <VendorBreadcrumbs vendorID={vendorID}>
        <Text>New Model</Text>
      </VendorBreadcrumbs>
      <Container size="sm">
        <Title mb="lg">New Model</Title>
        <NewVendorModelForm vendorID={vendorID} />
      </Container>
    </>
  );
}
