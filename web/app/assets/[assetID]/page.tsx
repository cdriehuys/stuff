import AssetDetail from "@/features/assets/AssetDetail";

interface Props {
  params: Promise<{ assetID: string }>;
}

export default async function AssetDetailPage({ params }: Props) {
  const assetID = parseInt((await params).assetID);

  return <AssetDetail assetID={assetID} />;
}
