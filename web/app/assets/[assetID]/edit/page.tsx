import Anchor from "@/features/Anchor";
import AssetEditForm from "@/features/assets/AssetEditForm";
import { Container, Title } from "@mantine/core";

interface Props {
  params: Promise<{ assetID: string }>;
}

export default async function AssetEditPage({ params }: Props) {
  const assetID = parseInt((await params).assetID);

  return (
    <>
      <Container size="sm">
        <Title mb="lg">Edit Asset</Title>
        <AssetEditForm
          assetID={assetID}
          actions={<Anchor href={`/assets/${assetID}`}>Cancel</Anchor>}
        />
      </Container>
    </>
  );
}
