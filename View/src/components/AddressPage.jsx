import React, { useEffect, useState } from 'react'
import '../css/address.css'
import { useCookies } from 'react-cookie';
import { Snackbar, Alert } from '@mui/material';

function AddressPage() {
    const [address, setAddress] = useState({
        title: "",
        name: "",
        surname: "",
        phoneNumber: "",
        city: "",
        ilce: "",
        adres: "",
        isRegister: "0",
    });
    const [cookie, setCookie] = useCookies(['Address']);
    const [cities, setCities] = useState([]);
    const [districts, setDistricts] = useState([]);
    const [responseMessage, setResponseMessage] = useState('');
    const [open, setOpen] = useState(false);
    const [snackbarMessage, setSnackbarMessage] = useState('');
    const [snackbarSeverity, setSnackbarSeverity] = useState('info');

    const handleSnackbarClose = () => {
        setOpen(false);
    };

    useEffect(() => {
        fetchCities();
    }, []);

    const fetchCities = () => {
        const formDataObj = new FormData();
        formDataObj.append('city', "allcity");

        fetch('destination service address', {
            method: 'POST',
            body: formDataObj,
        })
            .then(response => response.json())
            .then(data => {
                setCities(data);
                setResponseMessage(data.message);
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.');
            });
    };

    const handleChange = (e) => {
        const { name, value } = e.target;
        if (name == "isRegister") {
            if (e.target.checked == true) {
                setAddress({ ...address, [name]: "1" });
            } else {
                setAddress({ ...address, [name]: "0" });
            }
        } else {
            setAddress({ ...address, [name]: value });
            if (name == "city") {
                fetchDistrict(value)
            }
        }

    };

    const fetchDistrict = (value) => {
        const formDataObj = new FormData();
        formDataObj.append('district', value);

        fetch('destination service address', {
            method: 'POST',
            body: formDataObj,
        })
            .then(response => response.json())
            .then(data => {
                setDistricts(data);
                setResponseMessage(data.message);
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.');
            });
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        const isValid = Object.values(address).every(value => value && value.trim() !== '');
        if (!isValid) {
            setSnackbarMessage("Lütfen gerekli alanları doldurun.");
            setSnackbarSeverity('error');
            setOpen(true);
        } else {
            setCookie('Address', address, { path: '/' });
            setSnackbarMessage("Adres kaydedilmiştir.");
            setSnackbarSeverity('success');
            setOpen(true);
        }
    };

    return (
        <div className="form-container">
            <div className="form-header">
                <h2>Yeni Adres Ekle</h2>
            </div>
            <div className="form-group">
                <label htmlFor="title">Adres Başlığı</label>
                <input type="text" name="title" value={address.title} onChange={handleChange} />
            </div>
            <div className="form-group-new">
                <div>
                    <label htmlFor="name">Ad</label>
                    <input type="text" name="name" value={address.name} onChange={handleChange} />
                </div>
                <div>
                    <label htmlFor="surname">Soyad</label>
                    <input type="text" name="surname" value={address.surname} onChange={handleChange} />
                </div>
            </div>
            <div className="form-group">
                <label htmlFor="phoneNumber">Cep Telefonu</label>
                <input type="text" name="phoneNumber" value={address.phoneNumber} onChange={handleChange} />
            </div>
            <div className="form-group-new">
                <div>
                    <label htmlFor="city">İl</label>
                    <select name="city" value={address.city} onChange={handleChange} >
                        <option value="">Seçin</option>
                        {cities.length > 0 ? (
                            cities.map((city) => (
                                <option value={city}>{city}</option>
                            ))
                        ) : (
                            <option value="">No Data</option>
                        )}
                    </select>
                </div>
                <div>
                    <label htmlFor="ilce">İlçe</label>
                    <select name="ilce" value={address.ilce} onChange={handleChange}>
                        <option value="">Seçin</option>
                        {districts.length > 0 ? (
                            districts.map((district) => (
                                <option value={district}>{district}</option>
                            ))
                        ) : (
                            <option value="">No Data</option>
                        )}
                    </select>
                </div>
            </div>
            <div className="form-group">
                <label htmlFor="adres">Adres</label>
                <textarea name="adres" value={address.adres} onChange={handleChange}></textarea>
                <label>
                    <input
                        type="checkbox"
                        name="isRegister"
                        onChange={handleChange}
                    />
                    Adres bilgilerimin kayıt edilmesini istiyorum
                </label>
            </div>
            <button className="checkout-btn" onClick={handleSubmit}>Adresi Kaydet</button>
            <Snackbar
                open={open}
                message="Adres Başarı ile kaydedilmiştir."
                autoHideDuration={6000}
                onClose={handleSnackbarClose}
                anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
            >
                <Alert severity={snackbarSeverity} sx={{ width: '100%' }}>
                    {snackbarMessage}
                </Alert>
            </Snackbar>
        </div >
    )
}

export default AddressPage

