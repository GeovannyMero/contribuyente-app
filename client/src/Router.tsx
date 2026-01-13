import { createBrowserRouter, Route, createRoutesFromElements } from 'react-router';
import Layout from "./Components/Layout";
import HomePage from "./pages/HomePage";
import List from './pages/Contribuyente/List';

const router = createBrowserRouter(
    createRoutesFromElements(
        <Route 
            path="/"
            element={<Layout />}>
            <Route index element={<HomePage />} />
            <Route 
                path='contribuyentes/:name'
                element={<List></List>}
            />
        </Route>
    )
)

export default router;