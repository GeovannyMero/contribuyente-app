import { BrowserRouter, Route, RouterProvider, Routes } from "react-router";
import "./App.css"
import Layout from "./Components/Layout";
import HomePage from "./pages/HomePage";
import List from "./pages/Contribuyente/List";
import router from './Router'



function App() {
  return (
    <RouterProvider router={router} />
  );
}

export default App
