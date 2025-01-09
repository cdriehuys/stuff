import { NavLink as BaseNavLink } from "@mantine/core";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { ReactNode } from "react";

interface Props {
  exact?: boolean;
  href: string;
  label: ReactNode;
}

/**
 * A navigation link that ties together Mantine's NavLink with NextJS' Link.
 */
export default function NavLink({ exact, href, label }: Props) {
  const pathname = usePathname();

  let active = pathname === href;
  if (!active && !exact) {
    if (!href.endsWith("/")) {
      active = pathname.startsWith(`${href}/`);
    } else {
      active = pathname.startsWith(href);
    }
  }

  return (
    <BaseNavLink component={Link} href={href} label={label} active={active} />
  );
}
