import { contribuyente } from "../Types/Contribuyentes.d";
import { ProvinceProps } from "../Types/ProvinceProps.d";

async function Get(provinceName: string, page: number): Promise<contribuyente[]> {
    try {
        var httpResponse = await fetch(`http://127.0.0.1:3000/api/v1/contribuyentes?provincia=${provinceName.toUpperCase()}&page=${page}&limit=10`)

        if (!httpResponse.ok)
            throw new Error(`Error HTTP: ${httpResponse.status}`);

        // 3. Convertir la respuesta a JSON
        const datos = await httpResponse.json();

        return datos['Value'] as contribuyente[]
    } catch (e) {
        console.error(e);
        return []
    }
}

async function GetGrouping(): Promise<ProvinceProps[]> {
    try {
        var httpResponse = await fetch(`http://127.0.0.1:3000/api/v1/contribuyentes/group`)

        if (!httpResponse.ok)
            throw new Error(`Error HTTP: ${httpResponse.status}`);

        // 3. Convertir la respuesta a JSON
        const datos = await httpResponse.json();

        return datos['Value'] as ProvinceProps[]
    } catch (e) {
        console.error(e);
        return []
    }
}

export { Get, GetGrouping }