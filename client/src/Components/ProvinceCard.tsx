import React from 'react';
import { useNavigate } from 'react-router'
import { ProvinceProps } from '../Types/ProvinceProps.d';

// interface ProvinceProps {
//     name: string;
//     total: number
// }

const ProvinceCard: React.FC<ProvinceProps> = ({ codigo_juridiccion, Total }: ProvinceProps) => {
    let navigate = useNavigate();
    return (
        <>
            <div className='card card-side bg-base-100 shadow-md hover:shadow-2xl transition-shadow duration-300 overflow-hidden border-l-4 border-primary'>
                <div className='card-body py-4 px-6'>
                    <h2 className='text-lg font-bold text-base-content/70'>{codigo_juridiccion}</h2>
                    <div className='flex items-baseline gap-2'>
                        <p className='text-2xl font-black italic'>{Total}</p>
                        <span className='text-xs opacity-50'>Registros</span>
                    </div>
                    <div className='card-actions justify-end'>
                        <button className="btn btn-ghost btn-xs text-primary hover:bg-primary/10"
                        onClick={() => {navigate(`/contribuyentes/${codigo_juridiccion.toLowerCase()}`)}}
                        >
                            Ver detalles →
                        </button>
                    </div>
                </div>
            </div>
        </>

    )
}


export default ProvinceCard;