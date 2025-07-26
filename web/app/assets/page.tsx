import AssetList from "@/features/assets/AssetList";
import { Title } from "@mantine/core";

export default function AssetListPage() {
  return (
    <>
      <Title mb="lg" order={1}>
        Assets
      </Title>
      <AssetList />;
    </>
  );
}
