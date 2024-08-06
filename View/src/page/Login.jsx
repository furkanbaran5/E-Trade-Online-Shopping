import React, { useState } from 'react';
import Header from '../components/Header';
import '../css/login.css';
import { useNavigate } from 'react-router-dom';
import { useCookies } from 'react-cookie';

function Login() {
    const navigate = useNavigate();
    const [activeTab, setActiveTab] = useState('login');
    const [isCheckedPromotion, setIsCheckedPromotion] = useState(false);
    const [isCheckedRecommendation, setIsCheckedRecommendation] = useState(false);
    const [responseMessage, setResponseMessage] = useState('');
    const [cookies, setCookie] = useCookies(['customerData']);
    const [errors, setErrors] = useState({});

    const openTab = (tabName) => {
        setActiveTab(tabName);
    }

    const [customer, setCustomer] = useState({});

    const handleInputRegister = (e) => {
        const { name, value } = e.target;
        setCustomer(prevState => ({
            ...prevState,
            [name]: value,
        }));
    }

    const allFieldsFilled = () => {
        return (
            customer.name &&
            customer.surname &&
            customer.phoneNumber &&
            customer.mail &&
            customer.password
        );
    }

    const validateInputs = () => {
        let tempErrors = {};
        if (!/^\d{3}\d{3}\d{4}$/.test(customer.phoneNumber)) {
            tempErrors.phoneNumber = "Geçerli bir telefon numarası giriniz.";
        }
        if (!/\S+@\S+\.\S+/.test(customer.mail)) {
            tempErrors.mail = "Geçerli bir e-posta adresi giriniz.";
        }
        setErrors(tempErrors);
        return Object.keys(tempErrors).length === 0;
    }

    const register = (e) => {
        e.preventDefault();
        if (validateInputs()) {
            registerWithDB();
        } else {
            setResponseMessage("Lütfen tüm alanları doğru doldurun.");
        }
    }

    const registerWithDB = () => {
        const formDataObj = new FormData();
        formDataObj.append('register', JSON.stringify(customer));
        fetch('destination service address', {
            method: 'POST',
            body: formDataObj,
        })
            .then(response => response.json())
            .then(data => {
                if (data.Ret == 1) {
                    setResponseMessage('Kayıt Başarılı. Giriş Ekranına Yönlendiriliyorsunuz...');
                    setTimeout(() => {
                        window.location.reload();
                    }, 2000);
                } else {
                    setResponseMessage(data.Text);
                }
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.');
            });
    }

    const [loginInput, setloginInput] = useState({
        phoneNumberOrMail: '',
        password: '',
    });

    const handleInputLogin = (e) => {
        const { name, value } = e.target;
        setloginInput(prevState => ({
            ...prevState,
            [name]: value,
        }));
    }

    const allFieldsFilled2 = () => {
        return (
            loginInput.phoneNumberOrMail &&
            loginInput.password
        );
    }

    const login = (e) => {
        e.preventDefault();
    }

    const loginWithDB = () => {
        const formDataObj = new FormData();
        formDataObj.append('login', JSON.stringify(loginInput));
        fetch('destination service address', {
            method: 'POST',
            body: formDataObj,
        })
            .then(response => response.json())
            .then(data => {
                if (data == null) {
                    setResponseMessage('Kullanıcı adı veya Hatalı şifre');
                } else {
                    setCookie('customerData', data, { path: '/' })
                    setResponseMessage("Giriş Başarılı Yönlendiriliyorsunuz...");
                    setTimeout(() => {
                        navigate("/")
                    }, 2000);
                }
            })
            .catch(error => {
                console.error('Error sending data:', error);
                setResponseMessage('Error occurred while sending data.');
            });
    }

    return (
        <div>
            <Header />

            <div className="container-login">
                <div className="tab-menu">
                    <button
                        className={`tab-button ${activeTab === 'login' ? 'active' : ''}`}
                        onClick={() => openTab('login')}
                    >
                        Giriş Yap
                    </button>
                    <button
                        className={`tab-button ${activeTab === 'signup' ? 'active' : ''}`}
                        onClick={() => openTab('signup')}
                    >
                        Üye Ol
                    </button>
                </div>
                <div id="login" className={`tab-content ${activeTab === 'login' ? 'active' : ''}`}>

                    <form onSubmit={login}>
                        <div style={{ textAlign: "center" }}>
                            {responseMessage}
                        </div>
                        <input
                            type="text"
                            name="phoneNumberOrMail"
                            placeholder="E-Posta Adresi veya Telefon"
                            value={loginInput.phoneNumberOrMail}
                            onChange={handleInputLogin}
                        />
                        <input
                            type="password"
                            name="password"
                            placeholder="Şifre"
                            value={loginInput.password}
                            onChange={handleInputLogin}
                        />
                        <button
                            className="btn-submit"
                            type="submit"
                            disabled={!allFieldsFilled2()}
                            onClick={loginWithDB}
                        >
                            GİRİŞ YAP
                        </button>
                    </form>
                </div>
                <div id="signup" className={`tab-content ${activeTab === 'signup' ? 'active' : ''}`}>
                    {responseMessage && (
                        <div style={{ textAlign: "center" }}>
                            {responseMessage}
                        </div>
                    )}
                    <form onSubmit={register}>
                        <input
                            type="text"
                            name="name"
                            placeholder="İsim"
                            value={customer.name}
                            onChange={handleInputRegister}
                        />
                        <input
                            type="text"
                            name="surname"
                            placeholder="Soyisim"
                            value={customer.surname}
                            onChange={handleInputRegister}
                        />
                        <input
                            type="text"
                            name="phoneNumber"
                            placeholder="(5xx) xxx xx xx "
                            value={customer.phoneNumber}
                            onChange={handleInputRegister}
                        />
                        {errors.phoneNumber && <div className="error-message">{errors.phoneNumber}</div>}
                        <input
                            type="email"
                            name="mail"
                            placeholder="E-Posta Adresi"
                            value={customer.mail}
                            onChange={handleInputRegister}
                        />
                        {errors.mail && <div className="error-message">{errors.mail}</div>}
                        <input
                            type="password"
                            name="password"
                            placeholder="Şifre"
                            value={customer.password}
                            onChange={handleInputRegister}
                        />

                        <div className="checkbox-group">
                            <label>
                                <input
                                    type="checkbox"
                                    checked={isCheckedPromotion}
                                    onChange={(e) => setIsCheckedPromotion(e.target.checked)}
                                />
                                Ürün tanıtım ve kampanyalardan haberdar olmak için elektronik ileti almak istiyorum.
                            </label>
                            <label>
                                <input
                                    type="checkbox"
                                    checked={isCheckedRecommendation}
                                    onChange={(e) => setIsCheckedRecommendation(e.target.checked)}
                                />
                                Tercihlerime özel ürün önerilmesini ve tanıtılmasını kabul ediyorum.
                            </label>
                        </div>
                        <div className="info-textx">
                            Kişisel verileriniz Aydınlatma Metni uyarınca işlenecektir. "Üye Ol" butonuna basarak Üyelik Sözleşmesi'ni kabul ediyorsunuz.
                        </div>
                        <button
                            className="btn-submit"
                            type="submit"
                            disabled={!allFieldsFilled()}
                        >
                            ÜYE OL
                        </button>
                    </form>
                </div>
            </div>
        </div>
    );
}

export default Login;
