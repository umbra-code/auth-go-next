import { useTheme } from "@/hooks/use-theme";
import React from "react";

export default function ThemeToggler() {
  const { theme, toggleTheme } = useTheme();

  return (
    <button
      onClick={toggleTheme}
      className='p-2 rounded-md bg-gray-200 dark:bg-gray-700 text-gray-900 dark:text-gray-100'
    >
      {theme === "dark" ? "Set light" : "Set dark"}
    </button>
  );
}
