import ModelBreadcrumbs from "@/components/ModelBreadcrumbs";
import ModelDetail from "@/components/ModelDetail";

interface Props {
  params: Promise<{ modelID: string }>;
}

export default async function ModelDetailPage({ params }: Props) {
  const modelID = parseInt((await params).modelID);

  return (
    <>
      <ModelBreadcrumbs modelID={modelID} />
      <ModelDetail modelID={modelID} />
    </>
  );
}
