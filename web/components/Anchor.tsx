import { Anchor as BaseAnchor, AnchorProps, ElementProps } from "@mantine/core";
import Link from "next/link";

interface Props extends AnchorProps, ElementProps<"a", keyof AnchorProps> {
  href: string;
}

/**
 * An anchor that utilizes NextJS' {@link Link}.
 */
export default function Anchor(props: Props) {
  return <BaseAnchor component={Link} {...props} />;
}
