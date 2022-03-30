import { useState, useEffect } from "react";
import {
  Box,
  Button,
  FormControl,
  FormLabel,
  Input,
  Stack,
  Text,
} from "@chakra-ui/react";
import useUser from "../lib/useUser";
import { useRouter } from "next/router";

export default function LoginBox() {
  const router = useRouter();
  const { user, login } = useUser();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  useEffect(() => {
    const mapToRoute = {
      org1: "/docentes",
      org2: "/estudiantes",
    };
    console.log(user);
    if (user === undefined) {
      return;
    }

    router.push(mapToRoute[user.organization]);
  }, [user, router]);

  const handleUserChange = (event) => {
    setUsername(event.target.value);
  };

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };

  const handleLoginSubmit = (event) => {
    event.preventDefault();
    login(username, password);

    setUsername("");
    setPassword("");
  };

  return (
    <Box padding={6} w={64} borderRadius="xl" boxShadow="xl">
      <form onSubmit={handleLoginSubmit}>
        <Stack direction="column" w="full" spacing={4}>
          <Text fontSize="md" fontWeight="bold">
            Datos de usuario
          </Text>
          <FormControl isRequired>
            <FormLabel>Nombre de usuario</FormLabel>
            <Input value={username} onChange={handleUserChange} />
          </FormControl>
          <FormControl isRequired>
            <FormLabel>Contraseña</FormLabel>
            <Input value={password} onChange={handlePasswordChange} />
          </FormControl>
          <Button type="submit" colorScheme="messenger">
            Iniciar Sesión
          </Button>
        </Stack>
      </form>
    </Box>
  );
}
