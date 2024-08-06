import React, { useEffect, useState } from 'react';
import '../css/basket.css';
import { useCookies } from 'react-cookie';

function AllOrders() {
    const [cookies, setCookies, removeCookies] = useCookies(['customerData']);
    const [responseMessage, setResponseMessage] = useState('');
    const [orders, setOrders] = useState([]);
    const [products, setProducts] = useState([{
        Id: '',
        Type: '',
        Brand: '',
        Model: '',
        Color: '',
        Price: '',
        ImageUrl: [],
        Size: [],
    }]);

    const getProduct = (id) => {
        if (!products[id]) {
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
                    setProducts(prevProducts => ({
                        ...prevProducts,
                        [id]: data
                    }));
                    setResponseMessage('Data fetched successfully');
                })
                .catch(error => {
                    console.error('Error sending data:', error);
                    setResponseMessage('Error occurred while sending data.');
                });
        }
    };

    useEffect(() => {
        const formDataObj = new FormData();
        formDataObj.append('getOrder', JSON.stringify(cookies.customerData.Id));

        fetch('destination service address', {
            method: 'POST',
            body: formDataObj,
        })
            .then(response => response.json())
            .then(data => {
                setOrders(data);
                setResponseMessage('Data fetched successfully');
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.');
            });
    }, [cookies.customerData.Id]);
    return (
        <div>
            <header>
                <h1>Tüm Siparişlerim</h1>
            </header>

            {orders.length > 0 ? (
                orders.map((order) => (
                    <div key={order.Date}>
                        <h4>Tarih: {order.Date}</h4>
                        <hr />
                        {order.Baskets.map((basket) => {
                            getProduct(basket.Id);
                            const product = products[basket.Id];
                            if (product != null) {
                                return (
                                    <div className="cart-item">
                                        <div className="cart-item-info">
                                            <div className="basket-info">
                                                <div><img src={product.ImageUrl[0]}></img></div>
                                                <div className="basket-details">
                                                    <div className="basket-name">{product.Brand}</div>
                                                    <div>{product.Model}</div>
                                                    <div className="basket-size">Beden: {basket.Size}</div>
                                                    <div></div>
                                                </div>
                                            </div>
                                        </div>
                                        <div className="cart-item-actions">
                                            <div className="quantity">
                                                <span className="quantity-number">Adet: {basket.Amount}</span>
                                            </div>
                                            <div className="price">Ücret: {basket.Amount * product.Price}</div>
                                        </div>
                                    </div>
                                );
                            }

                        })}
                        <hr />
                    </div>
                ))
            ) : (
                <p>Geçmiş Sipariş Bulunmamaktadır</p>
            )}
        </div>
    );
}

export default AllOrders;
