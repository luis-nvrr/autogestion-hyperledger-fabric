import { Text, Stack, Heading, Grid } from "@chakra-ui/react";
import Layout from "../components/layout";
import ActionCard from "../components/actionCard";
import withAuth from "../lib/withAuth";

const actions = [{ id: 1, name: "Ver notas", image: "/images/books.jpg" }];

function Alumnos() {
  return (
    <Layout>
      <Stack
        direction="row"
        justifyContent="flex-start"
        alignItems="center"
        w="6xl"
        paddingY={6}
      >
        <Heading fontSize="xl">Alumnos</Heading>
      </Stack>

      <Grid columns={3} spacing={6}>
        {actions.map((action) => (
          <ActionCard key={action.id} name={action.name} image={action.image} />
        ))}
      </Grid>
    </Layout>
  );
}

export default withAuth(Alumnos);
