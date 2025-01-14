import Anchor from "@/features/Anchor";
import ModelBreadcrumbs from "@/features/models/ModelBreadcrumbs";
import ModelEditForm from "@/features/models/ModelEditForm";
import { Container, Text, Title } from "@mantine/core";

interface Props {
  params: Promise<{ modelID: string; vendorID: string }>;
}

export default async function ModelEditPage({ params }: Props) {
  const resolvedParams = await params;
  const modelID = parseInt(resolvedParams.modelID);
  const vendorID = parseInt(resolvedParams.vendorID);

  return (
    <>
      <ModelBreadcrumbs modelID={modelID}>
        <Text>Edit</Text>
      </ModelBreadcrumbs>
      <Container size="sm">
        <Title mb="lg">Edit Model</Title>
        <ModelEditForm
          modelID={modelID}
          vendorID={vendorID}
          actions={
            <Anchor href={`/vendors/${vendorID}/models/${modelID}`}>
              Cancel
            </Anchor>
          }
        />
      </Container>
    </>
  );
}
