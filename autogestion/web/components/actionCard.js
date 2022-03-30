import { Center, Stack, Text, Image, AspectRatio } from "@chakra-ui/react";

export default function ActionCard({ name, image, handleClick }) {
  return (
    <Center
      padding={6}
      boxShadow="lg"
      borderRadius="xl"
      onClick={handleClick}
      cursor="pointer"
    >
      <Stack spacing={4}>
        <Text fontSize="md" fontWeight="bold">
          {name}
        </Text>
        {image !== "" && (
          <AspectRatio ratio={16 / 9} w="xs">
            <Image
              src={image}
              alt="notes"
              objectFit="cover"
              borderRadius="md"
            />
          </AspectRatio>
        )}
      </Stack>
    </Center>
  );
}
