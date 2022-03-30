import create from "zustand";
import { userService } from "../services/user.service";

const useUser = create((set) => ({
  user: undefined,
  isLoggedIn: false,
  login: async (username, password) => {
    const user = await userService.login(username, password);
    set({ user: user, isLoggedIn: true });
  },
}));

export default useUser;
