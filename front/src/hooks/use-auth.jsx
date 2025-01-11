import { useState, useEffect } from "react";
import { useRouter } from "next/router";
import { useToast } from "@/contexts/toast-context";
import {
  login as apiLogin,
  register as apiRegister,
  getUserProfile,
  clearAuthState,
} from "@/utils/api";

export default function useAuth() {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const router = useRouter();
  const { showToast } = useToast();

  useEffect(() => {
    async function loadUserFromToken() {
      try {
        const { data, status } = await getUserProfile();
        if (status !== 200) throw new Error(data.message);
        setUser(data);
      } catch (error) {
        console.error("Failed to load user:", error);
        if (error.isNetworkError) {
          showToast(error.message);
        } else {
          clearAuthState();
        }
      } finally {
        setLoading(false);
      }
    }

    loadUserFromToken();
  }, [showToast]);

  const login = async (username, password) => {
    try {
      const loginResponse = await apiLogin(username, password);
      if (loginResponse.status === 200) {
        const { status, data } = await getUserProfile();
        if (status !== 200) throw new Error(data.message);
        setUser(data);
        return data;
      } else {
        console.log("login failed...", loginResponse);
        throw new Error(loginResponse.message || "Login failed");
      }
    } catch (error) {
      console.log("Login failed:", error);
      showToast(error.message);
    }
  };

  const register = async (username, password) => {
    try {
      const registerResponse = await apiRegister(username, password);
      if (registerResponse.success) {
        showToast("Registration successful. Please log in.", "success");
      } else {
        throw new Error(registerResponse.message || "Registration failed");
      }
    } catch (error) {
      console.error("Registration failed:", error);
      if (error.isNetworkError) {
        showToast(error.message);
      } else {
        throw new Error(
          error.message || "Registration failed. Please try again."
        );
      }
    }
  };

  const logout = () => {
    clearAuthState();
    setUser(null);
    router.push("/auth/login");
  };

  return { user, loading, login, register, logout };
}
