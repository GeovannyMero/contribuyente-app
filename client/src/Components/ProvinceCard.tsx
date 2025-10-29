import React from 'react';
import {useNavigate} from 'react-router'
import { ProvinceProps } from '../Types/ProvinceProps.d';

// interface ProvinceProps {
//     name: string;
//     total: number
// }

const ProvinceCard: React.FC<ProvinceProps> = ({ name, total }: ProvinceProps) => {
    let navigate = useNavigate();
    return (
        <>
            {/* <div className="stats shadow"> */}
                <div className="stat">
                    <div className="stat-title">Provincia:</div>
                    <div className="stat-value">{name}</div>
                    <div className="stat-desc">{total} contribuyentes</div>
                    <div className='stat-actions'>
                        <button 
                            className="btn btn-xs btn-accent" 
                            onClick={() => {navigate(`/contribuyentes/${name.toLowerCase()}`)}}
                        >Visualizar
                        </button>
                    </div>
                </div>
            {/* </div> */}
        </>

    )
}


export default ProvinceCard;