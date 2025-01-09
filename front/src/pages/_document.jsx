import { Html, Head, Main, NextScript } from "next/document";

export default function Document() {
  return (
    <Html lang='en'>
      <Head />
      <body className='antialiased bg-white dark:bg-neutral-900 text-gray-800 dark:text-white transition ease-in-out'>
        <Main />
        <NextScript />
      </body>
    </Html>
  );
}
