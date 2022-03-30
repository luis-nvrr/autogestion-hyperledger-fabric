import { Container, Flex } from "@chakra-ui/react";
import Head from "next/head";
import Navbar from "./navbar";

export const siteTitle = "Autogestion blockchain";

const Layout = ({ children }) => (
  <Flex flex={1} direction="column" backgroundColor="whiteAlpha.500">
    <Head>
      <meta name="description" content="Hyperledger fabric app" />
      <meta name="og:title" content={siteTitle} />
    </Head>
    <Navbar />
    <Container
      centerContent
      minH="100vh"
      maxW="100vw"
      padding={0}
      marginBottom={6}
    >
      {children}
    </Container>
  </Flex>
);

export default Layout;
