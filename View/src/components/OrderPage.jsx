import React, { useEffect, useState } from 'react';
import { useCookies } from 'react-cookie';
import { Snackbar, Alert } from '@mui/material';

function OrderPage() {
    const [responseMessage, setResponseMessage] = useState('');
    const [allStorageData, setAllStorageData] = useState({});
    const [cookie, setCookie, removeCookie] = useCookies(['Address']);
    const [cookies] = useCookies(['customerData']);
    const [cId, setCId] = useState(-1);
    const [guest, setGuest] = useState(0);
    const [openSnackbar, setOpenSnackbar] = useState(false); // Snackbar'ın açık/kapalı durumu
    const [snackbarMessage, setSnackbarMessage] = useState(''); // Snackbar'da gösterilecek mesaj
    const [snackbarSeverity, setSnackbarSeverity] = useState('info'); // Snackbar mesajının türü (success, error, warning, info)

    const [sendBasket, setSendBasket] = useState({
        Address: {},
        Baskets: [],
        CustomerId: 0,
        Date: '',
        IsGuest: 0,
    });

    const getCurrentDate = () => {
        const today = new Date();
        const date = today.getFullYear() + '-' + (today.getMonth() + 1) + '-' + today.getDate();
        const time = today.getHours() + '-' + today.getMinutes() + '-' + today.getSeconds();
        return date + ' ' + time;
    };

    useEffect(() => {
        setAllStorageData(getAllLocalStorage());
    }, []);

    useEffect(() => {
        if (sendBasket.Baskets.length > 0) {
            sendBaskets();
        }
    }, [sendBasket]);

    useEffect(() => {
        setSendBasket(prevState => ({
            ...prevState,
            CustomerId: cId,
            IsGuest: guest,
        }));
    }, [cId, guest]);

    const getAllLocalStorage = () => {
        let keys = Object.keys(localStorage);
        let localStorageData = {};

        keys.forEach(key => {
            localStorageData[key] = JSON.parse(localStorage.getItem(key));
        });
        return localStorageData;
    };

    const clearLocalStorage = () => {
        localStorage.clear();
        setAllStorageData({});
    };

    const changeSetBasket = () => {
        const newItems = [];
        Object.keys(allStorageData).forEach((key) => {
            allStorageData[key].forEach(item => {
                const newItem = {
                    Id: item.object.Id.toString(),
                    Size: item.size,
                    Amount: item.amount.toString(),
                };
                newItems.push(newItem);
            });
        });

        if (cookies.customerData == null) {
            setCId(0);
            setGuest(1);
        } else {
            setCId(cookies.customerData.Id);
            setGuest(0);
        }

        setSendBasket(prevState => ({
            ...prevState,
            Address: cookie.Address,
            Baskets: newItems,
            Date: getCurrentDate(),
        }));
    };

    const sendBaskets = () => {
        if (sendBasket.CustomerId !== -1) {
            const formDataObj = new FormData();
            console.log(sendBasket)
            formDataObj.append('baskets', JSON.stringify(sendBasket));
            fetch('destination service address', {
                method: 'POST',
                body: formDataObj,
            })
                .then(response => response.json())
                .then(data => {
                    setSnackbarMessage(data.Text);
                    setSnackbarSeverity(data.Ret === 1 ? 'success' : 'error');
                    setOpenSnackbar(true);
                    if (data.Ret === 1) {
                        clearLocalStorage();
                    }
                    setResponseMessage('Data fetched successfully');
                })
                .catch(error => {
                    console.error('Error sending data:', error);
                    setResponseMessage('Error occurred while sending data.');
                });
        }
    };

    const handleClick = () => {
        changeSetBasket();
    };

    const handleSnackbarClose = () => {
        setOpenSnackbar(false);
    };

    return (
        <div>
            <button className="checkout-btn" onClick={handleClick}>SİPARİŞ VER</button>
            <Snackbar
                open={openSnackbar}
                autoHideDuration={6000}
                onClose={handleSnackbarClose}
                anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
            >
                <Alert onClose={handleSnackbarClose} severity={snackbarSeverity} sx={{ width: '100%' }}>
                    {snackbarMessage}
                </Alert>
            </Snackbar>
        </div>
    );
}

export default OrderPage;
