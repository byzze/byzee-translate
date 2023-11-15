import { useState, useEffect } from 'react';
import logo from './assets/images/logo-universal.png';
import { BrowserRouter } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import { useTheme } from 'next-themes';

import Translate from './window/Translate';
import Screenshot from './window/Screenshot';
import { useConfig } from './hooks';
import { EventsOn } from '../wailsjs/runtime';
import './i18n';
import './style.css';


function App() {
    const [appTheme] = useConfig('app_theme', 'system');
    const [appLanguage] = useConfig('app_language', 'zh_cn');
    const { setTheme } = useTheme();
    const { i18n } = useTranslation();
    const [isScreenshot, setIsScreenshot] = useState(false);

    useEffect(() => {
        EventsOn("ocrShow", (result) => {
            setIsScreenshot(result)
        })
    }, [])

    useEffect(() => {
        if (appTheme !== null) {
            if (appTheme !== 'system') {
                setTheme(appTheme);
            } else {
                try {
                    if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
                        setTheme('dark');
                    } else {
                        setTheme('light');
                    }
                    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
                        if (e.matches) {
                            setTheme('dark');
                        } else {
                            setTheme('light');
                        }
                    });
                } catch {
                    warn("Can't detect system theme.");
                }
            }
        }
        if (appLanguage !== null) {
            i18n.changeLanguage(appLanguage);
        }
    }, [appTheme, appLanguage]);

    return <BrowserRouter>
        {isScreenshot ? (
            <Screenshot />
        ) : (
            <Translate />
        )}
    </BrowserRouter>;
}

export default App
