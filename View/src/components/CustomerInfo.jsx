import React, { useEffect, useState } from 'react';
import { useCookies } from 'react-cookie';

function CustomerInfo() {
    const [cookies] = useCookies(['customerData']);
    const [formData, setFormData] = useState({});
    const [responseMessage, setResponseMessage] = useState('');
    const [editableData, setEditableData] = useState({
        id: cookies.customerData.Id,
        mail: '',
        phoneNumber: '',
        name: '',
        surname: '',
    });

    useEffect(() => {
        const formDataObj = new FormData();
        formDataObj.append('customerInfo', cookies.customerData.Id);
        fetch('destination service address', {
            method: 'POST',
            body: formDataObj,
        })
            .then(response => {
                return response.json();
            })
            .then(data => {
                setFormData(data);  // Set initial editable data
                setResponseMessage('Data fetched successfully');
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.');
            });
    }, [cookies.customerData.Id]);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setEditableData({ ...editableData, [name]: value });
    };
    const allFieldsFilled = () => {
        return (
            editableData.name &&
            editableData.surname &&
            editableData.phoneNumber &&
            editableData.mail
        );
    }
    const handleSubmit = (e) => {
        e.preventDefault();
        const formDataObj = new FormData();
        formDataObj.append('updateCustomerInfo', JSON.stringify(editableData));
        fetch('destination service address', {
            method: 'POST',
            body: formDataObj,
        })
            .then(response => response.json())
            .then(data => {
                setResponseMessage('Data updated successfully');
            })
            .catch(error => {
                console.error('Error updating data:', error);
                setResponseMessage('Error occurred while updating data.');
            });
    };

    return (
        <div className="containery">
            <div className="user-info">
                <h2>Mevcut Kullanıcı Bilgileri</h2>
                <p><strong>Mail:</strong> {formData.mail}</p>
                <p><strong>Telefon Numarası:</strong> {formData.phoneNumber}</p>
                <p><strong>İsim:</strong> {formData.name}</p>
                <p><strong>Soyisim:</strong> {formData.surname}</p>
            </div>
            <div className="edit-form">
                <h2>Bilgileri Güncelle</h2>
                <form onSubmit={handleSubmit}>
                    <label>
                        İsim:
                        <input
                            type="text"
                            name="name"
                            onChange={handleChange}
                        />
                    </label>
                    <label>
                        Soyisim:
                        <input
                            type="text"
                            name="surname"
                            onChange={handleChange}
                        />
                    </label>
                    <label>
                        Mail:
                        <input
                            type="email"
                            name="mail"
                            onChange={handleChange}
                        />
                    </label>
                    <label>
                        Telefon Numarası:
                        <input
                            type="text"
                            name="phoneNumber"
                            onChange={handleChange}
                        />
                    </label>
                    <button type="submit" disabled={!allFieldsFilled()}>Değişiklikleri Kaydet</button>
                </form>
            </div>
        </div>
    );
}

export default CustomerInfo;
