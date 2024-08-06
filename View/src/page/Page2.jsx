import React from 'react'
import Header from '../components/Header'
import DataUploadWithId from '../components/DataUploadWithId'
import { useParams } from 'react-router-dom'

const Shoes = () => {
    let { type } = useParams();
    return (
        <div className='main-body'>
            <div>
                <Header />
            </div>
            <div>
                <DataUploadWithId content={type} />
            </div>
        </div>
    )
}

export default Shoes