import {  useEffect, useState } from "react";
import ProvinceCard from "../Components/ProvinceCard";
import { GetGrouping } from '../../src/Services/ContribuyenteService.ts';
import { ProvinceProps } from "../Types/ProvinceProps.d";

// const dataProvinces = [{ 'id': 1, 'name': 'Carchi', 'total': 10 }, { 'id': 2, 'name': 'Imbabura', 'total': 20 }, { 'id': 3, 'name': 'Pichincha', 'total': 20 }, { 'id': 4, 'name': 'Cotopaxi', 'total': 20 }, { 'id': 5, 'name': 'Tungurahua', 'total': 20 }, { 'id': 6, 'name': 'Chimborazo', 'total': 20 }, { 'id': 7, 'name': 'Manabí', 'total': 20 }, { 'id': 8, 'name': 'Bolivar', 'total': 20 }, { 'id': 9, 'name': 'Cañar', 'total': 20 }, { 'id': 10, 'name': 'Azuay', 'total': 20 }]

// console.log(dataProvinces);

const HomePage: React.FC = () => {

    const [dataProvinces, setDataProvinces] = useState<ProvinceProps[]>([]);

    async function data(): Promise<ProvinceProps[]> {
        try {
            const d: ProvinceProps[] = await GetGrouping();
            return d;
        } catch (error) {
            console.error('Hubo un problema con la petición fetch:', error);
            return []
        }
    }


    useEffect(() => {
        data()
        .then(d => {
            setDataProvinces(d);
        })
        .catch(e => console.error(e))

    }, [])

    return (
        <div className="card p-4 md:p-8 bg-base-100">
            <div className="card-body">
                <h2 className="card-title">Provincias del Ecuador:</h2>
                <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-3 gap-4">
                    {dataProvinces.length > 0 ?
                        (
                            dataProvinces.map(function (e) {
                                return <ProvinceCard 
                                        codigo_juridiccion={e.codigo_juridiccion} 
                                        Total={e.Total} 
                                        key={e.codigo_juridiccion} 
                                    />
                            })
                            // <ProvinceCard name="Guayas" total={9898} />
                        ) :
                        (
                            <h1>Sin datos</h1>
                        )
                    }

                </div>

            </div>
        </div>
    )
}

export default HomePage