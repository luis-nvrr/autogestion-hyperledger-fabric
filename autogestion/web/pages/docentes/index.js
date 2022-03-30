import Layout from "../../components/layout";
import { Heading, SimpleGrid as Grid, Text, Stack } from "@chakra-ui/react";
import ActionCard from "../../components/actionCard";
import withAuth from "../../lib/withAuth";
import { useRouter } from "next/router";

function Docentes() {
  const router = useRouter();

  const handleRegistrarNota = () => {
    console.log("registro nota");
    router.push("/docentes/registrar-nota");
  };

  const actions = [
    {
      id: 1,
      name: "Registrar nota",
      image: "/images/notes.jpg",
      handleClick: handleRegistrarNota,
    },
    {
      id: 2,
      name: "Ver alumnos",
      image: "/images/students.jpg",
      handleClick: () => {},
    },
    {
      id: 3,
      name: "Ver transacciones",
      image: "/images/books.jpg",
      handleClick: () => {},
    },
  ];

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
          <ActionCard
            key={action.id}
            name={action.name}
            image={action.image}
            handleClick={action.handleClick}
          />
        ))}
      </Grid>
    </Layout>
  );
}

export default withAuth(Docentes);
