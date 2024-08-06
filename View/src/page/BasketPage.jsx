import React, { useEffect, useState } from 'react'
import Header from '../components/Header'
import '../css/basket.css'
import AddressPage from '../components/AddressPage';
import { useCookies } from 'react-cookie';
import OrderPage from '../components/OrderPage';
import AddressPage2 from '../components/AddressPage2';
function Page3() {

    const [responseMessage, setResponseMessage] = useState('');
    const [allStorageData, setAllStorageData] = useState({});
    const [totalPrice, setTotalPrice] = useState(0);
    const [cargoPrice, setCargoPrice] = useState(0);
    const [cargoPriceDisplay, setCargoPriceDisplay] = useState('');
    const [component, setComponent] = useState('');
    const [component2, setComponent2] = useState('');
    const [component3, setComponent3] = useState('');
    const [cookie, setCookie, removeCookie] = useCookies(['Address']);


    const openAddressPage = (control) => {
        if (control != 1) {
            setComponent(<AddressPage />)
        }
        setComponent2(<OrderPage />)
    }

    const openAddressPage2 = () => {
        setComponent3(<AddressPage2 onNewAddressClick={openAddressPage} />)
    }

    const sendOrder = () => {

    }
    useEffect(() => {
        setAllStorageData(getAllLocalStorage());
    }, []);

    const getAllLocalStorage = () => {
        let keys = Object.keys(localStorage);
        let localStorageData = {};

        keys.forEach(key => {
            localStorageData[key] = JSON.parse(localStorage.getItem(key));
        });
        return localStorageData;
    };

    useEffect(() => {
        let newTotalPrice = 0;
        Object.keys(allStorageData).forEach((key) => {
            allStorageData[key].forEach(item => {
                newTotalPrice += (item.amount * item.object.Price);
            });
        });
        setTotalPrice(newTotalPrice);
    }, [allStorageData]);

    useEffect(() => {
        if (totalPrice > 600) {
            setCargoPrice(0);
            setCargoPriceDisplay(<span>Ücretsiz <del>50₺</del></span>);
        } else {
            setCargoPriceDisplay(<span>50₺</span>);
            setCargoPrice(50);
        }
    }, [totalPrice])

    const removeItemFromLocalStorage = (key) => {
        localStorage.removeItem(key);
        setAllStorageData(prevState => {
            const updatedState = { ...prevState };
            delete updatedState[key];
            return updatedState;
        });
    };

    const clearLocalStorage = () => {
        localStorage.clear();
        setAllStorageData({});
    };

    const increaseAmount = (item) => {
        const existingData = JSON.parse(localStorage.getItem(item.object.Id));

        let flag = 1;
        for (let i = 0; i < existingData.length; i++) {
            if (existingData[i].size == item.size) {//Aynı beden ise sayı arttır
                if (existingData[i].amount != 5) {
                    existingData[i].amount += 1;
                    flag = 0;
                }
            }
        }
        if (flag == 0) {
            localStorage.setItem(item.object.Id, JSON.stringify(existingData));
            setAllStorageData(getAllLocalStorage());
        }

    };

    const decreaseAmount = (item) => {
        const existingData = JSON.parse(localStorage.getItem(item.object.Id));
        let flag = 1;
        for (let i = 0; i < existingData.length; i++) {
            if (existingData[i].size == item.size) {//Aynı beden ise sayı arttır
                if (existingData[i].amount != 1) {
                    existingData[i].amount -= 1;
                    flag = 0;
                }
            }
        }
        if (flag == 0) {
            localStorage.setItem(item.object.Id, JSON.stringify(existingData));
            setAllStorageData(getAllLocalStorage());
        }
    };



    return (
        <div className="main-body">
            <Header />
            <div className="container">
                <div className="cart-header">
                    <h1>Sepetim</h1>
                </div>
                {Object.keys(allStorageData).map((key) => (
                    allStorageData[key].map(item => (
                        <div className="cart-item">
                            <div className="cart-item-info">
                                <div className="basket-info">
                                    <img src={item.object.ImageUrl}></img>
                                    <div className="basket-details">
                                        <div className="basket-name">{item.object.Brand}    {item.object.Model}</div>
                                        <div className="basket-size">Beden: {item.size}</div>
                                        <div className="delivery">Tahmini Kargoya Teslimat: 13 - 16 Temmuz</div>
                                    </div>
                                </div>
                            </div>
                            <div className="cart-item-actions">
                                <div className="quantity">
                                    {item.amount > 1 && (<button className="quantity-btn" onClick={() => decreaseAmount(item)}>-</button>)}
                                    <span className="quantity-number">{item.amount}</span>
                                    {item.amount < 5 && (<button className="quantity-btn" onClick={() => increaseAmount(item)}>+</button>)}
                                </div>
                                <div className="price">{item.amount * item.object.Price}₺</div>
                                <button className="delete-btn" onClick={() => removeItemFromLocalStorage(item.object.Id)}>Sil</button>
                            </div>
                        </div>

                    ))
                ))}

                <div className="cart-summary">
                    <div className="summary-item">
                        <span>Ürünler</span>
                        <span>{totalPrice}₺</span>
                    </div>
                    <div className="summary-item">
                        <span>Kargo</span>
                        {cargoPriceDisplay}
                    </div>
                    <div className="summary-total">
                        <span>Toplam</span>
                        <span>{cargoPrice + totalPrice}₺</span>
                    </div>
                    <div>
                        <button className="checkout-btn" onClick={openAddressPage2}>SEPETİ ONAYLA</button>

                    </div>
                    <div >
                        <button className="checkout2-btn" onClick={clearLocalStorage}>SEPETİ SİL</button>
                    </div>
                </div>
                <div>
                    {component3}
                </div>
                <div>
                    {component}
                    {component2}
                </div>
            </div>
        </div >
    )
}

export default Page3