import React from "react";
import Navbar from "./navbar";
import Footer from "./footer";

export default function Layout({ children }) {
  return (
    <div className='flex flex-col min-h-screen'>
      {/* Navbar */}
      <Navbar />

      {/* Main content area */}
      <main className='flex-grow overflow-y-auto'>
        <div className='container mx-auto p-4'>{children}</div>
      </main>

      {/* Footer */}
      <Footer />
    </div>
  );
}
