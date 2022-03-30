import { useState } from "react";
import Layout from "../../components/layout";
import {
  FormControl,
  FormLabel,
  Input,
  Stack,
  Box,
  Button,
  Select,
  Heading,
} from "@chakra-ui/react";
import { useRouter } from "next/router";
import withAuth from "../../lib/withAuth";
import { SingleDatepicker } from "chakra-dayzed-datepicker";
import { gradeService } from "../../services/grade.service";

function RegistrarNota() {
  const router = useRouter();
  const [date, setDate] = useState(new Date());
  const [nota, setNota] = useState();
  const [legajo, setLegajo] = useState();
  const [nombre, setNombre] = useState();
  const [instancia, setInstancia] = useState();
  const [observaciones, setObservaciones] = useState();
  const [cursado, setCursado] = useState();
  const [apellido, setApellido] = useState();

  const handleFormSubmit = (event) => {
    event.preventDefault();
    const data = {
      grade: nota,
      date: date,
      student: {
        id: legajo,
        name: nombre,
        lastName: apellido,
        year: cursado,
      },
      instance: instancia,
      observations: observaciones,
    };

    gradeService.registerGrade(data);
  };

  return (
    <Layout>
      <Stack mt={8} direction="row" justifyContent="flex-end" w="6xl">
        <Button onClick={() => router.push("/docentes")}>Regresar</Button>
      </Stack>
      <Box padding={3} borderRadius="lg" boxShadow="2xl" w="xl">
        <Heading marginLeft={6} fontSize="2xl">
          Registrar nota de alumno
        </Heading>
        <form onSubmit={handleFormSubmit}>
          <Stack w="full">
            <Stack padding={6} spacing={6} direction="row">
              <Stack direction="column" spacing={3}>
                <FormControl isRequired>
                  <FormLabel>Nota</FormLabel>
                  <Input
                    type="number"
                    value={nota}
                    onChange={(event) => setNota(event.target.value)}
                  />
                </FormControl>
                <FormControl isRequired>
                  <FormLabel>Legajo</FormLabel>
                  <Input
                    type="number"
                    value={legajo}
                    onChange={(event) => setLegajo(event.target.value)}
                  />
                </FormControl>
                <FormControl isRequired>
                  <FormLabel>Nombre</FormLabel>
                  <Input
                    value={nombre}
                    onChange={(event) => setNombre(event.target.value)}
                    type="text"
                  />
                </FormControl>
                <FormControl isRequired maxWidth="2xs">
                  <FormLabel>Instancia de nota</FormLabel>
                  <Select
                    value={instancia}
                    onChange={(event) => setInstancia(event.target.value)}
                  >
                    <option value="exam">Parcial</option>
                    <option value="lab">TP</option>
                    <option value="presentation">Presentación</option>
                  </Select>
                </FormControl>
                <FormControl isRequired width="2xs">
                  <FormLabel>Observaciones</FormLabel>
                  <Input
                    value={observaciones}
                    onChange={(event) => setObservaciones(event.target.value)}
                    type="text"
                    h="16"
                  />
                </FormControl>
                <Button type="submit" colorScheme="messenger">
                  Cargar Nota
                </Button>
              </Stack>
              <Stack direction="column" spacing={3}>
                <FormControl>
                  <FormLabel>Fecha</FormLabel>
                  <SingleDatepicker
                    name="date-input"
                    date={date}
                    onDateChange={setDate}
                  />
                </FormControl>
                <FormControl isRequired>
                  <FormLabel>Año de cursado</FormLabel>
                  <Input
                    value={cursado}
                    onChange={(event) => setCursado(event.target.value)}
                  />
                </FormControl>
                <FormControl isRequired>
                  <FormLabel>Apellido</FormLabel>
                  <Input
                    value={apellido}
                    onChange={(event) => setApellido(event.target.value)}
                  />
                </FormControl>
              </Stack>
            </Stack>
          </Stack>
        </form>
      </Box>
    </Layout>
  );
}

export default withAuth(RegistrarNota);
