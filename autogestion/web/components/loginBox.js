import { useState } from "react";
import {
  Box,
  Button,
  FormControl,
  FormLabel,
  Input,
  Stack,
  Text,
} from "@chakra-ui/react";
import { userService } from "../services";
import useUser from "../lib/useUser";

export default function LoginBox() {
  const { mutateUser } = useUser({
    redirectIfFound: true,
  });

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleUserChange = (event) => {
    setUsername(event.target.value);
  };

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };

  const handleLoginSubmit = (event) => {
    event.preventDefault();
    userService.login(username, password).then(mutateUser).catch(console.log);

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
