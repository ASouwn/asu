"use client";

import { useState } from "react";

type Props = {
  modules: string
  data: any
}

export default function Page() {
  const [modules, setModules] = useState([]);
  const [props, setProps] = useState<Props>();
  return (
    <div className="flex h-screen items-center justify-center bg-gray-100">
      {}
      <button className="rounded bg-blue-500 px-4 py-2 text-white hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        onClick={() => { }}>
        addModule
      </button>
    </div >
  )
}