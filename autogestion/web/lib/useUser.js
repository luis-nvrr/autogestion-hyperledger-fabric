import { useEffect } from "react";
import Router from "next/router";
import useSWR from "swr";

export default function useUser({ redirectIfFound = false }) {
  const { data: user, mutate: mutateUser } = useSWR("/api/user");

  useEffect(() => {
    const mapToRoute = {
      org1: "/docentes",
      org2: "/estudiantes",
    };
    console.log(user);
    // if no redirect needed, just return (example: already on /dashboard)
    // if user data not yet there (fetch in progress, logged in or not) then don't do anything yet
    if (!user) return;

    if (
      // If redirectIfFound is also set, redirect if the user was found
      redirectIfFound &&
      user?.username
    ) {
      Router.push(mapToRoute[user.organization]);
    }
  }, [user, redirectIfFound]);

  return { user, mutateUser };
}
