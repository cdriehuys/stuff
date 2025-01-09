"use client";

import NavLink from "@/components/NavLink";
import {
  AppShell as AppShellBase,
  Burger,
  Group,
  Title
} from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import React from "react";

interface Props {
  children: React.ReactNode;
}

export default function AppShell({ children }: Props) {
  const [opened, { toggle }] = useDisclosure();

  return (
    <AppShellBase
      header={{ height: 60 }}
      navbar={{ width: 300, breakpoint: "sm", collapsed: { mobile: !opened } }}
      padding="md"
    >
      <AppShellBase.Header>
        <Group h="100%" px="md">
          <Burger opened={opened} onClick={toggle} hiddenFrom="sm" size="sm" />
          <Title>Stuff</Title>
        </Group>
      </AppShellBase.Header>
      <AppShellBase.Navbar p="md">
        <NavLink exact href="/" label="Home" />
        <NavLink href="/vendors" label="Vendors" />
      </AppShellBase.Navbar>
      <AppShellBase.Main>{children}</AppShellBase.Main>
    </AppShellBase>
  );
}
