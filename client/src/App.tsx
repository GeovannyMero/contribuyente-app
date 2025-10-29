import { BrowserRouter, Route, Routes } from "react-router";
import "./App.css"
import Layout from "./Components/Layout";
import HomePage from "./pages/HomePage";
import List from "./pages/Contribuyente/List";


function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/"
          element={<Layout />}>
          <Route index element={<HomePage />} />
        </Route>
        <Route
          path="/contribuyentes/:province_name"
          element={<List />}
        >
        </Route>
        {/* <Route path="/:name" element={<DashboardPage />} /> */}
      </Routes>
    </BrowserRouter>
  );
}

export default App
