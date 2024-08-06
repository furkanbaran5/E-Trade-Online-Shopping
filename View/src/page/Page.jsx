import React, { useEffect, useState } from 'react'
import Header from '../components/Header'
import '../css/product.css'
import { useParams } from 'react-router-dom'
import { Snackbar, Alert } from '@mui/material';

function Page() {
    let { id } = useParams();
    const [formData, setFormData] = useState({
        Id: '',
        Type: '',
        Brand: '',
        Model: '',
        Color: '',
        Price: '',
        ImageUrl: [],
        Size: [],
    });

    const [responseMessage, setResponseMessage] = useState('');
    const [selectedSize, setSelectedSize] = useState('');
    const [open, setOpen] = useState(false);
    const [snackbarMessage, setSnackbarMessage] = useState('');
    const [snackbarSeverity, setSnackbarSeverity] = useState('info');

    useEffect(() => {
        const formDataObj = new FormData();
        formDataObj.append('id', id);
        fetch('destination service address', {
            method: 'POST',
            body: formDataObj,
        })
            .then(response => {
                return response.json();
            })
            .then(data => {
                setFormData(data);
                setResponseMessage('Data fetched successfully');
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.');
            });
    }, []);

    const handleSnackbarClose = () => {
        setOpen(false);
    };
    const handleSizeChange = (event) => {
        setSelectedSize(event.target.value);
    };

    const saveToLocalStorage = () => {
        if (selectedSize != '') {
            const existingData = JSON.parse(localStorage.getItem(formData.Id));
            if (existingData) {//Var ise
                let flag = 1;
                for (let i = 0; i < existingData.length; i++) {
                    if (existingData[i].size == selectedSize) {//Aynı beden ise sayı arttır
                        existingData[i].amount += 1;
                        flag = 0;
                    }
                }
                if (flag == 1) {//Farklı beden ise üstüne ekle
                    existingData.push({ object: formData, size: selectedSize, amount: 1 })
                    console.log(existingData.size)
                }
                localStorage.setItem(formData.Id, JSON.stringify(existingData));
            } else {//Yoksa yeni aç
                const newBasket = [{ object: formData, size: selectedSize, amount: 1 }];
                localStorage.setItem(formData.Id, JSON.stringify(newBasket));
            }
            setSnackbarMessage("Sepete eklenmiştir.");
            setSnackbarSeverity('success');
            setOpen(true);
        } else {
            setSnackbarMessage("Lütfen Beden Seçiniz.");
            setSnackbarSeverity('error');
            setOpen(true);
        }
    };

    return (
        <div>
            <Header />
            <div className="product-detail">
                <div className="product-image2">
                    {(formData.ImageUrl).map(name => (
                        <img src={name} alt="Ürün Resmi"></img>
                    ))}
                </div>
                <div className="product-info">
                    <div className="product-title">{(formData.Brand).toUpperCase()}</div>

                    <div className="product-description">

                        <p>{formData.Model.toUpperCase()}</p>
                        <p>Renk:{formData.Color.toUpperCase()}</p>
                    </div>
                    <div className="size-info">
                        <a href="#" className="size-chart">Beden Tablosu</a>
                    </div>
                    <div className="sizes" >
                        {
                            (formData.Size) != null ? (
                                (formData.Size).map((size, index) => (
                                    <div>
                                        <input type="radio" id={size} name="size" value={size} checked={selectedSize === size} onChange={handleSizeChange}  ></input>
                                        <label htmlFor={size} className="size">{size}</label>
                                    </div>
                                ))
                            ) : (
                                <div>
                                    <h2>Stokta yok</h2>
                                </div>
                            )}

                    </div>
                    <div className="product-price">{formData.Price}₺</div>
                    <div className="product-actions">
                        <button onClick={saveToLocalStorage}>Sepete Ekle</button>
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
                    </div>
                </div>
            </div>
        </div >
    )
}

export default Page