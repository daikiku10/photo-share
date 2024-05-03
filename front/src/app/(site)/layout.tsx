import { Navigation } from "@/app/(site)/_components/Navigation";
import { Container } from "@/components/Layout/Container";
import { Footer } from "@/components/Layout/Footer";
import { Header } from "@/components/Layout/Header";
import { Main } from "@/components/Layout/Main";
import { Root } from "@/components/Layout/Root";

type Props = {
  children: React.ReactNode;
};

export default function SiteLayout({ children }: Props) {
  return (
    <Root>
      <Header />
      <Container>
        <Navigation />
        <Main>{children}</Main>
      </Container>
      <Footer />
    </Root>
  );
}
