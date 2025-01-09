import dayjs from "dayjs";
import React from "react";

export default function Footer() {
  return (
    <footer className='bg-gray-800 text-white p-4'>
      <div className='container mx-auto text-center'>
        © {dayjs().year()} Mi Sitio. Todos los derechos reservados.
      </div>
    </footer>
  );
}
