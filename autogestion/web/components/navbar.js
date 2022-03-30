import { Box, Container, Stack, Image, Text } from "@chakra-ui/react";
import { useRouter } from "next/router";

const Navbar = () => {
  const router = useRouter();

  return (
    <Box backgroundColor="gray.700" boxShadow="md">
      <Container maxWidth="6xl">
        <Stack
          alignItems="center"
          as="nav"
          direction="row"
          justifyContent="space-between"
          paddingY={3}
        >
          <Stack
            alignItems="flex-end"
            direction="row"
            spacing={3}
            onClick={() => router.push("/")}
            cursor="pointer"
          >
            <Image w={10} h={10} src="/images/block.png" alt="block" />
            <Text fontSize="xl" fontWeight="bold" color="whiteAlpha.900">
              Autogestion Smart
            </Text>
          </Stack>
          <Stack
            alignItems="flex-end"
            color="gray.600"
            direction="row"
            spacing={3}
            h={10}
          ></Stack>
        </Stack>
      </Container>
    </Box>
  );
};

export default Navbar;
