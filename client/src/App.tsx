import { BrowserRouter, Route, Routes } from "react-router";
import "./App.css"
import Layout from "./Components/Layout";
import HomePage from "./pages/HomePage";


function App() {

  // const [contrinuyentes, setContribuyentes] = useState<Contribuyente[]>([])

  // useEffect(() => {
  //   async function loadData() {
  //     const respopnse = await fetch("http://localhost:3000/contribuyentes")
  //     const result = await respopnse.json();
  //     const data: Contribuyente[] = result.data;
  //     console.log(data)
  //     setContribuyentes(data)
  //   }
  //   loadData()
  // }, [])

  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/"
          element={<Layout />}>
          <Route index element={<HomePage />} />
        </Route>

        {/* <Route path="/:name" element={<DashboardPage />} /> */}
      </Routes>
    </BrowserRouter>
  );
}

export default App
