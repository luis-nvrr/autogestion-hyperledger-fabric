import axios from "axios";

const baseUrl = "http://localhost:8081/api/grades";

const registerGrade = async (data) => {
  console.log(data);
  const response = await axios.post(`${baseUrl}`, data);
  return response.data;
};

export const gradeService = {
  registerGrade,
};
