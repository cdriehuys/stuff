import NewModelAssetForm from "@/features/assets/NewModelAssetForm";
import ModelBreadcrumbs from "@/features/models/ModelBreadcrumbs";
import { Container, Text, Title } from "@mantine/core";

interface Props {
  params: Promise<{ modelID: string; vendorID: string }>;
}

export default async function NewModelAssetPage({ params }: Props) {
  const resolvedParams = await params;
  const modelID = parseInt(resolvedParams.modelID);
  const vendorID = parseInt(resolvedParams.vendorID);

  return (
    <>
      <ModelBreadcrumbs modelID={modelID}>
        <Text>New Asset</Text>
      </ModelBreadcrumbs>
      <Container size="sm">
        <Title mb="lg">New Asset</Title>
        <NewModelAssetForm modelID={modelID} vendorID={vendorID} />
      </Container>
    </>
  );
}
