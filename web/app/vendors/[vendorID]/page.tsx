import VendorBreadcrumbs from "@/components/VendorBreadcrumbs";
import VendorDetail from "@/components/VendorDetail";

interface Props {
  params: Promise<{ vendorID: string }>;
}

export default async function VendorDetailPage({ params }: Props) {
  const vendorID = parseInt((await params).vendorID);

  return (
    <>
      <VendorBreadcrumbs vendorID={vendorID} />
      <VendorDetail vendorID={vendorID} />
    </>
  );
}
