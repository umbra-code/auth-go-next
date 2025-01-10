import { useState } from "react";
import { useRouter } from "next/router";
import useAuth from "@/hooks/use-auth";

export default function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();
  const { login } = useAuth();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await login(username, password);
      router.push("/");
    } catch (error) {
      console.error("Login failed:", error);
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className='border dark:border-none dark:bg-black rounded-lg shadow p-5 m-auto flex flex-col gap-3 max-w-sm'
    >
      <input
        type='text'
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        className='border dark:border-black rounded p-2 dark:bg-neutral-900'
        placeholder='Username'
        required
      />
      <input
        type='password'
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        className='border dark:border-black rounded p-2 dark:bg-neutral-900'
        placeholder='Password'
        required
      />
      <button
        type='submit'
        className='bg-indigo-400 p-2 text-white rounded hover:bg-indigo-500 transition dark:bg-indigo-700 dark:hover:bg-indigo-800'
      >
        Login
      </button>
    </form>
  );
}
