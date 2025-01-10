import { useState, useEffect } from "react";
import { useRouter } from "next/router";
import {
  login as apiLogin,
  register as apiRegister,
  refreshToken,
  getUserProfile,
} from "@/utils/api";

export default function useAuth() {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const router = useRouter();

  useEffect(() => {
    async function loadUserFromToken() {
      try {
        const userData = await getUserProfile();
        setUser(userData);
      } catch (error) {
        console.error("Failed to load user:", error);
      } finally {
        setLoading(false);
      }
    }

    loadUserFromToken();
  }, []);

  const login = async (username, password) => {
    const userData = await apiLogin(username, password);
    setUser(userData);
  };

  const register = async (username, password) => {
    await apiRegister(username, password);
  };

  const logout = () => {
    setUser(null);
    // Aquí podrías agregar lógica para eliminar las cookies si es necesario
    router.push("/login");
  };

  return { user, loading, login, register, logout };
}
