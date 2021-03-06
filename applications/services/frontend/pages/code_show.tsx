import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import React, { useEffect, useState } from 'react'
import Countdown from 'react-countdown'
import styles from '../styles/Home.module.css'
import Link from 'next/link'
import Menu from '../components/menu'

const CodeShow: NextPage = () => {
    const [code, setCode] = useState({code: 0, unix: 999999999999999});

    useEffect(() => {
        setCode(JSON.parse(localStorage.getItem("code") || '{}'))
    }, []);

    const renderer = ({ minutes, seconds, completed }: { minutes: number, seconds: number, completed: boolean }) => {
        if (completed) {
            return <span className="badge bg-danger"><h1>TIMES UP</h1></span>;
        } else {
            return <span className="badge bg-success"><h1>{minutes}:{seconds}</h1></span>;
        }
    };

    return (
        <div className={styles.container}>
            <Head>
                <title>AA - Code Show</title>
                <link rel="icon" href="/favicon.ico" />
            </Head>
            
            <Menu />

            <main className={styles.main}>
                <h1 className={styles.title}>
                    <Link href="/code_show"><a>{code.code}</a></Link>
                </h1>

                <p className={styles.description}>
                    <Countdown date={code.unix} renderer={renderer} />
                </p>
            </main>
        </div>
    )
}

export default CodeShow