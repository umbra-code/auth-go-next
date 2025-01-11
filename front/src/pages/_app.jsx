import Layout from "@/components/layout/layout";
import { ToastProvider } from "@/contexts/toast-context";
import "@/styles/globals.css";

export default function App({ Component, pageProps }) {
  return (
    <ToastProvider>
      <Layout>
        <Component {...pageProps} />
      </Layout>
    </ToastProvider>
  );
}
