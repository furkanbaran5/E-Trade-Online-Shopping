import React, { useState } from 'react'
import Header from '../components/Header'
import '../css/customer.css'
import { useCookies } from 'react-cookie';
import AdressInfo from '../components/AdressInfo';
import AllOrders from '../components/AllOrders';
import CardInfo from '../components/CardInfo';
import CommunicationInfo from '../components/CommunicationInfo';
import CustomerInfo from '../components/CustomerInfo';

function Customer() {
    const [cookies, setCookies, removeCookies] = useCookies(['customerData']);
    const [component, setComponent] = useState(<AllOrders />);
    const [activeTab, setActiveTab] = useState('AllOrders');

    const openTab = (tabName, component) => {
        setActiveTab(tabName);
        setComponent(component)
    }

    const handleRemoveCookie = () => {
        removeCookies('customerData', { path: '/' });
    };

    return (
        <div className="main-body">
            <Header />
            <div className="containerx">
                <aside className="sidebarx">
                    <nav>
                        <ul>
                            <li
                                className={`${activeTab === 'AllOrders' ? 'active' : ''}`}
                                onClick={() => openTab('AllOrders', <AllOrders />)}
                            >
                                <a >Tüm Siparişlerim</a>
                            </li>
                            <li
                                className={`${activeTab === 'CustomerInfo' ? 'active' : ''}`}
                                onClick={() => openTab('CustomerInfo', <CustomerInfo />)}
                            >
                                <a>Kullanıcı Bilgilerim</a>
                            </li>
                            <li
                                className={`${activeTab === 'AdressInfo' ? 'active' : ''}`}
                                onClick={() => openTab('AdressInfo', <AdressInfo />)}
                            >
                                <a >Adres Bilgilerim</a>
                            </li>
                            <li
                                className={`${activeTab === 'CardInfo' ? 'active' : ''}`}
                                onClick={() => openTab('CardInfo', <CardInfo />)}
                            >
                                <a >Kayıtlı Kartlarım<span className="new">Yeni</span></a>
                            </li>
                            <li
                                className={`${activeTab === 'CommunicationInfo' ? 'active' : ''}`}
                                onClick={() => openTab('CommunicationInfo', <CommunicationInfo />)}
                            >
                                <a >İletişim Tercihlerim</a>
                            </li>
                        </ul>
                    </nav>
                    <a href="/login"><button className="logout" onClick={handleRemoveCookie} >Çıkış</button></a>
                </aside>
                <main className="content" >
                    {component}
                </main>

            </div>
        </div>
    )
}

export default Customer