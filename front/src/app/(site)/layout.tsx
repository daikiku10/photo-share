import { Root } from "@/components/Layout/Root";

type Props = {
  children: React.ReactNode;
};

export default function SiteLayout({ children }: Props) {
  return <Root>{children}</Root>;
}
