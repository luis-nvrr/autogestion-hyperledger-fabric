import Layout from "../components/layout";
import { Heading, SimpleGrid as Grid, Text, Stack } from "@chakra-ui/react";
import ActionCard from "../components/actionCard";
import withAuth from "../lib/withAuth";

const actions = [
  { id: 1, name: "Registrar nota", image: "/images/notes.jpg" },
  { id: 2, name: "Ver alumnos", image: "/images/students.jpg" },
  { id: 3, name: "Ver transacciones", image: "/images/books.jpg" },
];

function Docentes() {
  return (
    <Layout>
      <Stack
        direction="row"
        justifyContent="flex-start"
        alignItems="center"
        w="6xl"
        paddingY={6}
      >
        <Heading fontSize="xl">Docentes</Heading>
      </Stack>

      <Grid columns={3} spacing={6}>
        {actions.map((action) => (
          <ActionCard key={action.id} name={action.name} image={action.image} />
        ))}
      </Grid>
    </Layout>
  );
}

export default withAuth(Docentes);
