import React, { useEffect, useState } from 'react'
import { useCookies } from 'react-cookie';

function AdressInfo() {
    const [address, setAddress] = useState([]);
    const [responseMessage, setResponseMessage] = useState('');
    const [cookies] = useCookies(['customerData']);

    useEffect(() => {
        const formDataObj = new FormData();
        formDataObj.append('address', cookies.customerData.Id);

        fetch('destination service address', {
            method: 'POST',
            body: formDataObj,
        })
            .then(response => response.json())
            .then(data => {
                setAddress(data);
                setResponseMessage(data.message);
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.');
            });
    }, [cookies.customerData.Id]);
    return (
        <div >
            <header>
                <h1>Adres Bilgileri</h1>
            </header>
            {address === null ? (
                <p>Kayıtlı adres bilgisi yok</p>
            ) : (
                address.map((adres, index) => (
                    <div key={index} className="address-card">
                        <label>
                            <h3>{adres.title}</h3>
                            <p>{adres.city} / {adres.ilce}</p>
                            <p>{adres.adres}</p>
                            <p>{adres.name} {adres.surname} - {adres.phoneNumber}</p>
                        </label>
                    </div>
                ))
            )}
        </div>
    )
}

export default AdressInfo