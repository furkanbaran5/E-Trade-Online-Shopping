import React, { useState, useEffect } from 'react';
import '../css/style.css';

function DataUploadWithId({ content }) {

    const [formDataArray, setFormDataArray] = useState('');
    const [responseMessage, setResponseMessage] = useState('');
    useEffect(() => {
        const formDataObj = new FormData();
        formDataObj.append('name', content);
        fetch('destination service address', {
            method: 'POST',
            body: formDataObj,
        })
            .then(response => {
                return response.json();
            })
            .then(data => {
                setFormDataArray(data);
                setResponseMessage('Data fetched successfully');
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.');
            });
    }, [content]);

    return (
        <div className="product-container">
            {formDataArray.length > 0 ? (
                formDataArray.map((item, index) => (
                    <a href={`/page/${item.Id}`} >
                        <div className="product-card">
                            <img src={item.ImageUrl[0]} alt="Product Image" className="product-image" />
                            <h2 className="product-title">
                                {item.Brand}
                            </h2>
                            {item.Model}
                            <div style={{ display: "flex", justifyContent: "end", margin: "20px" }}>
                                <div style={{ color: "red", fontSize: "25px", fontWeight: "bold" }}>
                                    {item.Price}₺
                                </div>
                            </div>
                        </div>
                    </a>
                ))
            ) : (
                <p>No data available</p>
            )
            }
        </div >
    );
}

export default DataUploadWithId;
