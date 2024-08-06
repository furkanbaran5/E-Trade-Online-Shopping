import React, { useEffect, useState } from 'react'
import Logo from '../img/logo.png'
import { useCookies } from 'react-cookie';

const Header = () => {
    const [userData, setUserData] = useState(null);
    const [text, setText] = useState('');
    const [cookies, setCookies, removeCookies] = useCookies(['customerData']);

    useEffect(() => {
        if (cookies.customerData == null) {
            setText(<a href="/login">Giriş</a>)
        } else {
            setText(<div>
                <a href="/customer">{cookies.customerData.name}</a>
            </div>
            )
        }
    }, [])
    return (
        <div>
            <div className="top-nav">
                <img src={Logo} alt="Logo" className="logo"></img>
                <nav className="main-menu">
                    <a href="#">Kadın</a>
                    <a href="#">Erkek</a>
                    <a href="#">Çocuk</a>
                </nav>
                <div className="search-bar">
                    <input type="text" placeholder="Aradığınız ürün, marka veya kategoriyi yazınız" />
                    <button><i className="search-icon"></i></button>
                </div>
                <div className="user-cart">
                    {text}

                </div>
                <div>
                    <a href="/basketPage" >Sepetim</a>
                </div>
            </div>
            <nav className="bottom-nav">
                <a href="/page2/shoes" >Ayakkabı</a>
                <a href="/page2/pants" >Pantolon</a>
                <a href="/page2/tshirt" >Tişört</a>
                <a href="/page2/hat" >Şapka</a>
                <a href="/page2/shorts" >Şort</a>
                <a href="/">Tüm Kategoriler</a>
            </nav>
        </div>
    )
}

export default Header