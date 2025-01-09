import "@mantine/core/styles.css";

import AppShell from "@/layout/AppShell";
import ClientProviders from "@/layout/ClientProviders";
import {
  ColorSchemeScript,
  MantineProvider,
  mantineHtmlProps,
} from "@mantine/core";

export const metadata = {
  title: "My Mantine app",
  description: "I have followed setup instructions carefully",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" {...mantineHtmlProps}>
      <head>
        <ColorSchemeScript />
      </head>
      <body>
        <ClientProviders>
          <MantineProvider>
            <AppShell>{children}</AppShell>
          </MantineProvider>
        </ClientProviders>
      </body>
    </html>
  );
}
