import React, { useEffect, useState } from 'react';
import { useCookies } from 'react-cookie';
import '../css/style.css';

function AddressPage2({ onNewAddressClick }) {
    const [address, setAddress] = useState([]);
    const [cookies] = useCookies(['customerData']);
    const [cookie, setCookie] = useCookies(['Address']);
    const [responseMessage, setResponseMessage] = useState('');


    useEffect(() => {
        if (cookies.customerData != null) {
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
        }
    }, [cookies.customerData]);

    const handleAddressSelect = (adres) => {
        setCookie('Address', JSON.stringify(adres), { path: '/' });
        onNewAddressClick(1);
    };

    return (
        <div className="delivery-options">
            <h2>Teslimat Seçenekleri</h2>
            <div className="delivery-option">
                {address === null ? (
                    <p>Kayıtlı adres bilgisi yok</p>
                ) : (
                    address.map((adres, index) => (
                        <div key={index} className="address-card">
                            <input
                                type="radio"
                                id={`address${index}`}
                                name="address"
                                onChange={() => handleAddressSelect(adres)}
                            />
                            <label htmlFor={`address${index}`}>
                                <h3>{adres.title}</h3>
                                <p>{adres.city} / {adres.ilce}</p>
                                <p>{adres.adres}</p>
                                <p>{adres.name} {adres.surname} - {adres.phoneNumber}</p>
                            </label>
                        </div>
                    ))
                )}
            </div>
            <button className="checkout-btn" onClick={onNewAddressClick}>Yeni Adres Ekle</button>
        </div>
    );
}

export default AddressPage2;
