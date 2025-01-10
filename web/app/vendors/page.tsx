import VendorList from "@/components/VendorList";
import { Button, Flex, Title } from "@mantine/core";
import styles from "./style.module.css";
import { IconPlus } from "@tabler/icons-react";
import Link from "next/link";

export default function VendorPage() {
  return (
    <>
      <Flex className={styles.header}>
        <Title className={styles.title}>Vendors</Title>
        <Button component={Link} href="/vendors/new" leftSection={<IconPlus />}>
          Create
        </Button>
      </Flex>
      <VendorList />
    </>
  );
}
