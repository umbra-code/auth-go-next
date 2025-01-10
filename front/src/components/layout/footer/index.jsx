import dayjs from "dayjs";
import React from "react";

export default function Footer() {
  return (
    <footer className='p-4 opacity-70 text-sm'>
      <div className='container mx-auto text-center'>
        © {dayjs().year()} Mi Sitio. Todos los derechos reservados.
      </div>
    </footer>
  );
}
