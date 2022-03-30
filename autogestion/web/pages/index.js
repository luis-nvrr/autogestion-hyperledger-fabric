import Layout from "../components/layout";
import { Text, Stack, Link, Flex, Container, Center } from "@chakra-ui/react";
import { ExternalLinkIcon } from "@chakra-ui/icons";
import LoginBox from "../components/loginBox";

export default function Home() {
  return (
    <Layout>
      <Stack direction="column" spacing={6} w="full">
        <Flex
          alignItems="flex-end"
          direction="row"
          backgroundImage="/images/header.jpg"
          backgroundPosition="center"
          backgroundSize="cover"
          justifyContent="flex-start"
          minH={64}
          padding={6}
        >
          <Container w="lg" m={0} p={0}>
            <Text color="whiteAlpha.900" fontSize="2xl" fontWeight="bold">
              Aplicación basada en Hyperledger Fabric para el registro de notas
              de alumnos utilizando smart contracts.
              <br />
              Podes visitar el código en{" "}
              <Link
                href="https://github.com/luis-nvrr/autogestion-hyperledger-fabric"
                isExternal
              >
                Github <ExternalLinkIcon mx="2px" />
              </Link>
            </Text>
          </Container>
        </Flex>
        <Center>
          <LoginBox />
        </Center>
      </Stack>
    </Layout>
  );
}
