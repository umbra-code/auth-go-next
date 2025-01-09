import { useEffect, useState } from "react";

export function useTheme() {
  const [theme, setTheme] = useState("dark"); // Tema inicial predeterminado
  const [isMounted, setIsMounted] = useState(false); // Para manejar el montaje

  useEffect(() => {
    // Detectar tema inicial desde localStorage o preferencia del sistema
    const storedTheme = localStorage.getItem("theme");
    const preferredTheme = storedTheme
      ? storedTheme
      : window.matchMedia("(prefers-color-scheme: dark)").matches
      ? "dark"
      : "light";

    setTheme(preferredTheme);
    setIsMounted(true); // Marcar como montado
  }, []);

  useEffect(() => {
    if (!isMounted) return; // Evitar manipular el DOM antes de estar montado
    if (theme === "dark") {
      document.documentElement.classList.add("dark");
      localStorage.setItem("theme", "dark");
    } else {
      document.documentElement.classList.remove("dark");
      localStorage.setItem("theme", "light");
    }
  }, [theme, isMounted]);

  const toggleTheme = () => {
    setTheme((prevTheme) => (prevTheme === "dark" ? "light" : "dark"));
  };

  return { theme, toggleTheme };
}
