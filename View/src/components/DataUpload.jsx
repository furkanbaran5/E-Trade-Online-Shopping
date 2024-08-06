import React, { useEffect, useState } from 'react'
function DataUpload({ content }) {
    const [formData, setFormData] = useState("");
    const [responseMessage, setResponseMessage] = useState('');
    useEffect(() => {
        const formDataObj = new FormData();
        formDataObj.append('name', content);

        fetch('destination service address', {
            method: 'POST',
            body: formDataObj,
        })
            .then(response => response.json())
            .then(data => {
                setFormData(data);
                setResponseMessage(data.message);
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.');
            });
    }, []);

    return (
        <div className="product-container">
            {Object.keys(formData).map(key => (
                formData[key].map(item => (
                    <a href={`/page/${item.Id}`} >
                        <div className="product-card">
                            <img src={item.ImageUrl[0]} alt="Product Image" className="product-image"></img>
                            <h2 className="product-title">
                                {item.Brand}
                            </h2>
                            {item.Model}
                            <div style={{ display: "flex", justifyContent: "end", margin: "20px" }}>
                                <div style={{ color: "red", fontSize: "30px", fontWeight: "bold" }}>
                                    {item.Price}₺
                                </div>
                            </div>
                        </div>
                    </a>
                ))
            ))}
        </div>


    );
};

export default DataUpload