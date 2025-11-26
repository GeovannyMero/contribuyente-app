import { useEffect } from "react";
import { useParams } from "react-router"
import { contribuyente } from "../../Types/Contribuyentes.d";

type ListaContri = contribuyente[]

const List: React.FC = () => {
    const params = useParams()
    console.log("http://127.0.0.1:3000/api/v1/contribuyentes?page=1&limit=5")
    const provinceName = params.province_name;


    async function data (): Promise<ListaContri> {
        

        try {
            var httpResponse = await fetch("http://127.0.0.1:3000/api/v1/contribuyentes?page=1&limit=5")

            if (!httpResponse.ok)
                throw new Error(`Error HTTP: ${httpResponse.status}`);

            // 3. Convertir la respuesta a JSON
            const datos = await httpResponse.json();

            const d: ListaContri = datos['value'] as ListaContri;
            // 4. Usar los datos
            console.log('Datos recibidos:', datos);
            console.log('Datos recibidos:', datos['values']);

            return d;
        } catch (error) {
            console.error('Hubo un problema con la petición fetch:', error);
            return []
        }


    }

    useEffect(() => {
        data().then(d => {
            console.log(d)
        })
    }, [])


    return (
        <>
            <h1>
                {provinceName}
            </h1>
        </>
    )
}

export default List;