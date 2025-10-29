import { useParams } from "react-router"


const List: React.FC = () => {
    const params = useParams()
    console.log(params)
    const provinceName = params.province_name;
    return (
        <>
            <h1>
                {provinceName}
            </h1>
        </>
    )
}

export default List;