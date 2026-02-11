import { useEffect, useState } from "react";
import { useParams } from "react-router"
import { contribuyente } from "../../Types/Contribuyentes.d";

type ListaContri = contribuyente[]

const List: React.FC = () => {
    console.log('Listado')
    const params = useParams()
    console.log(params)
    console.log("http://127.0.0.1:3000/api/v1/contribuyentes?page=1&limit=5")
    // const provinceName = params.province_name;

    const [person, setPerson] = useState<contribuyente[]>([])
    // const [provinceName, setProvinceName] = useState(params.name);
    const [loading, setLoading] = useState<boolean>(true);
    const [page, setPage] = useState<number>(1)


    async function data(): Promise<ListaContri> {
        try {
            var httpResponse = await fetch(`http://127.0.0.1:3000/api/v1/contribuyentes?provincia=${params.name?.toUpperCase()}&page=${page}&limit=10`)

            if (!httpResponse.ok)
                throw new Error(`Error HTTP: ${httpResponse.status}`);

            // 3. Convertir la respuesta a JSON
            const datos = await httpResponse.json();

            const d: ListaContri = datos['Value'] as ListaContri;
            // 4. Usar los datos
            console.log('Datos recibidos:', datos);
            console.log('Datos recibidos:', datos['Value']);

            return d;
        } catch (error) {
            console.error('Hubo un problema con la petición fetch:', error);
            return []
        }


    }

    useEffect(() => {
        data().then(d => {
            console.log(d)
            setPerson(d)
            setLoading(false);
        })
    }, [page])

    // 1. Estado de Carga (Loading)
    if (loading) {
        return (
            <div className="flex flex-col items-center justify-center p-20">
                <span className="loading loading-spinner loading-lg text-primary"></span>
                <p className="mt-4 text-sm animate-pulse">Cargando información...</p>
            </div>
        );
    }

    return (
        <>
            {/* {person.map(e => {
            return <p>{e.razon_social}</p>
        })} */}

            <div className="card w-full h-full bg-base-100 shadow-sm">
                <div className="card-body">
                    <div className="card-title text-primary">
                        {`Provincia de ${params.name?.toUpperCase()}`}
                    </div>
                    <div className="overflow-x-auto rounded-box border border-base-content/5 bg-base-100">
                        <table className="table table-zebra">
                            <thead>
                                <tr>
                                    <th>Ruc</th>
                                    <th>Rasón Social</th>
                                    <th>Tipo</th>
                                    <th>Actividad Economíca</th>
                                    <th>Estado</th>
                                    <th>Aciones</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    (person != undefined && person.length > 0) ?
                                        person.map(e => {
                                            return (
                                                <tr key={e.ruc}>
                                                    <th>{e.ruc}</th>
                                                    <td>{e.razon_social}</td>
                                                    <td>{e.tipo_contribuyente}</td>
                                                    <td>{`${e.codigo_ciiu} - ${e.actividad_economica}`}</td>
                                                    <td>{e.estado_contribuyente}</td>
                                                    <td>
                                                        <button
                                                            className="btn btn-xs btn-outline btn-info"
                                                            onClick={() => alert(e.ruc)}>
                                                            Ver
                                                        </button>
                                                    </td>
                                                </tr>
                                            )
                                        }) :
                                        <tr>
                                            <td colSpan={6} className="text-center py-10">
                                                <div className="flex flex-col items-center gap-2">
                                                    <span className="text-gray-400 italic">No hay datos disponibles</span>
                                                    <div className="badge badge-ghost text-xs">Aún no se ha cargado información</div>
                                                </div>
                                            </td>
                                        </tr>
                                }
                            </tbody>
                        </table>
                    </div>
                    <div className="join">
                        <button className="join-item btn" onClick={() => setPage(prevPage => prevPage - 1)}>«</button>
                        <button className="join-item btn">Página {page.toString()}</button>
                        <button className="join-item btn" onClick={() => setPage(prevPage => prevPage + 1)}>»</button>
                    </div>
                </div>
            </div>
        </>
    )
}

export default List;