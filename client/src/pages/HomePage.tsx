import ProvinceCard from "../Components/ProvinceCard";

const dataProvinces = [{ 'id': 1, 'name': 'Carchi', 'total': 10 }, { 'id': 2, 'name': 'Imbabura', 'total': 20 }, { 'id': 3, 'name': 'Pichincha', 'total': 20 }, { 'id': 4, 'name': 'Cotopaxi', 'total': 20 }, { 'id': 5, 'name': 'Tungurahua', 'total': 20 }, { 'id': 6, 'name': 'Chimborazo', 'total': 20 }, { 'id': 7, 'name': 'Manabí', 'total': 20 }, { 'id': 8, 'name': 'Bolivar', 'total': 20 }, { 'id': 9, 'name': 'Cañar', 'total': 20 }, { 'id': 10, 'name': 'Azuay', 'total': 20 }]

console.log(dataProvinces);

const HomePage = () => {
    return (
        <div className="card p-4 md:p-8 bg-base-200">
            <div className="card-body">
                <h2 className="card-title">Provincias del Ecuador:</h2>
                <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4 shadow-xl rounded-box p-0 bg-base-100">
                    {dataProvinces.length > 0 ?
                        (
                            dataProvinces.map(function (e) {
                                return <ProvinceCard name={e.name} total={e.total} key={e.id} />
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