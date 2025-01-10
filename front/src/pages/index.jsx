import { useEffect } from "react";
import { useRouter } from "next/router";
import useAuth from "@/hooks/use-auth";

export default function Home() {
  const router = useRouter();
  const { user, loading } = useAuth();

  useEffect(() => {
    if (!loading && !user) {
      router.push("/login");
    }
  }, [user, loading, router]);

  if (loading) return <div>Loading...</div>;
  if (!user) return null;

  return (
    <div>
      <h1>Welcome, {user.username}!</h1>
      <p>This is a protected page.</p>
    </div>
  );
}
