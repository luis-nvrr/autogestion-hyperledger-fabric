import axios from "axios";

const baseUrl = "http://localhost:8080/api/users";

const login = async (username, password) => {
  const response = await axios.post(`${baseUrl}/auth`, {
    username: username,
    password: password,
  });

  return response.data;
};

export const userService = {
  login,
};
